// +build darwin

package desktop

func browserOpenURI(s string) {
	w := NSWorkspaceSharedWorkspace()
	defer w.Release()
	n := NSURLNew(s)
	defer n.Release()
	w.OpenURL(n)
}
