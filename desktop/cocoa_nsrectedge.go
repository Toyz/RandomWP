// +build darwin

package desktop

import (
	"unsafe"
)

type NSRectEdge uint32

const NSMaxXEdge = NSRectEdgeMaxX
const NSMaxYEdge = NSRectEdgeMaxY
const NSMinXEdge = NSRectEdgeMinX
const NSMinYEdge = NSRectEdgeMinY
const NSRectEdgeMaxX = CGRectMaxXEdge
const NSRectEdgeMaxY = CGRectMaxYEdge
const NSRectEdgeMinX = CGRectMinXEdge
const NSRectEdgeMinY = CGRectMinYEdge

// CGGeometry.h enum CGRectEdge
const CGRectMinXEdge NSRectEdge = 0
const CGRectMinYEdge NSRectEdge = 1
const CGRectMaxXEdge NSRectEdge = 2
const CGRectMaxYEdge NSRectEdge = 3

func (m NSRectEdge) Pointer() unsafe.Pointer {
	return Int2Pointer(int(m))
}
