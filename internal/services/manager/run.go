package manager

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"os/exec"
	"strconv"
)

func (m *ApiManager) Run(name string, nameModule string) error {
	result, err := m.Db.GetName(name)

	if err != nil {
		fmt.Print("Error Connection Db: ", err)
		return err
	}

	if result.Pid == 0 {

		rootFolder := result.PathFolder
		pythonPath := utils.SearchPythonExe(rootFolder)

		// Sirve para que Go lea los argumentos por separado
		uvicornArgs := []string{
			"-m", "uvicorn",
			nameModule,
			"--host", result.Host,
			"--port", strconv.Itoa(result.Port),
		}

		uvicornCmd := exec.Command(pythonPath, uvicornArgs...)

		uvicornCmd.Dir = rootFolder

		logFile, err := m.Logger.GetLogFile(result.Name)
		if err != nil {
			return err
		}

		// Escribe en los archivos que le pidio al logger
		uvicornCmd.Stdout = logFile
		uvicornCmd.Stderr = logFile

		// Iniciamos el proceso sin bloquear el programa
		err = uvicornCmd.Start()
		if err != nil {
			fmt.Println("Error al iniciar el proceso:", err)
			return err
		}

		m.Db.UpdatePID(uvicornCmd.Process.Pid, result.Id)
		m.Db.UpdateState("run", result.Id)
		fmt.Println("API ejecutada exitosamente")
	} else {
		fmt.Println("La API ya se esta ejecutnado")
	}

	return nil
}
