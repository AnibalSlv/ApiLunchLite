package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func SearchPythonExe(path string) string {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("ERROR: ", err)
		return ""
	}

	pythonBinari := "python3"
	if runtime.GOOS == "Windows" {
		pythonBinari = "python"
	}
	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())

		// Se saltan las carpetas grandes
		if entry.IsDir() && (entry.Name() == "Lib" || entry.Name() == "Include") {
			continue
		}

		if entry.IsDir() {
			resultado := SearchPythonExe(fullPath)

			if resultado != "" {
				return resultado
			}
		} else {
			if entry.Name() == pythonBinari {
				return fullPath
			}
		}
	}

	pythonPathGlobal, err := exec.LookPath(pythonBinari) // O simplemente "python" en Windows
	if err != nil {
		fmt.Println("No se encontró Python en el PATH global del sistema.")
		return ""
	}

	return pythonPathGlobal
}
