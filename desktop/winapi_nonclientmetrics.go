// +build windows

package desktop

// WINVER < 0x0600
type NONCLIENTMETRICS_XP struct {
	cbSize           UINT
	iBorderWidth     int32
	iScrollWidth     int32
	iScrollHeight    int32
	iCaptionWidth    int32
	iCaptionHeight   int32
	lfCaptionFont    LOGFONT
	iSmCaptionWidth  int32
	iSmCaptionHeight int32
	lfSmCaptionFont  LOGFONT
	iMenuWidth       int32
	iMenuHeight      int32
	lfMenuFont       LOGFONT
	lfStatusFont     LOGFONT
	lfMessageFont    LOGFONT
}

// Windows Vista+
type NONCLIENTMETRICS struct {
	NONCLIENTMETRICS_XP 
	iPaddedBorderWidth int32
}
