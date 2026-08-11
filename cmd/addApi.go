package cmd

import (
	"fmt"
	"strconv"

	"apiLunchLite/internal/utils"
	"apiLunchLite/models"

	"github.com/spf13/cobra"
)

var addApiCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega una Api",
	Long:  `Utiliza addApi para agregar una api`,
	Run: func(cmd *cobra.Command, args []string) {
		item := models.ApiConfig{
			Name:       Api.Name,
			Host:       Api.Host,
			Port:       Api.Port,
			PathFolder: Api.PathFolder,
		}

		apiMgr.AddApi(item)

		fmt.Printf("\n%s\n", utils.Green("Api Agregada:"))
		fmt.Printf("Nombre: %s\n", Api.Name)
		fmt.Printf("Host: %s\n", Api.Host)
		fmt.Printf("Port: %s\n", utils.Yellow(strconv.Itoa(Api.Port)))
		fmt.Printf("Path Folder: %s\n\n", Api.PathFolder)
	},
}

func init() {
	rootCmd.AddCommand(addApiCmd)

	addApiCmd.Flags().StringVarP(&Api.Name, "name", "n", "API", "Agreagarle un nombre a la API")
	addApiCmd.Flags().StringVarP(&Api.Host, "host", "H", "localhost", "El Host para el server")
	addApiCmd.Flags().IntVarP(&Api.Port, "port", "p", 8080, "Agregar un puerto")
	addApiCmd.Flags().StringVarP(&Api.PathFolder, "folder", "f", "", "Coloca la dirrecion de la carpeta de tu API")

}
