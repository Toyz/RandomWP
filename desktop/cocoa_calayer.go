// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/reference/appkit/nsviewcontroller?language=objc

var CALayerClass unsafe.Pointer = Runtime_objc_lookUpClass("CALayer")
var CALayerSetBackgroundColor unsafe.Pointer = Runtime_sel_getUid("setBackgroundColor:")

type CALayer struct {
	NSObject
}

func CALayerPointer(p unsafe.Pointer) CALayer {
	return CALayer{NSObjectPointer(p)}
}

func (m CALayer) SetBackgroundColor(c CGColor) {
	Runtime_objc_msgSend(m.Pointer, CALayerSetBackgroundColor, c.Pointer)
}
