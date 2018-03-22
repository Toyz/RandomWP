// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

CGSize NSImageSize(id to, SEL sel) {
  return ((CGSize(*)(id, SEL, ...))objc_msgSend)(to, sel);
}
*/
import "C"

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ApplicationKit/Classes/NSImage_Class

var NSImageClass unsafe.Pointer = Runtime_objc_lookUpClass("NSImage")
var NSImageInitWithData unsafe.Pointer = Runtime_sel_getUid("initWithData:")
var NSImageSize unsafe.Pointer = Runtime_sel_getUid("size")

type NSImage struct {
	NSObject
}

func Image2Bytes(p image.Image) []byte {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := png.Encode(w, p)
	// it we have correct image on memory we have to write it down.
	// error would mean out of memroy or similar error. so panic.
	if err != nil {
		panic(err)
	}
	w.Flush()
	return b.Bytes()
}

func NSImageNew() NSImage {
	return NSImagePointer(Runtime_class_createInstance(NSImageClass, 0))
}

func NSImageData(p NSData) NSImage {
	var m NSImage = NSImageNew()
	r := Runtime_objc_msgSend(m.Pointer, NSImageInitWithData, p.Pointer)
	m.Assert(r)
	return m
}

func NSImageImage(p image.Image) NSImage {
	var d NSData = NSDataNew(Image2Bytes(p))
	defer d.Release()
	return NSImageData(d)
}

func NSImagePointer(p unsafe.Pointer) NSImage {
	return NSImage{NSObjectPointer(p)}
}

func (m NSImage) GetSize() NSSize {
	p := C.NSImageSize(m.Pointer, NSImageSize)
	return NSSizeC(p)
}
