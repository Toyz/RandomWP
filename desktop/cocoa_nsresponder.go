// +build darwin

package desktop

import (
	"unsafe"
)

var NSResponderClass unsafe.Pointer = Runtime_objc_lookUpClass("NSResponder")
var NSResponderInit unsafe.Pointer = Runtime_sel_getUid("init")

type NSResponder struct {
	NSObject
}

func NSResponderNew() NSResponder {
	var m = NSResponderPointer(Runtime_class_createInstance(NSResponderClass, 0))
	p := Runtime_objc_msgSend(m.Pointer, NSResponderInit)
	m.Assert(p)
	return m
}

func NSResponderPointer(p unsafe.Pointer) NSResponder {
	return NSResponder{NSObjectPointer(p)}
}
