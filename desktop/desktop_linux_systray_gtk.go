// +build linux

package desktop

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"unsafe"
)

var SpaceIcon image.Image = image.NewRGBA(image.Rect(0, 0, 1, 1))

type DesktopSysTrayGtk struct {
	M             *DesktopSysTray
	Icon          image.Image
	GtkStatusIcon GtkWidget
	GtkMenu       GtkWidget

	// WebPopup
	WebWnd     GtkWidget
	WebView    GtkWidget
	WebCurrent *WebPopup
	HidePop    *GSourceFunc
	Pol        *GSourceFunc

	IconActivate *GSourceFunc
	IconPopup    *GSourceFunc

	ShowInvokeVar    *GSourceFunc
	HideInvokeVar    *GSourceFunc
	UpdateInvokeVar  *GSourceFunc
	SetIconInvokeVar *GSourceFunc

	GSourceFuncs []*GSourceFunc
}

func DesktopSysTrayGtkNew(m *DesktopSysTray) *DesktopSysTrayGtk {
	GtkMessageLoopInc()

	os := &DesktopSysTrayGtk{}

	os.M = m

	os.ShowInvokeVar = GSourceFuncNew(func() {
		os.UpdateMenus()
		if os.GtkStatusIcon == nil {
			os.GtkStatusIcon = os.CreateGStatusIcon()
		}
		gtk_status_icon_set_visible(os.GtkStatusIcon, true)
	})

	os.HideInvokeVar = GSourceFuncNew(func() {
		gtk_status_icon_set_visible(os.GtkStatusIcon, false)
	})

	os.UpdateInvokeVar = GSourceFuncNew(func() {
		os.UpdateMenus()

		if os.GtkStatusIcon != nil {
			gtk_status_icon_set_from_gicon(os.GtkStatusIcon, ConvertMenuImage(os.Icon))
			gtk_status_icon_set_tooltip_text(os.GtkStatusIcon, m.Title)
		}
	})

	os.SetIconInvokeVar = GSourceFuncNew(func() {
		if os.GtkStatusIcon != nil {
			gtk_status_icon_set_from_gicon(os.GtkStatusIcon, ConvertMenuImage(os.Icon))
			gtk_status_icon_set_tooltip_text(os.GtkStatusIcon, m.Title)
		}
	})

	os.IconActivate = GSourceFuncNew(func() {
		for l := range m.Listeners {
			l.MouseLeftClick()
		}
	})

	os.IconPopup = GSourceFuncNew(func() {
		if m.Menu != nil {
			gtk_menu_popup(os.GtkMenu, nil, nil, gtk_status_icon_position_menu, nil, 1, gtk_get_current_event_time())
		}
		if m.WebPopup != nil {
			os.ShowWebPopup()
		}
	})

	return os
}

func ConvertMenuImage(icon image.Image) GIcon {
	var menubarHeigh uint = 64

	c := resize.Resize(menubarHeigh, menubarHeigh, icon, resize.Lanczos3)

	var b bytes.Buffer
	err := png.Encode(&b, c)
	if err != nil {
		panic(err)
	}

	buf := b.Bytes()
	gb := g_bytes_new(buf, len(buf))
	gi := g_bytes_icon_new(gb)
	return gi
}

func (os *DesktopSysTrayGtk) CreateMenuItem(item *Menu) GtkWidget {
	img := item.Icon

	if img == nil {
		img = SpaceIcon
	}

	spacing := 6

	box := gtk_hbox_new(false, spacing)
	wicon := gtk_image_new_from_gicon(ConvertMenuImage(img), GTK_ICON_SIZE_MENU)
	gtk_box_pack_start(box, wicon, false, false, spacing)
	label := gtk_label_new(item.Name)

	var menu GtkWidget

	switch item.Type {
	case MenuCheckBox:
		menu = gtk_check_menu_item_new()
		gtk_check_menu_item_set_active(menu, item.State)
	case MenuItem:
		menu = gtk_menu_item_new()
	}

	if !item.Enabled {
		gtk_widget_set_sensitive(menu, false)
	}

	gtk_box_pack_start(box, label, false, false, spacing)
	gtk_container_add(menu, box)
	gtk_widget_show_all(menu)

	if item.Menu == nil {
		var fn *GSourceFunc = GSourceFuncNew(func() {
			item.Action(item)
		})
		os.GSourceFuncs = append(os.GSourceFuncs, fn)
		g_signal_connect_activate(menu, fn)
	}

	return menu
}

func (os *DesktopSysTrayGtk) CreateSubMenu(mm []Menu) GtkWidget {
	gmenu := gtk_menu_new()

	for i := range mm {
		mn := &mm[i]

		switch mn.Type {
		case MenuItem, MenuCheckBox:
			if mn.Menu != nil {
				sub := os.CreateSubMenu(mn.Menu)
				item := os.CreateMenuItem(mn)
				gtk_menu_item_set_submenu(item, sub)
				gtk_menu_shell_append(gmenu, item)
			} else {
				item := os.CreateMenuItem(mn)
				gtk_menu_shell_append(gmenu, item)
			}
		case MenuSeparator:
			item := gtk_separator_menu_item_new()
			gtk_menu_shell_append(gmenu, item)
		}
	}

	return gmenu
}

