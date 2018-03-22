// +build linux

package desktop

/*
#include <stdlib.h>

#include "linux_appindicator.h"

extern void* appindicator_fallback(void*);

#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
*/
import "C"

import (
	"unsafe"
)

const (
	APP_INDICATOR_CATEGORY_APPLICATION_STATUS = 0
	APP_INDICATOR_CATEGORY_COMMUNICATIONS     = 1
	APP_INDICATOR_CATEGORY_SYSTEM_SERVICES    = 2
	APP_INDICATOR_CATEGORY_HARDWARE           = 3
	APP_INDICATOR_CATEGORY_OTHER              = 4

	APP_INDICATOR_STATUS_PASSIVE   = 0
	APP_INDICATOR_STATUS_ACTIVE    = 1
	APP_INDICATOR_STATUS_ATTENTION = 2
)

type AppIndicator unsafe.Pointer

type GTypeInstanceStruct struct {
	g_class GPointer
}

type GObjectStruct struct {
	g_type_instance GTypeInstanceStruct
	ref_count       int
	qdata           GPointer
}

type AppIndicatorInstanceStruct struct {
	parent GObjectStruct
	priv   GPointer
}

type GTypeClassStruct struct {
	g_type int
}

type GObjectClassStruct struct {
	g_type_class                GTypeClassStruct
	construct_properties        GPointer
	constructor                 GPointer
	set_property                GPointer
	get_property                GPointer
	dispose                     GPointer
	finalize                    GPointer
	dispatch_properties_changed GPointer
	notify                      GPointer
	constructed                 GPointer
	flags                       int
	dummy1                      GPointer
	dummy2                      GPointer
	dummy3                      GPointer
	dummy4                      GPointer
	dummy5                      GPointer
	dummy6                      GPointer
}

type AppIndicatorClassStruct struct {
	parent_class GObjectClassStruct

	new_icon                   GPointer
	new_attention_icon         GPointer
	new_status                 GPointer
	new_icon_theme             GPointer
	new_label                  GPointer
	connection_changed         GPointer
	scroll_event               GPointer
	app_indicator_reserved_ats GPointer
	fallback                   GPointer
	unfallback                 GPointer
	app_indicator_reserved_1   GPointer
	app_indicator_reserved_2   GPointer
	app_indicator_reserved_3   GPointer
	app_indicator_reserved_4   GPointer
	app_indicator_reserved_5   GPointer
	app_indicator_reserved_6   GPointer
}

//
// AppIndicatorFallback
//

type AppIndicatorFallback func() GtkWidget

func (m *AppIndicatorFallback) Set(app AppIndicator) {
	appindicator_fallback_map[app] = m

	inst := (*AppIndicatorInstanceStruct)(app)
	class := (*AppIndicatorClassStruct)(inst.parent.g_type_instance.g_class)
	class.fallback = GPointer(C.appindicator_fallback)
}

func (m *AppIndicatorFallback) Close(app AppIndicator) {
	delete(appindicator_fallback_map, app)
}

var appindicator_fallback_map = make(map[AppIndicator]*AppIndicatorFallback)

//export appindicator_fallback
func appindicator_fallback(p unsafe.Pointer) unsafe.Pointer {
	fn := *appindicator_fallback_map[AppIndicator(p)]
	r := fn()
	return Arg(r)
}

//
// app indicator
//

func app_indicator_new(id string, icon_name string, category int) AppIndicator {
	n := C.CString(id)
	defer C.free(unsafe.Pointer(n))
	k := C.CString(icon_name)
	defer C.free(unsafe.Pointer(k))
	return AppIndicator(C.app_indicator_new(n, k, (C.int)(category)))
}

func app_indicator_set_icon_theme_path(app AppIndicator, path string) {
	n := C.CString(path)
	defer C.free(unsafe.Pointer(n))
	C.app_indicator_set_icon_theme_path(Arg(app), n)
}

func app_indicator_set_menu(app AppIndicator, menu GtkWidget) {
	C.app_indicator_set_menu(Arg(app), Arg(menu))
}

func app_indicator_set_icon_full(app AppIndicator, name string, desc string) {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	k := C.CString(desc)
	defer C.free(unsafe.Pointer(k))
	C.app_indicator_set_icon_full(Arg(app), n, k)
}

func app_indicator_set_title(app AppIndicator, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.app_indicator_set_title(Arg(app), n)
}

func app_indicator_set_status(app AppIndicator, status int) {
	C.app_indicator_set_status(Arg(app), (C.int)(status))
}
