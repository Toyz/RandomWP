// +build darwin

package desktop

/*
#cgo LDFLAGS: -lobjc -framework WebKit

#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

id WebViewInitWithFrame(id to, SEL sel, CGRect rect, SEL frame, SEL group) {
  return objc_msgSend(to, sel, rect, frame, group);
}
*/
import "C"

import (
	"unsafe"
)

var WebViewClass unsafe.Pointer = Runtime_objc_lookUpClass("WebView")
var WebViewInitWithFrame unsafe.Pointer = Runtime_sel_getUid("initWithFrame:frameName:groupName:")
var WebViewMainFrame unsafe.Pointer = Runtime_sel_getUid("mainFrame")
var WebViewSetPolicyDelegate unsafe.Pointer = Runtime_sel_getUid("setPolicyDelegate:")

type WebView struct {
	NSView
}

func WebViewNew(rect NSRect, frame string, group string) WebView {
	f := NSStringNew(frame)
	defer f.Release()
	g := NSStringNew(group)
	defer g.Release()
	var m = WebViewPointer(Runtime_class_createInstance(WebViewClass, 0))
	p := unsafe.Pointer(C.WebViewInitWithFrame((*C.struct_objc_object)(m.Pointer), (*C.struct_objc_selector)(WebViewInitWithFrame), rect.C(), f.Pointer, g.Pointer))
	m.Assert(p)
	return m
}

func WebViewPointer(p unsafe.Pointer) WebView {
	return WebView{NSViewPointer(p)}
}

func (m WebView) GetMainFrame() WebFrame {
	return WebFramePointer(Runtime_objc_msgSend(m.Pointer, WebViewMainFrame))
}

func (m WebView) SetPolicyDelegate(p NSObject) {
	Runtime_objc_msgSend(m.Pointer, WebViewSetPolicyDelegate, p.Pointer)
}
