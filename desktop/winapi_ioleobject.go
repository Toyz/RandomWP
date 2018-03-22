// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	IID_IOleObject = GUIDNew("{00000112-0000-0000-C000-000000000046}")
)

type OLECLOSE DWORD

const (
	OLECLOSE_SAVEIFDIRTY OLECLOSE = 0
	OLECLOSE_NOSAVE               = 1
	OLECLOSE_PROMPTSAVE           = 2
)

type IOleObject struct {
	IUnknown
}

type IOleObjectVtbl struct {
	IUnknownVtbl
	SetClientSite    uintptr
	GetClientSite    uintptr
	SetHostNames     uintptr
	Close            uintptr
	SetMoniker       uintptr
	GetMoniker       uintptr
	InitFromData     uintptr
	GetClipboardData uintptr
	DoVerb           uintptr
}

func (v *IOleObject) VTable() *IOleObjectVtbl {
	return (*IOleObjectVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IOleObject) SetClientSite(s *IOleClientSite) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().SetClientSite,
		2,
		Arg(m),
		Arg(s),
		0,
	))
}

func (m *IOleObject) GetClientSite(s **IOleClientSite) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().GetClientSite,
		2,
		Arg(m),
		Arg(s),
		0,
	))
}

func (m *IOleObject) DoVerb(iVerb int, lpmsg *MSG, cli *IOleClientSite, lindex int, hwndParent HWND, rect *RECT) HRESULT {
	return HRESULTPtr(syscall.Syscall9(
		m.VTable().DoVerb,
		7,
		Arg(m),
		Arg(iVerb),
		Arg(lpmsg),
		Arg(cli),
		Arg(lindex),
		Arg(hwndParent),
		Arg(rect),
		0,
		0,
	))
}

func (m *IOleObject) Close(save DWORD) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Close,
		2,
		Arg(m),
		Arg(save),
		0,
	))
}
