// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>

extern BOOL CanBecomeKeyWindow(void*, void*);

extern void DesktopOSXSysTrayViewActionMap(void*, void*, void*);
*/
import "C"

import (
	"unsafe"
)

// NSWindow key hack
var NSWindowPrev unsafe.Pointer

var NSWindowSelected bool

//export CanBecomeKeyWindow
func CanBecomeKeyWindow(id unsafe.Pointer, sel unsafe.Pointer) C.BOOL {
	if NSWindowSelected {
		return 1
	} else {
		return 0 // (*(*func(id unsafe.Pointer, sel unsafe.Pointer) C.BOOL)(NSWindowPrev))(id, sel)
	}
}

func NSWindowRegister() bool {
	NSWindowPrev = Runtime_class_replaceMethod(NSWindowClass, NSWindowCanBecomeKeyWindow, C.CanBecomeKeyWindow, "B@:")
	return true
}

var NSWindowRegistred = NSWindowRegister()

//
// register
//

var DesktopOSXSysTrayViewClassReg unsafe.Pointer = Runtime_objc_allocateClassPair(NSImageViewClass, "DesktopOSXSysTrayViewClass", 0)
var DesktopOSXSysTrayViewMouseUpReg unsafe.Pointer = Runtime_sel_registerName("mouseUp:")
var DesktopOSXSysTrayViewMouseDownReg unsafe.Pointer = Runtime_sel_registerName("mouseDown:")

var DesktopOSXSysTrayViewMap = make(map[unsafe.Pointer]*DesktopOSXSysTrayView)

//export DesktopOSXSysTrayViewActionMap
func DesktopOSXSysTrayViewActionMap(id unsafe.Pointer, sel unsafe.Pointer, sender unsafe.Pointer) {
	if sel == DesktopOSXSysTrayViewMouseUpReg {
		DesktopOSXSysTrayViewMap[id].MouseUp(sender)
	}
	if sel == DesktopOSXSysTrayViewMouseDownReg {
		DesktopOSXSysTrayViewMap[id].MouseDown(sender)
	}
}

func DesktopOSXSysTrayViewRegister() bool {
	if !Runtime_class_addMethod(DesktopOSXSysTrayViewClassReg, DesktopOSXSysTrayViewMouseUpReg, C.DesktopOSXSysTrayViewActionMap, "v@::") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayViewClassReg, DesktopOSXSysTrayViewMouseDownReg, C.DesktopOSXSysTrayViewActionMap, "v@::") {
		panic("problem initalizing class")
	}
	Runtime_objc_registerClassPair(DesktopOSXSysTrayViewClassReg)

	return true
}

var DesktopOSXSysTrayViewRegistred bool = DesktopOSXSysTrayViewRegister()

//
// object
//

var DesktopOSXSysTrayViewClass unsafe.Pointer = Runtime_objc_lookUpClass("DesktopOSXSysTrayViewClass")
var DesktopOSXSysTrayViewMouseUp unsafe.Pointer = Runtime_sel_getUid("mouseUp:")
var DesktopOSXSysTrayViewMouseDown unsafe.Pointer = Runtime_sel_getUid("mouseDown:")

type DesktopOSXSysTrayView struct {
	NSImageView

	Selected bool

	s *DesktopOSXSysTray
}

func DesktopOSXSysTrayViewNew(s *DesktopOSXSysTray, i NSImage) *DesktopOSXSysTrayView {
	size := i.GetSize()

	m := DesktopOSXSysTrayViewPointer(Runtime_class_createInstance(DesktopOSXSysTrayViewClass, 0))

	rect := NSRectSize(size)
	m.NSImageView.Init(rect)

	m.SetWantsLayer(true)

	m.s = s

	return m
}

func DesktopOSXSysTrayViewPointer(p unsafe.Pointer) *DesktopOSXSysTrayView {
	m := DesktopOSXSysTrayView{NSImageViewPointer(p), false, nil}

	DesktopOSXSysTrayViewMap[m.Pointer] = &m

	return &m
}

func (m *DesktopOSXSysTrayView) MouseDown(sender unsafe.Pointer) {
	if m.s.Popover.Pointer != nil {
		NSWindowSelected = true
		m.GetWindow().BecomeKeyWindow()

		m.s.Popover.ShowRelativeToRect(m.GetBounds(), m.NSView, NSMaxYEdge)
	}
}

func (m *DesktopOSXSysTrayView) MouseUp(sender unsafe.Pointer) {
}

func (m *DesktopOSXSysTrayView) Update() {
	l := m.GetLayer()
	var c CGColor
	if m.Selected {
		c = CGColorRGB(0, 0, 1, 1)
	} else {
		NSWindowSelected = false
		c = CGColorRGB(0, 0, 0, 0)
	}
	l.SetBackgroundColor(c)
}

func (m *DesktopOSXSysTrayView) Release() {
	delete(DesktopOSXSysTrayViewMap, m.Pointer)
	m.NSObject.Release()
}
