// +build darwin

package desktop

import (
	"image"
	"runtime"

	"github.com/nfnt/resize"
)

func FontHeight(f NSFont) uint {
	defer f.Release()

	fd := f.FontDescriptor()
	defer fd.Release()

	n := NSNumberPointer(fd.ObjectForKey(NSFontSizeAttribute))
	defer n.Release()

	return uint(n.IntValue())
}

func convertTrayIcon(i image.Image) NSImage {
	menubarHeight := uint(22) // FontHeight(NSFontMenuBarFontOfSize(0))

	c := resize.Resize(menubarHeight, menubarHeight, i, resize.Lanczos3)

	return NSImageImage(c)
}

func convertMenuIcon(i image.Image) NSImage {
	menubarHeight := FontHeight(NSFontMenuFontOfSize(0))

	c := resize.Resize(menubarHeight, menubarHeight, i, resize.Lanczos3)

	return NSImageImage(c)
}

func createSubMenu(mm []Menu) NSMenu {
	var mn NSMenu = NSMenuNew()

	for i := range mm {
		m := &mm[i]

		switch m.Type {
		case MenuItem:
			var icon NSImage
			var menu NSMenuItem = NSMenuItemNew()
			if m.Icon != nil {
				icon = convertMenuIcon(m.Icon)
				defer icon.Release()
			}
			menu.SetTitle(m.Name)
			menu.SetImage(icon)
			menu.SetEnabled(m.Enabled)

			if m.Action != nil {
				a := DesktopOSXSysTrayActionNew(m)
				menu.SetTarget(a.Pointer)
				menu.SetAction(DesktopOSXSysTrayActionAction)
			}

			if m.Menu != nil {
				sub := createSubMenu(m.Menu)
				defer sub.Release()
				menu.SetSubmenu(sub)
			}
			mn.AddItem(menu)
		case MenuSeparator:
			menu := NSMenuItemSeparatorItem()
			defer menu.Release()
			mn.AddItem(menu)
		case MenuCheckBox:
			var icon NSImage
			var menu NSMenuItem = NSMenuItemNew()
			if m.Icon != nil {
				icon = convertMenuIcon(m.Icon)
				defer icon.Release()
			}
			menu.SetTitle(m.Name)
			menu.SetImage(icon)
			menu.SetEnabled(m.Enabled)
			menu.SetState((map[bool]int{true: NSOnState, false: NSOffState})[m.State])

			if m.Action != nil {
				a := DesktopOSXSysTrayActionNew(m)
				menu.SetTarget(a.Pointer)
				menu.SetAction(DesktopOSXSysTrayActionAction)
			}

			mn.AddItem(menu)
		}
	}

	mn.SetAutoenablesItems(false)
	return mn
}

type DesktopOSXSysTray struct {
	statusbar  NSStatusBar
	statusitem NSStatusItem
	image      NSImage

	// html menu
	Popover    NSPopover
	View       *DesktopOSXSysTrayView
	WebView    WebView
	Controller *DesktopOSXSysTrayController
	Policy     *DesktopOSXSysTrayPolicy
	WebCurrent *WebPopup // current url loaded
}

func desktopSysTrayNew() *DesktopSysTray {
	// locket for thread message pump, will be unlocked after
	// desktopMainClose
	//
	// [NSUndoManager(NSInternal) _endTopLevelGroupings] is only safe to invoke on the main thread.
	runtime.LockOSThread()

	return &DesktopSysTray{os: &DesktopOSXSysTray{statusbar: NSStatusBarSystemStatusBar()}}
}

func (m *DesktopSysTray) update() {
	d := m.os.(*DesktopOSXSysTray)

	if d.statusitem.Pointer == nil {
		d.statusitem = d.statusbar.StatusItemWithLength(NSVariableStatusItemLength)
	}

	if m.Menu != nil {
		mn := createSubMenu(m.Menu)
		defer mn.Release()
		d.statusitem.SetMenu(mn)
	}

	if m.WebPopup != nil {
		m.createWebView()
		if m.WebPopup.Html != "" {
			d.View.SetImage(d.image)
		}
		if m.WebPopup.Url != "" {
			d.View.SetImage(d.image)
		}
	} else { // we use statisitem view or custom view
		d.statusitem.SetToolTip(m.Title)
		d.statusitem.SetHighlightMode(true)
		d.statusitem.SetImage(d.image)
	}
}

func (m *DesktopSysTray) createWebView() {
	d := m.os.(*DesktopOSXSysTray)
	if d.View == nil {
		d.View = DesktopOSXSysTrayViewNew(d, d.image)
		d.statusitem.SetView(d.View.NSView)
	}
	if d.Policy == nil {
		d.Policy = DesktopOSXSysTrayPolicyNew(m)
	}
	if d.WebView.Pointer == nil {
		width, height := m.WebPopup.Size()
		d.WebView = WebViewNew(NSRectSize(NSSizeNew(float64(width), float64(height))), "", "")
		d.WebView.SetPolicyDelegate(d.Policy.NSObject)
	}
	if d.Controller == nil {
		d.Controller = DesktopOSXSysTrayControllerNew(m)
		d.Controller.SetView(d.WebView.NSView)
	}
	if d.Popover.Pointer == nil {
		d.Popover = NSPopoverNew()
		d.Popover.SetContentViewController(d.Controller.NSViewController)
		d.Popover.SetBehavior(NSPopoverBehaviorTransient)
	}
}

func (m *DesktopSysTray) setIcon(icon image.Image) {
	d := m.os.(*DesktopOSXSysTray)

	if d.image.Pointer != nil {
		d.image.Release()
		d.image.Pointer = nil
	}

	d.image = convertTrayIcon(icon)

	if d.statusitem.Pointer != nil {
		d.statusitem.SetImage(d.image)
	}

	if d.View != nil {
		d.View.SetImage(d.image)
	}
}

func (m *DesktopSysTray) show() {
	m.update()
}

func (m *DesktopSysTray) hide() {
	m.close()
}

func (m *DesktopSysTray) close() {
	d := m.os.(*DesktopOSXSysTray)

	if d.statusitem.Pointer != nil {
		d.statusbar.RemoveStatusItem(d.statusitem)
		d.statusitem.Release()
		d.statusitem.Pointer = nil
	}

	if d.statusbar.Pointer != nil {
		d.statusbar.Release()
		d.statusbar.Pointer = nil
	}

	if d.image.Pointer != nil {
		d.image.Release()
		d.image.Pointer = nil
	}

	desktopMainClose()
}
