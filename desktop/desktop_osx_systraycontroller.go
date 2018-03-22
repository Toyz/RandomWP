// +build darwin

package desktop

/*
extern void DesktopOSXSysTrayControllerActionMap(void*, void*);
*/
import "C"

import (
	"unsafe"
)

//
// register
//

var DesktopOSXSysTrayControllerClassReg unsafe.Pointer = Runtime_objc_allocateClassPair(NSViewControllerClass, "DesktopOSXSysTrayControllerClass", 0)
var DesktopOSXSysTrayControllerViewWillAppearReg unsafe.Pointer = Runtime_sel_registerName("viewWillAppear")
var DesktopOSXSysTrayControllerViewDidAppearReg unsafe.Pointer = Runtime_sel_registerName("viewDidAppear")
var DesktopOSXSysTrayControllerViewDidDisappearReg unsafe.Pointer = Runtime_sel_registerName("viewDidDisappear")

var DesktopOSXSysTrayControllerMap = make(map[unsafe.Pointer]*DesktopOSXSysTrayController)

//export DesktopOSXSysTrayControllerActionMap
func DesktopOSXSysTrayControllerActionMap(id unsafe.Pointer, sel unsafe.Pointer) {
	if sel == DesktopOSXSysTrayControllerViewWillAppearReg {
		DesktopOSXSysTrayControllerMap[id].ViewWillAppear()
	}
	if sel == DesktopOSXSysTrayControllerViewDidAppearReg {
		DesktopOSXSysTrayControllerMap[id].ViewDidAppear()
	}
	if sel == DesktopOSXSysTrayControllerViewDidDisappearReg {
		DesktopOSXSysTrayControllerMap[id].ViewDidDisappear()
	}
}

func DesktopOSXSysTrayControllerRegister() bool {
	if !Runtime_class_addMethod(DesktopOSXSysTrayControllerClassReg, DesktopOSXSysTrayControllerViewWillAppearReg, C.DesktopOSXSysTrayControllerActionMap, "v@:") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayControllerClassReg, DesktopOSXSysTrayControllerViewDidAppearReg, C.DesktopOSXSysTrayControllerActionMap, "v@:") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayControllerClassReg, DesktopOSXSysTrayControllerViewDidDisappearReg, C.DesktopOSXSysTrayControllerActionMap, "v@:") {
		panic("problem initalizing class")
	}
	Runtime_objc_registerClassPair(DesktopOSXSysTrayControllerClassReg)
	return true
}

var DesktopOSXSysTrayControllerRegistred bool = DesktopOSXSysTrayControllerRegister()

//
// object
//

var DesktopOSXSysTrayControllerClass unsafe.Pointer = Runtime_objc_lookUpClass("DesktopOSXSysTrayControllerClass")
var DesktopOSXSysTrayControllerViewWillAppear unsafe.Pointer = Runtime_sel_getUid("viewWillAppear")
var DesktopOSXSysTrayControllerViewDidAppear unsafe.Pointer = Runtime_sel_getUid("viewDidAppear")
var DesktopOSXSysTrayControllerViewDidDisappear unsafe.Pointer = Runtime_sel_getUid("viewDidDisappear")

type DesktopOSXSysTrayController struct {
	NSViewController

	s *DesktopSysTray
	e *NSEvent
}

func DesktopOSXSysTrayControllerNew(s *DesktopSysTray) *DesktopOSXSysTrayController {
	m := DesktopOSXSysTrayControllerPointer(Runtime_class_createInstance(DesktopOSXSysTrayControllerClass, 0))

	m.NSViewController.Init()

	m.s = s

	return m
}

func DesktopOSXSysTrayControllerPointer(p unsafe.Pointer) *DesktopOSXSysTrayController {
	m := DesktopOSXSysTrayController{NSViewControllerPointer(p), nil, nil}

	DesktopOSXSysTrayControllerMap[m.Pointer] = &m

	return &m
}

func (m *DesktopOSXSysTrayController) ViewWillAppear() {
}

func (m *DesktopOSXSysTrayController) ViewDidAppear() {
	d := m.s.os.(*DesktopOSXSysTray)

	d.View.Selected = true
	d.View.Update()
	e := NSEventAddGlobalMonitor(NSLeftMouseDown|NSRightMouseDown, m.Handler)
	m.e = &e

	if d.WebCurrent != m.s.WebPopup {
		d.WebCurrent = m.s.WebPopup
		if m.s.WebPopup.Url != "" {
			d.WebView.GetMainFrame().LoadRequest(m.s.WebPopup.Url)
		}
		if m.s.WebPopup.Html != "" {
			d.WebView.GetMainFrame().LoadHTMLString("", m.s.WebPopup.Html)
		}
	}
}

func (m *DesktopOSXSysTrayController) Handler(e NSEvent) {
	d := m.s.os.(*DesktopOSXSysTray)
	if m.e != nil {
		d.Popover.PerformClose()
		NSEventRemoveMonitor(*m.e)
		m.e = nil
	}
}

func (m *DesktopOSXSysTrayController) ViewDidDisappear() {
	d := m.s.os.(*DesktopOSXSysTray)
	d.View.Selected = false
	d.View.Update()
}

func (m *DesktopOSXSysTrayController) Release() {
	delete(DesktopOSXSysTrayControllerMap, m.Pointer)
	m.NSObject.Release()
}
