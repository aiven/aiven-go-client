// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	// ServiceIntegrationEndpoint represents a service integration endpoint,
	// like parameters for integration to Datadog
	ServiceIntegrationEndpoint struct {
		EndpointID     string                 `json:"endpoint_id"`
		EndpointName   string                 `json:"endpoint_name"`
		EndpointType   string                 `json:"endpoint_type"`
		UserConfig     map[string]interface{} `json:"user_config"`
		EndpointConfig map[string]interface{} `json:"endpoint_config"`
	}

	// ServiceIntegrationEndpointsHandler is the client that interacts
	// with the Service Integration Endpoints API endpoints on Aiven.
	ServiceIntegrationEndpointsHandler struct {
		client *Client
	}

	// CreateServiceIntegrationEndpointRequest are the parameters to create
	// a Service Integration Endpoint.
	CreateServiceIntegrationEndpointRequest struct {
		EndpointName string                 `json:"endpoint_name"`
		EndpointType string                 `json:"endpoint_type"`
		UserConfig   map[string]interface{} `json:"user_config,omitempty"`
	}

	// UpdateServiceIntegrationEndpointRequest are the parameters to update
	// a Service Integration Endpoint.
	UpdateServiceIntegrationEndpointRequest struct {
		UserConfig map[string]interface{} `json:"user_config,omitempty"`
	}

	// ServiceIntegrationEndpointResponse represents the response from Aiven
	// after interacting with the Service Integration Endpoints API.
	ServiceIntegrationEndpointResponse struct {
		APIResponse
		ServiceIntegrationEndpoint *ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
	}

	// ServiceIntegrationEndpointListResponse represents the response from Aiven
	// for listing service integration endpoints.
	ServiceIntegrationEndpointListResponse struct {
		APIResponse
		ServiceIntegrationEndpoints []*ServiceIntegrationEndpoint `json:"service_integration_endpoints"`
	}
)

// Create the given Service Integration Endpoint on Aiven.
func (h *ServiceIntegrationEndpointsHandler) Create(
	project string,
	req CreateServiceIntegrationEndpointRequest,
) (*ServiceIntegrationEndpoint, error) {
	path := buildPath("project", project, "integration_endpoint")
	rsp, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationEndpointResponse(rsp)
}

// Get a specific service integration endpoint from Aiven.
func (h *ServiceIntegrationEndpointsHandler) Get(project, endpointID string) (*ServiceIntegrationEndpoint, error) {
	// There's no API for getting integration endpoint by ID. List all endpoints
	// and pick the correct one instead. (There shouldn't ever be many endpoints.)
	endpoints, err := h.List(project)
	if err != nil {
		return nil, err
	}

	for _, endpoint := range endpoints {
		if endpoint.EndpointID == endpointID {
			return endpoint, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Integration endpoint with ID %v not found", endpointID), Status: 404}
	return nil, err
}

// Update the given service integration endpoint with the given parameters.
func (h *ServiceIntegrationEndpointsHandler) Update(
	project string,
	endpointID string,
	req UpdateServiceIntegrationEndpointRequest,
) (*ServiceIntegrationEndpoint, error) {
	path := buildPath("project", project, "integration_endpoint", endpointID)
	rsp, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationEndpointResponse(rsp)
}

// Delete the given service integration endpoint from Aiven.
func (h *ServiceIntegrationEndpointsHandler) Delete(project, endpointID string) error {
	path := buildPath("project", project, "integration_endpoint", endpointID)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// List all service integration endpoints for a given project.
func (h *ServiceIntegrationEndpointsHandler) List(project string) ([]*ServiceIntegrationEndpoint, error) {
	path := buildPath("project", project, "integration_endpoint")
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var response *ServiceIntegrationEndpointListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.ServiceIntegrationEndpoints, nil
}

func parseServiceIntegrationEndpointResponse(rsp []byte) (*ServiceIntegrationEndpoint, error) {
	var response *ServiceIntegrationEndpointResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.ServiceIntegrationEndpoint, nil
}
