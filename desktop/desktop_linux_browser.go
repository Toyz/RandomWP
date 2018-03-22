// +build linux

package desktop

import (
	"fmt"
	"os/exec"
	"strings"
)

type RunErr struct {
	error

	Cmd string
	Out string
}

func (m *RunErr) Error() string {
	return fmt.Sprintf("# %s\n%s\n%s", m.Cmd, m.Out, m.error.Error())
}

func Run(args ...string) *RunErr {
	head := args[0]
	tail := args[1:]

	err := exec.Command(head, tail...).Start()

	if err != nil {
		return &RunErr{err, strings.Join(args, " "), ""}
	}

	return nil
}

func Shell(s string) *RunErr {
	err := Run("sh", "-c", s)
	if err != nil {
		return &RunErr{err.error, s, err.Out}
	}
	return nil
}

func browserOpenURI(s string) {
	Run("xdg-open", s)
}
