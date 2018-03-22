// +build linux

package desktop

/*
#include <stdlib.h>
#include <stdio.h>

#include "linux_gtk.h"

extern void* gsourcefunc(void*);
extern void signal_activate(void*, void*);
extern void signal_focus_out(void*, void*, void*);
extern void signal_popup_menu(void*, void*, void*, void*);
extern int signal_navigation_policy_decision_requested(void*, void*, void*, void*, void*, void*);

#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
*/
import "C"

import (
	"sync"
	"unsafe"
)

const (
	GTK_ORIENTATION_HORIZONTAL  = 0
	GTK_ORIENTATION_VERTICAL    = 1
	GTK_ICON_SIZE_INVALID       = 0
	GTK_ICON_SIZE_MENU          = 1
	GTK_ICON_SIZE_SMALL_TOOLBAR = 2
	GTK_ICON_SIZE_LARGE_TOOLBAR = 3
	GTK_ICON_SIZE_BUTTON        = 4
	GTK_ICON_SIZE_DND           = 5
	GTK_ICON_SIZE_DIALOG        = 6

	GTK_WINDOW_TOPLEVEL = 0
	GTK_WINDOW_POPUP    = 1

	GTK_WIN_POS_NONE   = 0
	GTK_WIN_POS_CENTER = 1
	GTK_WIN_POS_MOUSE  = 2

	GDK_HINT_POS      = 0
	GDK_HINT_MIN_SIZE = 1
	GDK_HINT_MAX_SIZE = 2

	GDK_FOCUS_CHANGE_MASK = 1 << 14
)

type GObject unsafe.Pointer
type GtkWidget unsafe.Pointer
type GMainLoop unsafe.Pointer
type GMainContext unsafe.Pointer
type GPointer unsafe.Pointer
type GtkStatusIcon unsafe.Pointer
type GIcon unsafe.Pointer
type GBytes unsafe.Pointer

var gsource_mu sync.Mutex
var gsource_index int = 0
var gsource_map = make(map[int]*GSourceFunc)

type GSourceFunc struct {
	i int
	f interface{}
}

func GSourceFuncNew(f interface{}) *GSourceFunc {
	gsource_mu.Lock()
	defer gsource_mu.Unlock()
	m := &GSourceFunc{}
	gsource_index++
	for gsource_map[gsource_index] != nil {
		gsource_index++
	}
	m.i = gsource_index
	m.f = f
	gsource_map[gsource_index] = m
	return m
}

func (m *GSourceFunc) Arg() unsafe.Pointer {
	return Arg(m.i)
}

func (m *GSourceFunc) Close() {
	gsource_mu.Lock()
	defer gsource_mu.Unlock()
	delete(gsource_map, m.i)
}

func gsource_lookup(p unsafe.Pointer) interface{} {
	gsource_mu.Lock()
	defer gsource_mu.Unlock()
	i := int(uintptr(p))
	return gsource_map[i].f
}

//export gsourcefunc
func gsourcefunc(p unsafe.Pointer) unsafe.Pointer {
	f := (gsource_lookup(p)).(func())
	f()
	return Arg(false)
}

//export signal_activate
func signal_activate(p, p1 unsafe.Pointer) {
	f := (gsource_lookup(p1)).(func())
	f()
}

//export signal_focus_out
func signal_focus_out(p, p1, p2 unsafe.Pointer) {
	f := (gsource_lookup(p2)).(func())
	f()
}

//export signal_navigation_policy_decision_requested
func signal_navigation_policy_decision_requested(p, p1, p2, p3, p4, p5 unsafe.Pointer) C.int {
	f := (gsource_lookup(p5)).(func(p0, p1, p2, p3, p4, p5 unsafe.Pointer) bool)
	return C.int(Bool2Int[f(p, p1, p2, p3, p4, p5)])
}

//export signal_popup_menu
func signal_popup_menu(p, p1, p2, p3 unsafe.Pointer) {
	f := (gsource_lookup(p3)).(func())
	f()
}

func gtk_init() {
	C.gtk_init(0, 0)
}

func g_signal_connect_navigation_policy_decision_requested(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("new-window-policy-decision-requested")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_navigation_policy_decision_requested), fn.Arg(), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_connect_focus_out(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("focus-out-event")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_focus_out), fn.Arg(), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_connect_activate(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("activate")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_activate), fn.Arg(), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_connect_popup(item GtkWidget, fn *GSourceFunc) {
	n := C.CString("popup-menu")
	defer C.free(unsafe.Pointer(n))
	i := C.g_signal_connect_data(Arg(item), n, Arg(C.signal_popup_menu), fn.Arg(), NULL, C.int(0))
	if i <= 0 {
		panic("unable to connect")
	}
}

func g_signal_emit_by_name(item GtkWidget, action string) {
	n := C.CString(action)
	defer C.free(unsafe.Pointer(n))
	C.g_signal_emit_by_name(Arg(item), n)
}

