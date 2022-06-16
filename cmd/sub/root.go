package sub

import "github.com/spf13/cobra"

var cfgFile string
var output string
var format string

var rootCmd = &cobra.Command{
	Use: "godeng",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
