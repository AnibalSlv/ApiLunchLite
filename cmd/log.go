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

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Muestra los log de la API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			Api.Name = args[0]

			result, err := database.GetName(Api.DbConn, Api.Name)

			if err != nil {
				fmt.Print("Error Connection Db: ", err)
				return
			}

			fileLog, err := os.ReadFile("internal/logs/" + result.Name + ".log")

			if err != nil {
				fmt.Println("Error Read Log: ", err)
				return
			}

			fmt.Print(string(fileLog))
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
