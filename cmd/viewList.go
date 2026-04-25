/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"apiLunchLite/internal/database"
	"apiLunchLite/internal/utils"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var viewListCmd = &cobra.Command{
	Use:   "viewList",
	Short: "Muestra las APIs guardadas",
	Long:  `Muestra las APIs guardadas y su estado`,
	Run: func(cmd *cobra.Command, args []string) {

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		result, err := database.GetAll(Api.DbConn)

		if err != nil {
			log.Fatal("Error:", err)
		}

		if len(result) == 0 {
			fmt.Println("No hay APIs registradas. Usa el comando 'addApi' para empezar.")
			return
		}

		t.AppendHeader(table.Row{"ID", "Nombre", "Host", "State", "PID"})

		for _, api := range result {
			displayState := utils.Capitalize(api.State)
			switch api.State {
			case "stop":
				displayState = color.RedString(displayState)
			case "run":
				displayState = color.GreenString(displayState)
			default:
				displayState = color.YellowString(displayState)
			}

			t.AppendRow(table.Row{
				api.Id,
				api.Name,
				fmt.Sprintf("%s:%d", api.Host, api.Port),
				displayState,
				api.Pid,
			})
		}

		t.SetStyle(table.StyleRounded)

		// Convierte el texto en el header de la tabla en negrita
		t.Style().Color.Header = text.Colors{text.Bold}
		t.Render()
	},
}

func init() {
	rootCmd.AddCommand(viewListCmd)

}
