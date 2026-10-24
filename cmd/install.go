/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var installLibVersion string
var installLib string

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Instala una o varias librerias",
	Long: `apl install [libreria] [version]
		   apl install [libreria1] [libreria2] [...]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Printf("%s Debes especificar al menos una api y una librería para instalar.\n", utils.Red("[Error]:"))
			return
		}

		Api.Name = args[0]
		libList := []string{}
		libList = append(libList, args[1:]...)

		if installLibVersion != "latest" && len(libList) == 1 {
			apiMgr.InstallOneLib(Api.Name, libList[0], installLibVersion)
			return
		}

		if installLibVersion != "latest" && len(libList) > 1 {
			fmt.Printf("%s Para descargar una version especifica debes colocar una sola libreria y la version.\n", utils.Red("[Error]:"))
			return
		}

		apiMgr.InstallVariousLib(Api.Name, libList)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	installCmd.Flags().StringVarP(&installLibVersion, "version", "v", "latest", "Instala una version especifica de una libreria")

}
