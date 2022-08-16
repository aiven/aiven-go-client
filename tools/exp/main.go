package main

import (
	"github.com/aiven/aiven-go-client/tools/exp/cmd"
	"github.com/aiven/aiven-go-client/tools/exp/util"
)

// main is the entrypoint for the application.
func main() {
	var logger = &util.Logger{}

	util.SetupLogger(logger)

	cmd.Execute(logger)
}
