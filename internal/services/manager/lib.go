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
	cmdOut := exec.Command(pythonPath, "-m", "pip", "list", "--outdated", "--format=json")
	cmdUpt := exec.Command(pythonPath, "-m", "pip", "list", "--uptodate", "--format=json")

	filePathOut := "internal/temp/requerimentOut.json"
	filePathUpt := "internal/temp/requerimentUpt.json"

	verify.GetJson(pythonPath, filePathOut, cmdOut)
	verify.GetJson(pythonPath, filePathUpt, cmdUpt)

	// var fileLib []verify.FileLib <-- Este es el tipo de dato
	fileLibOut := verify.ReadJson(filePathOut)
	fileLibUpt := verify.ReadJson(filePathUpt)

	for _, api := range fileLibUpt {
		fmt.Printf("\n%s %s\n", utils.Green(api.Name), api.Version)

	}

	for _, api := range fileLibOut {
		fmt.Printf("\n%s %s -> %s\n", utils.Red(api.Name), utils.Yellow(api.Version), utils.Green(api.Latest_version))
	}

	return nil
}
