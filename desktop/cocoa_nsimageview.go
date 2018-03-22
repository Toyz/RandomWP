// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

id NSImageViewInitWithFrame(id to,  SEL sel, CGRect rect) {
  return objc_msgSend(to, sel, rect);
}
*/
import "C"

import (
	"unsafe"
)

var NSImageViewClass unsafe.Pointer = Runtime_objc_lookUpClass("NSImageView")
var NSImageViewInitWithFrame unsafe.Pointer = Runtime_sel_getUid("initWithFrame:")
var NSImageViewSetImage unsafe.Pointer = Runtime_sel_getUid("setImage:")

type NSImageView struct {
	NSControl
}

func NSImageViewNew(rect NSRect) NSImageView {
	m := NSImageViewPointer(Runtime_class_createInstance(NSImageViewClass, 0))
	m.Init(rect)
	return m
}

func (m NSImageView) Init(rect NSRect) {
	p := unsafe.Pointer(C.NSImageViewInitWithFrame(m.Pointer, NSImageViewInitWithFrame, rect.C()))
	m.Assert(p)
}

func NSImageViewPointer(p unsafe.Pointer) NSImageView {
	return NSImageView{NSControlPointer(p)}
}

func (m NSImageView) SetImage(i NSImage) {
	Runtime_objc_msgSend(m.Pointer, NSImageViewSetImage, i.Pointer)
}
