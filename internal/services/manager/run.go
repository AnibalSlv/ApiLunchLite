package manager

import (
	"os/exec"
	"strconv"

	"apiLunchLite/internal/utils"
)

func (m *ApiManager) Run(name string, nameModule string) (string, error, bool) {
	inExecution := false

	result, err := m.Db.GetName(name)
	if err != nil {
		return "", err, inExecution
	}

	if result.Pid == 0 {

		rootFolder := result.PathFolder

		pythonPath := utils.SearchPythonExe(rootFolder)

		vDependecy := utils.VerifyDependency(pythonPath)

		if !vDependecy {
			return "", nil, inExecution
		}

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
			return "", err, inExecution
		}

		// Escribe en los archivos que le pidio al logger
		uvicornCmd.Stdout = logFile
		uvicornCmd.Stderr = logFile

		// Inicia el proceso sin bloquear el programa
		err = uvicornCmd.Start()
		if err != nil {
			return "", err, inExecution
		}

		m.Db.UpdatePID(uvicornCmd.Process.Pid, result.Id)
		m.Db.UpdateState("run", result.Id)
		return "API ejecutada exitosamente", nil, inExecution

	} else {
		inExecution = true
		return "La API ya se esta ejecutnado", nil, inExecution
	}
}
