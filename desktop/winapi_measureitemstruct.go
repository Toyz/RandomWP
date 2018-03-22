// +build windows

package desktop

import (
	"unsafe"
)

type MEASUREITEMSTRUCT struct {
	CtlType    UINT
	CtlID      UINT
	itemID     UINT
	itemWidth  UINT
	itemHeight UINT
	itemData   ULONG_PTR
}

func MEASUREITEMSTRUCTPtr(p uintptr) *MEASUREITEMSTRUCT {
	return (*MEASUREITEMSTRUCT)(unsafe.Pointer(p))
}
