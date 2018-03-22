// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSFont_Class

var NSNumberFloatValue unsafe.Pointer = Runtime_sel_getUid("floatValue")
var NSNumberIntValue unsafe.Pointer = Runtime_sel_getUid("intValue")
var NSNumberDoubleValue unsafe.Pointer = Runtime_sel_getUid("doubleValue")
var NSNumberStringValue unsafe.Pointer = Runtime_sel_getUid("stringValue")

type NSNumber struct {
	NSObject
}

func NSNumberPointer(p unsafe.Pointer) NSNumber {
	return NSNumber{NSObjectPointer(p)}
}

func (m NSNumber) FloatValue() float64 {
	return Pointer2Float(Runtime_objc_msgSend(m.Pointer, NSNumberFloatValue))
}

func (m NSNumber) DoubleValue() float64 {
	return Pointer2Float(Runtime_objc_msgSend(m.Pointer, NSNumberDoubleValue))
}

func (m NSNumber) IntValue() int {
	return Pointer2Int(Runtime_objc_msgSend(m.Pointer, NSNumberIntValue))
}

func (m NSNumber) StringValue() string {
	return NSStringPointer2String(Runtime_objc_msgSend(m.Pointer, NSNumberStringValue))
}
