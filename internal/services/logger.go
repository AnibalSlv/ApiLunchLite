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

	file, err := os.Create(pathFile)

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
