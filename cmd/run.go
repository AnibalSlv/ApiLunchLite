/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var nameModule string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Ejecuta una API",
	Long: `Ejecuta una API seleccionada por su nombre
			apl run [nombre] -n [modulo:instancia]`,
	Args: cobra.ExactArgs(1), // Captura el primer argumento automaticamente
	Run: func(cmd *cobra.Command, args []string) {

		Api.Name = args[0]

		fmt.Println("Intentando ejecutar la API")

		result, err, inExecute := apiMgr.Run(Api.Name, nameModule)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		if !inExecute {
			fmt.Printf("\n%s\n", utils.Green(result))
		} else {
			fmt.Printf("\n%s\n", utils.Yellow(result))
		}

	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&nameModule, "name", "n", "main:app", "Módulo e instancia de la aplicación")
}
