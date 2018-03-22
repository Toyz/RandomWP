// +build darwin

package desktop

import (
	"os"
)

// user application data folder
func getAppDataFolder() string {
	return path(NSApplicationSupportDirectory, NSUserDomainMask)
}

// user home "/home/user"
func getHomeFolder() string {
	return os.Getenv("HOME")
}

// user my documents "~/Documents"
func getDocumentsFolder() string {
	return path(NSDocumentDirectory, NSUserDomainMask)
}

// user downloads "~/Downloads"
func getDownloadsFolder() string {
	return path(NSDownloadsDirectory, NSUserDomainMask)
}

// user desktop "~/Desktop"
func getDesktopFolder() string {
	return path(NSDesktopDirectory, NSUserDomainMask)
}

func path(d int, dd int) string {
	f := NSFileManagerNew()
	defer f.Release()

	a := f.URLsForDirectoryInDomains(d, dd)
	defer a.Release()

	if a.Count() != 1 {
		return ""
	}

	var u NSURL = NSURLPointer(a.ObjectAtIndex(0))
	defer u.Release()

	return u.Path()
}
