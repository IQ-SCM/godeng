package sub

import (
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var output string
var format string
var count int64

var rootCmd = &cobra.Command{
	Use: "godeng",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "stdout", "output")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "json", "format")
	rootCmd.PersistentFlags().Int64VarP(&count, "count", "n", 100, "count")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() {
}
