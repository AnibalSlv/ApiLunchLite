/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var updateAllLib bool

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Permite actualizar las librerias de una API",
	Long: `apl update [API] [lib1] [lib2] [lib...]  <- Actualiza las librerias especificadas
		   apl update [API] --all                   <- Actualiza todas las librerias
	`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Printf("\n%s\n\n", utils.Red("[Error] Necesita colocar una API y una libreria"))
		} else {
			Api.Name = args[0]
			libList := []string{}
			libList = append(libList, args[1:]...)

			fmt.Println("Intentando actualizar librerias...")
			apiMgr.UpdateLib(Api.Name, libList, updateAllLib)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	libCmd.Flags().BoolVar(&updateAllLib, "all", false, "Actualiza todas las librerias de la API a su ultima version disponible")
}
