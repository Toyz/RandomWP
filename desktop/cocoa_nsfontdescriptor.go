// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSFont_Class

const NSFontNameAttribute = "NSFontNameAttribute"
const NSFontFamilyAttribute = "NSFontFamilyAttribute"
const NSFontSizeAttribute = "NSFontSizeAttribute"
const NSFontMatrixAttribute = "NSFontMatrixAttribute"
const NSFontCharacterSetAttribute = "NSFontCharacterSetAttribute"
const NSFontTraitsAttribute = "NSFontTraitsAttribute"
const NSFontFaceAttribute = "NSFontFaceAttribute"
const NSFontFixedAdvanceAttribute = "NSFontFixedAdvanceAttribute"
const NSFontVisibleNameAttribute = "NSFontVisibleNameAttribute"

var NSFontDescriptorObjectForKey unsafe.Pointer = Runtime_sel_getUid("objectForKey:")

type NSFontDescriptor struct {
	NSObject
}

func NSFontDescriptorPointer(p unsafe.Pointer) NSFontDescriptor {
	return NSFontDescriptor{NSObjectPointer(p)}
}

func (m NSFontDescriptor) ObjectForKey(key string) unsafe.Pointer {
	n := NSStringNew(key)
	defer n.Release()
	return Runtime_objc_msgSend(m.Pointer, NSFontDescriptorObjectForKey, n.Pointer)
}
