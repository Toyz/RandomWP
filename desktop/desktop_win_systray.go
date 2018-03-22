// +build windows

package desktop

import (
	"image"
	"runtime"
	"unsafe"

	"github.com/nfnt/resize"
)

const (
	SPACE_ICONS = 2
)

func ConvertMenuIcon(i image.Image) *BitmapImage {
	menubarHeigh := GetSystemMenuImageSize()

	c := resize.Resize(menubarHeigh, menubarHeigh, i, resize.Lanczos3)

	return BitmapImageNew(c)
}

func GetSystemMenuImageSize() uint {
	return uint(UINTPtr(GetSystemMetrics.Call(Arg(SM_CYMENUCHECK))))
}

func SystemMenuFontNew() HFONT {
	nm := NONCLIENTMETRICS{}
	nm.cbSize = UINT(unsafe.Sizeof(nm))

	if IsWindowsXP() {
		nm.cbSize = UINT(unsafe.Sizeof(NONCLIENTMETRICS_XP{}))
	}

	if !BOOLPtr(SystemParametersInfo.Call(Arg(SPI_GETNONCLIENTMETRICS), NULL, Arg(&nm), NULL)).Bool() {
		panic(GetLastErrorString())
	}

	h := HFONTPtr(CreateFontIndirect.Call(Arg(&nm.lfMenuFont)))
	if h == 0 {
		panic(GetLastErrorString())
	}

	return h
}

//
// MenuWin
//

type MenuItemWin struct {
	Menu  *Menu
	Image *BitmapImage
}

func (m *MenuItemWin) Close() {
	m.Image.Close()
}

//
// DesktopSysTrayWin
//

type DesktopSysTrayWin struct {
	MainMenu  HMENU
	MenuItems []*MenuItemWin
	Icon      HICON
	MainWnd   *Window

	Checked_png   *BitmapImage
	Unchecked_png *BitmapImage

	TaskbarCreated    WString
	WM_TASKBARCREATED UINT

	// WebView
	WebWnd     *Window
	browser    *IOleObject
	WebCurrent *WebPopup // current url loaded
}

func desktopSysTrayNew() *DesktopSysTray {
	// locket for thread message pump, will be unlocked after
	// desktopMain() exits
	runtime.LockOSThread()

	d := &DesktopSysTrayWin{}
	m := &DesktopSysTray{os: d}

	d.TaskbarCreated = WStringNew("TaskbarCreated")
	d.WM_TASKBARCREATED = UINTPtr(RegisterWindowMessage.Call(Arg(d.TaskbarCreated)))
	d.MainWnd = WindowNew(WNDPROCNew(m.WndProc), "SystrayIcon", 0, WS_OVERLAPPEDWINDOW, 0, 0)

	d.Checked_png = ConvertMenuIcon(DecodeImageString(checked_png))
	d.Unchecked_png = ConvertMenuIcon(DecodeImageString(unchecked_png))
	return m
}

