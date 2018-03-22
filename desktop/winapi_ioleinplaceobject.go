// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

//
// IOleInPlaceObject
//

var IID_IOleInPlaceObject = GUIDNew("{00000113-0000-0000-C000-000000000046}")

type IOleInPlaceObject struct {
	IUnknown
}

type IOleInPlaceObjectVtbl struct {
	IUnknownVtbl
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
	InPlaceDeactivate    uintptr
	UIDeactivate         uintptr
	SetObjectRects       uintptr
}

func (v *IOleInPlaceObject) VTable() *IOleInPlaceObjectVtbl {
	return (*IOleInPlaceObjectVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IOleInPlaceObject) SetObjectRects(rect *CRECT, rect2 *CRECT) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().SetObjectRects,
		3,
		Arg(m),
		Arg(rect),
		Arg(rect2),
	))
}
