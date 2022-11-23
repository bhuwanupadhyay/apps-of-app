package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var tiltCmd = &cobra.Command{
	Use:   "tilt",
	Short: "Tilt commands",
}

var tiltUpCmd = buildTiltCmd("up")
var tiltDownCmd = buildTiltCmd("down")
var tiltDir = filepath.Join(baseDir, "compose", "Tiltfiles")
var tiltListCmd = &cobra.Command{
	Use:   "list <app>",
	Short: "List tilt available apps",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		_ = filepath.Walk(tiltDir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Println("Available Applications")
				fmt.Printf("- %s\n", strings.TrimSuffix(info.Name(), ".Tiltfile"))
			}
			return nil
		})
	},
}

func buildTiltCmd(command string) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("%s <app_name>", command),
		Short: fmt.Sprintf("To run tilt %s for app <app_name>", command),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires <app_name> arguments")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			tiltFile := filepath.Join(tiltDir, fmt.Sprintf("%s.Tiltfile", args[0]))
			ExecuteCommand(CommandRequest{"tilt", []string{command, "--file", tiltFile}})
		},
	}
}

func init() {
	rootCmd.AddCommand(tiltCmd)
	tiltCmd.AddCommand(tiltUpCmd)
	tiltCmd.AddCommand(tiltDownCmd)
	tiltCmd.AddCommand(tiltListCmd)
}
