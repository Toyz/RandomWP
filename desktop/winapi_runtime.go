// +build windows

package desktop

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

func Arg(d interface{}) uintptr {
	switch d.(type) {
	case bool:
		return uintptr(Bool2Int[d.(bool)])
	}

	v := reflect.ValueOf(d)
	UIntPtr := reflect.TypeOf((uintptr)(0))

	if v.Type().ConvertibleTo(UIntPtr) {
		vv := v.Convert(UIntPtr)
		return vv.Interface().(uintptr)
	} else {
		return v.Pointer()
	}
}

func LOWORD(d DWORD) WORD {
	return WORD(d & 0xFFFF)
}

func HIWORD(d DWORD) WORD {
	return WORD((d >> 16) & 0xFFFF)
}

/*
https://msdn.microsoft.com/en-us/library/ms724832(VS.85).aspx
Windows 10					10.0*
Windows Server 2016			10.0*
Windows 8.1					6.3*
Windows Server 2012 R2		6.3*
Windows 8					6.2
Windows Server 2012			6.2
Windows 7					6.1
Windows Server 2008 R2		6.1
Windows Server 2008			6.0
Windows Vista				6.0
Windows Server 2003 R2		5.2
Windows Server 2003			5.2
Windows XP 64-Bit Edition	5.2
Windows XP					5.1
Windows 2000				5.0
*/
func GetVersion() (int, int, int) {
	v, err := syscall.GetVersion()
	if err != nil {
		panic(err)
	}
	return int(uint8(v)), int(uint8(v >> 8)), int(uint16(v >> 16))
}

func IsWindowsXP() bool {
	v1, v2, _ := GetVersion()
	return v1 == 5 && v2 == 1
}

func WString2String(p uintptr) string {
	var rr []uint16 = make([]uint16, 0, MAX_PATH)
	for p := uintptr(unsafe.Pointer(p)); ; p += 2 {
		u := *(*uint16)(unsafe.Pointer(p))
		if u == 0 {
			return string(utf16.Decode(rr))
		}
		rr = append(rr, u)
	}
	panic("No zero at end of the string")
}

func WArray2String(rr []uint16) string {
	return string(utf16.Decode(rr))
}

// copy last error from last syscall
var LastError uintptr

func GetLastErrorString() string {
	return HRESULT(LastError).String()
}

var NULL = Arg(0)
var TRUE = Arg(1)
var FALSE = Arg(0)

var Bool2Int = map[bool]int{
	true:  1,
	false: 0,
}

var Int2Bool = map[int]bool{
	1: true,
	0: false,
}

type WPARAM uintptr
type LPARAM uintptr
type PPVOID *[1]uintptr
type PVOID uintptr
type ULONG_PTR uintptr
type BYTE byte
type WORD uint16
type LPCTSTR uintptr
type LPTSTR uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type USHORT uint16
type LONG uint32
type HACCEL HANDLE
type OLECHAR WCHAR
type WCHAR TCHAR
type TCHAR uint16
type LPOLESTR uintptr
type LPCOLESTR uintptr
type HOLEMENU HANDLE

//
// compound types
//

type WString uintptr

func WStringPtr(r1, r2 uintptr, err error) WString {
	LastError = uintptr(err.(syscall.Errno))
	return WString(r1)
}

func WStringNew(s string) WString {
	u := utf16.Encode([]rune(s + "\x00"))

	size := len(u) * int(unsafe.Sizeof(u[0]))

	m := WStringPtr(GlobalAlloc.Call(Arg(GMEM_FIXED), Arg(size)))
	if m == 0 {
		panic(GetLastErrorString())
	}

	to := (*(*[1 << 30]byte)(unsafe.Pointer(m)))[:(size)]
	from := (*(*[1 << 30]byte)(unsafe.Pointer(&u[0])))[:(size)]

	for i := range to {
		to[i] = from[i]
	}

	return m
}

