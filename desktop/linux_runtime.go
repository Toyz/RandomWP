// +build linux

package desktop

/*
#include <dlfcn.h>
#include <stdlib.h>
#cgo LDFLAGS:-ldl
*/
import "C"

import (
	"errors"
	"reflect"
	"unsafe"
)

const (
	RTLD_LAZY    = 0x0001
	RTLD_NOW     = 0x0002
	RTLD_GLOBAL  = 0x0100
	RTLD_LOCAL   = 0x0000
	RTLD_NOSHARE = 0x1000
	RTLD_EXE     = 0x2000
	RTLD_SCRIPT  = 0x4000

	RTLD_DEFAULT = 0
)

var NULL = Arg(0)

var Bool2Int = map[bool]int{
	true:  1,
	false: 0,
}

func Arg(d interface{}) unsafe.Pointer {
	switch d.(type) {
	case bool:
		return unsafe.Pointer(uintptr(Bool2Int[d.(bool)]))
	}

	v := reflect.ValueOf(d)
	UIntPtr := reflect.TypeOf((uintptr)(0))

	if v.Type().ConvertibleTo(UIntPtr) {
		vv := v.Convert(UIntPtr)
		return unsafe.Pointer(vv.Interface().(uintptr))
	} else {
		return unsafe.Pointer(v.Pointer())
	}
}

func dlopen(lib string, flags uint) (uintptr, error) {
	n := C.CString(lib)
	defer C.free(unsafe.Pointer(n))
	u := C.dlopen(n, (C.int)(flags))
	if u == nil {
		err := errors.New(C.GoString(C.dlerror()))
		return 0, err
	}
	return uintptr(u), nil
}

func dlmust(lib string) uintptr {
	u, err := dlopen(lib, RTLD_LAZY|RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	return u
}

func dlclose(p uintptr) {
	C.dlclose(unsafe.Pointer(p))
}
