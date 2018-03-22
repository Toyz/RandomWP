// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>
*/
import "C"

type CGSize struct {
	width  CGFloat
	height CGFloat
}

func CGSizeNew(width, height float64) CGSize {
	return CGSize{CGFloat(width), CGFloat(height)}
}

func CGSizeC(i C.CGSize) CGSize {
	r := CGSize{}
	r.width = CGFloat(i.width)
	r.height = CGFloat(i.height)
	return r
}

func (m CGSize) C() C.CGSize {
	return C.CGSize{C.CGFloat(m.width), C.CGFloat(m.height)}
}
