package verify

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"os"
	"os/exec"
)

func GetJson(pythonPath, filePath string, cmd *exec.Cmd) error {

	file, err := os.Create(filePath)

	if err != nil {
		fmt.Printf("\n%s: %v\n", utils.Red("[Error] No se pudo crear el archivo tempRequeriment"), err)
		return nil
	}

	cmd.Stdout = file

	err = cmd.Run()

	if err != nil {
		fmt.Printf("\n%s: %v\n", utils.Red("[Error] No se pudo obtener el json"), err)
		return err
	}

	file.Close()

	return nil
}
