// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>
*/
import "C"

type NSSize CGSize

func NSSizeNew(x, y float64) NSSize {
	p := NSSize{CGFloat(x), CGFloat(y)}
	return p
}

func NSSizeC(i C.CGSize) NSSize {
	return NSSize(CGSizeC(i))
}

func (m NSSize) C() C.CGSize {
	return CGSize(m).C()
}
