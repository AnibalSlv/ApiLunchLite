//go:build windows

package platform

import (
	"os/exec"
	"syscall"
)

// startDaemon aplica las banderas nativas de Windows
func StartDaemon(clon *exec.Cmd) {
	clon.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: 0x00000008, // DETACHED_PROCESS
	}
}