func (m *DesktopSysTray) WndProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	switch msg {
	case WM_SHELLNOTIFY:
		switch lParam {
		case WM_LBUTTONUP:
			for l := range m.Listeners {
				l.MouseLeftClick()
			}
		case WM_LBUTTONDBLCLK:
			for l := range m.Listeners {
				l.MouseLeftDoubleClick()
			}
		case WM_RBUTTONUP:
			m.showContextMenu()
		}
	case WM_COMMAND:
		i := int(wParam)
		mn := d.MenuItems[i]
		if mn.Menu.Action != nil && mn.Menu.Enabled {
			mn.Menu.Action(mn.Menu)
		}
	case WM_MEASUREITEM:
		ms := MEASUREITEMSTRUCTPtr(uintptr(lParam))

		i := int(ms.itemData)
		mn := d.MenuItems[i]

		hdc := HDCPtr(GetDC.Call(Arg(d.MainWnd.Wnd)))
		defer ReleaseDC.Call(Arg(d.MainWnd.Wnd), Arg(hdc))
		font := SystemMenuFontNew()
		defer font.Close()
		fontold := HFONTPtr(SelectObject.Call(Arg(hdc), Arg(font)))
		size := SIZE{}
		w := WStringNew(mn.Menu.Name)
		GetTextExtentPoint32.Call(Arg(hdc), Arg(w), Arg(w.Size()), Arg(&size))
		SelectObject.Call(Arg(hdc), Arg(fontold))
		size.cx += LONG(GetSystemMenuImageSize()+SPACE_ICONS) * 2
		ms.itemWidth = UINT(size.cx)
		ms.itemHeight = UINT(size.cy)
	case WM_DRAWITEM:
		di := (*DRAWITEMSTRUCT)(unsafe.Pointer(lParam))

		i := int(di.itemData)
		mn := d.MenuItems[i]

		if !mn.Menu.Enabled {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_GRAYTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENU)))))
		} else if (di.itemState & ODS_SELECTED) == ODS_SELECTED {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_HIGHLIGHTTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_HIGHLIGHT)))))
		} else {
			SetTextColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENUTEXT)))))
			SetBkColor.Call(Arg(di.hDC), Arg(COLORREFPtr(GetSysColor.Call(Arg(COLOR_MENU)))))
		}

		x := di.rcItem.left
		y := di.rcItem.top
		//w := di.rcItem.right - di.rcItem.left
		//h := di.rcItem.bottom - di.rcItem.top

		x += LONG(GetSystemMenuImageSize()+SPACE_ICONS) * 2

		font := SystemMenuFontNew()
		defer font.Close()
		SelectObject.Call(Arg(di.hDC), Arg(font))
		w := WStringNew(mn.Menu.Name)
		defer w.Close()
		ExtTextOut.Call(Arg(di.hDC), Arg(x), Arg(y), Arg(ETO_OPAQUE), Arg(&di.rcItem), Arg(w), Arg(w.Size()), NULL)

		x = di.rcItem.left

		if mn.Menu.Type == MenuCheckBox {
			if mn.Menu.State {
				d.Checked_png.Draw(x, y, di.hDC)
			} else {
				d.Unchecked_png.Draw(x, y, di.hDC)
			}
		}

		x += LONG(GetSystemMenuImageSize() + SPACE_ICONS)
		if mn.Image != nil {
			mn.Image.Draw(x, y, di.hDC)
		}
	case WM_QUIT:
		PostMessage.Call(Arg(d.MainWnd.Wnd), Arg(WM_QUIT), NULL, NULL)
	}

	if msg == d.WM_TASKBARCREATED {
		m.show()
	}

	return d.MainWnd.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (m *DesktopSysTray) WebWndProc(hWnd HWND, msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var d *DesktopSysTrayWin = m.os.(*DesktopSysTrayWin)

	switch msg {
	case WM_ACTIVATE:
		if wParam == WA_INACTIVE {
			if d.WebWnd != nil {
				if lParam == 0 {
					ShowWindow.Call(Arg(d.WebWnd.Wnd), SW_HIDE)
				}
			}
		}
	case WM_SIZE:
		width := LOWORD(DWORD(lParam))
		height := HIWORD(DWORD(lParam))
		var webbrowser *IWebBrowser2
		d.browser.QueryInterface(&IID_IWebBrowser2, &webbrowser).S_OK()
		defer webbrowser.Release()
		webbrowser.put_Width(LONG(width))
		webbrowser.put_Height(LONG(height))
	case WM_CREATE:
		HRESULTPtr(OleInitialize.Call(NULL)).S_OK()
		m.createWebBrowser(hWnd)
	case WM_DESTROY:
		d.browser.Close(OLECLOSE_NOSAVE)
		d.browser.Release()
		d.browser = nil
		d.WebWnd = nil
	}

	return d.WebWnd.DefWindowProc(hWnd, msg, wParam, lParam)
}

func (m *DesktopSysTray) setIcon(i image.Image) {
	d := m.os.(*DesktopSysTrayWin)

	bm := BitmapImageNew(i)
	defer bm.Close()

	if d.Icon != 0 {
		d.Icon.Close()
	}
	d.Icon = HICONNew(bm.hbm)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_MODIFY), Arg(n))).Bool() {
		// no icon
	}
}

func (m *DesktopSysTray) show() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	n.SetTooltip(m.Title)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_ADD), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
	m.createWebMenu()
}

func (m *DesktopSysTray) hide() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_DELETE), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
}

func (m *DesktopSysTray) update() {
	d := m.os.(*DesktopSysTrayWin)

	n := NOTIFYICONDATANew()
	n.hWnd = d.MainWnd.Wnd
	n.SetCallback(WM_SHELLNOTIFY)
	n.SetIcon(d.Icon)
	n.SetTooltip(m.Title)
	if !BOOLPtr(Shell_NotifyIcon.Call(Arg(NIM_MODIFY), Arg(n))).Bool() {
		panic(GetLastErrorString())
	}
	m.createWebMenu()
}

func (m *DesktopSysTray) createWebMenu() {
	d := m.os.(*DesktopSysTrayWin)
	if m.WebPopup != nil {
		if d.WebWnd == nil {
			w, h := m.WebPopup.Size()
			d.WebWnd = WindowNew(WNDPROCNew(m.WebWndProc), "WebView", WS_EX_TOOLWINDOW, WS_POPUPWINDOW, w, h)
		}
	}
}

