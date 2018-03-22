// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSWorkspace_Class/

var NSWorkspaceClass unsafe.Pointer = Runtime_objc_lookUpClass("NSWorkspace")
var NSWorkspaceSharedWorkspaceSel unsafe.Pointer = Runtime_sel_getUid("sharedWorkspace")
var NSWorkspaceOpenURL unsafe.Pointer = Runtime_sel_getUid("openURL:")

type NSWorkspace struct {
	NSObject
}

func NSWorkspaceSharedWorkspace() NSWorkspace {
	return NSWorkspace{NSObjectPointer(Runtime_objc_msgSend(NSWorkspaceClass, NSWorkspaceSharedWorkspaceSel))}
}

func (m NSWorkspace) OpenURL(u NSURL) bool {
	return Pointer2Bool(Runtime_objc_msgSend(m.Pointer, NSWorkspaceOpenURL, u.Pointer))
}
