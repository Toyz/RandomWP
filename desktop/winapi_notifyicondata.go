// +build windows

package desktop

import (
	"unsafe"
)

const (
	NIF_ICON    = 0x01
	NIF_MESSAGE = 0x02
	NIF_TIP     = 0x04
	NIF_INFO    = 0x10
)

type NOTIFYICONDATA struct {
	cbSize DWORD
	hWnd   HWND
	// id of icon - passed back in wParam of message
	uID    UINT
	uFlags UINT
	// notification message, pass to hwnd WM_USER + 1
	uCallbackMessage UINT
	hIcon            HICON
	szTip            [128]TCHAR
	dwState          DWORD
	dwStateMask      DWORD
	szInfo           [256]TCHAR
	union            UINT // {UINT uTimeout; UINT uVersion;};
	szInfoTitle      [64]TCHAR
	dwInfoFlags      DWORD
	guidItem         GUID
	hBalloonIcon     HICON
}

func NOTIFYICONDATANew() *NOTIFYICONDATA {
	m := &NOTIFYICONDATA{}
	m.cbSize = DWORD(unsafe.Sizeof(*m))
	return m
}

func (m *NOTIFYICONDATA) SetIcon(i HICON) {
	m.hIcon = i
	m.uFlags |= NIF_ICON
}

func (m *NOTIFYICONDATA) SetCallback(i UINT) {
	m.uCallbackMessage = i
	m.uFlags |= NIF_MESSAGE
}

func (m *NOTIFYICONDATA) SetTooltip(s string) {
	m.uFlags |= NIF_TIP
	//m.uFlags |= NIF_INFO

	p := WStringNew(s)
	defer p.Close()

	i := 0
	for p := uintptr(unsafe.Pointer(p)); ; p += 2 {
		u := *(*TCHAR)(unsafe.Pointer(p))
		if u == 0 {
			return
		}
		m.szTip[i] = u
		//m.szInfo[i] = u
		//m.szInfoTitle[i] = u
		i++
	}
}
