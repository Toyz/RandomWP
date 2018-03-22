// +build windows

package desktop

import (
	"syscall"
)

const (
	FORMAT_MESSAGE_FROM_SYSTEM = 0x00001000
	GMEM_FIXED                 = 0
)

var Kernel32Dll = syscall.MustLoadDLL("Kernel32.dll")
var GetLastError = Kernel32Dll.MustFindProc("GetLastError")
var FormatMessage = Kernel32Dll.MustFindProc("FormatMessageW")
var GetModuleHandle = Kernel32Dll.MustFindProc("GetModuleHandleW")
var GlobalAlloc = Kernel32Dll.MustFindProc("GlobalAlloc")
var GlobalFree = Kernel32Dll.MustFindProc("GlobalFree")
var lstrlen = Kernel32Dll.MustFindProc("lstrlenW")
var GetCurrentThread = Kernel32Dll.MustFindProc("GetCurrentThread")
var GetCurrentThreadId = Kernel32Dll.MustFindProc("GetCurrentThreadId")
