/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateLib bool

var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Verifica todas las librerias de tu API",
	Long:  `apl lib [API]`,
	Run: func(cmd *cobra.Command, args []string) {

		Api.Name = args[0]

		fmt.Println("Verificando si hay librerias por actualizar")

		apiMgr.LibAll(Api.Name)

		fmt.Printf("\nLibrerias de %s actualizadas\n\n", Api.Name)
	},
}

func init() {
	rootCmd.AddCommand(libCmd)

	libCmd.Flags().BoolVar(&updateLib, "update", false, "Actualiza todas las librerias de la API a su ultima version disponible")
}
