// +build linux

package desktop

import (
	"runtime"
	"sync"
)

type GtkMessageLoop struct {
	Count    int
	Lock     *sync.Cond
	MainLoop GMainLoop
	Context  GMainContext
}

var messageloop *GtkMessageLoop

func GtkMessageLoopInc() {
	if messageloop == nil {
		messageloop = GtkMessageLoopNew()
	}
	messageloop.Count++
}

func GtkMessageLoopDec() {
	if messageloop != nil {
		messageloop.Count--
		if messageloop.Count == 0 {
			messageloop.Close()
			messageloop = nil
		}
	}
}

func GtkMessageLoopInvoke(fn *GSourceFunc) {
	g_main_context_invoke(messageloop.Context, fn)
}

func GtkMessageLoopThreadsNew() *GtkMessageLoop {
	m := &GtkMessageLoop{}
	m.Lock = sync.NewCond(&sync.Mutex{})

	go func() {
		m.Init()

		m.Lock.L.Lock()
		m.Lock.Broadcast()
		m.Lock.L.Unlock()

		m.Main()
	}()

	m.Lock.L.Lock()
	m.Lock.Wait()
	m.Lock.L.Unlock()

	return m
}

func GtkMessageLoopNew() *GtkMessageLoop {
	m := &GtkMessageLoop{}

	m.Init()

	return m
}

func (m *GtkMessageLoop) Init() {
	runtime.LockOSThread()

	gtk_init()

	m.MainLoop = g_main_loop_new(nil, false)
	m.Context = g_main_loop_get_context(m.MainLoop)
}

func (m *GtkMessageLoop) Main() {
	g_main_loop_run(m.MainLoop)

	runtime.UnlockOSThread()
}

func (m *GtkMessageLoop) Close() {
	if m.MainLoop != nil {
		g_main_loop_quit(m.MainLoop)
		m.MainLoop = nil
	}
}

func desktopGtkMain() {
	messageloop.Main()
}
