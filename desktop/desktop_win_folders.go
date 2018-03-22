// +build windows

package desktop

import (
	"os"
)

func getAppDataFolder() string {
	return path(CSIDL_LOCAL_APPDATA)
}

func getHomeFolder() string {
	return os.Getenv("USERPROFILE")
}

func getDocumentsFolder() string {
	return path(CSIDL_PERSONAL)
}

func getDownloadsFolder() string {
	if IsWindowsXP() {
		// xp has no default downloads folder. so be it ~/Documents :)
		return getDocumentsFolder()
	} else {
		// vista+ has a download folder
		//
		// http://stackoverflow.com/questions/7672774/how-do-i-determine-the-windows-download-folder-path
		//
		var guid GUID = GUIDNew("374DE290-123F-4565-9164-39C4925E467B")
		return knowpath(guid)
	}
}

func getDesktopFolder() string {
	return path(CSIDL_DESKTOPDIRECTORY)
}

func knowpath(guid GUID) string {
	var pszPath uintptr
	HRESULTPtr(SHGetKnownFolderPath.Call(Arg(&guid), Arg(SHGFP_TYPE_CURRENT), NULL, Arg(&pszPath))).S_OK()
	path := WString2String(pszPath)
	CoTaskMemFree.Call(pszPath)
	return path
}

func path(nFolder int) string {
	pszPath := [MAX_PATH]uint16{}
	HRESULTPtr(SHGetFolderPath.Call(NULL, Arg(nFolder), NULL, Arg(SHGFP_TYPE_CURRENT), Arg(&pszPath[0]))).S_OK()
	return WString2String(Arg(&pszPath[0]))
}
