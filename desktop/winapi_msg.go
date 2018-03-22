// +build windows

package desktop

type MSG struct {
	hwnd    HWND
	message UINT
	wParam  WPARAM
	lParam  LPARAM
	time    DWORD
	pt      POINT
}

type LPMSG *MSG
