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
var loop bool
var file string

var rootCmd = &cobra.Command{
	Use: "godeng",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "", "stdout", "output")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "", "json", "format")
	rootCmd.PersistentFlags().Int64VarP(&count, "count", "", 100, "count")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "", "", "http request url")
	rootCmd.PersistentFlags().StringVarP(&tablename, "tablename", "", "", "tablename")
	rootCmd.PersistentFlags().Int64VarP(&sleep, "sleep", "", 0, "sleep")
	rootCmd.PersistentFlags().BoolVarP(&loop, "loop", "", false, "loop")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "", "", "output file")
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
	godeng.MakeGoDeng(config, output, format, count, loop, sleep, url, tablename, file).Start()
}
