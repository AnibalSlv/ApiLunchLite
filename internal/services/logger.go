package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	FolderLogs string
}

func (l *Logger) GetLogFile(name string) (*os.File, error) {

	pathFile := filepath.Join("internal/logs", name+".log")

	// os.O_APPEND: Escribe al final del archivo.
	// os.O_CREATE: Crea el archivo si no existe.
	// os.O_WRONLY: Abre en modo solo escritura.
	// 0644: Permisos de lectura/escritura para el propietario y lectura para otros.
	file, err := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Error Create File: ", err)
		return nil, err
	}

	return file, nil
}

func (l *Logger) ReadLog(fileName string) error {
	fileLog, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error Read Log: ", err)
		return err
	}

	fmt.Print(string(fileLog))
	return nil
}

func (l *Logger) LiveLog(fileName string) error {

	fileLive, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	defer fileLive.Close()

	reader := bufio.NewReader(fileLive)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			fmt.Println("Error al leer: ", err)
			break
		}

		fmt.Print(line)
	}

	return nil
}
