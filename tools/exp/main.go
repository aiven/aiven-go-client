package main

import (
	"github.com/aiven/aiven-go-client/tools/exp/cmd"
	"github.com/aiven/aiven-go-client/tools/exp/util"
	"github.com/spf13/cobra"
)

// logger is the logger of the application.
var logger = &util.Logger{}

// rootCmd is the root command for the application.
var rootCmd *cobra.Command

// setup sets up the application.
func setup() {
	util.SetupLogger(logger)

	rootCmd = cmd.NewCmdRoot(logger)
}

// main is the entrypoint for the application.
func main() {
	setup()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
