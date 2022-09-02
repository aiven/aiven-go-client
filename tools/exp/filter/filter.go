package filter

import (
	"errors"

	"github.com/aiven/aiven-go-client"
	"github.com/mitchellh/copystructure"
	"golang.org/x/exp/maps"
)

// errUnexpected is the error that is returned when an unexpected error occurs.
var errUnexpected = errors.New("unexpected filtering error")

// supportedServiceTypes is a pseudo-constant map of supported service types.
var supportedServiceTypes = func() map[string]struct{} {
	return map[string]struct{}{
		"cassandra":                       {},
		"clickhouse":                      {},
		"datadog":                         {},
		"elasticsearch":                   {},
		"external_aws_cloudwatch_logs":    {},
		"external_aws_cloudwatch_metrics": {},
		"external_elasticsearch_logs":     {},
		"external_google_cloud_logging":   {},
		"external_kafka":                  {},
		"external_opensearch_logs":        {},
		"external_postgresql":             {},
		"flink":                           {},
		"grafana":                         {},
		"influxdb":                        {},
		"jolokia":                         {},
		"kafka":                           {},
		"kafka_connect":                   {},
		"kafka_mirrormaker":               {},
		"m3aggregator":                    {},
		"m3coordinator":                   {},
		"m3db":                            {},
		"mysql":                           {},
		"opensearch":                      {},
		"pg":                              {},
		"prometheus":                      {},
		"redis":                           {},
		"rsyslog":                         {},
		"signalfx":                        {},
	}
}

// ServiceTypes is a filter that filters out service types that are not supported.
func ServiceTypes(f map[string]aiven.ServiceType) (map[string]aiven.ServiceType, error) {
	cf, err := copystructure.Copy(f)
	if err != nil {
		return nil, err
	}

	acf, ok := cf.(map[string]aiven.ServiceType)
	if !ok {
		return nil, errUnexpected
	}

	maps.DeleteFunc(acf, func(k string, _ aiven.ServiceType) bool {
		_, ok = supportedServiceTypes()[k]
		return !ok
	})

	return acf, nil
}

// IntegrationTypes is a filter that filters out integration types that are not supported.
func IntegrationTypes(f []aiven.IntegrationType) ([]aiven.IntegrationType, error) {
	cf, err := copystructure.Copy(f)
	if err != nil {
		return nil, err
	}

	acf, ok := cf.([]aiven.IntegrationType)
	if !ok {
		return nil, errUnexpected
	}

	var nf []aiven.IntegrationType

	for _, v := range acf {
		{
			var dst []string

			for _, vn := range v.DestServiceTypes {
				if _, ok = supportedServiceTypes()[vn]; ok {
					dst = append(dst, vn)
				}
			}

			if len(dst) == 0 {
				break
			}

			v.DestServiceTypes = dst
		}

		{
			var sst []string

			for _, vn := range v.SourceServiceTypes {
				if _, ok = supportedServiceTypes()[vn]; ok {
					sst = append(sst, vn)
				}
			}

			if len(sst) == 0 {
				break
			}

			v.SourceServiceTypes = sst
		}

		nf = append(nf, v)
	}

	return nf, nil
}

// IntegrationEndpointTypes is a filter that filters out integration endpoint types that are not supported.
func IntegrationEndpointTypes(f []aiven.IntegrationEndpointType) ([]aiven.IntegrationEndpointType, error) {
	cf, err := copystructure.Copy(f)
	if err != nil {
		return nil, err
	}

	acf, ok := cf.([]aiven.IntegrationEndpointType)
	if !ok {
		return nil, errUnexpected
	}

	var nf []aiven.IntegrationEndpointType

	for _, v := range acf {
		var st []string

		for _, vn := range v.ServiceTypes {
			if _, ok = supportedServiceTypes()[vn]; ok {
				st = append(st, vn)
			}
		}

		if len(st) == 0 {
			break
		}

		v.ServiceTypes = st

		nf = append(nf, v)
	}

	return nf, nil
}
