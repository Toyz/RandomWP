// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

id NSViewInitWithFrame(id to, SEL sel, CGRect rect) {
  return objc_msgSend(to, sel, rect);
}

CGRect NSViewFrame(id to, SEL sel) {
  return ((CGRect(*)(id, SEL, ...))objc_msgSend_stret)(to, sel);
}

CGRect NSViewBounds(id to, SEL sel) {
  return ((CGRect(*)(id, SEL, ...))objc_msgSend_stret)(to, sel);
}
*/
import "C"

import (
	"unsafe"
)

var NSViewClass unsafe.Pointer = Runtime_objc_lookUpClass("NSView")
var NSViewInitWithFrame unsafe.Pointer = Runtime_sel_getUid("initWithFrame:")
var NSViewSetFrameOrigin unsafe.Pointer = Runtime_sel_getUid("setFrameOrigin:")
var NSViewSetFrameSize unsafe.Pointer = Runtime_sel_getUid("setFrameSize:")
var NSViewFrame unsafe.Pointer = Runtime_sel_getUid("frame")
var NSViewBounds unsafe.Pointer = Runtime_sel_getUid("bounds")
var NSViewSetWantsLayer unsafe.Pointer = Runtime_sel_getUid("setWantsLayer:")
var NSViewGetWantsLayer unsafe.Pointer = Runtime_sel_getUid("wantsLayer")
var NSViewGetLayer unsafe.Pointer = Runtime_sel_getUid("layer")
var NSViewGetWindow unsafe.Pointer = Runtime_sel_getUid("window")

type NSView struct {
	NSResponder
}

func NSViewNew(rect NSRect) NSView {
	var m = NSViewPointer(Runtime_class_createInstance(NSViewClass, 0))
	p := unsafe.Pointer(C.NSViewInitWithFrame((*C.struct_objc_object)(m.Pointer), (*C.struct_objc_selector)(NSViewInitWithFrame), rect.C()))
	m.Assert(p)
	return m
}

func NSViewPointer(p unsafe.Pointer) NSView {
	return NSView{NSResponderPointer(p)}
}

func (m NSView) SetFrameOrigin(p NSPoint) {
	Runtime_objc_msgSend(m.Pointer, NSViewSetFrameOrigin,
		p.x.Pointer(), p.y.Pointer(),
	)
}

func (m NSView) SetFrameSize(p NSSize) {
	Runtime_objc_msgSend(m.Pointer, NSViewSetFrameSize, p.width.Pointer(), p.height.Pointer())
}

func (m NSView) GetFrame() NSRect {
	p := C.NSViewFrame(m.Pointer, NSViewFrame)
	return NSRectC(p)
}

func (m NSView) GetBounds() NSRect {
	p := C.NSViewBounds(m.Pointer, NSViewBounds)
	return NSRectC(p)
}

func (m NSView) GetWindow() NSWindow {
	p := Runtime_objc_msgSend(m.Pointer, NSViewGetWindow)
	return NSWindowPointer(p)
}

func (m NSView) SetWantsLayer(b bool) {
	Runtime_objc_msgSend(m.Pointer, NSViewSetWantsLayer, Bool2Pointer(b))
}

func (m NSView) GetWantsLayer() bool {
	return Pointer2Bool(Runtime_objc_msgSend(m.Pointer, NSViewGetWantsLayer))
}

func (m NSView) GetLayer() CALayer {
	return CALayerPointer(Runtime_objc_msgSend(m.Pointer, NSViewGetLayer))
}
