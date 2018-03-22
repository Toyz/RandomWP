// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

// http://www.codeproject.com/Articles/3365/Embed-an-HTML-control-in-your-own-window-using-pla

type IOleInPlaceUIWindow IUnknown
type IDataObject IUnknown
type OLECONTAINER IUnknown

type OLEINPLACEFRAMEINFO struct {
	cb            UINT
	fMDIApp       BOOL
	hwndFrame     HWND
	haccel        HACCEL
	cAccelEntries UINT
}

const (
	DOCHOSTUIFLAG_NO3DBORDER = 0x4
	DOCHOSTUIDBLCLK_DEFAULT  = 0
)

type DOCHOSTUIINFO struct {
	cbSize        ULONG
	dwFlags       DWORD
	dwDoubleClick DWORD
	pchHostCss    *OLECHAR
	pchHostNS     *OLECHAR
}

//
// IOleClientSite
//

type IMoniker IUnknown

var IID_IOleClientSite = GUIDNew("{00000118-0000-0000-C000-000000000046}")

type IOleClientSite struct {
	IUnknown
	inplace *IOleInPlaceSite
	frame   *IOleInPlaceFrame
	ui      *IDocHostUIHandler
}

type IOleClientSiteVtbl struct {
	IUnknownVtbl
	SaveObject             uintptr
	GetMoniker             uintptr
	GetContainer           uintptr
	ShowObject             uintptr
	OnShowWindow           uintptr
	RequestNewObjectLayout uintptr
}

func IOleClientSiteNew(hwnd HWND, browser *IOleObject, s *DesktopSysTray) *IOleClientSite {
	m := &IOleClientSite{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IOleClientSiteVtbl{}))
	m.New()
	m.frame = IOleInPlaceFrameNew(hwnd)
	m.inplace = IOleInPlaceSiteNew(hwnd, m.frame, browser)
	m.ui = IDocHostUIHandlerNew(hwnd, browser, s)
	m.AddInterface(&IID_IOleInPlaceFrame, m.frame)
	m.AddInterface(&IID_IOleInPlaceSite, m.inplace)
	m.AddInterface(&IID_IDocHostUIHandler, m.ui)
	return m
}

func (m *IOleClientSite) New() {
	m.IUnknown.New()
	m.VTable().SaveObject = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().GetMoniker = syscall.NewCallback(func(m *IOleClientSite, dwAssign DWORD, dwWhichMoniker DWORD, ppmk *[1]*IMoniker) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().GetContainer = syscall.NewCallback(func(m *IOleClientSite, ppContainer *[1]*OLECONTAINER) HRESULT {
		ppContainer[0] = nil
		return E_NOINTERFACE
	})
	m.VTable().ShowObject = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return S_OK
	})
	m.VTable().OnShowWindow = syscall.NewCallback(func(m *IOleClientSite, fShow BOOL) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().RequestNewObjectLayout = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return E_NOTIMPL
	})
	m.AddInterface(&IID_IOleClientSite, m)
}

func (v *IOleClientSite) VTable() *IOleClientSiteVtbl {
	return (*IOleClientSiteVtbl)(unsafe.Pointer(v.RawVTable))
}

//
// IID_IOleInPlaceSite
//

var IID_IOleInPlaceSite = GUIDNew("{00000119-0000-0000-C000-000000000046}")

type IOleInPlaceSite struct {
	IUnknown
	hwnd    HWND
	frame   *IOleInPlaceFrame
	browser *IOleObject
}

type IOleInPlaceSiteVtbl struct {
	IUnknownVtbl
	GetWindow            uintptr
	ContextSensitiveHelp uintptr
	CanInPlaceActivate   uintptr
	OnInPlaceActivate    uintptr
	OnUIActivate         uintptr
	GetWindowContext     uintptr
	Scroll               uintptr
	OnUIDeactivate       uintptr
	OnInPlaceDeactivate  uintptr
	DiscardUndoState     uintptr
	DeactivateAndUndo    uintptr
	OnPosRectChange      uintptr
}

func IOleInPlaceSiteNew(hwnd HWND, frame *IOleInPlaceFrame, browser *IOleObject) *IOleInPlaceSite {
	m := &IOleInPlaceSite{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IOleInPlaceSiteVtbl{}))
	m.hwnd = hwnd
	m.frame = frame
	m.browser = browser
	m.New()
	return m
}

