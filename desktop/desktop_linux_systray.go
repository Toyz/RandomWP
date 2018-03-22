// +build linux

package desktop

/*
SysTray Linux have a lot of way to go. You can use AppIndicator, Gtk, XEvent or something else.
Not all distributions can have all/any installed so we have to check it dynamically on build or on run.

1) using cgo.

we can write code by dynamic linking. cgo #include "appindicator.h" but different distributions has different versions.
it can be appindicator3.so or appindicator.so or does not esitst at all.

we need project which can on fly detect propper libraries installed and build against it or chooise right library
at runtime.

can we link againt undefended symbols and then call dlopen for the right library?

2) dynamic loading libraries.

go does not have dynamic library loading for linux.

https://github.com/rainycape/dl

build only against 64 systems.
*/

import (
	"image"
)

var APPINDICATOR = true
var GTK = true

func desktopSysTrayNew() *DesktopSysTray {
	m := &DesktopSysTray{}

	for _, s := range []string{"libwebkitgtk-3.0.so", "libwebkitgtk-1.0.so", "libwebkitgtk-3.0.so.0"} {
		_, err := dlopen(s, RTLD_LAZY|RTLD_GLOBAL)
		if err == nil {
			break
		}
	}

	if APPINDICATOR {
		for _, s := range []string{"libappindicator3.so.1", "libappindicator3.so", "libappindicator.so"} {
			_, err := dlopen(s, RTLD_LAZY|RTLD_GLOBAL)
			if err == nil {
				m.os = DesktopSysTrayAppIndicatorNew(m)
				return m
			}
		}
	}

	if GTK {
		for _, s := range []string{"libgtk-3.so", "libgtk-3.so.0", "libgtk-x11-2.0.so", "libgtk-x11-2.0.so.0"} {
			_, err := dlopen(s, RTLD_LAZY|RTLD_GLOBAL)
			if err == nil {
				m.os = DesktopSysTrayGtkNew(m)
				return m
			}
		}
	}

	panic("unable find systray interface")
}

func (m *DesktopSysTray) show() {
	if os, ok := m.os.(*DesktopSysTrayAppIndicator); ok {
		os.show()
		return
	}
	if os, ok := m.os.(*DesktopSysTrayGtk); ok {
		os.show()
		return
	}
	panic("broken systray interface")
}

func (m *DesktopSysTray) hide() {
	if os, ok := m.os.(*DesktopSysTrayAppIndicator); ok {
		os.hide()
		return
	}
	if os, ok := m.os.(*DesktopSysTrayGtk); ok {
		os.hide()
		return
	}
	panic("broken systray interface")
}

func (m *DesktopSysTray) update() {
	if os, ok := m.os.(*DesktopSysTrayAppIndicator); ok {
		os.update()
		return
	}
	if os, ok := m.os.(*DesktopSysTrayGtk); ok {
		os.update()
		return
	}
	panic("broken systray interface")
}

func (m *DesktopSysTray) close() {
	if os, ok := m.os.(*DesktopSysTrayAppIndicator); ok {
		os.close()
		return
	}
	if os, ok := m.os.(*DesktopSysTrayGtk); ok {
		os.close()
		return
	}
	panic("broken systray interface")
}

func (m *DesktopSysTray) setIcon(i image.Image) {
	if os, ok := m.os.(*DesktopSysTrayAppIndicator); ok {
		os.setIcon(i)
		return
	}
	if os, ok := m.os.(*DesktopSysTrayGtk); ok {
		os.setIcon(i)
		return
	}
	panic("broken systray interface")
}
