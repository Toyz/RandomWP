// +build windows

package desktop

import (
	"syscall"
)

const ()

var OleAut32Dll = syscall.MustLoadDLL("OleAut32.dll")
var VariantTimeToSystemTime = OleAut32Dll.MustFindProc("VariantTimeToSystemTime")
var VariantClear = OleAut32Dll.MustFindProc("VariantClear")
var VariantInit = OleAut32Dll.MustFindProc("VariantInit")
var SysAllocStringByteLen = OleAut32Dll.MustFindProc("SysAllocStringByteLen")
var SysFreeString = OleAut32Dll.MustFindProc("SysFreeString")
var SysAllocStringLen = OleAut32Dll.MustFindProc("SysAllocStringLen")
var SafeArrayCreate = OleAut32Dll.MustFindProc("SafeArrayCreate")
var SafeArrayAccessData = OleAut32Dll.MustFindProc("SafeArrayAccessData")
var SafeArrayDestroy = OleAut32Dll.MustFindProc("SafeArrayDestroy")
var SafeArrayUnaccessData = OleAut32Dll.MustFindProc("SafeArrayUnaccessData")
