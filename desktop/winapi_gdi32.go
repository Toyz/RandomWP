// +build windows

package desktop

import (
	"syscall"
)

const (
	BI_RGB = 0

	DIB_RGB_COLORS = 0

	ETO_OPAQUE = 2
	SRCCOPY    = 0xCC0020
)

var Gdi32Dll = syscall.MustLoadDLL("Gdi32.dll")
var DeleteObject = Gdi32Dll.MustFindProc("DeleteObject")
var CreateCompatibleDC = Gdi32Dll.MustFindProc("CreateCompatibleDC")
var DeleteDC = Gdi32Dll.MustFindProc("DeleteDC")
var CreateDIBSection = Gdi32Dll.MustFindProc("CreateDIBSection")
var SelectObject = Gdi32Dll.MustFindProc("SelectObject")
var GetTextExtentPoint32 = Gdi32Dll.MustFindProc("GetTextExtentPoint32W")
var SetTextColor = Gdi32Dll.MustFindProc("SetTextColor")
var SetBkColor = Gdi32Dll.MustFindProc("SetBkColor")
var ExtTextOut = Gdi32Dll.MustFindProc("ExtTextOutW")
var CreateFontIndirect = Gdi32Dll.MustFindProc("CreateFontIndirectW")
