package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of CLI",
	Long:  `Print the version number of CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1 -- HEAD")
	},
}
