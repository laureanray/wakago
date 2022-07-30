/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"wakago/api"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get stats (goals, editors, etc)",
	Long: `Get statistics from WakaTime API
  See: https://wakatime.com/developers#introduction
  `,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Option on how should this be look
// Example: 1-Liner
//          2-Liner
//          Pretty
//          Custom? [Pass formatting string]
var getGoalsCmd = &cobra.Command{
	Use:   "goals [output type]",
	Short: "Print to the terminal your current goals",
	Long:  `Print to the terminal your current goals and choose the look and feel of the output`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Get goals")
		wt := api.GetInstance()
		fmt.Print(wt.GetGoals())
	},
}

func init() {
	getCmd.AddCommand(getGoalsCmd)
	rootCmd.AddCommand(getCmd)
}
