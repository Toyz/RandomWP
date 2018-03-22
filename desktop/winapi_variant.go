// +build windows

package desktop

import (
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

const (
	VT_EMPTY            = 0
	VT_NULL             = 1
	VT_I2               = 2
	VT_I4               = 3
	VT_R4               = 4
	VT_R8               = 5
	VT_CY               = 6
	VT_DATE             = 7
	VT_BSTR             = 8
	VT_DISPATCH         = 9
	VT_ERROR            = 10
	VT_BOOL             = 11
	VT_VARIANT          = 12
	VT_UNKNOWN          = 13
	VT_DECIMAL          = 14
	VT_I1               = 16
	VT_UI1              = 17
	VT_UI2              = 18
	VT_UI4              = 19
	VT_I8               = 20
	VT_UI8              = 21
	VT_INT              = 22
	VT_UINT             = 23
	VT_VOID             = 24
	VT_HRESULT          = 25
	VT_PTR              = 26
	VT_SAFEARRAY        = 27
	VT_CARRAY           = 28
	VT_USERDEFINED      = 29
	VT_LPSTR            = 30
	VT_LPWSTR           = 31
	VT_RECORD           = 36
	VT_INT_PTR          = 37
	VT_UINT_PTR         = 38
	VT_FILETIME         = 64
	VT_BLOB             = 65
	VT_STREAM           = 66
	VT_STORAGE          = 67
	VT_STREAMED_OBJECT  = 68
	VT_STORED_OBJECT    = 69
	VT_BLOB_OBJECT      = 70
	VT_CF               = 71
	VT_CLSID            = 72
	VT_VERSIONED_STREAM = 73
	VT_BSTR_BLOB        = 0xfff
	VT_VECTOR           = 0x1000
	VT_ARRAY            = 0x2000
	VT_BYREF            = 0x4000
)

type VARIANT_BOOL uint16

var Bool2VariantBool = map[bool]VARIANT_BOOL{
	true:  0xFFFF,
	false: 0,
}

var VariantBool2Bool = map[VARIANT_BOOL]bool{
	0xFFFF: true,
	0:      false,
}

type VARTYPE WORD

type VARIANT struct {
	VT         VARTYPE
	wReserved1 WORD
	wReserved2 WORD
	wReserved3 WORD
	Val        uintptr
	pRecInfo   uintptr
}

func VARIANTNew(v interface{}) VARIANT {
	switch vv := v.(type) {
	case bool:
		return VARIANT{VT: VT_BOOL, Val: Arg(Bool2VariantBool[vv])}
	case *bool:
		return VARIANT{VT: VT_BOOL | VT_BYREF, Val: Arg(v)}
	case uint8:
		return VARIANT{VT: VT_I1, Val: Arg(v)}
	case *uint8:
		return VARIANT{VT: VT_I1 | VT_BYREF, Val: Arg(v)}
	case int8:
		return VARIANT{VT: VT_I1, Val: Arg(v)}
	case *int8:
		return VARIANT{VT: VT_I1 | VT_BYREF, Val: Arg(v)}
	case int16:
		return VARIANT{VT: VT_I2, Val: Arg(v)}
	case *int16:
		return VARIANT{VT: VT_I2 | VT_BYREF, Val: Arg(v)}
	case uint16:
		return VARIANT{VT: VT_UI2, Val: Arg(v)}
	case *uint16:
		return VARIANT{VT: VT_UI2 | VT_BYREF, Val: Arg(v)}
	case int32:
		return VARIANT{VT: VT_I4, Val: Arg(v)}
	case *int32:
		return VARIANT{VT: VT_I4 | VT_BYREF, Val: Arg(v)}
	case uint32:
		return VARIANT{VT: VT_UI4, Val: Arg(v)}
	case *uint32:
		return VARIANT{VT: VT_UI4 | VT_BYREF, Val: Arg(v)}
	case int64:
		return VARIANT{VT: VT_I8, Val: Arg(v)}
	case *int64:
		return VARIANT{VT: VT_I8 | VT_BYREF, Val: Arg(v)}
	case uint64:
		return VARIANT{VT: VT_UI8, Val: Arg(v)}
	case *uint64:
		return VARIANT{VT: VT_UI8 | VT_BYREF, Val: Arg(v)}
	case int:
		return VARIANT{VT: VT_I4, Val: Arg(v)}
	case *int:
		return VARIANT{VT: VT_I4 | VT_BYREF, Val: Arg(v)}
	case uint:
		return VARIANT{VT: VT_UI4, Val: Arg(v)}
	case *uint:
		return VARIANT{VT: VT_UI4 | VT_BYREF, Val: Arg(v)}
	case float32:
		return VARIANT{VT: VT_R4, Val: Arg(v)}
	case *float32:
		return VARIANT{VT: VT_R4 | VT_BYREF, Val: Arg(v)}
	case float64:
		return VARIANT{VT: VT_R8, Val: Arg(v)}
	case *float64:
		return VARIANT{VT: VT_R8 | VT_BYREF, Val: Arg(v)}
	case string:
		return VARIANT{VT: VT_BSTR, Val: Arg(SysAllocString(v.(string)))}
	case time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		return VARIANT{VT: VT_BSTR, Val: Arg(SysAllocString(s))}
	case *time.Time:
		s := vv.Format("2006-01-02 15:04:05")
		return VARIANT{VT: VT_BSTR, Val: Arg(SysAllocString(s))}
	case *IUnknown:
		return VARIANT{VT: VT_UNKNOWN, Val: Arg(v)}
	case **IUnknown:
		return VARIANT{VT: VT_UNKNOWN | VT_BYREF, Val: Arg(v)}
	case *IDispatch:
		return VARIANT{VT: VT_DISPATCH, Val: Arg(v)}
	case **IDispatch:
		return VARIANT{VT: VT_DISPATCH | VT_BYREF, Val: Arg(v)}
	case nil:
		return VARIANT{VT: VT_NULL, Val: Arg(0)}
	case *VARIANT:
		return VARIANT{VT: VT_VARIANT | VT_BYREF, Val: Arg(v)}
	default:
		panic("unknown type")
	}
}

func SysAllocString(v string) uintptr {
	utf16 := utf16.Encode([]rune(v + "\x00"))
	ptr := &utf16[0]
	w := WStringPtr(SysAllocStringLen.Call(Arg(ptr), Arg(len(utf16)-1)))
	if w == 0 {
		panic(GetLastErrorString())
	}
	return Arg(w)
}

func (v VARIANT) Value() interface{} {
	switch v.VT {
	case VT_I1:
		return int8(v.Val)
	case VT_UI1:
		return uint8(v.Val)
	case VT_I2:
		return int16(v.Val)
	case VT_UI2:
		return uint16(v.Val)
	case VT_I4:
		return int32(v.Val)
	case VT_UI4:
		return uint32(v.Val)
	case VT_I8:
		return int64(v.Val)
	case VT_UI8:
		return uint64(v.Val)
	case VT_INT:
		return int(v.Val)
	case VT_UINT:
		return uint(v.Val)
	case VT_INT_PTR:
		return uintptr(v.Val) // TODO
	case VT_UINT_PTR:
		return uintptr(v.Val)
	case VT_R4:
		return *(*float32)(unsafe.Pointer(&v.Val))
	case VT_R8:
		return *(*float64)(unsafe.Pointer(&v.Val))
	case VT_BSTR:
		return WString2String(v.Val)
	case VT_DATE:
		// VT_DATE type will either return float64 or time.Time.
		d := float64(v.Val)
		var st syscall.Systemtime
		hr := HRESULTPtr(VariantTimeToSystemTime.Call(Arg(&d), Arg(&st)))
		if hr != S_OK {
			return time.Date(int(st.Year), time.Month(st.Month), int(st.Day), int(st.Hour), int(st.Minute), int(st.Second), int(st.Milliseconds/1000), nil)
		}
		return d
	case VT_UNKNOWN:
		return (*IUnknown)(unsafe.Pointer(v.Val))
	case VT_DISPATCH:
		return (*IDispatch)(unsafe.Pointer(v.Val))
	case VT_BOOL:
		return v.Val != 0
	}
	return nil
}

func (m VARIANT) Clear() {
	HRESULTPtr(VariantClear.Call(Arg(&m))).S_OK()
}

//
// SafeArray

type SAFEARRAYBOUND struct {
	cElements ULONG
	lLbound   LONG
}

type SAFEARRAY struct {
	cDims      USHORT
	fFeatures  USHORT
	cbElements ULONG
	cLocks     ULONG
	pvData     PVOID
	rgsabound  [1]SAFEARRAYBOUND
}

func SAFEARRAYPtr(r1, r2 uintptr, err error) *SAFEARRAY {
	LastError = uintptr(err.(syscall.Errno))
	return (*SAFEARRAY)(unsafe.Pointer(r1))
}

func SAFEARRAYString(str string) *SAFEARRAY {
	var ArrayBound = SAFEARRAYBOUND{1, 0}
	m := SAFEARRAYPtr(SafeArrayCreate.Call(Arg(VT_VARIANT), Arg(1), Arg(&ArrayBound)))
	var v *VARIANT
	HRESULTPtr(SafeArrayAccessData.Call(Arg(m), Arg(&v))).S_OK()
	h := VARIANTNew(str)
	*v = h
	HRESULTPtr(SafeArrayUnaccessData.Call(Arg(m))).S_OK()
	return m
}

func (m *SAFEARRAY) Destory() {
	HRESULTPtr(SafeArrayDestroy.Call(Arg(m))).S_OK()
}
