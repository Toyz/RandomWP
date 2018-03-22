// +build windows

package desktop

import (
	"syscall"
)

const (
	S_OK               = 0
	E_FAIL             = 0x80004005
	E_NOINTERFACE      = 0x80004002
	E_NOTIMPL          = 0x80004001
	S_FALSE            = 1
	DISP_E_UNKNOWNNAME = 0x80020006

	CLSCTX_INPROC         = (CLSCTX_INPROC_SERVER | CLSCTX_INPROC_HANDLER)
	CLSCTX_INPROC_SERVER  = 0x1
	CLSCTX_LOCAL_SERVER   = 0x4
	CLSCTX_INPROC_HANDLER = 0x2
	CLSCTX_REMOTE_SERVER  = 0x10

	OLEIVERB_INPLACEACTIVATE = -5
)

var Ole32Dll = syscall.MustLoadDLL("Ole32.dll")
var CoTaskMemFree = Ole32Dll.MustFindProc("CoTaskMemFree")
var OleInitialize = Ole32Dll.MustFindProc("OleInitialize")
var OleUninitialize = Ole32Dll.MustFindProc("OleUninitialize")
var CoCreateInstance = Ole32Dll.MustFindProc("CoCreateInstance")
