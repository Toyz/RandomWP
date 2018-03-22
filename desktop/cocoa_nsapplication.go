// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/reference/appkit/nsapplication?language=objc

var NSApplicationClass unsafe.Pointer = Runtime_objc_lookUpClass("NSApplication")
var NSApplicationSharedApplicationSel unsafe.Pointer = Runtime_sel_getUid("sharedApplication")
var NSApplicationRun unsafe.Pointer = Runtime_sel_getUid("run")
var NSApplicationNextEventMatchingMaskUntilDateInModeDequeue unsafe.Pointer = Runtime_sel_getUid("nextEventMatchingMask:untilDate:inMode:dequeue:")
var NSApplicationUpdateWindows unsafe.Pointer = Runtime_sel_getUid("updateWindows")
var NSApplicationSendEvent unsafe.Pointer = Runtime_sel_getUid("sendEvent:")
var NSApplicationTerminate unsafe.Pointer = Runtime_sel_getUid("terminate:")
var NSApplicationForceTerminate unsafe.Pointer = Runtime_sel_getUid("forceTerminate")

type NSApplication struct {
	NSObject
}

func NSApplicationSharedApplication() NSApplication {
	return NSApplication{NSObjectPointer(Runtime_objc_msgSend(NSApplicationClass, NSApplicationSharedApplicationSel))}
}

func NSApplicationPointer(p unsafe.Pointer) NSApplication {
	return NSApplication{NSObjectPointer(p)}
}

func (m NSApplication) Run() {
	Runtime_objc_msgSend(m.Pointer, NSApplicationRun)
}

func (m NSApplication) NextEventMatchingMaskUntilDateInModeDequeue(mask int, expiration NSDate, mode string, flag bool) NSEvent {
	n := NSStringNew(mode)
	defer n.Release()
	return NSEventPointer(Runtime_objc_msgSend(m.Pointer, NSApplicationNextEventMatchingMaskUntilDateInModeDequeue, Int2Pointer(mask), expiration.Pointer, n.Pointer, Bool2Pointer(flag)))
}

func (m NSApplication) UpdateWindows() {
	Runtime_objc_msgSend(m.Pointer, NSApplicationUpdateWindows)
}

func (m NSApplication) SendEvent(e NSEvent) {
	Runtime_objc_msgSend(m.Pointer, NSApplicationSendEvent, e.Pointer)
}

func (m NSApplication) Terminate() {
	Runtime_objc_msgSend(m.Pointer, NSApplicationTerminate, m.Pointer)
}

func (m NSApplication) ForceTerminate() {
	Runtime_objc_msgSend(m.Pointer, NSApplicationForceTerminate)
}
