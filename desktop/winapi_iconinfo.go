// +build windows

package desktop

type ICONINFO struct {
	fIcon    BOOL
	xHotspot DWORD
	yHotspot DWORD
	hbmMask  HBITMAP
	hbmColor HBITMAP
}
