// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
)

type (
	// NewServiceIntegration defines partial set of service integration fields used
	// when passing integration as part of service creation call
	NewServiceIntegration struct {
		DestinationEndpointID *string                `json:"dest_endpoint_id"`
		DestinationService    *string                `json:"dest_service"`
		IntegrationType       string                 `json:"integration_type"`
		SourceService         *string                `json:"source_service"`
		SourceEndpointID      *string                `json:"source_endpoint_id"`
		UserConfig            map[string]interface{} `json:"user_config"`
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
		IntegrationType       string                 `json:"integration_type"`
		SourceService         *string                `json:"source_service,omitempty"`
		SourceEndpointID      *string                `json:"source_endpoint_id,omitempty"`
		UserConfig            map[string]interface{} `json:"user_config,omitempty"`
	}

	// UpdateServiceIntegrationRequest are the parameters to update a Service Integration.
	UpdateServiceIntegrationRequest struct {
		UserConfig map[string]interface{} `json:"user_config,omitempty"`
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
	project string,
	req CreateServiceIntegrationRequest,
) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration")
	rsp, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// Get a specific service integration endpoint from Aiven.
func (h *ServiceIntegrationsHandler) Get(project, integrationID string) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration", integrationID)
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// Update the given service integration with the given parameters.
func (h *ServiceIntegrationsHandler) Update(
	project string,
	integrationID string,
	req UpdateServiceIntegrationRequest,
) (*ServiceIntegration, error) {
	path := buildPath("project", project, "integration", integrationID)
	rsp, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// Delete the given service integration from Aiven.
func (h *ServiceIntegrationsHandler) Delete(project, integrationID string) error {
	path := buildPath("project", project, "integration", integrationID)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// List all service integration for a given project and service.
func (h *ServiceIntegrationsHandler) List(project, service string) ([]*ServiceIntegration, error) {
	path := buildPath("project", project, "service", service, "integration")
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var response *ServiceIntegrationListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.ServiceIntegrations, nil
}

func parseServiceIntegrationResponse(rsp []byte) (*ServiceIntegration, error) {
	var response *ServiceIntegrationResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.ServiceIntegration, nil
}
