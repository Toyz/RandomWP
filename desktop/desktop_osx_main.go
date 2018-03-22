// +build darwin

package desktop

import (
	"runtime"
)

var App NSApplication

func desktopMain() {
	App = NSApplicationSharedApplication()
	defer App.Release()
	App.Run()
}

func desktopMainClose() {
	// get locked when window created in desktop
	runtime.UnlockOSThread()

	App.Terminate()
}
