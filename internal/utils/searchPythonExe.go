package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func SearchPythonExe(path string) string {
	entries, err := os.ReadDir(path)

	if err != nil {
		return ""
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
			if entry.Name() == "python.exe" {
				return fullPath
			}

		}
	}

	pythonPathGlobal, err := exec.LookPath("python3") // O simplemente "python" en Windows
	if err != nil {
		fmt.Println("No se encontró Python en el PATH global del sistema.")
		return ""
	}

	return pythonPathGlobal
}
