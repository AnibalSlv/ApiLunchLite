/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var coreCmd = &cobra.Command{
	Use:    "core",
	Short:  "Ejecuta el daemon del reverse proxy",
	Long:   `Este comando corre en segundo plano y mantiene vivo el proxy.`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		proxy, err := apiMgr.Core()

		if err != nil {
			fmt.Print("Error: ", err)
		}

		log.Println("Reverse Proxy escuchando en :8058...")
		log.Fatal(http.ListenAndServe(":8058", proxy))
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)
}
