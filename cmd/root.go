package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "JavaScript Module analyzer",
	Short: "Js dependency analyzer",
	Long:  `A dependency analyzer for javascript code`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags can be defined here
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fileapp.yaml)")
}


// add flag for file path
// add flag for depth analysis
