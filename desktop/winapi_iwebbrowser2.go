// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	CLSID_WebBrowser = GUIDNew("8856F961-340A-11D0-A96B-00C04FD705A2")
	IID_IWebBrowser2 = GUIDNew("D30C1661-CDAF-11d0-8A3E-00C04FC9E26E")
)

type READYSTATE int

const (
	READYSTATE_UNINITIALIZED READYSTATE = 0
	READYSTATE_LOADING                  = 1
	READYSTATE_LOADED                   = 2
	READYSTATE_INTERACTIVE              = 3
	READYSTATE_COMPLETE                 = 4
)

type IWebBrowser2 struct {
	IDispatch
}

type IWebBrowser2Vtbl struct {
	IDispatchVtbl
	GoBack                uintptr
	GoForward             uintptr
	GoHome                uintptr
	GoSearch              uintptr
	Navigate              uintptr
	Refresh               uintptr
	Refresh2              uintptr
	Stop                  uintptr
	get_Application       uintptr
	get_Parent            uintptr
	get_Container         uintptr
	get_Document          uintptr
	get_TopLevelContainer uintptr
	get_Type              uintptr
	get_Left              uintptr
	put_Left              uintptr
	get_Top               uintptr
	put_Top               uintptr
	get_Width             uintptr
	put_Width             uintptr
	get_Height            uintptr
	put_Height            uintptr
	get_LocationName      uintptr
	get_LocationURL       uintptr
	get_Busy              uintptr
	Quit                  uintptr
	ClientToWindow        uintptr
	PutProperty           uintptr
	GetProperty           uintptr
	get_Name              uintptr
	get_HWND              uintptr
	get_FullName          uintptr
	get_Path              uintptr
	get_Visible           uintptr
	put_Visible           uintptr
	get_StatusBar         uintptr
	put_StatusBar         uintptr
	get_StatusText        uintptr
	put_StatusText        uintptr
	get_ToolBar           uintptr
	put_ToolBar           uintptr
	get_MenuBar           uintptr
	put_MenuBar           uintptr
	get_FullScreen        uintptr
	put_FullScreen        uintptr
	Navigate2             uintptr
	QueryStatusWB         uintptr
	ExecWB                uintptr
	ShowBrowserBar        uintptr
	get_ReadyState        uintptr
}

func (v *IWebBrowser2) VTable() *IWebBrowser2Vtbl {
	return (*IWebBrowser2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IWebBrowser2) Navigate2(url *VARIANT, Flags *VARIANT, TargetFrameName *VARIANT, PostData *VARIANT, Headers *VARIANT) HRESULT {
	return HRESULTPtr(syscall.Syscall6(
		m.VTable().Navigate2,
		6,
		Arg(m),
		Arg(url),
		Arg(Flags),
		Arg(TargetFrameName),
		Arg(PostData),
		Arg(Headers),
	))
}

func (m *IWebBrowser2) put_Left(i LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().put_Left,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) put_Top(i LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().put_Top,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) put_Height(i LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().put_Height,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) put_Width(i LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().put_Width,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) get_Width(i *LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().get_Width,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) get_Height(i *LONG) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().get_Height,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) put_Visible(i VARIANT_BOOL) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().put_Visible,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) get_Visible(i *VARIANT_BOOL) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().get_Visible,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) get_ReadyState(i *READYSTATE) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().get_ReadyState,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) get_Document(i **IDispatch) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().get_Document,
		2,
		Arg(m),
		Arg(i),
		0,
	))
}

func (m *IWebBrowser2) Stop() HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Stop,
		1,
		Arg(m),
		0,
		0,
	))
}
