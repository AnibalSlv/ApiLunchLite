package manager

import (
	"apiLunchLite/internal/utils"
	"apiLunchLite/internal/utils/verify"
	"fmt"
	"os/exec"
)

func (m *ApiManager) LibAll(name string) error {

	result, err := m.Db.GetName(name)
	if err != nil {
		return err
	}

	pythonPath := utils.SearchPythonExe(result.PathFolder)
	cmd := exec.Command(pythonPath, "-m", "pip", "list", "--outdated", "--format=json")

	filePath := "internal/temp/requeriment.json"
	verify.GetJson(pythonPath, filePath, cmd)

	// var fileLib []verify.FileLib <-- Este es el tipo de dato
	fileLib := verify.ReadJson(filePath)

	for _, api := range fileLib {
		if api.Version != api.Latest_version {
			fmt.Printf("\n%s requiere actualziacion %s -> %s\n", api.Name, utils.Yellow(api.Version), utils.Green(api.Latest_version))
		} else {
			fmt.Printf("\n%s ya tiene la ultima actualizacion\n", utils.Green(api.Name))
		}
	}
	return nil
}
