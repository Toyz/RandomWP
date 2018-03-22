// +build windows

package desktop

import (
	"syscall"
)

type HBITMAP uintptr

func HBITMAPPtr(r1, r2 uintptr, err error) HBITMAP {
	LastError = uintptr(err.(syscall.Errno))
	return HBITMAP(r1)
}

func (m HBITMAP) Close() {
	DeleteObject.Call(Arg(m))
}
