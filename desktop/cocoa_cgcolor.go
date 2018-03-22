// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGColor.h>
*/
import "C"

import (
	"unsafe"
)

// https://developer.apple.com/reference/appkit/nsviewcontroller?language=objc

type CGColor struct {
	Pointer unsafe.Pointer
}

func CGColorRGB(r, g, b, a int) CGColor {
	return CGColor{unsafe.Pointer(C.CGColorCreateGenericRGB(C.CGFloat(r), C.CGFloat(g), C.CGFloat(b), C.CGFloat(a)))}
}
