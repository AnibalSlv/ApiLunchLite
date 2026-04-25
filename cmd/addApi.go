/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"apiLunchLite/internal/database"
	"apiLunchLite/models"

	"github.com/spf13/cobra"
)

var addApiCmd = &cobra.Command{
	Use:   "addApi",
	Short: "Agrega una Api",
	Long:  `Utiliza addApi para agregar una api`,
	Run: func(cmd *cobra.Command, args []string) {
		item := models.ApiConfig{
			Name:       Api.Name,
			Host:       Api.Host,
			Port:       Api.Port,
			PathFolder: Api.PathFolder,
		}

		database.Save(Api.DbConn, item)

		fmt.Println("Api Agregada:")
		fmt.Printf("Nombre: %s\n", Api.Name)
		fmt.Printf("Host: %s\n", Api.Host)
		fmt.Printf("Port: %d\n", Api.Port)
		fmt.Printf("Path: %s\n", Api.PathFolder)

	},
}

func init() {
	rootCmd.AddCommand(addApiCmd)

	addApiCmd.Flags().StringVarP(&Api.Name, "name", "n", "API", "Agreagarle un nombre a la API")
	addApiCmd.Flags().StringVarP(&Api.Host, "host", "H", "localhost", "El Host para el server")
	addApiCmd.Flags().IntVarP(&Api.Port, "port", "p", 8080, "Agregar un puerto")
	addApiCmd.Flags().StringVarP(&Api.PathFolder, "folder", "f", "", "Coloca la dirrecion de la carpeta de tu API")

}
