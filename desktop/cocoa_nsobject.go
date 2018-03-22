// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSObject_Class/index.html#//apple_ref/occ/cl/NSObject

var NSObjectClass unsafe.Pointer = Runtime_objc_lookUpClass("NSObject")
var NSObjectAlloc unsafe.Pointer = Runtime_sel_getUid("alloc")
var NSObjectRetain unsafe.Pointer = Runtime_sel_getUid("retain")
var NSObjectRelease unsafe.Pointer = Runtime_sel_getUid("release")

type NSObject struct {
	Pointer unsafe.Pointer
}

func NSObjectNew() NSObject {
	return NSObjectPointer(Runtime_objc_msgSend(NSObjectClass, NSObjectAlloc))
}

func NSObjectPointer(p unsafe.Pointer) NSObject {
	if p == nil {
		panic("NSObject nil")
	}

	var m NSObject = NSObject{p}
	m.Retain()
	return m
}

// for selectors with '- (instancetype)someSelector:'
func (m NSObject) Assert(p unsafe.Pointer) {
	if p == nil {
		panic("failed to create object")
	}
	if m.Pointer != p {
		panic("new object has different refference")
	}
}

func (m NSObject) Retain() {
	Runtime_objc_msgSend(m.Pointer, NSObjectRetain)
}

func (m NSObject) Release() {
	Runtime_objc_msgSend(m.Pointer, NSObjectRelease)
}
