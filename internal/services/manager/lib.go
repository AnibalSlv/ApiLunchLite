package manager

import (
	"apiLunchLite/internal/utils"
	"apiLunchLite/internal/utils/verify"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func (m *ApiManager) LibAll(name string) (error, bool) {

	inExecution := false

	result, err := m.Db.GetName(name)
	if err != nil {
		return err, inExecution
	}

	if result.Pid != 0 {
		err := m.Stop(name, false)
		if err != nil {
			return nil, inExecution
		}
	}

	pythonPath := utils.SearchPythonExe(result.PathFolder)
	cmdUpdate := exec.Command(pythonPath, "-m", "pip", "list", "--outdated", "--format=json")

	filePath := "internal/temp/requeriment.json"
	verify.GetJson(pythonPath, filePath, cmdUpdate)

	readJson(filePath)

	return nil, false
}

func readJson(filePath string) {

	// El tag `json:"..."` sirve para traducir cuando se crea un archivo json es decir
	// en vez de "Name":"Api1" --> "name":"Api" (por tema de que en json se suele trabajar en minusculas)
	type FileLib struct {
		Name            string `json:"name"`
		Version         string `json:"version"`
		Latest_version  string `json:"latest_version"`
		Latest_filetype string `json:"latest_filetype"`
	}
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("[Error] No se pudo abrir el archivo json", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var fileLib []FileLib

	decoder.Decode(&fileLib)

	for _, api := range fileLib {
		if api.Version != api.Latest_version {
			fmt.Printf("\n%s requiere actualziacion %s -> %s\n", api.Name, utils.Yellow(api.Version), utils.Green(api.Latest_version))
		} else {
			fmt.Printf("\n%s ya tiene la ultima actualizacion\n", utils.Green(api.Name))
		}
	}
}
