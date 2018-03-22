// +build darwin

package desktop

import (
	"unsafe"
)

var NSControlClass unsafe.Pointer = Runtime_objc_lookUpClass("NSControl")
var NSControlInitWithFrame unsafe.Pointer = Runtime_sel_getUid("initWithFrame")

type NSControl struct {
	NSView
}

func NSControlNew() NSControl {
	var m = NSControlPointer(Runtime_class_createInstance(NSControlClass, 0))
	p := Runtime_objc_msgSend(m.Pointer, NSControlInitWithFrame)
	m.Assert(p)
	return m
}

func NSControlPointer(p unsafe.Pointer) NSControl {
	return NSControl{NSViewPointer(p)}
}
