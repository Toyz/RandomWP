// +build windows

package desktop

import (
	"syscall"
)

const (
	MAX_PATH DWORD = 260

	// Local Settings\Application Data
	CSIDL_LOCAL_APPDATA = 0x001c

	// ~/My Documents
	CSIDL_PERSONAL = 0x005

	// ~/Desktop
	CSIDL_DESKTOPDIRECTORY = 0x10

	SHGFP_TYPE_CURRENT = 0
	SHGFP_TYPE_DEFAULT = 1
	S_FILE_NOT_FOUND   = 0x80070002

	NIM_ADD    DWORD = 0
	NIM_MODIFY DWORD = 1
	NIM_DELETE DWORD = 2

	SW_SHOWNORMAL = 1
)

var Shell32Dll = syscall.MustLoadDLL("Shell32.dll")
var SHGetKnownFolderPath, _ = Shell32Dll.FindProc("SHGetKnownFolderPath")
var SHGetFolderPath = Shell32Dll.MustFindProc("SHGetFolderPathW")
var Shell_NotifyIcon = Shell32Dll.MustFindProc("Shell_NotifyIconW")
var ShellExecute = Shell32Dll.MustFindProc("ShellExecuteW")
