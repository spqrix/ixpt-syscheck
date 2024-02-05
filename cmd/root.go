/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	// "log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ixpt-syscheck [...commands]",
	Short: "A simple way to check for dependency applications on your system.",
	Long: `Sometimes our apps rely on other apps in order to function properly.
Check to ensure they are on the system, as well as configs and other files.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		operating_system := runtime.GOOS
		if operating_system != "linux" {
			err := errors.New("This command is for Linux only.")
			fmt.Println(err)
			os.Exit(1)
		}
		for _, c := range args {
			script_command := fmt.Sprintf("type -ap %s", c)
			command := exec.Command("sh", "-c", script_command)
			var out strings.Builder
			command.Stdout = &out
			err := command.Run()
			if err != nil {
				fmt.Printf("Failed to locate file path for: %s\n", c)
			} else {
				if cmd.Flags().Lookup("verbose").Changed {
					fmt.Printf("Path for %s:\n\t%s", c,out.String())
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ixpt-syscheck.yaml)")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose output")
}


