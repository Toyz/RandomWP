// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/reference/appkit/nsviewcontroller?language=objc

var NSViewControllerClass unsafe.Pointer = Runtime_objc_lookUpClass("NSViewController")
var NSViewControllerSetView unsafe.Pointer = Runtime_sel_getUid("setView:")
var NSViewControllerInit unsafe.Pointer = Runtime_sel_getUid("init")
var NSViewControllerInitWithNibName unsafe.Pointer = Runtime_sel_getUid("initWithNibName:bundle:")
var NSViewControllerIsViewLoaded unsafe.Pointer = Runtime_sel_getUid("isViewLoaded")

type NSViewController struct {
	NSResponder
}

func NSViewControllerNew() NSViewController {
	var m = NSViewControllerPointer(Runtime_class_createInstance(NSViewControllerClass, 0))
	m.Init()
	return m
}

func (m NSViewController) Init() {
	p := Runtime_objc_msgSend(m.Pointer, NSViewControllerInit)
	m.Assert(p)
}

func NSViewControllerPointer(p unsafe.Pointer) NSViewController {
	return NSViewController{NSResponderPointer(p)}
}

func (m NSViewController) SetView(i NSView) {
	Runtime_objc_msgSend(m.Pointer, NSViewControllerSetView, i.Pointer)
}

func (m NSViewController) GetViewLoaded() bool {
	return Pointer2Bool(Runtime_objc_msgSend(m.Pointer, NSViewControllerIsViewLoaded))
}
