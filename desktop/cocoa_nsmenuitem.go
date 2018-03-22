// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSStatusBar_Class

var NSMenuItemClass unsafe.Pointer = Runtime_objc_lookUpClass("NSMenuItem")
var NSMenuItemSetTitle unsafe.Pointer = Runtime_sel_getUid("setTitle:")
var NSMenuItemSetImage unsafe.Pointer = Runtime_sel_getUid("setImage:")
var NSMenuItemSetEnabled unsafe.Pointer = Runtime_sel_getUid("setEnabled:")
var NSMenuItemSeparatorItemSel unsafe.Pointer = Runtime_sel_getUid("separatorItem")
var NSMenuItemSetSubmenu unsafe.Pointer = Runtime_sel_getUid("setSubmenu:")
var NSMenuItemSetState unsafe.Pointer = Runtime_sel_getUid("setState:")
var NSMenuItemSetTarget unsafe.Pointer = Runtime_sel_getUid("setTarget:")
var NSMenuItemSetAction unsafe.Pointer = Runtime_sel_getUid("setAction:")
var NSMenuItemSubmenu unsafe.Pointer = Runtime_sel_getUid("submenu")
var NSMenuItemTag unsafe.Pointer = Runtime_sel_getUid("tag")
var NSMenuItemTitle unsafe.Pointer = Runtime_sel_getUid("title")
var NSMenuItemInitWithTitleActionKeyEquivalentSel unsafe.Pointer = Runtime_sel_getUid("initWithTitle:action:keyEquivalent:")

type NSMenuItem struct {
	NSObject
}

func NSMenuItemNew() NSMenuItem {
	return NSMenuItemPointer(Runtime_class_createInstance(NSMenuItemClass, 0))
}

func NSMenuItemSeparatorItem() NSMenuItem {
	return NSMenuItem{NSObjectPointer(Runtime_objc_msgSend(NSMenuItemClass, NSMenuItemSeparatorItemSel))}
}

func NSMenuItemInitWithTitleActionKeyEquivalent(name string, action unsafe.Pointer, code string) NSMenuItem {
	var m NSMenuItem = NSMenuItemNew()
	n := NSStringNew(name)
	defer n.Release()
	k := NSStringNew(code)
	defer k.Release()
	Runtime_objc_msgSend(m.Pointer, NSMenuItemInitWithTitleActionKeyEquivalentSel, n.Pointer, action, k.Pointer)
	return m
}

func NSMenuItemPointer(p unsafe.Pointer) NSMenuItem {
	return NSMenuItem{NSObjectPointer(p)}
}

func (m NSMenuItem) SetTitle(s string) {
	n := NSStringNew(s)
	defer n.Release()
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetTitle, n.Pointer)
}

func (m NSMenuItem) SetImage(i NSImage) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetImage, i.Pointer)
}

func (m NSMenuItem) SetEnabled(b bool) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetEnabled, Bool2Pointer(b))
}

func (m NSMenuItem) SetSubmenu(i NSMenu) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetSubmenu, i.Pointer)
}

func (m NSMenuItem) SetState(i int) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetState, Int2Pointer(i))
}

func (m NSMenuItem) SetTarget(p unsafe.Pointer) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetTarget, p)
}

func (m NSMenuItem) SetAction(o unsafe.Pointer) {
	Runtime_objc_msgSend(m.Pointer, NSMenuItemSetAction, o)
}

func (m NSMenuItem) Submenu() NSMenu {
	return NSMenuPointer(Runtime_objc_msgSend(m.Pointer, NSMenuItemSubmenu))
}

func (m NSMenuItem) Tag() int {
	return Pointer2Int(Runtime_objc_msgSend(m.Pointer, NSMenuItemTag))
}

func (m NSMenuItem) Title() string {
	return NSStringPointer2String(Runtime_objc_msgSend(m.Pointer, NSMenuItemTitle))
}
