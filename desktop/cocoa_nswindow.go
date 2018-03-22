// +build darwin

package desktop

import (
	"unsafe"
)

var NSWindowClass unsafe.Pointer = Runtime_objc_lookUpClass("NSWindow")
var NSWindowBecomeKeyWindow unsafe.Pointer = Runtime_sel_getUid("becomeKeyWindow")
var NSWindowCanBecomeKeyWindow unsafe.Pointer = Runtime_sel_getUid("canBecomeKeyWindow")

type NSWindow struct {
	NSResponder
}

func NSWindowPointer(p unsafe.Pointer) NSWindow {
	return NSWindow{NSResponderPointer(p)}
}

func (m NSWindow) BecomeKeyWindow() {
	Runtime_objc_msgSend(m.Pointer, NSWindowBecomeKeyWindow)
}

func (m NSWindow) CanBecomeKeyWindow() bool {
	return Pointer2Bool(Runtime_objc_msgSend(m.Pointer, NSWindowCanBecomeKeyWindow))
}
