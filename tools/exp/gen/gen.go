package gen

import (
	"os"
	"strings"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/aiven-go-client/tools/exp/filter"
	"github.com/aiven/aiven-go-client/tools/exp/util"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v3"
)

const (
	// generating is a part of the message that is printed when the generation process starts.
	generating = "generating %s"
)

const (
	// serviceTypesFilename is the name of the service types file.
	serviceTypesFilename = "service_types.yml"

	// integrationTypesFilename is the name of the integration types file.
	integrationTypesFilename = "integration_types.yml"

	// integrationEndpointTypesFilename is the name of the integration endpoint types file.
	integrationEndpointTypesFilename = "integration_endpoint_types.yml"
)

// logger is a pointer to the logger.
var logger *util.Logger

// flags is a pointer to the flags.
var flags *pflag.FlagSet

// env is a map of environment variables.
var env util.EnvMap

// client is a pointer to the Aiven client.
var client *aiven.Client

// write is a function that writes the generated file to the specified path.
func write(filename string, in interface{}) error {
	outputDir, err := flags.GetString("output-dir")
	if err != nil {
		return err
	}

	f, err := os.Create(strings.Join([]string{outputDir, filename}, string(os.PathSeparator)))
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err = f.Close()
	}(f)

	e := yaml.NewEncoder(f)

	defer func(e *yaml.Encoder) {
		err = e.Close()
	}(e)

	e.SetIndent(2)

	if err = e.Encode(in); err != nil {
		return err
	}

	return err
}

// serviceTypes generates the service types file.
func serviceTypes() error {
	defer util.MeasureExecutionTime(logger)()

	logger.Info.Printf(generating, serviceTypesFilename)

	result, err := client.Projects.ServiceTypes(env[util.EnvAivenProjectName])
	if err != nil {
		return err
	}

	filtered, err := filter.ServiceTypes(result)
	if err != nil {
		return err
	}

	out := map[string]interface{}{}

	for k, v := range filtered {
		out[k] = v.UserConfigSchema
	}

	return write(serviceTypesFilename, out)
}

// integrationTypes generates the integration types file.
func integrationTypes() error {
	defer util.MeasureExecutionTime(logger)()

	logger.Info.Printf(generating, integrationTypesFilename)

	result, err := client.Projects.IntegrationTypes(env[util.EnvAivenProjectName])
	if err != nil {
		return err
	}

	filtered, err := filter.IntegrationTypes(result)
	if err != nil {
		return err
	}

	out := map[string]interface{}{}

	for _, v := range filtered {
		out[v.IntegrationType] = v.UserConfigSchema
	}

	return write(integrationTypesFilename, out)
}

// integrationEndpointTypes generates the integration endpoint types file.
func integrationEndpointTypes() error {
	defer util.MeasureExecutionTime(logger)()

	logger.Info.Printf(generating, integrationEndpointTypesFilename)

	result, err := client.Projects.IntegrationEndpointTypes(env[util.EnvAivenProjectName])
	if err != nil {
		return err
	}

	filtered, err := filter.IntegrationEndpointTypes(result)
	if err != nil {
		return err
	}

	out := map[string]interface{}{}

	for _, v := range filtered {
		out[v.EndpointType] = v.UserConfigSchema
	}

	return write(integrationEndpointTypesFilename, out)
}

// setup sets up the generation process.
func setup(l *util.Logger, f *pflag.FlagSet, e util.EnvMap, c *aiven.Client) {
	logger = l
	flags = f
	env = e
	client = c
}

// Run executes the generation process.
func Run(ctx context.Context, logger *util.Logger, flags *pflag.FlagSet, env util.EnvMap, client *aiven.Client) error {
	setup(logger, flags, env, client)

	errs, _ := errgroup.WithContext(ctx)

	errs.Go(serviceTypes)
	errs.Go(integrationTypes)
	errs.Go(integrationEndpointTypes)

	return errs.Wait()
}
