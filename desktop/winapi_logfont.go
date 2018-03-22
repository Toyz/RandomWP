// +build windows

package desktop

const (
	LF_FACESIZE = 32
)

type LOGFONT struct {
	lfHeight         LONG
	lfWidth          LONG
	lfEscapement     LONG
	lfOrientation    LONG
	lfWeight         LONG
	lfItalic         BYTE
	lfUnderline      BYTE
	lfStrikeOut      BYTE
	lfCharSet        BYTE
	lfOutPrecision   BYTE
	lfClipPrecision  BYTE
	lfQuality        BYTE
	lfPitchAndFamily BYTE
	lfFaceName       [LF_FACESIZE]TCHAR
}
