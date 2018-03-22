typedef int bool;

void gtk_init(int a, int v);

// window
void* gtk_scrolled_window_new(void*p, void*p2);
void* gtk_window_new(int t);
void gtk_window_set_default_size(void *p, int w, int h);
void gtk_window_set_position(void *p, int t);
void gtk_window_set_decorated(void *p, bool b);
void gtk_widget_grab_focus(void *p);
void gtk_widget_set_size_request(void *p, int w, int h);
void gtk_window_set_skip_taskbar_hint(void *p, bool b);
void gtk_window_set_skip_pager_hint(void *p, bool b);
void gtk_widget_set_events(void *p, int mask);
void gtk_window_set_resizable(void *p, bool b);
void gtk_window_resize(void *p, int w, int h);
void* gtk_viewport_new(void*p, void*p2);

void* gtk_fixed_new();
void gtk_fixed_put(void *p, void *p2, int x, int y);

// webkit
void* webkit_web_view_new();
void webkit_web_view_load_uri(void *p, const char* s);
void webkit_web_view_load_html_string(void *p, const char* s, const char* ss);
const char* webkit_network_request_get_uri(void *p);

int g_signal_connect_data(void* item, const char* action, void* callback, void* data, void* destroy_data, int connect_flags);
void g_object_ref(void* p);
void g_object_unref(void* p);
void gtk_widget_destroy(void* p);
int gtk_get_current_event_time();
void g_signal_emit_by_name(void*, const char*p);

// menus
void* gtk_menu_new();
void gtk_menu_shell_append(void* menu, void* item);
void* gtk_separator_menu_item_new();
void* gtk_menu_item_new();
void* gtk_menu_item_new_with_label(const char* s);
void* gtk_check_menu_item_new_with_label(const char* s);
const char* gtk_menu_item_get_label(void* item);
void gtk_menu_item_set_submenu(void* menu, void* item);
void gtk_menu_popup(void* m, void* parent, void* parentitem, void* func, void* data, int button, int time);
void gtk_widget_show(void* item);
void* gtk_hbox_new(bool homogeneous, int spacing);
void gtk_box_pack_start(void* box, void* item, bool expand, bool fill, int padding);
void gtk_box_pack_end(void* box, void* item, bool expand, bool fill, int padding);
void* gtk_label_new(const char* s);
void gtk_label_set_text(void* label, const char* s);
const char* gtk_label_get_text(void* label);
void gtk_container_add(void* container, void* widget);
void gtk_widget_show_all(void* container);
void gtk_widget_hide(void* p);
void* gtk_check_menu_item_new();
void gtk_check_menu_item_set_active(void* menu, bool b);
void gtk_widget_set_sensitive(void* item, bool b);

// status icon
void* gtk_status_icon_new_from_gicon(void* icon);
void gtk_status_icon_set_from_gicon(void* s, void* i);
void gtk_status_icon_set_visible(void* icon, bool b);
void* gtk_image_new();
void* gtk_image_new_from_gicon(void* g, int size);
void gtk_status_icon_set_title(void* icon, const char* title);
const char* gtk_status_icon_get_title(void* icon);
void gtk_status_icon_set_tooltip_text(void* icon, const char* title);
const char* gtk_status_icon_get_tooltip_text(void* icon);
void gtk_status_icon_position_menu(void*, void *x, void *y, void *push, void* data);

// GBytes
void* g_bytes_new(void* buf, int size);
void* g_bytes_icon_new(void* bytes);
void g_bytes_unref(void* b);

// loop
void* g_main_loop_new(void* context, void* is_running);
void g_main_loop_run(void* loop);
void g_main_loop_quit(void* loop);
void* g_main_loop_get_context(void* loop);

// threads
void gdk_threads_init();
void gdk_threads_enter();
void gdk_threads_leave();
void g_main_context_invoke(void* c, void* func, void* data);

