// +build windows

package desktop

func browserOpenURI(s string) {
	w := WStringNew("open")
	defer w.Close()
	n := WStringNew(s)
	defer n.Close()
	ShellExecute.Call(NULL, Arg(w), Arg(n), NULL, NULL, Arg(SW_SHOWNORMAL))
}
