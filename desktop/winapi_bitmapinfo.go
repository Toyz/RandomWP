// +build windows

package desktop

type BITMAPINFO struct {
	bmiHeader BITMAPINFOHEADER
	bmiColors [1]RGBQUAD
}
