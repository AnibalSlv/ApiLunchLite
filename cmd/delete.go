/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Borrar una API",
	Long:  `apl delete [Name API]`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Api.Name = args[0]

		apiMgr.Delete(Api.Name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
