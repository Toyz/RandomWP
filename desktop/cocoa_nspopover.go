// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

id NSPopoverShowRelativeToRect(id to, SEL sel, CGRect rect, SEL ofView, SEL preferredEdge) {
  return objc_msgSend(to, sel, rect, ofView, preferredEdge);
}
*/
import "C"

import (
	"unsafe"
)

// https://developer.apple.com/reference/appkit/nspopover

const NSPopoverBehaviorApplicationDefined = 0
const NSPopoverBehaviorTransient = 1
const NSPopoverBehaviorSemitransient = 2

var NSPopoverClass unsafe.Pointer = Runtime_objc_lookUpClass("NSPopover")
var NSPopoverInit unsafe.Pointer = Runtime_sel_getUid("init")
var NSPopoverShowRelativeToRect unsafe.Pointer = Runtime_sel_getUid("showRelativeToRect:ofView:preferredEdge:")
var NSPopoverPerformClose unsafe.Pointer = Runtime_sel_getUid("performClose:")
var NSPopoverSetContentViewController unsafe.Pointer = Runtime_sel_getUid("setContentViewController:")
var NSPopoverSetBehavior unsafe.Pointer = Runtime_sel_getUid("setBehavior:")
var NSPopoverIsShown unsafe.Pointer = Runtime_sel_getUid("isShown")

type NSPopover struct {
	NSResponder
}

func NSPopoverNew() NSPopover {
	var m = NSPopoverPointer(Runtime_class_createInstance(NSPopoverClass, 0))
	p := Runtime_objc_msgSend(m.Pointer, NSPopoverInit)
	m.Assert(p)
	return m
}

func NSPopoverPointer(p unsafe.Pointer) NSPopover {
	return NSPopover{NSResponderPointer(p)}
}

func (m NSPopover) ShowRelativeToRect(positioningRect NSRect, ofView NSView, preferredEdge NSRectEdge) {
	C.NSPopoverShowRelativeToRect(m.Pointer,
		NSPopoverShowRelativeToRect,
		positioningRect.C(),
		ofView.Pointer,
		preferredEdge.Pointer(),
	)
}

func (m NSPopover) PerformClose() {
	Runtime_objc_msgSend(m.Pointer, NSPopoverPerformClose)
}

func (m NSPopover) SetContentViewController(i NSViewController) {
	Runtime_objc_msgSend(m.Pointer, NSPopoverSetContentViewController, i.Pointer)
}

func (m NSPopover) SetBehavior(i int) {
	Runtime_objc_msgSend(m.Pointer, NSPopoverSetBehavior, Int2Pointer(i))
}

func (m NSPopover) GetShown() bool {
	return Pointer2Bool(Runtime_objc_msgSend(m.Pointer, NSPopoverIsShown))
}
