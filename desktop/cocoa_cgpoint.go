// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>
*/
import "C"

type CGPoint struct {
	x CGFloat
	y CGFloat
}

func CGPointNew(x, y float64) CGPoint {
	return CGPoint{CGFloat(x), CGFloat(y)}
}

func CGPointC(i C.CGPoint) CGPoint {
	r := CGPoint{}
	r.x = CGFloat(i.x)
	r.y = CGFloat(i.y)
	return r
}

func (m CGPoint) C() C.CGPoint {
	return C.CGPoint{C.CGFloat(m.x), C.CGFloat(m.y)}
}