func (m *DesktopSysTray) close() {
	d := m.os.(*DesktopSysTrayWin)

	if d.MainWnd != nil {
		m.hide()
	}

	if d.Icon != 0 {
		d.Icon.Close()
		d.Icon = 0
	}

	if d.MainMenu != 0 {
		d.MainMenu.Close()
		d.MainMenu = 0
	}

	if d.TaskbarCreated != 0 {
		d.TaskbarCreated.Close()
		d.TaskbarCreated = 0
	}

	if d.MainWnd != nil {
		d.MainWnd.Close()
		d.MainWnd = nil
	}
}

func (m *DesktopSysTray) showContextMenu() {
	d := m.os.(*DesktopSysTrayWin)

	if len(m.Menu) > 0 {
		m.updateMenus()

		if !BOOLPtr(SetForegroundWindow.Call(Arg(d.MainWnd.Wnd))).Bool() {
			panic(GetLastErrorString())
		}

		var pos POINT
		if !BOOLPtr(GetCursorPos.Call(Arg(&pos))).Bool() {
			panic(GetLastErrorString())
		}

		for !BOOLPtr(TrackPopupMenu.Call(Arg(d.MainMenu), TPM_RIGHTBUTTON, Arg(pos.x), Arg(pos.y), NULL, Arg(d.MainWnd.Wnd), NULL)).Bool() {
			var hWnd HWND
			// in case popup menu lost focus, did not die, and user right clied icon again
			// we have to find pop up menu, kill it and show context menu again

			// 0x000005a6 - "Popup menu already active."
			if LastError == 0x000005a6 {
				for {
					// "#32768" - pop up menu window class
					w := WStringNew("#32768")
					defer w.Close()
					hWnd = HWNDPtr(FindWindowEx.Call(NULL, Arg(hWnd), Arg(w), NULL))
					if hWnd == 0 {
						break
					}
					SendMessage.Call(Arg(hWnd), Arg(WM_KEYDOWN), Arg(VK_ESCAPE), NULL)
				}
				// noting is working...
				// just return.
				return
			} else {
				panic(GetLastErrorString())
			}
		}
	}

	if m.WebPopup != nil {
		if !BOOLPtr(SetForegroundWindow.Call(Arg(d.WebWnd.Wnd))).Bool() {
			panic(GetLastErrorString())
		}

		var pos POINT
		if !BOOLPtr(GetCursorPos.Call(Arg(&pos))).Bool() {
			panic(GetLastErrorString())
		}

		w, h := m.WebPopup.Size()
		if !BOOLPtr(SetWindowPos.Call(Arg(d.WebWnd.Wnd), Arg(HWND_TOPMOST), Arg(int(pos.x)-w), Arg(int(pos.y)-h), Arg(w), Arg(h), NULL)).Bool() {
			panic(GetLastErrorString())
		}

		ShowWindow.Call(Arg(d.WebWnd.Wnd), SW_SHOW)
		UpdateWindow.Call(Arg(d.WebWnd.Wnd))

		if m.WebPopup.Url != "" {
			if d.WebCurrent != m.WebPopup { // do not reload page twice
				d.WebCurrent = m.WebPopup

				var webbrowser *IWebBrowser2
				d.browser.QueryInterface(&IID_IWebBrowser2, &webbrowser).S_OK()
				defer webbrowser.Release()

				u := VARIANTNew(m.WebPopup.Url)
				defer u.Clear()
				webbrowser.Navigate2(&u, nil, nil, nil, nil).S_OK()
			}
		}

		if m.WebPopup.Html != "" {
			if d.WebCurrent != m.WebPopup { // do not reload page twice
				d.WebCurrent = m.WebPopup

				var webbrowser *IWebBrowser2
				d.browser.QueryInterface(&IID_IWebBrowser2, &webbrowser).S_OK()
				defer webbrowser.Release()

				var cont *IConnectionPointContainer
				d.browser.QueryInterface(&IID_IConnectionPointContainer, &cont).S_OK()

				var conn *IConnectionPoint
				cont.FindConnectionPoint(&DIID_DWebBrowserEvents2, &conn).S_OK()

				var cookieDone DWORD
				events := DWebBrowserEvents2New()
				events.DocumentComplete = func() HRESULT {
					defer cont.Release()
					defer func() {
						if cookieDone != 0 {
							conn.Unadvise(cookieDone).S_OK()
							conn.Release()
						}
						cookieDone = 0xffffffff
					}()

					var disp *IDispatch
					webbrowser.get_Document(&disp).S_OK()
					defer disp.Release()

					var doc *IHTMLDocument2
					disp.QueryInterface(&IID_IHTMLDocument2, &doc).S_OK()
					defer doc.Release()

					sfArray := SAFEARRAYString(m.WebPopup.Html)
					defer sfArray.Destory()
					doc.Write(sfArray)
					doc.Close()
					return S_OK
				}
				var cookie DWORD
				conn.Advise(&events.IUnknown, &cookie).S_OK()
				if cookieDone == 0xffffffff { // alredy triggered
					conn.Unadvise(cookie)
					conn.Release()
				} else { // not yet triggered (hope no concuricy)
					cookieDone = cookie
				}

				u := VARIANTNew("about:blank")
				defer u.Clear()
				webbrowser.Navigate2(&u, nil, nil, nil, nil).S_OK()
			}
		}
	}
}

