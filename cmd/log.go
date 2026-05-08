package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var liveLog bool
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Muestra los log de la API",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		Api.Name = args[0]

		file, err := apiMgr.Logger.GetLogFile(Api.Name)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		if !liveLog {
			// Static Log

			apiMgr.Logger.ReadLog(file.Name())

		} else {
			// Live Log
			utils.Clear()

			green := color.New(color.FgGreen).SprintFunc()

			fmt.Printf("%s %s %s \n\n",
				green("Monitoreando log de:"),
				green(Api.Name),
				green("(Ctrl+C para salir y cerrar API)"),
			)

			apiMgr.Logger.LiveLog(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().BoolVarP(&liveLog, "live", "l", false, "Ver el log en vivo de una API")

}
