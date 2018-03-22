// +build darwin

package desktop

import (
	"unsafe"
)

var WebPolicyDelegateClass unsafe.Pointer = Runtime_objc_lookUpClass("WebPolicyDelegate")
var WebPolicyDelegateInit unsafe.Pointer = Runtime_sel_getUid("init")
var WebPolicyDelegateDecidePolicyForNavigationAction unsafe.Pointer = Runtime_sel_getUid("webView:decidePolicyForNavigationAction:request:frame:decisionListener:")
var WebPolicyDelegateDecidePolicyForNewWindowAction unsafe.Pointer = Runtime_sel_getUid("webView:decidePolicyForNewWindowAction:request:newFrameName:decisionListener:")
var WebPolicyDelegateUnableToImplementPolicyWithError unsafe.Pointer = Runtime_sel_getUid("webView:unableToImplementPolicyWithError:frame:")
var WebPolicyDelegateDecidePolicyForMIMEType unsafe.Pointer = Runtime_sel_getUid("webView:decidePolicyForMIMEType:request:frame:decisionListener:")

type WebPolicyDelegate struct {
	NSObject
}

func (m WebPolicyDelegate) Init() {
	print("pol init")
	p := Runtime_objc_msgSend(m.Pointer, WebPolicyDelegateInit)
	m.Assert(p)
}

func WebPolicyDelegatePointer(p unsafe.Pointer) WebPolicyDelegate {
	return WebPolicyDelegate{NSObjectPointer(p)}
}
