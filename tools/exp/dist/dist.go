package dist

import _ "embed"

var (
	// ServiceTypes is the service_types.yml file embedded in the binary.
	//
	//go:embed service_types.yml
	ServiceTypes []byte

	// IntegrationTypes is the integration_types.yml file embedded in the binary.
	//
	//go:embed integration_types.yml
	IntegrationTypes []byte

	// IntegrationEndpointTypes is the integration_endpoint_types.yml file embedded in the binary.
	//
	//go:embed integration_endpoint_types.yml
	IntegrationEndpointTypes []byte
)
