package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
)

type (
	// ServiceIntegration is the representation of a Service Integration in the Aiven API.
	ServiceIntegration struct {
		Active               bool                   `json:"active"`
		Description          string                 `json:"description"`
		DestEndpoint         string                 `json:"dest_endpoint"`
		DestEndpointID       string                 `json:"dest_endpoint_id"`
		DestProject          string                 `json:"dest_project"`
		DestService          string                 `json:"dest_service"`
		DestServiceType      string                 `json:"dest_service_type"`
		Enabled              bool                   `json:"enabled"`
		IntegrationStatus    interface{}            `json:"integration_status"`
		IntegrationType      string                 `json:"integration_type"`
		ServiceIntegrationID string                 `json:"service_integration_id"`
		SourceEndpoint       string                 `json:"source_endpoint"`
		SourceEndpointID     string                 `json:"source_endpoint_id"`
		SourceProject        string                 `json:"source_project"`
		SourceService        string                 `json:"source_service"`
		SourceServiceType    string                 `json:"source_service_type"`
		UserConfig           map[string]interface{} `json:"user_config"`
	}

	// ServiceIntegrationEndpoint is the representation of the Create Service Integration model on Aiven.
	ServiceIntegrationEndpoint struct {
		EndpointID   string                 `json:"endpoint_id"`
		EndpointName string                 `json:"endpoint_name"`
		EndpointType string                 `json:"endpoint_type"`
		UserConfig   map[string]interface{} `json:"user_config"`
	}

	// IntegrationEndpointTypes is the representation of Endpoint Types
	IntegrationEndpointTypes struct {
		EndpointType     string                 `json:"endpoint_type"`
		ServiceTypes     []string               `json:"service_types"`
		Title            string                 `json:"title"`
		UserConfigSchema map[string]interface{} `json:"user_config_schema,omitempty"`
	}

	// IntegrationTypes is a representation of service integration types
	IntegrationTypes struct {
		DestDescription    string                 `json:"dest_description"`
		DestServiceType    string                 `json:"dest_service_type"`
		IntegrationType    string                 `json:"integration_type"`
		SourceDescription  string                 `json:"source_description"`
		SourceServiceTypes []string               `json:"source_service_types"`
		UserConfigSchema   map[string]interface{} `json:"user_config_schema,omitempty"`
	}
	// ServiceIntegrationHandler is the client that interacts with the ServiceIntegration endpoints.
	ServiceIntegrationHandler struct {
		client *Client
	}

	// CreateServiceIntegrationRequest are the parameters required to create a ServiceIntegration
	CreateServiceIntegrationRequest struct {
		IntegrationType string                 `json:"integration_type"`
		SourceService   string                 `json:"source_service"`
		DestService     string                 `json:"dest_service"`
		UserConfig      map[string]interface{} `json:"user_config,omitempty"`
	}

	// CreateServiceIntegrationEndpointRequest are the parameters required to create a ServiceIntegration Endpoint.
	CreateServiceIntegrationEndpointRequest struct {
		EndpointName string `json:"endpoint_name"`
		// Allowed values: "dashboard", "datadog", "logs", "metrics", "mirrormaker"
		EndpointType string `json:"endpoint_type"`
	}

	// UpdateIntegrationRequest are the parameters to update a Service.
	UpdateIntegrationRequest struct {
		UserConfig map[string]interface{} `json:"user_config,omitempty"`
	}

	// ServiceIntegrationResponse represents the response after creating a ServiceIntegration.
	ServiceIntegrationResponse struct {
		APIResponse
		Integration *ServiceIntegration `json:"service_integration"`
	}

	// CreateServiceIntegrationResponse represents the response after creating a ServiceIntegration.
	CreateServiceIntegrationResponse struct {
		APIResponse
		IntegrationEndpoint ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
	}

	// ServiceIntegrationEndpointResponse represents the response for Integration Endpoint
	ServiceIntegrationEndpointResponse struct {
		APIResponse
		ServiceIntegrationEndpoint *ServiceIntegrationEndpoint `json:"service_integration_endpoint"`
	}

	// ServiceIntegrationEndpointListResponse represents the response for listing Integration Endpoints
	ServiceIntegrationEndpointListResponse struct {
		APIResponse
		ServiceIntegrationEndpoint []*ServiceIntegrationEndpoint `json:"service_integration_endpoints"`
	}

	// ServiceIntegrationListResponse represents the response for listing ServiceIntegrations
	ServiceIntegrationListResponse struct {
		APIResponse
		ServiceIntegrations []*ServiceIntegration `json:"service_integrations"`
	}

	// IntegrationEndpointTypesResponse represents the response for listing Endpoint Types
	IntegrationEndpointTypesResponse struct {
		APIResponse
		EndpointTypes []*IntegrationEndpointTypes `json:"endpoint_types"`
	}

	// IntegrationTypesResponse represents the response for listing service integration types
	IntegrationTypesResponse struct {
		APIResponse
		IntegrationTypes []*IntegrationTypes `json:"integration_types"`
	}
)