func (m *IOleInPlaceSite) New() {
	m.IUnknown.New()
	m.VTable().GetWindow = syscall.NewCallback(func(m *IOleInPlaceSite, lphwnd *[1]HWND) HRESULT {
		lphwnd[0] = m.hwnd
		return S_OK
	})
	m.VTable().ContextSensitiveHelp = syscall.NewCallback(func(m *IOleClientSite, fEnterMode BOOL) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().CanInPlaceActivate = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return S_OK
	})
	m.VTable().OnInPlaceActivate = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return S_OK
	})
	m.VTable().OnUIActivate = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return S_OK
	})
	m.VTable().GetWindowContext = syscall.NewCallback(func(m *IOleInPlaceSite, lplpFrame *[1]*IOleInPlaceFrame, lplpDoc *[1]*IOleInPlaceUIWindow, lprcPosRect *RECT, lprcClipRect *RECT, lpFrameInfo *OLEINPLACEFRAMEINFO) HRESULT {
		lplpDoc[0] = nil
		lplpFrame[0] = m.frame
		lpFrameInfo.fMDIApp = 0
		lpFrameInfo.hwndFrame = m.hwnd
		lpFrameInfo.haccel = 0
		lpFrameInfo.cAccelEntries = 0
		return S_OK
	})
	m.VTable().Scroll = syscall.NewCallback(func(m *IOleClientSite, scrollExtant uintptr) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().OnUIDeactivate = syscall.NewCallback(func(m *IOleClientSite, fUndoable BOOL) HRESULT {
		return S_OK
	})
	m.VTable().OnInPlaceDeactivate = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return S_OK
	})
	m.VTable().DiscardUndoState = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().DeactivateAndUndo = syscall.NewCallback(func(m *IOleClientSite) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().OnPosRectChange = syscall.NewCallback(func(m *IOleInPlaceSite, lprcPosRect *CRECT) HRESULT {
		var inplace *IOleInPlaceObject
		if m.browser.QueryInterface(&IID_IOleInPlaceObject, &inplace) == S_OK {
			defer inplace.Release()
			inplace.SetObjectRects(lprcPosRect, lprcPosRect).S_OK()
		}
		return S_OK
	})
	m.AddInterface(&IID_IOleInPlaceSite, m)
}

func (v *IOleInPlaceSite) VTable() *IOleInPlaceSiteVtbl {
	return (*IOleInPlaceSiteVtbl)(unsafe.Pointer(v.RawVTable))
}

//
// IOleInPlaceFrame
//

type LPCBORDERWIDTHS uintptr
type LPOLEMENUGROUPWIDTHS uintptr

var IID_IOleWindow = GUIDNew("{00000114-0000-0000-C000-000000000046}")
var IID_IOleInPlaceUIWindow = GUIDNew("{00000115-0000-0000-C000-000000000046}")
var IID_IOleInPlaceFrame = GUIDNew("{00000116-0000-0000-C000-000000000046}")

type IOleInPlaceFrame struct {
	IUnknown
	hwnd HWND
}

type IOleInPlaceFrameVtbl struct {
	IUnknownVtbl
	GetWindow            uintptr // IOleWindow
	ContextSensitiveHelp uintptr
	GetBorder            uintptr // IOleInPlaceUIWindow
	RequestBorderSpace   uintptr
	SetBorderSpace       uintptr
	SetActiveObject      uintptr
	InsertMenus          uintptr // IOleInPlaceFrame
	SetMenu              uintptr
	RemoveMenus          uintptr
	SetStatusText        uintptr
	EnableModeless       uintptr
	TranslateAccelerator uintptr
}

func IOleInPlaceFrameNew(hwnd HWND) *IOleInPlaceFrame {
	m := &IOleInPlaceFrame{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IOleInPlaceFrameVtbl{}))
	m.hwnd = hwnd
	m.New()
	return m
}

