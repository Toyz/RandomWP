// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSDate_Class

var NSDateClass unsafe.Pointer = Runtime_objc_lookUpClass("NSDate")
var NSDateDistantFutureSel unsafe.Pointer = Runtime_sel_getUid("distantFuture")

type NSDate struct {
	NSObject
}

func NSDateDistantFuture() NSDate {
	return NSDate{NSObjectPointer(Runtime_objc_msgSend(NSDateClass, NSDateDistantFutureSel))}
}

func NSDatePointer(p unsafe.Pointer) NSDate {
	return NSDate{NSObjectPointer(p)}
}
