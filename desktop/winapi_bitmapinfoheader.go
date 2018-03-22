// +build windows

package desktop

type BITMAPINFOHEADER struct {
	biSize          DWORD
	biWidth         LONG
	biHeight        LONG
	biPlanes        WORD
	biBitCount      WORD
	biCompression   DWORD
	biSizeImage     DWORD
	biXPelsPerMeter LONG
	biYPelsPerMeter LONG
	biClrUsed       DWORD
	biClrImportant  DWORD
}
