// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ApplicationKit/Classes/NSImage_Class

var NSDataClass unsafe.Pointer = Runtime_objc_lookUpClass("NSData")
var NSDataDataWithBytesLength unsafe.Pointer = Runtime_sel_getUid("dataWithBytes:length:")

type NSData struct {
	NSObject
}

func NSDataNew(b []byte) NSData {
	return NSData{NSObjectPointer(Runtime_objc_msgSend(NSDataClass, NSDataDataWithBytesLength, unsafe.Pointer(&b[0]), Int2Pointer(len(b))))}
}

func NSDataPointer(p unsafe.Pointer) NSData {
	return NSData{NSObjectPointer(p)}
}
