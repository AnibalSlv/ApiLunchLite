/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Verifica todas las librerias de tu API",
	Long:  `apl lib [API]`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 || len(args) > 1 {
			fmt.Printf("\n%s\n\n", utils.Red("[Error] Necesita colocar una API"))

		} else {
			Api.Name = args[0]

			fmt.Println("Verificando las librerias")

			apiMgr.LibAll(Api.Name)

		}
	},
}

func init() {
	rootCmd.AddCommand(libCmd)
}
