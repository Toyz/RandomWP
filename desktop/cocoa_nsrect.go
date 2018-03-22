// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>
*/
import "C"

type NSRect CGRect

func NSRectNew(x, y, xx, yy float64) NSRect {
	r := NSRect{}
	r.origin.x = CGFloat(x)
	r.origin.y = CGFloat(y)
	r.size.width = CGFloat(xx)
	r.size.height = CGFloat(yy)
	return r
}

func NSRectSize(size NSSize) NSRect {
	return NSRect(CGRectSize(CGSize(size)))
}

func NSRectC(i C.CGRect) NSRect {
	return NSRect(CGRectC(i))
}

func NSRectZero() NSRect {
	return NSRect{}
}

func (m NSRect) C() C.CGRect {
	return CGRect(m).C()
}
