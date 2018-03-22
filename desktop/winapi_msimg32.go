// +build windows

package desktop

import (
	"syscall"
)

var Msimg32Dll = syscall.MustLoadDLL("Msimg32.dll")
var AlphaBlend = Msimg32Dll.MustFindProc("AlphaBlend")
