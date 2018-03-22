// +build darwin

package desktop

import (
	"unsafe"
)

// http://developer.apple.com/library/mac/#documentation/Cocoa/Reference/Foundation/Classes/NSData_Class/Reference/Reference.html#//apple_ref/doc/c_ref/NSData

var NSArrayClass unsafe.Pointer = Runtime_objc_lookUpClass("NSArray")
var NSArrayCount unsafe.Pointer = Runtime_sel_getUid("count")
var NSArrayObjectAtIndex unsafe.Pointer = Runtime_sel_getUid("objectAtIndex:")

type NSArray struct {
	NSObject
}

func NSArrayPointer(p unsafe.Pointer) NSArray {
	var m NSArray = NSArray{NSObjectPointer(p)}
	return m
}

func (m NSArray) Count() int {
	return Pointer2Int(Runtime_objc_msgSend(m.Pointer, NSArrayCount))
}

func (m NSArray) ObjectAtIndex(i int) unsafe.Pointer {
	return Runtime_objc_msgSend(m.Pointer, NSArrayObjectAtIndex, Int2Pointer(i))
}
