package manager

import (
	"apiLunchLite/internal/utils"
	"apiLunchLite/internal/utils/verify"
	"fmt"
	"os/exec"
)

func (m *ApiManager) UpdateLib(name string, libList []string, allLib bool) (error, bool) {
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
	cmd := exec.Command(pythonPath, "-m", "pip", "list", "--outdated", "--format=json")

	filePath := "internal/temp/requeriment.json"
	verify.GetJson(pythonPath, filePath, cmd)

	// var fileLib []verify.FileLib <-- Este es el tipo de dato
	fileLib := verify.ReadJson(filePath)

	if allLib {
		for _, api := range fileLib {
			if api.Version != api.Latest_version {
				fmt.Printf("\nActualizando %s %s -> %s\n", api.Name, utils.Yellow(api.Version), utils.Green(api.Latest_version))
				cmdUpdate := exec.Command(pythonPath, "-m", "pip", "install", "--upgrade", api.Name)
				cmdUpdate.Run()
			} else {
				fmt.Printf("\n%s ya tiene la ultima actualizacion\n", utils.Green(api.Name))
			}
		}
	} else {
		// Crea un diccionario (map) donde la key es un string y en este caso seria `api.name`
		// el resto es la informacion interna de la api, ejemplo:
		// fileMap["fastapi"] = {Name: "fastapi", Version: "0.95.0", ...}
		fileMap := make(map[string]verify.FileLib)
		for _, api := range fileLib {
			fileMap[api.Name] = api
		}

		// Busca la libreria actualizar entre las librerias de la api
		for _, lib := range libList {
			api, exists := fileMap[lib]

			if !exists {
				fmt.Printf("%s\n", utils.Red(lib+" no encontrada"))
				continue
			}

			if api.Version != api.Latest_version {
				fmt.Printf("\nActualizando %s %s -> %s\n", api.Name, utils.Yellow(api.Version), utils.Green(api.Latest_version))
				cmdUpdate := exec.Command(pythonPath, "-m", "pip", "install", "--upgrade", api.Name)
				cmdUpdate.Run()
			} else {
				fmt.Printf("\n%s ya tiene la ultima actualizacion\n", utils.Green(api.Name))
			}
		}
	}

	return err, inExecution
}
