package desktop

import (
	"encoding/base64"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

//
// Desktop Folders
//

// Config folder
//
//  - osx: /Users/user/Library/Application Support
//  - windows: C:\Users\user\AppData\Local
//  - linux: /home/user/.config
func GetAppDataFolder() string {
	return getAppDataFolder()
}

// Home folder
//
//  - osx: /Users/user
//  - windows: C:\Users\user
//  - linux: /home/user
func GetHomeFolder() string {
	return getHomeFolder()
}

// Documents folder
//
//  - osx: /Users/user/Documents
//  - windows: C:\Users\user\Documents
//  - linux: /home/user/Documents
func GetDocumentsFolder() string {
	return getDocumentsFolder()
}

// Downloads folder
//
//  - osx: /Users/user/Downloads
//  - windows: C:\Users\user\Downloads
//  - linux: /home/user/Desktop
func GetDownloadsFolder() string {
	return getDownloadsFolder()
}

// Desktop folder
//
//  - osx: /Users/user/Desktop
//  - windows: C:\Users\user\Desktop
//  - linux: /home/user/Desktop
func GetDesktopFolder() string {
	return getDesktopFolder()
}

//
// Main function
//
// Need to keep messages loop running. Have to be run on main thread.
// All GUI applications need that. So if you plan to use SysTray call
// this function from main function.
//

func Main() {
	desktopMain()
}

//
// Browser Functions
//

func BrowserOpenURI(s string) {
	browserOpenURI(s)
}

//
// SysTrayIcon / NSStatusBar / Notification Area
//

const (
	MenuItem      = 1
	MenuSeparator = 2
	MenuCheckBox  = 3
)

type MenuAction func(*Menu)

type Menu struct {
	Menu    []Menu
	Action  MenuAction
	State   bool
	Type    int
	Enabled bool
	Name    string
	Icon    image.Image
}

// true url handled
type WebPopupHandler func(url string) bool

type WebPopup struct {
	Width   int             // window width
	Height  int             // window height
	Html    string          // html menu
	Url     string          // url loaded menu
	Handler WebPopupHandler // web mouse clicks handler
}

func (m WebPopup) Size() (int, int) {
	w := m.Width
	if w == 0 {
		w = 400
	}
	h := m.Height
	if h == 0 {
		h = 500
	}
	return w, h
}

type DesktopSysTray struct {
	Listeners map[DesktopSysTrayListener]bool
	Title     string
	Menu      []Menu
	WebPopup  *WebPopup

	// os specific structs
	os interface{}
}

type DesktopSysTrayListener interface {
	MouseLeftClick()

	MouseLeftDoubleClick()

	// We do not handle right clicks, because:
	//
	// 1) Icon is binded to context menu anyway.
	//
	// 2) On Windows if you call showContextMenu from java thread, HMENU bugged
	// and you can't use it.
	//
	// 3) Mac OSX does not support showing context menu programmatically.
}

func DesktopSysTrayNew() *DesktopSysTray {
	m := desktopSysTrayNew()
	m.Listeners = make(map[DesktopSysTrayListener]bool)
	return m
}

func (m *DesktopSysTray) AddListener(l DesktopSysTrayListener) {
	m.Listeners[l] = true
}

func (m *DesktopSysTray) RemoveListener(l DesktopSysTrayListener) {
	delete(m.Listeners, l)
}

func (m *DesktopSysTray) SetIcon(icon image.Image) {
	m.setIcon(icon)
}

func (m *DesktopSysTray) SetTitle(title string) {
	m.Title = title
}

func (m *DesktopSysTray) Show() {
	m.show()
}

func (m *DesktopSysTray) Update() {
	m.update()
}

func (m *DesktopSysTray) Hide() {
	m.hide()
}

func (m *DesktopSysTray) SetMenu(menu []Menu) {
	m.Menu = menu
}

func (m *DesktopSysTray) SetWebPopup(w WebPopup) {
	m.WebPopup = &w
}

func (m *DesktopSysTray) Close() {
	m.close()
}

//
// funcs
//

func DecodeImageString(s string) image.Image {
	i, _, err := image.Decode(base64.NewDecoder(base64.StdEncoding, strings.NewReader(s)))
	if err != nil {
		panic(err)
	}
	return i
}
