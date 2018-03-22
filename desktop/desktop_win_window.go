// +build windows

package desktop

type Window struct {
	WndClassEx *WNDCLASSEX
	Wnd        HWND
}

const (
	HWND_DESKTOP = 0
)

func WindowNew(w WNDPROC, wndclass string, exstyle uint32, style uint32, width, height int) *Window {
	m := &Window{}

	hinstance := HINSTANCEPtr(GetModuleHandle.Call(Arg(0)))

	m.WndClassEx = WNDCLASSEXNew(hinstance, w, wndclass)

	m.Wnd = HWNDPtr(CreateWindowEx.Call(Arg(exstyle),
		Arg(m.WndClassEx.lpszClassName), Arg(m.WndClassEx.lpszClassName),
		Arg(style),
		Arg(0), Arg(0), Arg(width), Arg(height),
		HWND_DESKTOP, NULL, Arg(hinstance), NULL))

	if m.Wnd == 0 {
		panic(GetLastErrorString())
	}

	return m
}

func (m *Window) DefWindowProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	return LRESULTPtr(DefWindowProc.Call(Arg(hWnd), Arg(msg), Arg(wParam), Arg(lParam)))
}

func (m *Window) Close() {
	PostMessage.Call(Arg(m.Wnd), Arg(WM_QUIT), NULL, NULL)
	m.Wnd.Close()
	m.WndClassEx.Close()
}
