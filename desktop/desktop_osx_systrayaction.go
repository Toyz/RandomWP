// +build darwin

package desktop

/*
extern void DesktopOSXSysTrayActionActionMap(void*, void*);
*/
import "C"

import (
	"unsafe"
)

//
// register
//

var DesktopOSXSysTrayActionClassReg unsafe.Pointer = Runtime_objc_allocateClassPair(NSObjectClass, "DesktopOSXSysTrayActionClass", 0)
var DesktopOSXSysTrayActionActionReg unsafe.Pointer = Runtime_sel_registerName("action")

var DesktopOSXSysTrayActionMap = make(map[unsafe.Pointer]*DesktopOSXSysTrayAction)

//export DesktopOSXSysTrayActionActionMap
func DesktopOSXSysTrayActionActionMap(id unsafe.Pointer, sel unsafe.Pointer) {
	if sel == DesktopOSXSysTrayActionActionReg {
		DesktopOSXSysTrayActionMap[id].Action()
	}
}

func DesktopOSXSysTrayActionRegister() bool {
	if !Runtime_class_addMethod(DesktopOSXSysTrayActionClassReg, DesktopOSXSysTrayActionActionReg, C.DesktopOSXSysTrayActionActionMap, "v@:") {
		panic("problem initalizing class")
	}
	Runtime_objc_registerClassPair(DesktopOSXSysTrayActionClassReg)

	return true
}

var DesktopOSXSysTrayActionRegistred bool = DesktopOSXSysTrayActionRegister()

//
// object
//

var DesktopOSXSysTrayActionClass unsafe.Pointer = Runtime_objc_lookUpClass("DesktopOSXSysTrayActionClass")
var DesktopOSXSysTrayActionAction unsafe.Pointer = Runtime_sel_getUid("action")

type DesktopOSXSysTrayAction struct {
	NSObject

	Menu *Menu
}

func DesktopOSXSysTrayActionNew(mn *Menu) *DesktopOSXSysTrayAction {
	m := DesktopOSXSysTrayActionPointer(Runtime_class_createInstance(DesktopOSXSysTrayActionClass, 0))

	m.Menu = mn

	return m
}

func DesktopOSXSysTrayActionPointer(p unsafe.Pointer) *DesktopOSXSysTrayAction {
	m := &DesktopOSXSysTrayAction{NSObjectPointer(p), nil}

	DesktopOSXSysTrayActionMap[m.Pointer] = m

	return m
}

func (m *DesktopOSXSysTrayAction) Action() {
	m.Menu.Action(m.Menu)
}

func (m *DesktopOSXSysTrayAction) Release() {
	delete(DesktopOSXSysTrayActionMap, m.Pointer)
	m.NSObject.Release()
}
