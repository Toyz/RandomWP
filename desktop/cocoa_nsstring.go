// +build darwin

package desktop

// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/Foundation/Classes/NSString_Class/Reference/NSString.html

var NSStringClass unsafe.Pointer = Runtime_objc_lookUpClass("NSString")
var NSStringStringWithUTF8String unsafe.Pointer = Runtime_sel_getUid("stringWithUTF8String:")
var NSStringUTF8String unsafe.Pointer = Runtime_sel_getUid("UTF8String")

type NSString struct {
	NSObject
}

func NSStringNew(s string) NSString {
	p := unsafe.Pointer(&[]byte(s + "\x00")[0])
	var m NSString = NSString{NSObjectPointer(Runtime_objc_msgSend(NSStringClass, NSStringStringWithUTF8String, p))}
	return m
}

func NSStringPointer(p unsafe.Pointer) NSString {
	return NSString{NSObjectPointer(p)}
}

func NSStringPointer2String(p unsafe.Pointer) string {
	s := NSStringPointer(p)
	defer s.Release()
	return s.String()
}

func (m NSString) String() string {
	return Pointer2String(Runtime_objc_msgSend(m.NSObject.Pointer, NSStringUTF8String))
}
