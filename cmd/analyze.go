package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var filePath string
var outputFormat string
var detectCircularDependency bool
var depth int16

func init() {
	analyzeCmd := &cobra.Command{
		Use:   "analyze",
		Short: "analyze file/folder dependencies",
		Long:  `Analyzes the javascript dependencies of the provided path.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from analyzer")
			// TODO - load git ignore file and parse the extensions

			// - go the provided path and traverse it to the provided depth while printing every visited file

			
		},
	}

	// Add flags to the open command
	analyzeCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the file to be analyzed (required)")
	analyzeCmd.Flags().StringVarP(&outputFormat, "format", "o", "text", "Output format (tree, UI)")
	analyzeCmd.Flags().BoolVarP(&detectCircularDependency, "circular", "c", false, "if set to true will output circular dependencies")
	analyzeCmd.Flags().Int16Var(&depth, "depth", 2, "Represents the depth on the analysis, if set to -1 no limit is applied")

	// Mark file flag as required
	analyzeCmd.MarkFlagRequired("file")

	// Add the open command to the root command
	rootCmd.AddCommand(analyzeCmd)
}
