package manager

import (
	"fmt"
	"os"
)

func (m *ApiManager) Stop(nameStop string, forceStop bool) error {
	result, err := m.Db.GetName(nameStop)

	if err != nil {
		fmt.Println("Error Connection Db: ", err)
		return err
	}

	if forceStop {
		m.Db.UpdatePID(0, result.Id)
		m.Db.UpdateState("stop", result.Id)
		return err
	}

	procces, err := os.FindProcess(result.Pid)

	if err != nil {
		fmt.Println("Error Find PID: ", err)
		return err
	}

	// ! Se deberia de cambiar el .Kill() por otro metodo que mate el proceso no tan bruscamente
	err = procces.Kill()

	if err != nil {
		fmt.Println("Error Close Process: ", err)
		return err
	}

	m.Db.UpdatePID(0, result.Id)
	m.Db.UpdateState("stop", result.Id)

	fmt.Println("Proceso detenido exitosamente")

	return nil
}
