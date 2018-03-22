// +build windows

package desktop

type RECT struct {
	left   LONG
	top    LONG
	right  LONG
	bottom LONG
}

type LPRECT *RECT
type CRECT RECT
type LPCRECT *CRECT
