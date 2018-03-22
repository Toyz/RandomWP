// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSStatusBar_Class

const NSVariableStatusItemLength = -1
const NSSquareStatusItemLength = -2

var NSStatusBarClass unsafe.Pointer = Runtime_objc_lookUpClass("NSStatusBar")
var NSStatusBarSystemStatusBarSel unsafe.Pointer = Runtime_sel_getUid("systemStatusBar")
var NSStatusBarStatusItemWithLength unsafe.Pointer = Runtime_sel_getUid("statusItemWithLength:")
var NSStatusBarRemoveStatusItem unsafe.Pointer = Runtime_sel_getUid("removeStatusItem:")
var NSStatusBarThickness unsafe.Pointer = Runtime_sel_getUid("thickness")

type NSStatusBar struct {
	NSObject
}

func NSStatusBarSystemStatusBar() NSStatusBar {
	var m NSStatusBar = NSStatusBar{NSObjectPointer(Runtime_objc_msgSend(NSStatusBarClass, NSStatusBarSystemStatusBarSel))}
	return m
}

func (m NSStatusBar) StatusItemWithLength(i int) NSStatusItem {
	return NSStatusItemPointer(Runtime_objc_msgSend(m.Pointer, NSStatusBarStatusItemWithLength, Float2Pointer(float64(i))))
}

func (m NSStatusBar) RemoveStatusItem(i NSStatusItem) {
	Runtime_objc_msgSend(m.Pointer, NSStatusBarRemoveStatusItem, i.Pointer)
}

func (m NSStatusBar) Thickness() CGFloat {
	return CGFloatPointer(Runtime_objc_msgSend(m.Pointer, NSStatusBarThickness))
}