func (m *DesktopSysTray) updateMenus() {
	d := m.os.(*DesktopSysTrayWin)

	if d.MainMenu != 0 {
		d.MainMenu.Close()
	}

	d.MainMenu = m.createSubMenu(m.Menu)
}

func (m *DesktopSysTray) createSubMenu(mm []Menu) HMENU {
	d := m.os.(*DesktopSysTrayWin)

	hmenu := HMENUPtr(CreatePopupMenu.Call())
	if hmenu == 0 {
		panic(GetLastErrorString())
	}

	for i := range mm {
		mn := &mm[i]

		switch mn.Type {
		case MenuItem, MenuCheckBox:
			menuwin := &MenuItemWin{}
			menuwin.Menu = mn

			if mn.Icon != nil {
				menuwin.Image = ConvertMenuIcon(mn.Icon)
			}

			id := len(d.MenuItems)
			d.MenuItems = append(d.MenuItems, menuwin)

			if mn.Menu != nil {
				sub := m.createSubMenu(mn.Menu)
				// seems like you dont have to free this menu, since it already attached
				// to main HMENU handler
				if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MF_POPUP|MFT_OWNERDRAW), Arg(sub), NULL)).Bool() {
					panic(GetLastErrorString())
				}
				mi := MENUITEMINFO{}
				mi.cbSize = UINT(unsafe.Sizeof(mi))
				if !BOOLPtr(GetMenuItemInfo.Call(Arg(hmenu), Arg(sub), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
				mi.dwItemData = ULONG_PTR(id)
				mi.fMask |= MIIM_DATA
				if !BOOLPtr(SetMenuItemInfo.Call(Arg(hmenu), Arg(sub), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
			} else {
				if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MFT_OWNERDRAW), Arg(id), NULL)).Bool() {
					panic(GetLastErrorString())
				}
				mi := MENUITEMINFO{}
				mi.cbSize = UINT(unsafe.Sizeof(mi))
				if !BOOLPtr(GetMenuItemInfo.Call(Arg(hmenu), Arg(id), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
				mi.dwItemData = ULONG_PTR(id)
				mi.fMask |= MIIM_DATA
				if !BOOLPtr(SetMenuItemInfo.Call(Arg(hmenu), Arg(id), Arg(false), Arg(&mi))).Bool() {
					panic(GetLastErrorString())
				}
			}
		case MenuSeparator:
			if !BOOLPtr(AppendMenu.Call(Arg(hmenu), Arg(MF_SEPARATOR), NULL, NULL)).Bool() {
				panic(GetLastErrorString())
			}
		}
	}

	return hmenu
}

func (m *DesktopSysTray) createWebBrowser(hwnd HWND) {
	d := m.os.(*DesktopSysTrayWin)

	var webbrowser *IWebBrowser2
	HRESULTPtr(CoCreateInstance.Call(Arg(&CLSID_WebBrowser), NULL, CLSCTX_INPROC, Arg(&IID_IWebBrowser2), Arg(&webbrowser))).S_OK()
	defer webbrowser.Release()

	webbrowser.QueryInterface(&IID_IOleObject, &d.browser).S_OK()

	o := IOleClientSiteNew(hwnd, d.browser, m)
	d.browser.SetClientSite(o).S_OK()

	rect := RECT{}
	GetClientRect.Call(Arg(hwnd), Arg(&rect))
	d.browser.DoVerb(OLEIVERB_INPLACEACTIVATE, nil, o, 0, hwnd, &rect).S_OK()
	webbrowser.put_Left(0).S_OK()
	webbrowser.put_Top(0).S_OK()
	webbrowser.put_Width(rect.right - rect.left).S_OK()
	webbrowser.put_Height(rect.bottom - rect.top).S_OK()
}
