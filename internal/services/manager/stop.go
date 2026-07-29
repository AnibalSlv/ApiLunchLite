package manager

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
)

func (m *ApiManager) Stop(nameStop string, forceStop bool) error {
	result, err := m.Db.GetName(nameStop)

	if err != nil {
		fmt.Println("Error Connection Db: ", err)
		return err
	}

	procces, err := os.FindProcess(result.Pid)

	if err != nil {
		fmt.Println("No se encontró el proceso (ya podría estar detenido): ", err)
		m.Db.UpdatePID(0, result.Id)
		m.Db.UpdateState("stop", result.Id)
		return err
	}

	if forceStop {
		err = procces.Kill()
	} else {
		if runtime.GOOS == "windows" {
			err = procces.Kill()
		} else {
			// En Linux/macOS se usa la señal SIGTERM estándar
			err = procces.Signal(syscall.SIGTERM)
		}
	}

	if err != nil {
		fmt.Println("Error Close Process: ", err)
		return err
	}

	m.Db.UpdatePID(0, result.Id)
	m.Db.UpdateState("stop", result.Id)

	fmt.Println("Proceso detenido exitosamente")

	return nil
}
