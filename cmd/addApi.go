package cmd

import (
	"fmt"
	"log"
	"strconv"

	"apiLunchLite/internal/utils"
	"apiLunchLite/models"

	"github.com/spf13/cobra"
)

var addApiCmd = &cobra.Command{
	Use:   "add",
	Short: "Agrega una Api",
	Long: `apl -n [nombre del servicio] -f [direccion absoluta de la carpeta con el servicio] -H [Host] -p [Puerto]
	Ejemplo: apl -n API1 -f C:\users\...\backend -p 8100
	`,
	Run: func(cmd *cobra.Command, args []string) {
		item := models.ApiConfig{
			Name:       Api.Name,
			Host:       Api.Host,
			Port:       Api.Port,
			PathFolder: Api.PathFolder,
		}

		result, err, complete := apiMgr.AddApi(item)

		if err != nil {
			log.Fatal("Error: ", err)
		}

		if !complete {
			fmt.Printf("\n%s\n\n", utils.Red("Por favor coloque el path de la carpeta contenedora de su API"))
		} else {
			fmt.Printf("\n%s\n", utils.Green("Api Agregada:"))
			fmt.Printf("Nombre: %s\n", result.Name)
			fmt.Printf("Host: %s\n", result.Host)
			fmt.Printf("Port: %s\n", utils.Yellow(strconv.Itoa(result.Port)))
			fmt.Printf("Path Folder: %s\n\n", result.PathFolder)
		}
	},
}

func init() {
	rootCmd.AddCommand(addApiCmd)

	addApiCmd.Flags().StringVarP(&Api.Name, "name", "n", "", "Agreagarle un nombre a la API")
	addApiCmd.Flags().StringVarP(&Api.Host, "host", "H", "localhost", "El Host para el server")
	addApiCmd.Flags().IntVarP(&Api.Port, "port", "p", 8080, "Agregar un puerto")
	addApiCmd.Flags().StringVarP(&Api.PathFolder, "folder", "f", "", "Coloca la dirrecion de la carpeta de tu API")

}
