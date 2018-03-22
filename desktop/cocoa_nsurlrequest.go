// +build darwin

package desktop

import (
	"unsafe"
)

var NSURLRequestClass unsafe.Pointer = Runtime_objc_lookUpClass("NSURLRequest")
var NSURLRequestRequestWithURL unsafe.Pointer = Runtime_sel_getUid("requestWithURL:")
var NSURLRequestURL unsafe.Pointer = Runtime_sel_getUid("URL")

type NSURLRequest struct {
	NSObject
}

func NSURLRequestNew(url string) NSURLRequest {
	u := NSURLNew(url)
	defer u.Release()
	var m = NSURLRequestPointer(Runtime_objc_msgSend(NSURLRequestClass, NSURLRequestRequestWithURL, u.Pointer))
	return m
}

func NSURLRequestPointer(p unsafe.Pointer) NSURLRequest {
	return NSURLRequest{NSObjectPointer(p)}
}

func (m NSURLRequest) GetURL() NSURL {
	p := Runtime_objc_msgSend(m.Pointer, NSURLRequestURL)
	return NSURLPointer(p)
}
