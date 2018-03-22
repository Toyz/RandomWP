// +build windows

package desktop

import (
	"syscall"
	"unsafe"
)

var IID_IHTMLDocument2 = GUIDNew("{332c4425-26cb-11d0-b483-00c04fd90119}")

type IHTMLDocument2 struct {
	IDispatch
}

type IHTMLDocument2Vtbl struct {
	IDispatchVtbl
	get_Script           uintptr
	get_all              uintptr
	get_body             uintptr
	get_activeElement    uintptr
	get_images           uintptr
	get_applets          uintptr
	get_links            uintptr
	get_forms            uintptr
	get_anchors          uintptr
	put_title            uintptr
	get_title            uintptr
	get_scripts          uintptr
	put_designMode       uintptr
	get_designMode       uintptr
	get_selection        uintptr
	get_readyState       uintptr
	get_frames           uintptr
	get_embeds           uintptr
	get_plugins          uintptr
	put_alinkColor       uintptr
	get_alinkColor       uintptr
	put_bgColor          uintptr
	get_bgColor          uintptr
	put_fgColor          uintptr
	get_fgColor          uintptr
	put_linkColor        uintptr
	get_linkColor        uintptr
	put_vlinkColor       uintptr
	get_vlinkColor       uintptr
	get_referrer         uintptr
	get_location         uintptr
	get_lastModified     uintptr
	put_URL              uintptr
	get_URL              uintptr
	put_domain           uintptr
	get_domain           uintptr
	put_cookie           uintptr
	get_cookie           uintptr
	put_expando          uintptr
	get_expando          uintptr
	put_charset          uintptr
	get_charset          uintptr
	put_defaultCharset   uintptr
	get_defaultCharset   uintptr
	get_mimeType         uintptr
	get_fileSize         uintptr
	get_fileCreatedDate  uintptr
	get_fileModifiedDate uintptr
	get_fileUpdatedDate  uintptr
	get_security         uintptr
	get_protocol         uintptr
	get_nameProp         uintptr
	Write                uintptr
	writeln              uintptr
	open                 uintptr
	Close                uintptr
	clear                uintptr
}

func (v *IHTMLDocument2) VTable() *IHTMLDocument2Vtbl {
	return (*IHTMLDocument2Vtbl)(unsafe.Pointer(v.RawVTable))
}

func (m *IHTMLDocument2) Write(sa *SAFEARRAY) HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Write,
		2,
		Arg(m),
		Arg(sa),
		0,
	))
}

func (m *IHTMLDocument2) Close() HRESULT {
	return HRESULTPtr(syscall.Syscall(
		m.VTable().Close,
		1,
		Arg(m),
		0,
		0,
	))
}
