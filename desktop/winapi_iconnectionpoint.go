// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	IID_IConnectionPoint = GUIDNew("{B196B286-BAB4-101A-B69C-00AA00341D07}")
)

type IConnectionPoint struct {
	IUnknown
}

type IConnectionPointVtbl struct {
	IUnknownVtbl
	GetConnectionInterface      uintptr
	GetConnectionPointContainer uintptr
	Advise                      uintptr
	Unadvise                    uintptr
	EnumConnections             uintptr
}

func (v *IConnectionPoint) VTable() *IConnectionPointVtbl {
	return (*IConnectionPointVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IConnectionPoint) Advise(pUnkSink *IUnknown, pdwCookie *DWORD) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Advise,
		3,
		Arg(m),
		Arg(pUnkSink),
		Arg(pdwCookie),
	))
}

func (m *IConnectionPoint) Unadvise(cookie DWORD) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Unadvise,
		2,
		Arg(m),
		Arg(cookie),
		0,
	))
}
