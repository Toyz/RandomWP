// +build windows

package desktop

import (
	"syscall"
)

type HICON uintptr

func HICONPtr(r1, r2 uintptr, err error) HICON {
	LastError = uintptr(err.(syscall.Errno))
	return HICON(r1)
}

func HICONNew(bm HBITMAP) HICON {
	info := &ICONINFO{}
	info.fIcon = BOOL(Bool2Int[true])
	info.hbmMask = bm
	info.hbmColor = bm
	h := HICONPtr(CreateIconIndirect.Call(Arg(info)))
	if h == 0 {
		panic(GetLastErrorString())
	}
	return h
}

func (m HICON) Close() {
	DeleteObject.Call(Arg(m))
}
