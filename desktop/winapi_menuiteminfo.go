// +build windows

package desktop

import ()

const (
	MFS_CHECKED   = 0x00000008
	MFS_DEFAULT   = 0x00001000
	MFS_DISABLED  = 0x00000003
	MFS_ENABLED   = 0x00000000
	MFS_GRAYED    = 0x00000003
	MFS_HILITE    = 0x00000080
	MFS_UNCHECKED = 0x00000000
	MFS_UNHILITE  = 0x00000000
	MIIM_DATA     = 0x00000020
)

type MENUITEMINFO struct {
	cbSize        UINT
	fMask         UINT
	fType         UINT
	fState        UINT
	wID           UINT
	hSubMenu      HMENU
	hbmpChecked   HBITMAP
	hbmpUnchecked HBITMAP
	dwItemData    ULONG_PTR
	dwTypeData    LPTSTR
	cch           UINT
	hbmpItem      HBITMAP
}
