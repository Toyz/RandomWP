// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>
*/
import "C"

type CGRect struct {
	origin CGPoint
	size   CGSize
}

func CGRectSize(i CGSize) CGRect {
	return CGRect{CGPointNew(0, 0), i}
}

func CGRectC(i C.CGRect) CGRect {
	r := CGRect{}
	r.origin = CGPointC(i.origin)
	r.size = CGSizeC(i.size)
	return r
}

func (m CGRect) C() C.CGRect {
	return C.CGRect{m.origin.C(), m.size.C()}
}
