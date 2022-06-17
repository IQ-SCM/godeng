package sub

import (
	"log"
	"os"

	"github.com/chenjiayao/godeng"
	"github.com/spf13/cobra"
)

var cfgFile string
var output string
var format string
var count int64
var url string
var tablename string
var sleep int64
var forever bool
var file string

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
	rootCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "url")
	rootCmd.PersistentFlags().StringVarP(&tablename, "tablename", "t", "", "tablename")
	rootCmd.PersistentFlags().Int64VarP(&sleep, "sleep", "s", 0, "sleep")
	rootCmd.PersistentFlags().BoolVarP(&forever, "forever", "", false, "forever")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "", "", "file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() {
	config, err := godeng.Parser(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	godeng.MakeGoDeng(config, output, format, count, forever, sleep, url, tablename, file).Start()
}
