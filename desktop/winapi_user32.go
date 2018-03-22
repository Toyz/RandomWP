// +build windows

package desktop

import (
	"syscall"
)

const (
	WS_OVERLAPPED           uint32 = 0
	WS_OVERLAPPEDWINDOW            = 0x00cf0000
	SPI_GETNONCLIENTMETRICS        = 0x0029
	COLOR_MENU                     = 4
	COLOR_MENUTEXT                 = 7
	COLOR_HIGHLIGHTTEXT            = 14
	COLOR_HIGHLIGHT                = 13
	COLOR_GRAYTEXT                 = 17
	WM_QUIT                        = 0x0012
	SW_HIDE                        = 0
	SW_SHOW                        = 5
	HWND_TOP                       = 0
	GWL_STYLE                      = -16
	WS_VISIBLE                     = 0x10000000
	WS_DLGFRAME                    = 0x00400000
	WS_BORDER                      = 0x00800000
	WS_EX_TOOLWINDOW               = 0x00000080
	GWL_EXSTYLE                    = -20
	WM_ACTIVATE                    = 0x0006
	WA_ACTIVE                      = 1
	WA_INACTIVE                    = 0
	HWND_TOPMOST                   = -1

	PM_NOREMOVE = 0
	PM_REMOVE   = 1

	WS_POPUP       uint32 = 0x80000000
	WS_SYSMENU     uint32 = 0x00080000
	WS_POPUPWINDOW uint32 = (WS_POPUP | WS_BORDER | WS_SYSMENU)

	WM_LBUTTONDOWN   DWORD = 513
	WM_NCCREATE            = 129
	WM_NCCALCSIZE          = 131
	WM_CREATE              = 1
	WM_SIZE                = 5
	WM_MOVE                = 3
	WM_USER                = 1024
	WM_LBUTTONUP           = 0x0202
	WM_LBUTTONDBLCLK       = 515
	WM_RBUTTONUP           = 517
	WM_CLOSE               = 0x0010
	WM_NULL                = 0x0000
	WM_COMMAND             = 0x0111
	WM_SHELLNOTIFY         = WM_USER + 1
	WM_MEASUREITEM         = 44
	WM_DRAWITEM            = 43
	WM_CANCELMODE          = 0x001F
	VK_ESCAPE              = 0x1B
	WM_KEYDOWN             = 0x0100
	WM_KEYUP               = 0x0101
	WM_DESTROY             = 0x0002

	WM_CONTEXTMENU = 0x007B

	MF_ENABLED    = 0
	MF_DISABLED   = 0x00000002
	MF_CHECKED    = 0x00000008
	MF_UNCHECKED  = 0
	MF_GRAYED     = 0x00000001
	MF_STRING     = 0x00000000
	MFT_OWNERDRAW = 256
	MF_SEPARATOR  = 0x00000800
	MF_POPUP      = 0x00000010

	TPM_RECURSE     = 0x0001
	TPM_RIGHTBUTTON = 0x0002

	SM_CYMENUCHECK = 72
	SM_CYMENU      = 15
)

var User32Dll = syscall.MustLoadDLL("User32.dll")
var CreatePopupMenu = User32Dll.MustFindProc("CreatePopupMenu")
var RegisterClassEx = User32Dll.MustFindProc("RegisterClassExW")
var UnregisterClass = User32Dll.MustFindProc("UnregisterClassW")
var CreateWindowEx = User32Dll.MustFindProc("CreateWindowExW")
var GetMessage = User32Dll.MustFindProc("GetMessageW")
var DispatchMessage = User32Dll.MustFindProc("DispatchMessageW")
var DefWindowProc = User32Dll.MustFindProc("DefWindowProcW")
var DestroyWindow = User32Dll.MustFindProc("DestroyWindow")
var CreateIconIndirect = User32Dll.MustFindProc("CreateIconIndirect")
var GetDC = User32Dll.MustFindProc("GetDC")
var ReleaseDC = User32Dll.MustFindProc("ReleaseDC")
var DestroyMenu = User32Dll.MustFindProc("DestroyMenu")
var GetWindowText = User32Dll.MustFindProc("GetWindowTextW")
var GetClassName = User32Dll.MustFindProc("GetClassNameW")
var SetWindowText = User32Dll.MustFindProc("SetWindowTextW")
var SetWindowPos = User32Dll.MustFindProc("SetWindowPos")
var ShowWindow = User32Dll.MustFindProc("ShowWindow")
var TranslateMessage = User32Dll.MustFindProc("TranslateMessage")
var RegisterWindowMessage = User32Dll.MustFindProc("RegisterWindowMessageW")
var PostMessage = User32Dll.MustFindProc("PostMessageW")
var AppendMenu = User32Dll.MustFindProc("AppendMenuW")
var GetMenuItemInfo = User32Dll.MustFindProc("GetMenuItemInfoW")
var SetMenuItemInfo = User32Dll.MustFindProc("SetMenuItemInfoW")
var GetSystemMetrics = User32Dll.MustFindProc("GetSystemMetrics")
var SetForegroundWindow = User32Dll.MustFindProc("SetForegroundWindow")
var GetCursorPos = User32Dll.MustFindProc("GetCursorPos")
var TrackPopupMenu = User32Dll.MustFindProc("TrackPopupMenu")
var FindWindowEx = User32Dll.MustFindProc("FindWindowExW")
var SendMessage = User32Dll.MustFindProc("SendMessageW")
var SystemParametersInfo = User32Dll.MustFindProc("SystemParametersInfoW")
var GetSysColor = User32Dll.MustFindProc("GetSysColor")
var SetWindowLong = User32Dll.MustFindProc("SetWindowLongW")
var EnumWindows = User32Dll.MustFindProc("EnumWindows")
var GetClientRect = User32Dll.MustFindProc("GetClientRect")
var UpdateWindow = User32Dll.MustFindProc("UpdateWindow")
var PeekMessage = User32Dll.MustFindProc("PeekMessageW")
