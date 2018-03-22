// +build darwin

package desktop

/*
extern void DesktopOSXSysTrayPolicyActionMap(void*, void*, void*, void*, void*, void*, void*);
*/
import "C"

import (
	"unsafe"
)

//
// register
//

var DesktopOSXSysTrayPolicyClassReg unsafe.Pointer = Runtime_objc_allocateClassPair(NSObjectClass, "DesktopOSXSysTrayPolicyClass", 0)

var DesktopOSXSysTrayPolicyMap = make(map[unsafe.Pointer]*DesktopOSXSysTrayPolicy)

//export DesktopOSXSysTrayPolicyActionMap
func DesktopOSXSysTrayPolicyActionMap(id unsafe.Pointer, sel unsafe.Pointer, arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer, arg4 unsafe.Pointer) {
	if sel == WebPolicyDelegateDecidePolicyForNavigationAction {
		DesktopOSXSysTrayPolicyMap[id].DecidePolicyForNavigationAction(WebViewPointer(arg0), arg1, NSURLRequestPointer(arg2), WebFramePointer(arg3), WebPolicyDecisionListenerPointer(arg4))
	}
	if sel == WebPolicyDelegateDecidePolicyForNewWindowAction {
		DesktopOSXSysTrayPolicyMap[id].DecidePolicyForNewWindowAction(WebViewPointer(arg0), arg1, NSURLRequestPointer(arg2), WebFramePointer(arg3), WebPolicyDecisionListenerPointer(arg4))
	}
}

func DesktopOSXSysTrayPolicyRegister() bool {
	if !Runtime_class_addMethod(DesktopOSXSysTrayPolicyClassReg, WebPolicyDelegateDecidePolicyForNavigationAction, C.DesktopOSXSysTrayPolicyActionMap, "v@::::::") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayPolicyClassReg, WebPolicyDelegateDecidePolicyForNewWindowAction, C.DesktopOSXSysTrayPolicyActionMap, "v@::::::") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayPolicyClassReg, WebPolicyDelegateUnableToImplementPolicyWithError, C.DesktopOSXSysTrayPolicyActionMap, "v@::::") {
		panic("problem initalizing class")
	}
	if !Runtime_class_addMethod(DesktopOSXSysTrayPolicyClassReg, WebPolicyDelegateDecidePolicyForMIMEType, C.DesktopOSXSysTrayPolicyActionMap, "v@::::::") {
		panic("problem initalizing class")
	}
	Runtime_objc_registerClassPair(DesktopOSXSysTrayPolicyClassReg)
	return true
}

var DesktopOSXSysTrayPolicyRegistred bool = DesktopOSXSysTrayPolicyRegister()

//
// object
//

var DesktopOSXSysTrayPolicyClass unsafe.Pointer = Runtime_objc_lookUpClass("DesktopOSXSysTrayPolicyClass")

type DesktopOSXSysTrayPolicy struct {
	NSObject

	s *DesktopSysTray
}

func DesktopOSXSysTrayPolicyNew(s *DesktopSysTray) *DesktopOSXSysTrayPolicy {
	m := DesktopOSXSysTrayPolicyPointer(Runtime_class_createInstance(DesktopOSXSysTrayPolicyClass, 0))

	m.s = s

	return m
}

func DesktopOSXSysTrayPolicyPointer(p unsafe.Pointer) *DesktopOSXSysTrayPolicy {
	m := DesktopOSXSysTrayPolicy{NSObjectPointer(p), nil}

	DesktopOSXSysTrayPolicyMap[m.Pointer] = &m

	return &m
}

func (m *DesktopOSXSysTrayPolicy) DecidePolicyForNavigationAction(webview WebView, actionInformation unsafe.Pointer, request NSURLRequest, frame WebFrame, listener WebPolicyDecisionListener) {
	listener.Use()
}

func (m *DesktopOSXSysTrayPolicy) DecidePolicyForNewWindowAction(webview WebView, actionInformation unsafe.Pointer, request NSURLRequest, frame WebFrame, listener WebPolicyDecisionListener) {
	listener.Ignore()
	s := request.GetURL().AbsoluteString()
	if m.s.WebPopup.Handler != nil {
		if m.s.WebPopup.Handler(s) {
			return
		}
	}
	BrowserOpenURI(s) // m.s.WebView.GetMainFrame().LoadRequest(s)
}

func (m *DesktopOSXSysTrayPolicy) Release() {
	delete(DesktopOSXSysTrayPolicyMap, m.Pointer)
	m.NSObject.Release()
}