func g_object_ref(p GObject) {
	C.g_object_ref(Arg(p))
}

func g_object_unref(p GObject) {
	C.g_object_unref(Arg(p))
}

func gtk_widget_destroy(p GtkWidget) {
	C.gtk_widget_destroy(Arg(p))
}

func gtk_get_current_event_time() int {
	return int(C.gtk_get_current_event_time())
}

func gtk_widget_set_events(p GtkWidget, m int) {
	C.gtk_widget_set_events(Arg(p), C.int(m))
}

func gtk_window_set_skip_taskbar_hint(p GtkWidget, b bool) {
	C.gtk_window_set_skip_taskbar_hint(Arg(p), C.bool(Bool2Int[b]))
}

func gtk_window_set_skip_pager_hint(p GtkWidget, b bool) {
	C.gtk_window_set_skip_pager_hint(Arg(p), C.bool(Bool2Int[b]))
}

func gtk_window_set_resizable(p GtkWidget, b bool) {
	C.gtk_window_set_resizable(Arg(p), C.bool(Bool2Int[b]))
}

func gtk_window_resize(p GtkWidget, w int, h int) {
	C.gtk_window_resize(Arg(p), C.int(w), C.int(h))
}

func webkit_web_view_new() GtkWidget {
	return GtkWidget(C.webkit_web_view_new())
}

func webkit_web_view_load_uri(p GtkWidget, s string) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	C.webkit_web_view_load_uri(Arg(p), n)
}

func webkit_web_view_load_html_string(p GtkWidget, s string, ss string) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	nn := C.CString(ss)
	defer C.free(unsafe.Pointer(nn))
	C.webkit_web_view_load_html_string(Arg(p), n, nn)
}

func webkit_network_request_get_uri(p unsafe.Pointer) string {
	return C.GoString(C.webkit_network_request_get_uri(Arg(p)))
}

func gtk_widget_grab_focus(p GtkWidget) {
	C.gtk_widget_grab_focus(Arg(p))
}

func gtk_fixed_new() GtkWidget {
	return GtkWidget(C.gtk_fixed_new())
}

func gtk_fixed_put(p GtkWidget, p2 GtkWidget, x int, y int) {
	C.gtk_fixed_put(Arg(p), Arg(p2), C.int(x), C.int(y))
}

func gtk_widget_set_size_request(p GtkWidget, w int, h int) {
	C.gtk_widget_set_size_request(Arg(p), C.int(w), C.int(h))
}

func gtk_scrolled_window_new() GtkWidget {
	return GtkWidget(C.gtk_scrolled_window_new(Arg(0), Arg(0)))
}

func gtk_viewport_new() GtkWidget {
	return GtkWidget(C.gtk_viewport_new(Arg(0), Arg(0)))
}

func gtk_window_new(t int) GtkWidget {
	return GtkWidget(C.gtk_window_new(C.int(t)))
}

func gtk_window_set_default_size(p GtkWidget, w int, h int) {
	C.gtk_window_set_default_size(Arg(p), C.int(w), C.int(h))
}

func gtk_window_set_position(p GtkWidget, t int) {
	C.gtk_window_set_position(Arg(p), C.int(t))
}

func gtk_window_set_decorated(p GtkWidget, b bool) {
	C.gtk_window_set_decorated(Arg(p), C.bool(Bool2Int[b]))
}

func gtk_menu_new() GtkWidget {
	return GtkWidget(C.gtk_menu_new())
}

func gtk_menu_shell_append(menu GtkWidget, item GtkWidget) {
	C.gtk_menu_shell_append(Arg(menu), Arg(item))
}

func gtk_separator_menu_item_new() GtkWidget {
	return GtkWidget(C.gtk_separator_menu_item_new())
}

func gtk_menu_item_new() GtkWidget {
	return GtkWidget(C.gtk_menu_item_new())
}

func gtk_menu_item_new_with_label(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_menu_item_new_with_label(n))
}

func gtk_check_menu_item_new_with_label(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_check_menu_item_new_with_label(n))
}

func gtk_menu_item_get_label(item GtkWidget) string {
	return C.GoString(C.gtk_menu_item_get_label(Arg(item)))
}

func gtk_menu_item_set_submenu(menu GtkWidget, item GtkWidget) {
	C.gtk_menu_item_set_submenu(Arg(menu), Arg(item))
}

func gtk_menu_popup(m GtkWidget, parent GtkWidget, parentitem GtkWidget, fn GPointer, data GPointer, button int, time int) {
	C.gtk_menu_popup(Arg(m), Arg(parent), Arg(parentitem), Arg(fn), Arg(data), C.int(button), C.int(time))
}

var gtk_status_icon_position_menu = GPointer(C.gtk_status_icon_position_menu)

