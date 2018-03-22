// +build darwin

package desktop

import (
	"unsafe"
)

type CGFloat float64

func CGFloatPointer(p unsafe.Pointer) CGFloat {
	return CGFloat(Pointer2Float(p))
}

func (m CGFloat) Pointer() unsafe.Pointer {
	return Float2Pointer(float64(m))
}
