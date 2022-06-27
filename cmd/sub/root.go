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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "./dodeng.json", "config file")

	rootCmd.PersistentFlags().StringVarP(&output, "output", "", "stdout", "output")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "", "json", "output format")
	rootCmd.PersistentFlags().Int64VarP(&count, "count", "", 100, "count")
	rootCmd.PersistentFlags().StringVarP(&url, "url", "", "", "http request url,only used when output is http/https and format is json")
	rootCmd.PersistentFlags().StringVarP(&tablename, "tablename", "", "godeng", "tablename, only used when output is sql")
	rootCmd.PersistentFlags().Int64VarP(&sleep, "sleep", "", 0, "fix creation time interval for each log (second)")
	rootCmd.PersistentFlags().BoolVarP(&loop, "loop", "", false, "loop output forever until killed. if loop is set, then count is ignored")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "", "./godeng.out", "output file, only used when output is file")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() {
	log.SetFlags(0)
	config, err := godeng.Parser(cfgFile)
	if err != nil {
		log.Fatal(err)
	}
	godeng, err := godeng.MakeGoDeng(config, output, format, count, loop, sleep, url, tablename, file)
	if err != nil {
		log.Fatal(err)
	}
	godeng.Start()
}