// Create creates the given Service Integration on Aiven.
func (h *ServiceIntegrationHandler) Create(project string, req CreateServiceIntegrationRequest) (*ServiceIntegration, error) {
	rsp, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/integration", project), req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// CreateEndpoint creates the given Service Integration Endpoint on Aiven.
func (h *ServiceIntegrationHandler) CreateEndpoint(project string, req CreateServiceIntegrationEndpointRequest) (*ServiceIntegrationEndpoint, error) {
	rsp, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/integration_endpoint", project), req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationEndpointResponse(rsp)
}

// Get gets a specific service integration from Aiven.
func (h *ServiceIntegrationHandler) Get(project, integrationID string) (*ServiceIntegration, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/integration/%s", project, integrationID), nil)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// DeleteEndpoint will delete the given Service Integration Endpoint from Aiven.
func (h *ServiceIntegrationHandler) DeleteEndpoint(project, endpoint string) error {
	bts, err := h.client.doDeleteRequest(fmt.Sprintf("/project/%s/integration_endpoint/%s", project, endpoint), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// DeleteIntegration will delete the given Service Integration from Aiven.
func (h *ServiceIntegrationHandler) DeleteIntegration(project, integrationID string) error {
	bts, err := h.client.doDeleteRequest(fmt.Sprintf("/project/%s/integration/%s", project, integrationID), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// ListEndpoints will fetch all available integration endpoints for project
func (h *ServiceIntegrationHandler) ListEndpoints(project string) ([]*ServiceIntegrationEndpoint, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/integration_endpoint", project), nil)
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

	return response.ServiceIntegrationEndpoint, nil
}

// ListEndpointTypes will fetch all available integration endpoints types
func (h *ServiceIntegrationHandler) ListEndpointTypes(project string) ([]*IntegrationEndpointTypes, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/integration_endpoint_types", project), nil)
	if err != nil {
		return nil, err
	}
	var response *IntegrationEndpointTypesResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.EndpointTypes, nil
}

// ListIntegrationTypes will fetch available service integration types
func (h *ServiceIntegrationHandler) ListIntegrationTypes(project string) ([]*IntegrationTypes, error) {
	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/integration_types", project), nil)
	if err != nil {
		return nil, err
	}
	var response *IntegrationTypesResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.IntegrationTypes, nil
}

// Update will update the a service integration.
func (h *ServiceIntegrationHandler) Update(project, integrationID string, req UpdateIntegrationRequest) (*ServiceIntegration, error) {
	rsp, err := h.client.doPutRequest(fmt.Sprintf("/project/%s/integration/%s", project, integrationID), req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationResponse(rsp)
}

// UpdateEndpoint will update the a service integration endpoint.
func (h *ServiceIntegrationHandler) UpdateEndpoint(project, integrationID string, req UpdateIntegrationRequest) (*ServiceIntegrationEndpoint, error) {
	rsp, err := h.client.doPutRequest(fmt.Sprintf("/project/%s/integration/%s", project, integrationID), req)
	if err != nil {
		return nil, err
	}

	return parseServiceIntegrationEndpointResponse(rsp)
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

func parseServiceIntegrationResponse(rsp []byte) (*ServiceIntegration, error) {
	var response *ServiceIntegrationResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Integration, nil
}
