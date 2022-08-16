package cmd

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/aiven-go-client/tools/exp/gen"
	"github.com/aiven/aiven-go-client/tools/exp/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "exp",
	Short: "exp is a tool for generating and persisting user configuration option schemas from Aiven APIs.",
	Run:   run,
}

// logger is the logger of the application.
var logger = &util.Logger{}

// env is a map of environment variables.
var env = util.EnvMap{
	util.EnvAivenProjectName: "",
}

// client is a pointer to the Aiven client.
var client = &aiven.Client{}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(l *util.Logger) {
	logger = l

	rootCmd.Flags().StringP("output-dir", "o", "", "the output directory for the generated files")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

// setupOutputDir sets up the output directory.
func setupOutputDir(flags *pflag.FlagSet) error {
	outputDir, err := flags.GetString("output-dir")
	if err != nil {
		return err
	}

	if outputDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}

		outputDir = strings.Join([]string{wd, "tools/exp/dist"}, string(os.PathSeparator))

		err = flags.Set("output-dir", outputDir)
		if err != nil {
			return err
		}
	}

	fi, err := os.Stat(outputDir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return errors.New("output directory is not a directory")
	}

	return nil
}

// setup sets up the application.
func setup(flags *pflag.FlagSet) {
	logger.Info.Println("exp tool started")

	logger.Info.Println("setting up output directory")

	if err := setupOutputDir(flags); err != nil {
		logger.Error.Fatalf("error setting up output directory: %s", err)
	}

	logger.Info.Println("setting up environment variables")

	if err := util.SetupEnv(env); err != nil {
		logger.Error.Fatalf("error setting up environment variables: %s", err)
	}

	logger.Info.Println("setting up client")

	if err := util.SetupClient(client); err != nil {
		logger.Error.Fatalf("error setting up client: %s", err)
	}
}

// run is the function that is called when the rootCmd is executed.
func run(cmd *cobra.Command, _ []string) {
	setup(cmd.Flags())

	ctx := context.Background()

	logger.Info.Println("generating files")

	if err := gen.Run(ctx, logger, cmd.Flags(), env, client); err != nil {
		logger.Error.Fatalf("error generating files: %s", err)
	}

	logger.Info.Println("done")
}
