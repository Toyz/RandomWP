// +build darwin

package desktop

import (
	"unsafe"
)

var WebFrameClass unsafe.Pointer = Runtime_objc_lookUpClass("WebFrame")
var WebFrameLoadRequest unsafe.Pointer = Runtime_sel_getUid("loadRequest:")
var WebFrameLoadHTMLString unsafe.Pointer = Runtime_sel_getUid("loadHTMLString:baseURL:")

type WebFrame struct {
	NSObject
}

func WebFramePointer(p unsafe.Pointer) WebFrame {
	return WebFrame{NSObjectPointer(p)}
}

func (m WebFrame) LoadRequest(url string) {
	u := NSURLRequestNew(url)
	defer u.Release()
	Runtime_objc_msgSend(m.Pointer, WebFrameLoadRequest, u.Pointer)
}

func (m WebFrame) LoadHTMLString(base string, html string) {
	var b NSURL
	if base != "" {
		b := NSURLNew(base)
		defer b.Release()
	}
	u := NSStringNew(html)
	defer u.Release()
	Runtime_objc_msgSend(m.Pointer, WebFrameLoadHTMLString, u.Pointer, b.Pointer)
}