func (os *DesktopSysTrayGtk) CreateWebMenu() GtkWidget {
	gmenu := gtk_menu_new()

	var menu GtkWidget
	menu = gtk_menu_item_new()

	label := gtk_label_new("WebPopup")

	gtk_container_add(menu, label)
	gtk_widget_show_all(menu)

	var fn *GSourceFunc = GSourceFuncNew(func() {
		os.ShowWebPopup()
	})
	os.GSourceFuncs = append(os.GSourceFuncs, fn)
	g_signal_connect_activate(menu, fn)

	gtk_menu_shell_append(gmenu, menu)

	return gmenu
}

func (os *DesktopSysTrayGtk) ShowWebPopup() {
	m := os.M
	if os.WebCurrent != m.WebPopup {
		os.WebCurrent = m.WebPopup
		if m.WebPopup.Url != "" {
			webkit_web_view_load_uri(os.WebView, m.WebPopup.Url)
		}
		if m.WebPopup.Html != "" {
			webkit_web_view_load_html_string(os.WebView, m.WebPopup.Html, "")
		}
	}
	gtk_widget_grab_focus(os.WebWnd)
	gtk_widget_show_all(os.WebWnd)
}

func (os *DesktopSysTrayGtk) UpdateMenus() {
	m := os.M

	if os.GtkMenu != nil {
		gtk_widget_destroy(os.GtkMenu)
	}

	if os.GSourceFuncs != nil {
		for _, v := range os.GSourceFuncs {
			v.Close()
		}
	}
	os.GSourceFuncs = nil

	if m.Menu != nil {
		os.GtkMenu = os.CreateSubMenu(m.Menu)
	}
	if m.WebPopup != nil {
		if os.WebWnd == nil {
			os.WebWnd = gtk_window_new(GTK_WINDOW_TOPLEVEL)
			gtk_window_set_decorated(os.WebWnd, false)
			gtk_window_set_position(os.WebWnd, GTK_WIN_POS_MOUSE)
			w, h := m.WebPopup.Size()
			gtk_widget_set_size_request(os.WebWnd, w, h)
			gtk_window_set_resizable(os.WebWnd, false)
			gtk_window_set_skip_taskbar_hint(os.WebWnd, true)
			gtk_window_set_skip_pager_hint(os.WebWnd, true)
			gtk_widget_set_events(os.WebWnd, GDK_FOCUS_CHANGE_MASK)

			os.HidePop = GSourceFuncNew(func() {
				gtk_widget_hide(os.WebWnd)
			})
			g_signal_connect_focus_out(os.WebWnd, os.HidePop)

			os.WebView = webkit_web_view_new()

			os.Pol = GSourceFuncNew(func(p0, p1, p2, p3, p4, p5 unsafe.Pointer) bool {
				s := webkit_network_request_get_uri(p2)
				if os.WebCurrent.Url == s { // ignore handling currently loaded url
					return false
				}
				if m.WebPopup.Handler != nil {
					if m.WebPopup.Handler(s) {
						return true
					}
				}
				BrowserOpenURI(s)
				return true
			})
			g_signal_connect_navigation_policy_decision_requested(os.WebView, os.Pol)

			scroll := gtk_scrolled_window_new()
			gtk_container_add(scroll, os.WebView)
			gtk_container_add(os.WebWnd, scroll)
		}
		os.GtkMenu = os.CreateWebMenu()
	}
}

func (os *DesktopSysTrayGtk) CreateGStatusIcon() GtkWidget {
	m := os.M

	gicon := gtk_status_icon_new_from_gicon(ConvertMenuImage(os.Icon))

	g_signal_connect_activate(gicon, os.IconActivate)
	g_signal_connect_popup(gicon, os.IconPopup)

	gtk_status_icon_set_tooltip_text(gicon, m.Title)
	return gicon
}

func (os *DesktopSysTrayGtk) show() {
	GtkMessageLoopInvoke(os.ShowInvokeVar)
}

func (os *DesktopSysTrayGtk) hide() {
	GtkMessageLoopInvoke(os.HideInvokeVar)
}

func (os *DesktopSysTrayGtk) update() {
	GtkMessageLoopInvoke(os.UpdateInvokeVar)
}

func (os *DesktopSysTrayGtk) close() {
	if os.GtkMenu != nil {
		gtk_widget_destroy(os.GtkMenu)
		os.GtkMenu = nil
	}

	if os.GtkStatusIcon != nil {
		gtk_widget_destroy(os.GtkStatusIcon)
		os.GtkStatusIcon = nil
	}

	os.IconActivate.Close()
	os.IconPopup.Close()

	os.ShowInvokeVar.Close()
	os.HideInvokeVar.Close()
	os.UpdateInvokeVar.Close()
	os.SetIconInvokeVar.Close()
	os.HidePop.Close()
	os.Pol.Close()

	if os.GSourceFuncs != nil {
		for _, v := range os.GSourceFuncs {
			v.Close()
		}
	}
	os.GSourceFuncs = nil

	GtkMessageLoopDec()
}

func (os *DesktopSysTrayGtk) setIcon(i image.Image) {
	os.Icon = i

	GtkMessageLoopInvoke(os.SetIconInvokeVar)
}
