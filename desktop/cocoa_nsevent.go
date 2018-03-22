// +build darwin

package desktop

/*
#include <stdlib.h>
#include <objc/objc-runtime.h>
#include <CoreGraphics/CGGeometry.h>

extern void NSEventActionMap(void*);

inline id NSEventAddGlobalMonitor(id to, SEL sel, SEL mask, void(*p)(void*) ) {
  return objc_msgSend(to, sel, mask, ^(void* event) {
    p(event);
  });
}
*/
import "C"

import (
	"unsafe"
)

// https://developer.apple.com/library/mac/documentation/Cocoa/Reference/ApplicationKit/Classes/NSEvent_Class

type NSEventMask uint

func (m NSEventMask) Pointer() unsafe.Pointer {
	return UInt2Pointer(uint(m))
}

const (
	NSAnyEventMask = -1
)

const (
	NSLeftMouseDown         NSEventMask = 1
	NSLeftMouseUp                       = 2
	NSRightMouseDown                    = 3
	NSRightMouseUp                      = 4
	NSMouseMoved                        = 5
	NSLeftMouseDragged                  = 6
	NSRightMouseDragged                 = 7
	NSMouseEntered                      = 8
	NSMouseExited                       = 9
	NSKeyDown                           = 10
	NSKeyUp                             = 11
	NSFlagsChanged                      = 12
	NSAppKitDefined                     = 13
	NSSystemDefined                     = 14
	NSApplicationDefined                = 15
	NSPeriodic                          = 16
	NSCursorUpdate                      = 17
	NSScrollWheel                       = 22
	NSTabletPoint                       = 23
	NSTabletProximity                   = 24
	NSOtherMouseDown                    = 25
	NSOtherMouseUp                      = 26
	NSOtherMouseDragged                 = 27
	NSEventTypeGesture                  = 29
	NSEventTypeMagnify                  = 30
	NSEventTypeSwipe                    = 31
	NSEventTypeRotate                   = 18
	NSEventTypeBeginGesture             = 19
	NSEventTypeEndGesture               = 20
	NSEventTypeSmartMagnify             = 32
	NSEventTypeQuickLook                = 33
	NSEventTypePressure                 = 34
)

type NSEventHandler func(e NSEvent)

type NSEventMonitor struct {
	e    NSEvent
	mask NSEventMask
}

var NSEventMap = make(map[*NSEventMonitor]NSEventHandler)

//export NSEventActionMap
func NSEventActionMap(e unsafe.Pointer) {
	t := NSEventPointer(e).GetType()
	for k, v := range NSEventMap {
		if k.mask&t == t {
			v(NSEventPointer(e))
		}
	}
}

var NSEventClass unsafe.Pointer = Runtime_objc_lookUpClass("NSEvent")
var NSEventType unsafe.Pointer = Runtime_sel_getUid("type")
var NSEventAddGlobalMonitorSel unsafe.Pointer = Runtime_sel_getUid("addGlobalMonitorForEventsMatchingMask:handler:")
var NSEventRemoveMonitorSel unsafe.Pointer = Runtime_sel_getUid("removeMonitor:")

type NSEvent struct {
	NSObject
}

func NSEventNew() NSEvent {
	var m = NSEventPointer(Runtime_class_createInstance(NSEventClass, 0))
	return m
}

func NSEventAddGlobalMonitor(mask NSEventMask, handler NSEventHandler) NSEvent {
	monitor := &NSEventMonitor{mask: mask}
	NSEventMap[monitor] = handler
	m := NSEventPointer(unsafe.Pointer(C.NSEventAddGlobalMonitor(NSEventClass, NSEventAddGlobalMonitorSel, mask.Pointer(), (*[0]byte)(C.NSEventActionMap))))
	monitor.e = m
	return m
}

func NSEventRemoveMonitor(e NSEvent) {
	for k, _ := range NSEventMap {
		if k.e.Pointer == e.Pointer {
			Runtime_objc_msgSend(NSEventClass, NSEventRemoveMonitorSel, e.Pointer)
			delete(NSEventMap, k)
			return
		}
	}
}

func NSEventPointer(p unsafe.Pointer) NSEvent {
	return NSEvent{NSObjectPointer(p)}
}

func (m NSEvent) GetType() NSEventMask {
	return NSEventMask(Pointer2Int(Runtime_objc_msgSend(m.Pointer, NSEventType)))
}