func (m *IOleInPlaceFrame) New() {
	m.IUnknown.New()
	m.VTable().GetWindow = syscall.NewCallback(func(m *IOleInPlaceFrame, lphwnd *[1]HWND) HRESULT {
		lphwnd[0] = m.hwnd
		return S_OK
	})
	m.VTable().ContextSensitiveHelp = syscall.NewCallback(func(m *IOleClientSite, fEnterMode BOOL) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().GetBorder = syscall.NewCallback(func(m *IOleClientSite, lprectBorder LPRECT) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().RequestBorderSpace = syscall.NewCallback(func(m *IOleClientSite, pborderwidths LPCBORDERWIDTHS) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().SetBorderSpace = syscall.NewCallback(func(m *IOleClientSite, pborderwidths LPCBORDERWIDTHS) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().SetActiveObject = syscall.NewCallback(func(m *IOleClientSite, pActiveObject *IOleInPlaceActiveObject, pszObjName LPCOLESTR) HRESULT {
		return S_OK
	})
	m.VTable().InsertMenus = syscall.NewCallback(func(m *IOleClientSite, hmenuShared HMENU, lpMenuWidths LPOLEMENUGROUPWIDTHS) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().SetMenu = syscall.NewCallback(func(m *IOleClientSite, hmenuShared HMENU, holemenu HOLEMENU, hwndActiveObject HWND) HRESULT {
		return S_OK
	})
	m.VTable().RemoveMenus = syscall.NewCallback(func(m *IOleClientSite, hmenuShared HMENU) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().SetStatusText = syscall.NewCallback(func(m *IOleClientSite, pszStatusText LPCOLESTR) HRESULT {
		return S_OK
	})
	m.VTable().EnableModeless = syscall.NewCallback(func(m *IOleClientSite, fEnable BOOL) HRESULT {
		return S_OK
	})
	m.VTable().TranslateAccelerator = syscall.NewCallback(func(m *IOleClientSite, lpmsg LPMSG, wID WORD) HRESULT {
		return E_NOTIMPL
	})
	m.AddInterface(&IID_IOleWindow, m)
	m.AddInterface(&IID_IOleInPlaceUIWindow, m)
	m.AddInterface(&IID_IOleInPlaceFrame, m)
}

func (v *IOleInPlaceFrame) VTable() *IOleInPlaceFrameVtbl {
	return (*IOleInPlaceFrameVtbl)(unsafe.Pointer(v.RawVTable))
}

//
// IID_IDocHostUIHandler
//

type IOleInPlaceActiveObject IUnknown
type IOleCommandTarget IUnknown
type IDropTarget IUnknown

var IID_IDocHostUIHandler = GUIDNew("{bd3f23c0-d43e-11cf-893b-00aa00bdce1a}")

type IDocHostUIHandler struct {
	IUnknown
	hwnd    HWND
	browser *IOleObject
	s       *DesktopSysTray
}

type IDocHostUIHandlerVtbl struct {
	IUnknownVtbl
	ShowContextMenu       uintptr
	GetHostInfo           uintptr
	ShowUI                uintptr
	HideUI                uintptr
	UpdateUI              uintptr
	EnableModeless        uintptr
	OnDocWindowActivate   uintptr
	OnFrameWindowActivate uintptr
	ResizeBorder          uintptr
	TranslateAccelerator  uintptr
	GetOptionKeyPath      uintptr
	GetDropTarget         uintptr
	GetExternal           uintptr
	TranslateUrl          uintptr
	FilterDataObject      uintptr
}

func IDocHostUIHandlerNew(hwnd HWND, browser *IOleObject, s *DesktopSysTray) *IDocHostUIHandler {
	m := &IDocHostUIHandler{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IDocHostUIHandlerVtbl{}))
	m.hwnd = hwnd
	m.browser = browser
	m.s = s
	m.New()
	return m
}

func (m *IDocHostUIHandler) New() {
	m.IUnknown.New()
	m.VTable().ShowContextMenu = syscall.NewCallback(func(m *IDocHostUIHandler, dwID DWORD, ppt *POINT, pcmdtReserved *IUnknown, pdispReserved *IDispatch) HRESULT {
		var pt POINT
		GetCursorPos.Call(Arg(&pt))
		PostMessage.Call(Arg(m.hwnd), Arg(WM_CONTEXTMENU), Arg(pt.x), Arg(pt.y))
		return S_OK
	})
	m.VTable().GetHostInfo = syscall.NewCallback(func(m *IDocHostUIHandler, info *DOCHOSTUIINFO) HRESULT {
		info.cbSize = ULONG(unsafe.Sizeof(DOCHOSTUIINFO{}))
		info.dwFlags = DOCHOSTUIFLAG_NO3DBORDER
		info.dwDoubleClick = DOCHOSTUIDBLCLK_DEFAULT
		return S_OK
	})
	m.VTable().ShowUI = syscall.NewCallback(func(m *IDocHostUIHandler, dwID DWORD, pActiveObject *IOleInPlaceActiveObject, pCommandTarget *IOleCommandTarget, pFrame *IOleInPlaceFrame, pDoc *IOleInPlaceUIWindow) HRESULT {
		return S_OK
	})
	m.VTable().HideUI = syscall.NewCallback(func(m *IDocHostUIHandler) HRESULT {
		return S_OK
	})
	m.VTable().UpdateUI = syscall.NewCallback(func(m *IDocHostUIHandler) HRESULT {
		return S_OK
	})
	m.VTable().EnableModeless = syscall.NewCallback(func(m *IDocHostUIHandler, fEnable BOOL) HRESULT {
		return S_OK
	})
	m.VTable().OnDocWindowActivate = syscall.NewCallback(func(m *IDocHostUIHandler, fActivate BOOL) HRESULT {
		return S_OK
	})
	m.VTable().OnFrameWindowActivate = syscall.NewCallback(func(m *IDocHostUIHandler, fActivate BOOL) HRESULT {
		return S_OK
	})
	m.VTable().ResizeBorder = syscall.NewCallback(func(m *IDocHostUIHandler, prcBorder LPCRECT, pUIWindow *IOleInPlaceUIWindow, fRameWindow BOOL) HRESULT {
		return S_OK
	})
	m.VTable().TranslateAccelerator = syscall.NewCallback(func(m *IDocHostUIHandler, lpMsg LPMSG, pguidCmdGroup *GUID, nCmdID DWORD) HRESULT {
		return S_FALSE
	})
	m.VTable().GetOptionKeyPath = syscall.NewCallback(func(m *IDocHostUIHandler, pchKey *LPOLESTR, dw DWORD) HRESULT {
		return S_FALSE
	})
	m.VTable().GetDropTarget = syscall.NewCallback(func(m *IDocHostUIHandler, pDropTarget *IDropTarget, ppDropTarget *[1]*IDropTarget) HRESULT {
		return S_FALSE
	})
	m.VTable().GetExternal = syscall.NewCallback(func(m *IDocHostUIHandler, ppDispatch *[1]*IDispatch) HRESULT {
		ppDispatch[0] = nil
		return S_FALSE
	})
	m.VTable().TranslateUrl = syscall.NewCallback(func(m *IDocHostUIHandler, dwTranslate DWORD, pchURLIn *OLECHAR, ppchURLOut *[1]*OLECHAR) HRESULT {
		url := WString2String(uintptr(unsafe.Pointer(pchURLIn)))

		d := m.s.os.(*DesktopSysTrayWin)

		if d.WebCurrent.Url == url { // ignore translate currently loaded url
			return S_FALSE
		}

		if m.s.WebPopup.Handler != nil {
			if !m.s.WebPopup.Handler(url) {
				BrowserOpenURI(url)
			}
		}

		w := WStringNew("#")
		ppchURLOut[0] = (*OLECHAR)(unsafe.Pointer(w))
		return S_OK
	})
	m.VTable().FilterDataObject = syscall.NewCallback(func(m *IDocHostUIHandler, pDO *IDataObject, ppDORet *[1]*IDataObject) HRESULT {
		ppDORet[0] = nil
		return S_FALSE
	})
	m.AddInterface(&IID_IDocHostUIHandler, m)
}

func (v *IDocHostUIHandler) VTable() *IDocHostUIHandlerVtbl {
	return (*IDocHostUIHandlerVtbl)(unsafe.Pointer(v.RawVTable))
}

//
// DWebBrowserEvents2
//

const (
	DISPID_DOCUMENTCOMPLETE = 259 // new document goes ReadyState_Complete
)

type DWebBrowserEvents2 struct {
	IDispatch
	DocumentComplete func() HRESULT
}

var DIID_DWebBrowserEvents2 = GUIDNew("34A715A0-6587-11D0-924A-0020AFC7AC4D")

type DWebBrowserEvents2Vtbl struct {
	IDispatchVtbl
}

func DWebBrowserEvents2New() *DWebBrowserEvents2 {
	m := &DWebBrowserEvents2{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&DWebBrowserEvents2Vtbl{}))
	m.New()
	return m
}

func (m *DWebBrowserEvents2) New() {
	m.IDispatch.New()
	m.AddInterface(&DIID_DWebBrowserEvents2, m)
	m.AddMethod("DocumentComplete", DISPID_DOCUMENTCOMPLETE, syscall.NewCallback(func() HRESULT {
		return m.DocumentComplete()
	}))
}

func (v *DWebBrowserEvents2) VTable() *DWebBrowserEvents2Vtbl {
	return (*DWebBrowserEvents2Vtbl)(unsafe.Pointer(v.RawVTable))
}
