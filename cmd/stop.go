/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var forceStop bool

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Detiene el daemon de la API",
	Long: `Detiene el daemon de la API:
			apl stop [nombre]`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		Api.Name = args[0]

		apiMgr.Stop(Api.Name, forceStop)

	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	stopCmd.Flags().BoolVarP(&forceStop, "force", "f", false, "Fuerza el cierre de la Api (DEBUG)")
}
