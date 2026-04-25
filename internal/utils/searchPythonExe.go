package utils

import (
	"os"
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

	//  Si no se encuentran coincidencias:
	return ""
}
