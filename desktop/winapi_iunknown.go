// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var (
	IID_IUnknown = GUIDNew("00000000-0000-0000-C000-000000000046")
)

type IUnknown struct {
	RawVTable *interface{}
	refs      map[string]interface{} // interfaces map
}

type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func IUnknownNew() *IUnknown {
	m := &IUnknown{}
	m.RawVTable = (*interface{})(unsafe.Pointer(&IUnknownVtbl{}))
	m.New()
	return m
}

func (m *IUnknown) New() {
	m.VTable().QueryInterface = syscall.NewCallback(func(m *IUnknown, guid *GUID, ret PPVOID) HRESULT {
		if ref, ok := m.refs[guid.String()]; ok {
			ret[0] = Arg(ref)
			return S_OK
		} else {
			ret[0] = 0
			return E_NOINTERFACE
		}
	})
	m.VTable().AddRef = syscall.NewCallback(func(m *IUnknown) int {
		return 1
	})
	m.VTable().Release = syscall.NewCallback(func(m *IUnknown) int {
		return 1
	})
	m.refs = make(map[string]interface{})
	m.AddInterface(&IID_IUnknown, m)
}

func (v *IUnknown) VTable() *IUnknownVtbl {
	return (*IUnknownVtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IUnknown) AddInterface(guid *GUID, ref interface{}) {
	m.refs[guid.String()] = ref
}

func (m *IUnknown) QueryInterface(guid *GUID, ret interface{}) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().QueryInterface,
		3,
		Arg(m),
		Arg(guid),
		Arg(ret),
	))
}

func (m *IUnknown) AddRef() ULONG {
	return ULONGPtr(syscall.Syscall(
		m.VTable().AddRef,
		1,
		Arg(m),
		0,
		0,
	))
}

func (m *IUnknown) Release() ULONG {
	return ULONGPtr(syscall.Syscall(
		m.VTable().Release,
		1,
		Arg(m),
		0,
		0,
	))
}
