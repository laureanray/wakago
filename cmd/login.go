package cmd

import (
	"wakago/api"
	"wakago/server"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		wt := api.GetInstance()
		wt.Login()
		s := server.GetInstance()

		s.Init()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
