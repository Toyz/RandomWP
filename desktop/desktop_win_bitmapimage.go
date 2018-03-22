// +build windows

package desktop

import (
	"image"
	"image/color"
	"unsafe"
)

//
// Bitmap
//

type BitmapImage struct {
	hbm   HBITMAP
	image image.Image
}

func BitmapImageNew(i image.Image) *BitmapImage {
	m := &BitmapImage{}

	m.image = i

	w := m.Width()
	h := m.Height()

	screenDC := HDCPtr(GetDC.Call(NULL))
	if screenDC == 0 {
		panic(GetLastErrorString())
	}
	defer ReleaseDC.Call(NULL, Arg(screenDC))

	memDC := HDCPtr(CreateCompatibleDC.Call(Arg(screenDC)))
	defer DeleteDC.Call(Arg(memDC))
	if memDC == 0 {
		panic(GetLastErrorString())
	}

	bmi := &BITMAPINFO{}
	bmi.bmiHeader.biSize = DWORD(unsafe.Sizeof(bmi.bmiHeader))
	bmi.bmiHeader.biWidth = LONG(w)
	bmi.bmiHeader.biHeight = LONG(h)
	bmi.bmiHeader.biPlanes = 1
	bmi.bmiHeader.biBitCount = 32
	bmi.bmiHeader.biCompression = BI_RGB
	bmi.bmiHeader.biSizeImage = DWORD(w * h * 4)

	var ppvBits unsafe.Pointer

	m.hbm = HBITMAPPtr(CreateDIBSection.Call(Arg(memDC), Arg(bmi), Arg(DIB_RGB_COLORS), Arg(&ppvBits), NULL, NULL))
	if m.hbm == 0 {
		panic(GetLastErrorString())
	}

	buf := (*(*[1 << 28]uint32)(ppvBits))[:(w * h)] // 30 too large on 32-bit

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := i.At(i.Bounds().Min.X+x, i.Bounds().Min.Y+y)
			r, g, b, a := color.RGBAModel.Convert(c).RGBA()

			r = r & 0xff
			g = g & 0xff
			b = b & 0xff
			a = a & 0xff

			// windows images are upside down and BGRA encoded.
			// yes. it is not just slash \\
			wy := h - y - 1
			buf[x+wy*w] = (r << 16) | (g << 8) | (b << 0) | (a << 24)
		}
	}

	return m
}

func (m *BitmapImage) Close() {
	m.hbm.Close()
}

func (m *BitmapImage) Width() int {
	return m.image.Bounds().Max.X - m.image.Bounds().Min.X
}

func (m *BitmapImage) Height() int {
	return m.image.Bounds().Max.Y - m.image.Bounds().Min.Y
}

func (m *BitmapImage) Draw(x LONG, y LONG, hdcDst HDC) {
	cx := LONG(m.Width())
	cy := LONG(m.Height())

	hdcSrc := HDCPtr(CreateCompatibleDC.Call(Arg(hdcDst)))
	if hdcSrc == 0 {
		panic(GetLastErrorString())
	}
	old := HANDLEPtr(SelectObject.Call(Arg(hdcSrc), Arg(m.hbm)))

	bld := BLENDFUNCTION{}
	bld.BlendOp = AC_SRC_OVER
	bld.BlendFlags = 0
	bld.SourceConstantAlpha = 255
	bld.AlphaFormat = AC_SRC_ALPHA

	if !BOOLPtr(AlphaBlend.Call(Arg(hdcDst), Arg(x), Arg(y), Arg(cx), Arg(cy),
		Arg(hdcSrc), NULL, NULL, Arg(cx), Arg(cy),
		Arg(*(*uint32)(unsafe.Pointer(&bld))))).Bool() {
		panic(GetLastErrorString())
	}

	SelectObject.Call(Arg(hdcSrc), Arg(old))

	if !BOOLPtr(DeleteDC.Call(Arg(hdcSrc))).Bool() {
		panic(GetLastErrorString())
	}
}
