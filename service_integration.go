package aiven

import "context"

type (
	// NewServiceIntegration defines partial set of service integration fields used
	// when passing integration as part of service creation call
	NewServiceIntegration struct {
		DestinationEndpointID *string                `json:"dest_endpoint_id"`
		DestinationService    *string                `json:"dest_service"`
		IntegrationType       string                 `json:"integration_type"`
		SourceService         *string                `json:"source_service"`
		SourceEndpointID      *string                `json:"source_endpoint_id"`
		UserConfig            map[string]interface{} `json:"user_config,omitempty"`
	}

	// ServiceIntegration represents a service integration endpoint,
	// like parameters for integration to Datadog
	ServiceIntegration struct {
		Active                  bool                   `json:"active"`
		Description             string                 `json:"description"`
		DestinationProject      *string                `json:"dest_project"`
		DestinationService      *string                `json:"dest_service"`
		DestinationEndpointID   *string                `json:"dest_endpoint_id"`
		DestinationEndpointName *string                `json:"dest_endpoint"`
		DestinationServiceType  *string                `json:"dest_service_type"`
		Enabled                 bool                   `json:"enabled"`
		IntegrationType         string                 `json:"integration_type"`
		IntegrationStatus       map[string]interface{} `json:"integration_status"`
		ServiceIntegrationID    string                 `json:"service_integration_id"`
		SourceProject           *string                `json:"source_project"`
		SourceService           *string                `json:"source_service"`
		SourceEndpointID        *string                `json:"source_endpoint_id"`
		SourceEndpointName      *string                `json:"source_endpoint"`
		SourceServiceType       *string                `json:"source_service_type"`
		UserConfig              map[string]interface{} `json:"user_config"`
	}

	// ServiceIntegrationsHandler is the client that interacts
	// with the Service Integration Endpoints API endpoints on Aiven.
	ServiceIntegrationsHandler struct {
		client *Client
	}

	// CreateServiceIntegrationRequest are the parameters to create a Service Integration.
	CreateServiceIntegrationRequest struct {
		DestinationService    *string                `json:"dest_service,omitempty"`
		DestinationEndpointID *string                `json:"dest_endpoint_id,omitempty"`
		DestinationProject    *string                `json:"dest_project,omitempty"`
		IntegrationType       string                 `json:"integration_type"`
		SourceService         *string                `json:"source_service,omitempty"`
		SourceProject         *string                `json:"source_project,omitempty"`
		SourceEndpointID      *string                `json:"source_endpoint_id,omitempty"`
		UserConfig            map[string]interface{} `json:"user_config,omitempty"`
	}

	// UpdateServiceIntegrationRequest are the parameters to update a Service Integration.
	UpdateServiceIntegrationRequest struct {
		UserConfig map[string]interface{} `json:"user_config"`
	}

	// ServiceIntegrationResponse represents the response from Aiven
	// after interacting with the Service Integration API.
	ServiceIntegrationResponse struct {
		APIResponse
		ServiceIntegration *ServiceIntegration `json:"service_integration"`
	}

	// ServiceIntegrationListResponse represents the response from Aiven
	// for listing service integrations.
	ServiceIntegrationListResponse struct {
		APIResponse
		ServiceIntegrations []*ServiceIntegration `json:"service_integrations"`
	}
)

// Create the given Service Integration on Aiven.
func (h *ServiceIntegrationsHandler) Create(
	ctx context.Context,
	project string,
	req CreateServiceIntegrationRequest,
) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceIntegrationResponse
	errR := checkAPIResponse(bts, &r)

	return r.ServiceIntegration, errR
}

// Get a specific service integration endpoint from Aiven.
func (h *ServiceIntegrationsHandler) Get(ctx context.Context, project, integrationID string) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration", integrationID)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceIntegrationResponse
	errR := checkAPIResponse(bts, &r)

	return r.ServiceIntegration, errR
}

// Update the given service integration with the given parameters.
func (h *ServiceIntegrationsHandler) Update(
	ctx context.Context,
	project string,
	integrationID string,
	req UpdateServiceIntegrationRequest,
) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration", integrationID)
	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceIntegrationResponse
	errR := checkAPIResponse(bts, &r)

	return r.ServiceIntegration, errR
}

// Delete the given service integration from Aiven.
func (h *ServiceIntegrationsHandler) Delete(ctx context.Context, project, integrationID string) error {
	path := buildPath("project", project, "integration", integrationID)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List all service integration for a given project and service.
func (h *ServiceIntegrationsHandler) List(ctx context.Context, project, service string) ([]*ServiceIntegration, error) {
	path := buildPath("project", project, "service", service, "integration")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceIntegrationListResponse
	errR := checkAPIResponse(bts, &r)

	return r.ServiceIntegrations, errR
}
