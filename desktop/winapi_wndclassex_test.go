// +build windows

package desktop

import (
	"strconv"
	"testing"
	"unsafe"
)

func TestWNDCLASSEX(t *testing.T) {
	m := &WNDCLASSEX{}
	m.cbSize = UINT(unsafe.Sizeof(*m))

	if strconv.IntSize == 32 {
		if m.cbSize != 48 {
			t.Error("wrong WNDCLASSEX size", m.cbSize)
		}
	}

	if strconv.IntSize == 64 {
		if m.cbSize != 80 {
			t.Error("wrong WNDCLASSEX size", m.cbSize)
		}
	}
}
