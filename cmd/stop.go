/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/database"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var nameStop string
var forceStop bool

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Detiene el daemon de la API",
	Long: `Detiene el daemon de la API:
			apl stop [nombre]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {

			nameStop = args[0]

			result, err := database.GetName(Api.DbConn, nameStop)

			if err != nil {
				fmt.Println("Error Connection Db: ", err)
				return
			}

			if forceStop {
				database.UpdatePID(Api.DbConn, 0, result.Id)
				database.UpdateState(Api.DbConn, "stop", result.Id)
				return
			}

			procces, err := os.FindProcess(result.Pid)

			if err != nil {
				fmt.Println("Error Find PID: ", err)
				return
			}

			err = procces.Kill()

			if err != nil {
				fmt.Println("Error Close Process: ", err)
				return
			}

			database.UpdatePID(Api.DbConn, 0, result.Id)
			database.UpdateState(Api.DbConn, "stop", result.Id)

			fmt.Println("Proceso detenido exitosamente")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	stopCmd.Flags().BoolVarP(&forceStop, "force", "f", false, "Fuerza el cierre de la Api (DEBUG)")
}
