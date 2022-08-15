package exp

import (
	"errors"

	"github.com/aiven/aiven-go-client"
	"github.com/mitchellh/copystructure"
)

var (
	// errUnexpected is an error that is returned when an unexpected error occurs.
	errUnexpected = errors.New("unexpected filtering error")
)

// supportedServiceTypes is a pseudo-constant map of supported service types.
var supportedServiceTypes = func() map[string]struct{} {
	return map[string]struct{}{
		"cassandra":         {},
		"clickhouse":        {},
		"elasticsearch":     {},
		"flink":             {},
		"grafana":           {},
		"influxdb":          {},
		"kafka":             {},
		"kafka_connect":     {},
		"kafka_mirrormaker": {},
		"m3aggregator":      {},
		"m3db":              {},
		"mysql":             {},
		"opensearch":        {},
		"pg":                {},
		"redis":             {},
	}
}

// filterServiceTypeList is a filter that filters out service types that are not supported.
func filterServiceTypes(f map[string]aiven.ServiceType) (map[string]aiven.ServiceType, error) {
	cf, err := copystructure.Copy(f)
	if err != nil {
		return nil, err
	}

	acf, ok := cf.(map[string]aiven.ServiceType)
	if !ok {
		return nil, errUnexpected
	}

	for k := range acf {
		if _, ok = supportedServiceTypes()[k]; !ok {
			delete(acf, k)
		}
	}

	return acf, nil
}
