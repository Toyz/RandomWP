// +build windows

package desktop

import (
	"unsafe"
)

type WNDCLASSEX struct {
	cbSize        UINT
	style         UINT
	lpfnWndProc   WNDPROC
	cbClsExtra    UINT
	cbWndExtra    UINT
	hInstance     HINSTANCE
	hIcon         HICON
	hCursor       HCURSOR
	hbrBackground HBRUSH
	lpszMenuName  LPCTSTR
	lpszClassName LPCTSTR
	hIconSm       HICON
}

func WNDCLASSEXNew(hInstance HINSTANCE, WndProc WNDPROC, klass string) *WNDCLASSEX {
	m := &WNDCLASSEX{}
	m.cbSize = UINT(unsafe.Sizeof(*m))
	m.style = 0
	m.lpfnWndProc = WndProc
	m.cbClsExtra = 0
	m.cbWndExtra = 0
	m.hInstance = hInstance
	m.hIcon = HICON(0)
	m.hbrBackground = HBRUSH(0)
	m.lpszMenuName = LPCTSTR(0)
	m.lpszClassName = LPCTSTR(WStringNew(klass))

	a := ATOMPtr(RegisterClassEx.Call(Arg(m)))
	if a == 0 {
		panic(GetLastErrorString())
	}

	return m
}

func (m *WNDCLASSEX) Close() {
	if !BOOLPtr(UnregisterClass.Call(Arg(m.lpszClassName), Arg(m.hInstance))).Bool() {
		panic(GetLastErrorString())
	}

	w := WString(m.lpszClassName)
	w.Close()
}
