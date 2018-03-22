// +build windows

package desktop

const (
	ODT_BUTTON   = 4
	ODT_COMBOBOX = 3
	ODT_LISTBOX  = 2
	ODT_LISTVIEW = 102
	ODT_MENU     = 1
	ODT_STATIC   = 5
	ODT_TAB      = 101

	ODS_SELECTED = 1
)

type DRAWITEMSTRUCT struct {
	CtlType    UINT
	CtlID      UINT
	itemID     UINT
	itemAction UINT
	itemState  UINT
	hwndItem   HWND
	hDC        HDC
	rcItem     RECT
	itemData   ULONG_PTR
}
