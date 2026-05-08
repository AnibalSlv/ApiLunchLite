/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Muestra las APIs guardadas",
	Long:  `Muestra las APIs guardadas y su estado`,
	Run: func(cmd *cobra.Command, args []string) {
		apiMgr.Check()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

}
