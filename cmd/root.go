package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/spqrix/ixpt-syscheck/executables"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ixpt-syscheck [...commands]",
	Short: "A simple way to check for dependency executables (commands) on your system.",
	Long: `Sometimes our apps rely on other executables (commands) in order to function properly.
Check to ensure they are on the system.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		verb := cmd.Flags().Lookup("verbose").Changed
		missing_execs, err := executables.CheckForExecutables(args, verb)
		if err != nil {
			fmt.Printf("Error during system check: %q", err)
			os.Exit(1)
		}
		for _, c := range missing_execs {
			fmt.Printf("Failed to locate file path for: %s\n", c)
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


