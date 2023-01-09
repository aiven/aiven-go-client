package util

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/aiven/aiven-go-client"
)

const (
	// logSeparator is the separator for the log.
	logSeparator = " "

	// logFlag is the flag for the log.
	logFlag = log.LstdFlags | log.LUTC | log.Lmsgprefix
)

// EnvAivenProjectName is the environment variable for the Aiven project name.
const EnvAivenProjectName = "AIVEN_PROJECT_NAME"

const (
	// ServiceTypesFilename is the name of the service types file.
	ServiceTypesFilename = "service_types.yml"

	// IntegrationTypesFilename is the name of the integration types file.
	IntegrationTypesFilename = "integration_types.yml"

	// IntegrationEndpointTypesFilename is the name of the integration endpoint types file.
	IntegrationEndpointTypesFilename = "integration_endpoint_types.yml"
)

// Logger is a struct that holds the loggers for the application.
type Logger struct {
	// Info is the logger for info messages.
	Info *log.Logger

	// Error is the logger for error messages.
	Error *log.Logger
}

// EnvMap is a type for a map of environment variables.
type EnvMap map[string]string

// SetupLogger sets up the logger.
func SetupLogger(logger *Logger) {
	logger.Info = log.New(os.Stdout, "[INFO]"+logSeparator, logFlag)
	logger.Error = log.New(os.Stderr, "[ERROR]"+logSeparator, logFlag)
}

// SetupEnv populates the provided environment variables map.
func SetupEnv(env EnvMap) error {
	for k := range env {
		ev, ok := os.LookupEnv(k)
		if !ok {
			return fmt.Errorf("environment variable is not set: %s", k)
		}

		env[k] = ev
	}

	return nil
}

// SetupClient sets up the Aiven client.
func SetupClient(client *aiven.Client) error {
	c, err := aiven.SetupEnvClient("aiven-go-client/exp")
	if err != nil {
		return err
	}

	*client = *c

	return nil
}

// MeasureExecutionTime prints the execution time of the caller when deferred.
func MeasureExecutionTime(logger *Logger) func() {
	start := time.Now()

	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("runtime.Caller failed")
	}

	fn := runtime.FuncForPC(pc)

	return func() {
		logger.Info.Printf("%s took %dms", fn.Name(), time.Since(start).Milliseconds())
	}
}

// Ref returns the reference (pointer) of the provided value.
func Ref[T any](v T) *T {
	return &v
}
