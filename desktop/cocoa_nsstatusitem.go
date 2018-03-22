// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSStatusBar_Class

var NSStatusItemClass unsafe.Pointer = Runtime_objc_lookUpClass("NSStatusItem")
var NSStatusItemSetHighlightMode unsafe.Pointer = Runtime_sel_getUid("setHighlightMode:")
var NSStatusItemSetImage unsafe.Pointer = Runtime_sel_getUid("setImage:")
var NSStatusItemSetView unsafe.Pointer = Runtime_sel_getUid("setView:")
var NSStatusItemGetView unsafe.Pointer = Runtime_sel_getUid("view")
var NSStatusItemSetTitle unsafe.Pointer = Runtime_sel_getUid("setTitle:")
var NSStatusItemSetMenu unsafe.Pointer = Runtime_sel_getUid("setMenu:")
var NSStatusItemSetToolTip unsafe.Pointer = Runtime_sel_getUid("setToolTip:")

type NSStatusItem struct {
	NSObject
}

func NSStatusItemPointer(p unsafe.Pointer) NSStatusItem {
	return NSStatusItem{NSObjectPointer(p)}
}

func (m NSStatusItem) SetHighlightMode(b bool) {
	Runtime_objc_msgSend(m.Pointer, NSStatusItemSetHighlightMode, Bool2Pointer(b))
}

func (m NSStatusItem) SetImage(i NSImage) {
	Runtime_objc_msgSend(m.Pointer, NSStatusItemSetImage, i.Pointer)
}

func (m NSStatusItem) SetMenu(i NSMenu) {
	Runtime_objc_msgSend(m.Pointer, NSStatusItemSetMenu, i.Pointer)
}

func (m NSStatusItem) SetToolTip(s string) {
	n := NSStringNew(s)
	defer n.Release()
	Runtime_objc_msgSend(m.Pointer, NSStatusItemSetToolTip, n.Pointer)
}

func (m NSStatusItem) SetView(i NSView) {
	Runtime_objc_msgSend(m.Pointer, NSStatusItemSetView, i.Pointer)
}

func (m NSStatusItem) GetView() NSView {
	return NSViewPointer(Runtime_objc_msgSend(m.Pointer, NSStatusItemGetView))
}
