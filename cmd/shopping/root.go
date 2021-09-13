package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var withVersion bool

var version = "1.0"
var DATA_SOURCE_NAME string
var LOG_FILE string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "use",
	Short: "Passes a variable.",
	Long:  `Passes a variable.`,
	Run: func(c *cobra.Command, args []string) {

		fmt.Println("DATA_SOURCE_NAME:", DATA_SOURCE_NAME)
		fmt.Println("LOG_FILE:", LOG_FILE)

		if withVersion {
			fmt.Printf("Go Version %v\nApplication Version %v\n", runtime.Version(), version)
		}

		os.Setenv("DATA_SOURCE_NAME", DATA_SOURCE_NAME)
		os.Setenv("LOG_FILE", LOG_FILE)

	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	// Add long option `--version` and short option `-v`.
	// default is flase
	RootCmd.Flags().BoolVarP(&withVersion, "version", "v", false, "version")
	RootCmd.Flags().StringVar(&DATA_SOURCE_NAME, "DATA_SOURCE_NAME", "", "Connecting to mysql")
	RootCmd.Flags().StringVar(&LOG_FILE, "LOG_FILE", "", "LOG_FILE")

}
