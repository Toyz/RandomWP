// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSFont_Class

var NSFontClass unsafe.Pointer = Runtime_objc_lookUpClass("NSFont")
var NSFontMenuFontOfSizeSel unsafe.Pointer = Runtime_sel_getUid("menuFontOfSize:")
var NSFontMenuBarFontOfSizeSel unsafe.Pointer = Runtime_sel_getUid("menuBarFontOfSize:")
var NSFontPointSize unsafe.Pointer = Runtime_sel_getUid("pointSize")
var NSFontFontName unsafe.Pointer = Runtime_sel_getUid("fontName")
var NSFontDisplayName unsafe.Pointer = Runtime_sel_getUid("displayName")
var NSFontFontDescriptor unsafe.Pointer = Runtime_sel_getUid("fontDescriptor")

type NSFont struct {
	NSObject
}

func NSFontMenuFontOfSize(i int) NSFont {
	return NSFont{NSObjectPointer(Runtime_objc_msgSend(NSFontClass, NSFontMenuFontOfSizeSel, Float2Pointer(float64(i))))}
}

func NSFontMenuBarFontOfSize(i int) NSFont {
	return NSFont{NSObjectPointer(Runtime_objc_msgSend(NSFontClass, NSFontMenuBarFontOfSizeSel, Float2Pointer(float64(i))))}
}

func (m NSFont) PointSize() CGFloat {
	return CGFloatPointer(Runtime_objc_msgSend(m.Pointer, NSFontPointSize))
}

func (m NSFont) FontName() string {
	return NSStringPointer2String(Runtime_objc_msgSend(m.Pointer, NSFontFontName))
}

func (m NSFont) DisplayName() string {
	return NSStringPointer2String(Runtime_objc_msgSend(m.Pointer, NSFontDisplayName))
}

func (m NSFont) FontDescriptor() NSFontDescriptor {
	return NSFontDescriptorPointer(Runtime_objc_msgSend(m.Pointer, NSFontFontDescriptor))
}
