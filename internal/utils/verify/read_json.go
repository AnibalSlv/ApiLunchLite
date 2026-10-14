package verify

import (
	"encoding/json"
	"fmt"
	"os"
)

// El tag `json:"..."` sirve para traducir cuando se crea un archivo json es decir
// en vez de "Name":"Api1" --> "name":"Api" (por tema de que en json se suele trabajar en minusculas)
type FileLib struct {
	Name            string `json:"name"`
	Version         string `json:"version"`
	Latest_version  string `json:"latest_version"`
	Latest_filetype string `json:"latest_filetype"`
}

func ReadJson(filePath string) []FileLib {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("[Error] No se pudo abrir el archivo json", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var fileLib []FileLib

	decoder.Decode(&fileLib)

	return fileLib
}
