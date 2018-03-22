// +build darwin

package desktop

import (
	"unsafe"
)

var WebPolicyDecisionListenerClass unsafe.Pointer = Runtime_objc_lookUpClass("WebPolicyDecisionListener")
var WebPolicyDecisionListenerUse unsafe.Pointer = Runtime_sel_getUid("use")
var WebPolicyDecisionListenerDownload unsafe.Pointer = Runtime_sel_getUid("download")
var WebPolicyDecisionListenerIgnore unsafe.Pointer = Runtime_sel_getUid("ignore")

type WebPolicyDecisionListener struct {
	NSObject
}

func WebPolicyDecisionListenerPointer(p unsafe.Pointer) WebPolicyDecisionListener {
	return WebPolicyDecisionListener{NSObjectPointer(p)}
}

func (m WebPolicyDecisionListener) Use() {
	Runtime_objc_msgSend(m.Pointer, WebPolicyDecisionListenerUse)
}

func (m WebPolicyDecisionListener) Download() {
	Runtime_objc_msgSend(m.Pointer, WebPolicyDecisionListenerDownload)
}

func (m WebPolicyDecisionListener) Ignore() {
	Runtime_objc_msgSend(m.Pointer, WebPolicyDecisionListenerIgnore)
}
