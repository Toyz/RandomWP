void* app_indicator_new(const char* id, const char* icon_name, int category);
void app_indicator_set_icon_theme_path(void* app, const char* path);
void app_indicator_set_menu(void* app, void* menu);
void app_indicator_set_icon_full(void* app, const char* name, const char* desc);
void app_indicator_set_title(void* app, const char* title);
void app_indicator_set_status(void* app, int status);

