/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"apiLunchLite/internal/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Muestra las APIs guardadas",
	Long:  `Muestra las APIs guardadas y su estado`,
	Run: func(cmd *cobra.Command, args []string) {
		apiMgr.CheckLife()
		result, err := apiMgr.Check()

		if err != nil {
			log.Fatal("Error: ", err)
		}

		if len(result) == 0 {
			fmt.Println("No hay APIs registradas. Usa el comando 'addApi' para empezar.")
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		t.AppendHeader(table.Row{"ID", "Nombre", "Host", "State", "PID"})

		for _, api := range result {
			displayState := utils.Capitalize(api.State)
			switch api.State {
			case "stop":
				displayState = utils.Stop(displayState)
			case "run":
				displayState = utils.Run(displayState)
			default:
				displayState = utils.Yellow(displayState)
			}

			t.AppendRow(table.Row{
				api.Id,
				api.Name,
				fmt.Sprintf("%s:%s", api.Host, utils.Yellow(strconv.Itoa(api.Port))),
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
	rootCmd.AddCommand(checkCmd)

}
