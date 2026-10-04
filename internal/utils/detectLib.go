package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Verifica que uvicorn este instalado
func VerifyDependency(pythonPath string) bool {

	url := "internal/temp"

	if _, err := os.Stat(url); os.IsNotExist(err) {
		// Crear con permisos 0755 (lectura, escritura, ejecución para el dueño)
		err = os.Mkdir(url, 0755)
		if err != nil {
			fmt.Println("Error al crear:", err)
			return false
		}
		fmt.Println("Carpeta creada con éxito.")
	}

	tempFile, err := createTempFile(pythonPath)
	if err != nil {
		return false
	}

	defer os.Remove(tempFile)

	readFile, err := os.Open(tempFile)

	if err != nil {
		fmt.Printf("%s: %v", Red("[Error] Al abrir las dependencias"), err)
		return false
	}

	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	var found = false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "uvicorn") {
			found = true
			break
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("%s: %v", Red("[Error] No se pudo leer las dependencias: "), err)
		return false
	}

	if !found {
		fmt.Printf("%s", Yellow("Se necesita descargar uvicorn para poder ejecutar la API"))
		return false
	}

	return true
}

// Crea un archivo temporar con las librerias del lenguaje que el usuario selecciono
func createTempFile(pythonPath string) (string, error) {

	pathFile := "internal/temp/requeriment.txt"

	file, err := os.Create(pathFile)

	if err != nil {
		fmt.Printf("\n%s: %v\n", Red("[Error] No se pudo crear el archivo tempRequeriment"), err)
		return "", err
	}

	cmd := exec.Command(pythonPath, "-m", "pip", "freeze")

	cmd.Stdout = file

	err = cmd.Run()

	file.Close()

	if err != nil {
		fmt.Printf("\n%s: %v\n", Red("[Error] No se pudo verificar las dependencias:"), err)
		return "", err
	}

	return pathFile, nil
}