func gtk_widget_show(item GtkWidget) {
	C.gtk_widget_show(Arg(item))
}

func gtk_hbox_new(homogeneous bool, spacing int) GtkWidget {
	return GtkWidget(C.gtk_hbox_new(C.bool(Bool2Int[homogeneous]), C.int(spacing)))
}

func gtk_box_pack_start(box GtkWidget, item GtkWidget, expand bool, fill bool, padding int) {
	C.gtk_box_pack_start(Arg(box), Arg(item), C.bool(Bool2Int[expand]), C.bool(Bool2Int[fill]), C.int(padding))
}

func gtk_box_pack_end(box GtkWidget, item GtkWidget, expand bool, fill bool, padding int) {
	C.gtk_box_pack_end(Arg(box), Arg(item), C.bool(Bool2Int[expand]), C.bool(Bool2Int[fill]), C.int(padding))
}

func gtk_label_new(s string) GtkWidget {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	return GtkWidget(C.gtk_label_new(n))
}

func gtk_label_set_text(label GtkWidget, s string) {
	n := C.CString(s)
	defer C.free(unsafe.Pointer(n))
	C.gtk_label_set_text(Arg(label), n)
}

func gtk_label_get_text(label GtkWidget) string {
	return C.GoString(C.gtk_label_get_text(Arg(label)))
}

func gtk_container_add(container GtkWidget, widget GtkWidget) {
	C.gtk_container_add(Arg(container), Arg(widget))
}

func gtk_widget_show_all(container GtkWidget) {
	C.gtk_widget_show_all(Arg(container))
}

func gtk_widget_hide(p GtkWidget) {
	C.gtk_widget_hide(Arg(p))
}

func gtk_check_menu_item_new() GtkWidget {
	return GtkWidget(GtkWidget(C.gtk_check_menu_item_new()))
}

func gtk_check_menu_item_set_active(menu GtkWidget, b bool) {
	C.gtk_check_menu_item_set_active(Arg(menu), C.bool(Bool2Int[b]))
}

func gtk_widget_set_sensitive(item GtkWidget, b bool) {
	C.gtk_widget_set_sensitive(Arg(item), C.bool(Bool2Int[b]))
}

func gtk_status_icon_new_from_gicon(icon GIcon) GtkWidget {
	return GtkWidget(C.gtk_status_icon_new_from_gicon(Arg(icon)))
}

func gtk_status_icon_set_from_gicon(s GtkWidget, i GIcon) {
	C.gtk_status_icon_set_from_gicon(Arg(s), Arg(i))
}

func gtk_status_icon_set_visible(icon GtkWidget, b bool) {
	C.gtk_status_icon_set_visible(Arg(icon), C.bool(Bool2Int[b]))
}

func gtk_image_new() GtkWidget {
	return GtkWidget(C.gtk_image_new())
}

func gtk_image_new_from_gicon(g GIcon, size int) GtkWidget {
	return GtkWidget(C.gtk_image_new_from_gicon(Arg(g), C.int(size)))
}

func gtk_status_icon_set_title(icon GtkWidget, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.gtk_status_icon_set_title(Arg(icon), n)
}

func gtk_status_icon_get_title(icon GtkWidget) string {
	return C.GoString(C.gtk_status_icon_get_title(Arg(icon)))
}

func gtk_status_icon_set_tooltip_text(icon GtkWidget, title string) {
	n := C.CString(title)
	defer C.free(unsafe.Pointer(n))
	C.gtk_status_icon_set_tooltip_text(Arg(icon), n)
}

func gtk_status_icon_get_tooltip_text(icon GtkWidget) string {
	return C.GoString(C.gtk_status_icon_get_tooltip_text(Arg(icon)))
}

func g_bytes_new(buf []byte, size int) GBytes {
	return GBytes(C.g_bytes_new(Arg(buf), C.int(size)))
}

func g_bytes_icon_new(bytes GBytes) GIcon {
	return GIcon(C.g_bytes_icon_new(Arg(bytes)))
}

func g_bytes_unref(b GBytes) {
	C.g_bytes_unref(Arg(b))
}

func g_main_loop_new(context GMainContext, is_running bool) GMainLoop {
	return GMainLoop(C.g_main_loop_new(Arg(context), Arg(Bool2Int[is_running])))
}

func g_main_loop_run(loop GMainLoop) {
	C.g_main_loop_run(Arg(loop))
}

func g_main_loop_quit(loop GMainLoop) {
	C.g_main_loop_quit(Arg(loop))
}

func g_main_loop_get_context(loop GMainLoop) GMainContext {
	return GMainContext(C.g_main_loop_get_context(Arg(loop)))
}

func g_main_context_invoke(c GMainContext, fn *GSourceFunc) {
	C.g_main_context_invoke(Arg(c), Arg(C.gsourcefunc), fn.Arg())
}
