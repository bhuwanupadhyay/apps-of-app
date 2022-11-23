package cmd

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

var cfgFile string

var colorYellow = "\x1b[33;1m"
var colorGreen = "\x1b[32;1m"
var colorNormal = "\x1b[0m"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "aoa",
}

var baseDir string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(workingDir string) {
	baseDir = workingDir
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	initConfig()
	rootCmd.SilenceErrors = true
	rootCmd.SetUsageTemplate(colorYellow + `Usage:` + colorNormal + `{{if .Runnable}}
` + colorGreen + `{{.UseLine}}` + colorNormal + `{{end}}{{if .HasAvailableSubCommands}}
  ` + colorGreen + `{{.CommandPath}}` + colorNormal + ` [command]{{end}}{{if gt (len .Aliases) 0}}
` + colorYellow + `Aliases:` + colorNormal + `
  {{.NameAndAliases}}{{end}}{{if .HasExample}}
` + colorYellow + `Examples:` + colorNormal + `
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}
` + colorYellow + `Available Commands:` + colorNormal + `{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  ` + colorGreen + `{{rpad .Name .NamePadding }}` + colorNormal + ` {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}
` + colorYellow + `Flags:` + colorNormal + `
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
` + colorYellow + `Global Flags:` + colorNormal + `
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "~/.aoa/config.json", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

type CommandRequest struct {
	Command string
	Args    []string
}

func ExecuteCommand(cr CommandRequest) {
	cmd := exec.Command(cr.Command, cr.Args...)

	stdout, _ := cmd.StdoutPipe()
	_ = cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	_ = cmd.Wait()
}
