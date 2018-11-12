package dos

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func System(cmdline string) error {
	var buffer strings.Builder

	buffer.WriteString(`/S /C "`)
	buffer.WriteString(cmdline)
	buffer.WriteString(`"`)

	cmd1 := exec.Command(os.Getenv("COMSPEC"))
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	cmd1.Stdin = os.Stdin
	if cmd1.SysProcAttr == nil {
		cmd1.SysProcAttr = new(syscall.SysProcAttr)
	}
	cmd1.SysProcAttr.CmdLine = buffer.String()
	return cmd1.Run()
}
