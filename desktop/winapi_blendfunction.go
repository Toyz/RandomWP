// +build windows

package desktop

const (
	AC_SRC_OVER             = 0x00
	AC_SRC_ALPHA            = 0x01
	AC_SRC_NO_PREMULT_ALPHA = 0x01
	AC_SRC_NO_ALPHA         = 0x02
)

type BLENDFUNCTION struct {
	BlendOp             BYTE
	BlendFlags          BYTE
	SourceConstantAlpha BYTE
	AlphaFormat         BYTE
}
