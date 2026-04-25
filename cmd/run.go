/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"apiLunchLite/internal/database"
	"apiLunchLite/internal/utils"

	"github.com/spf13/cobra"
)

var nameModule string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Ejecuta una API",
	Long: `Ejecuta una API seleccionada por su nombre
			apl run [nombre] -n [modulo:instancia]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			Api.Name = args[0]

			result, err := database.GetName(Api.DbConn, Api.Name)

			if err != nil {
				fmt.Print("Error Connection Db: ", err)
				return
			}

			if result.Pid == 0 {

				pathFile := filepath.Join("internal/logs", result.Name+".log")

				file, err := os.Create(pathFile)

				if err != nil {
					fmt.Println("Error Create File: ", err)
					return
				}

				defer file.Close()

				rootFolder := result.PathFolder
				pythonPath := utils.SearchPythonExe(rootFolder)

				// Sirve para que Go lea los argumentos por separado
				uvicornArgs := []string{
					"-m", "uvicorn",
					nameModule,
					"--host", result.Host,
					"--port", strconv.Itoa(result.Port),
				}

				uvicornCmd := exec.Command(pythonPath, uvicornArgs...)

				uvicornCmd.Dir = rootFolder

				// Conecta los log de uvicorn con el archivo.log
				uvicornCmd.Stdout = file

				// Conecta los mensajes de errores de uvicorn con el archivo.log
				uvicornCmd.Stderr = file

				err = uvicornCmd.Start()
				if err != nil {
					fmt.Println("Error run:", err)
				}

				database.UpdatePID(Api.DbConn, uvicornCmd.Process.Pid, result.Id)
				database.UpdateState(Api.DbConn, "run", result.Id)
				fmt.Println("API ejecutada exitosamente")
			} else {
				fmt.Println("La API ya se esta ejecutnado")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&nameModule, "name", "n", "main:app", "Módulo e instancia de la aplicación")
}
