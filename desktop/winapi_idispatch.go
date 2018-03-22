// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	IID_IDispatch = GUIDNew("{00020400-0000-0000-C000-000000000046}")
)

const (
	DISP_E_MEMBERNOTFOUND = 0x80020003
)

type DISPPARAMS struct {
	rgvarg            uintptr
	rgdispidNamedArgs uintptr
	cArgs             uint32
	cNamedArgs        uint32
}

type IDispatch struct {
	IUnknown
	refs map[string]int  // func names map
	ids  map[int]uintptr // func names map
}

type IDispatchVtbl struct {
	IUnknownVtbl
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

func IDispatchNew() *IDispatch {
	m := &IDispatch{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IDispatchVtbl{}))
	m.New()
	return m
}

func (m *IDispatch) New() {
	m.IUnknown.New()
	m.VTable().GetIDsOfNames = syscall.NewCallback(func(m *IDispatch, guid *GUID, names *[1 << 28]uintptr, count int, lcid uintptr, ret *[1 << 28]int) HRESULT {
		for i := 0; i < count; i++ {
			name := WString2String(names[i])
			if p, ok := m.refs[name]; ok {
				ret[i] = p
			} else {
				return DISP_E_UNKNOWNNAME
			}
		}
		return S_OK
	})
	m.VTable().GetTypeInfoCount = syscall.NewCallback(func(m *IDispatch) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().GetTypeInfo = syscall.NewCallback(func(m *IDispatch) HRESULT {
		return E_NOTIMPL
	})
	m.VTable().Invoke = syscall.NewCallback(func(m *IDispatch, id int, guid *GUID, lcid uintptr, flags int, params *DISPPARAMS, ret uintptr, err uintptr, index PPVOID) HRESULT {
		if p, ok := m.ids[id]; ok {
			argv := (*[1 << 28]uintptr)(unsafe.Pointer(params.rgvarg))
			switch params.cArgs {
			case 0:
				return HRESULTPtr(syscall.Syscall(
					p,
					1,
					Arg(m),
					0,
					0,
				))
			case 1:
				return HRESULTPtr(syscall.Syscall(
					p,
					2,
					Arg(m),
					Arg(argv[0]),
					0,
				))
			case 2:
				return HRESULTPtr(syscall.Syscall(
					p,
					3,
					Arg(m),
					Arg(argv[0]),
					Arg(argv[1]),
				))
			default:
				panic("more agruments")
			}
		} else {
			return DISP_E_MEMBERNOTFOUND
		}
	})
	m.refs = make(map[string]int)
	m.ids = make(map[int]uintptr)

	// IUnknown can't add him self
	m.AddMethod("QueryInterface", 0x60000000, m.IUnknown.VTable().QueryInterface)
	m.AddMethod("AddRef", 0x60000001, m.IUnknown.VTable().AddRef)
	m.AddMethod("Release", 0x60000002, m.IUnknown.VTable().Release)

	m.AddInterface(&IID_IDispatch, m)
	m.AddMethod("GetTypeInfoCount", 0x60010000, m.VTable().GetTypeInfoCount)
	m.AddMethod("GetTypeInfo", 0x60010001, m.VTable().GetTypeInfo)
	m.AddMethod("GetIDsOfNames", 0x60010002, m.VTable().GetIDsOfNames)
	m.AddMethod("Invoke", 0x60010003, m.VTable().Invoke)
}

func (v *IDispatch) VTable() *IDispatchVtbl {
	return (*IDispatchVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IDispatch) AddMethod(name string, id int, p uintptr) {
	m.refs[name] = id
	m.ids[id] = p
}

func (m *IDispatch) GetIDsOfNames(guid *GUID, ret interface{}) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().GetIDsOfNames,
		3,
		Arg(m),
		Arg(guid),
		Arg(ret),
	))
}

func (m *IDispatch) GetTypeInfoCount(ret *UINT) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().GetTypeInfoCount,
		2,
		Arg(m),
		Arg(ret),
		0,
	))
}
