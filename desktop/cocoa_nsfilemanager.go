// +build darwin

package desktop

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/Foundation/Classes/NSFileManager_Class/

var NSFileManagerClass unsafe.Pointer = Runtime_objc_lookUpClass("NSFileManager")
var NSFileManagerDefaultManager unsafe.Pointer = Runtime_sel_getUid("defaultManager")
var NSFileManagerURLsForDirectoryInDomains unsafe.Pointer = Runtime_sel_getUid("URLsForDirectory:inDomains:")

// http://developer.apple.com/library/mac/#documentation/Cocoa/Reference/Foundation/Miscellaneous/Foundation_Constants/Reference/reference.html#//apple_ref/doc/c_ref/NSSearchPathDirectory

type NSSearchPathDirectory int

const (
	NSApplicationDirectory          = 1
	NSDemoApplicationDirectory      = 2
	NSDeveloperApplicationDirectory = 3
	NSAdminApplicationDirectory     = 4
	NSLibraryDirectory              = 5
	NSDeveloperDirectory            = 6
	NSUserDirectory                 = 7
	NSDocumentationDirectory        = 8
	NSDocumentDirectory             = 9
	NSCoreServiceDirectory          = 10
	NSAutosavedInformationDirectory = 11
	NSDesktopDirectory              = 12
	NSCachesDirectory               = 13
	NSApplicationSupportDirectory   = 14
	NSDownloadsDirectory            = 15
	NSInputMethodsDirectory         = 16
	NSMoviesDirectory               = 17
	NSMusicDirectory                = 18
	NSPicturesDirectory             = 19
	NSPrinterDescriptionDirectory   = 20
	NSSharedPublicDirectory         = 21
	NSPreferencePanesDirectory      = 22
	NSApplicationScriptsDirectory   = 23
	NSItemReplacementDirectory      = 99
	NSAllApplicationsDirectory      = 100
	NSAllLibrariesDirectory         = 101
	NSTrashDirectory                = 102
)

// http://developer.apple.com/library/mac/#documentation/Cocoa/Reference/Foundation/Miscellaneous/Foundation_Constants/Reference/reference.html#//apple_ref/doc/c_ref/NSSearchPathDomainMask

type NSSearchPathDomainMask int

const (
	NSUserDomainMask    = 1
	NSLocalDomainMask   = 2
	NSNetworkDomainMask = 4
	NSSystemDomainMask  = 8
	NSAllDomainsMask    = 0x0ffff
)

type NSFileManager struct {
	NSObject
}

func NSFileManagerNew() NSFileManager {
	return NSFileManager{NSObjectPointer(Runtime_objc_msgSend(NSFileManagerClass, NSFileManagerDefaultManager))}
}

func (m NSFileManager) URLsForDirectoryInDomains(directory int, domainMask int) NSArray {
	return NSArrayPointer(Runtime_objc_msgSend(m.Pointer, NSFileManagerURLsForDirectoryInDomains, Int2Pointer(directory), Int2Pointer(domainMask)))
}
