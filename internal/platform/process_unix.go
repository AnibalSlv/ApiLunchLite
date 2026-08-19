//go:build !windows

package platform

import (
	"os/exec"
	"syscall"
)

// startDaemon aplica el truco de Unix para desvincular el proceso
func StartDaemon(clon *exec.Cmd) {
	clon.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true, // Crea una nueva sesión independiente de la terminal
	}
}
