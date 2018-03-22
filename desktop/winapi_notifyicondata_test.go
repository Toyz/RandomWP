// +build windows

package desktop

import (
	"strconv"
	"testing"
)

func TestNOTIFYICONDATA(t *testing.T) {
	n := NOTIFYICONDATANew()

	if strconv.IntSize == 32 {
		if n.cbSize != 956 {
			t.Error("wrong NOTIFYICONDATA size", n.cbSize)
		}
	}

	if strconv.IntSize == 64 {
		if n.cbSize != 976 {
			t.Error("wrong NOTIFYICONDATA size", n.cbSize)
		}
	}
}