func (m WString) Size() int {
	return int(UINTPtr(lstrlen.Call(Arg(m))))
}

func (m WString) Close() {
	HRESULTPtr(GlobalFree.Call(Arg(m))).S_OK()
}

type HMENU uintptr

func HMENUPtr(r1, r2 uintptr, err error) HMENU {
	LastError = uintptr(err.(syscall.Errno))
	return HMENU(r1)
}

func (m HMENU) Close() {
	if !BOOLPtr(DestroyMenu.Call(Arg(m))).Bool() {
		panic(GetLastErrorString())
	}
}

type HRESULT uintptr

func HRESULTPtr(r1, r2 uintptr, err error) HRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return HRESULT(r1)
}

func (m HRESULT) S_OK() {
	if m != S_OK {
		panic(m.String())
	}
}

func (m HRESULT) String() string {
	msg := [1024]uint16{}
	FormatMessage.Call(Arg(FORMAT_MESSAGE_FROM_SYSTEM), NULL, Arg(m), NULL, Arg(&msg[0]), Arg(len(msg)), NULL)
	return fmt.Sprintf("HRESULT: 0x%08x [%s]", uintptr(m), strings.TrimSpace(WString2String(Arg(&msg[0]))))
}

type BOOL uint32

func BOOLPtr(r1, r2 uintptr, err error) BOOL {
	LastError = uintptr(err.(syscall.Errno))
	return BOOL(r1)
}

func (m BOOL) Bool() bool {
	return Int2Bool[int(m)]
}

type UINT uint32

func UINTPtr(r1, r2 uintptr, err error) UINT {
	LastError = uintptr(err.(syscall.Errno))
	return UINT(r1)
}

type ULONG uint32

func ULONGPtr(r1, r2 uintptr, err error) ULONG {
	LastError = uintptr(err.(syscall.Errno))
	return ULONG(r1)
}

type DWORD uint32

func DWORDPtr(r1, r2 uintptr, err error) DWORD {
	LastError = uintptr(err.(syscall.Errno))
	return DWORD(r1)
}

type HFONT uintptr

func HFONTPtr(r1, r2 uintptr, err error) HFONT {
	LastError = uintptr(err.(syscall.Errno))
	return HFONT(r1)
}

func (m HFONT) Close() {
	DeleteObject.Call(Arg(m))
}

type HWND uintptr

func HWNDPtr(r1, r2 uintptr, err error) HWND {
	LastError = uintptr(err.(syscall.Errno))
	return HWND(r1)
}

func (m HWND) Close() {
	if !BOOLPtr(DestroyWindow.Call(Arg(m))).Bool() {
		panic(GetLastErrorString())
	}
}

type WNDPROC uintptr

func WNDPROCNew(fn interface{}) WNDPROC {
	return WNDPROC(syscall.NewCallback(fn))
}

type HINSTANCE uintptr

func HINSTANCEPtr(r1, r2 uintptr, err error) HINSTANCE {
	LastError = uintptr(err.(syscall.Errno))
	return HINSTANCE(r1)
}

type HANDLE uintptr

func HANDLEPtr(r1, r2 uintptr, err error) HANDLE {
	LastError = uintptr(err.(syscall.Errno))
	return HANDLE(r1)
}

type COLORREF uintptr

func COLORREFPtr(r1, r2 uintptr, err error) COLORREF {
	LastError = uintptr(err.(syscall.Errno))
	return COLORREF(r1)
}

type LRESULT uintptr

func LRESULTPtr(r1, r2 uintptr, err error) LRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return LRESULT(r1)
}

type ATOM uintptr

func ATOMPtr(r1, r2 uintptr, err error) ATOM {
	LastError = uintptr(err.(syscall.Errno))
	return ATOM(r1)
}

type HDC uintptr

func HDCPtr(r1, r2 uintptr, err error) HDC {
	LastError = uintptr(err.(syscall.Errno))
	return HDC(r1)
}
