// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	IID_IConnectionPointContainer = GUIDNew("{B196B284-BAB4-101A-B69C-00AA00341D07}")
)

type IConnectionPointContainer struct {
	IUnknown
}

type IConnectionPointContainerVtbl struct {
	IUnknownVtbl
	EnumConnectionPoints uintptr
	FindConnectionPoint  uintptr
}

func (v *IConnectionPointContainer) VTable() *IConnectionPointContainerVtbl {
	return (*IConnectionPointContainerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IConnectionPointContainer) FindConnectionPoint(guid *GUID, ret **IConnectionPoint) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().FindConnectionPoint,
		3,
		Arg(m),
		Arg(guid),
		Arg(ret),
	))
}
