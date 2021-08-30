package aiven

import (
	"fmt"
	"net/http"
)

type (
	// KafkaConnectorsHandler Aiven go-client handler for Kafka Connectors
	KafkaConnectorsHandler struct {
		client *Client
	}

	// KafkaConnectorConfig represents a configuration for Kafka Connectors
	KafkaConnectorConfig map[string]string

	// KafkaConnector represents Kafka Connector
	KafkaConnector struct {
		Name   string `json:"name"`
		Config KafkaConnectorConfig
		Plugin KafkaConnectorPlugin
		Tasks  []KafkaConnectorTask
	}

	// KafkaConnectorTask represents Kafka Connector Task
	KafkaConnectorTask struct {
		Connector string
		Task      int
	}

	// KafkaConnectorPlugin represents Kafka Connector Plugin
	KafkaConnectorPlugin struct {
		Author           string `json:"author"`
		Class            string `json:"class"`
		DocumentationURL string `json:"docURL"`
		Title            string `json:"title"`
		Type             string `json:"type"`
		Version          string `json:"version"`
	}

	// KafkaConnectorsResponse represents Kafka Connectors API response
	KafkaConnectorsResponse struct {
		APIResponse
		Connectors []KafkaConnector
	}

	// KafkaConnectorResponse represents single Kafka Connector API response
	KafkaConnectorResponse struct {
		APIResponse
		Connector KafkaConnector
	}

	// KafkaConnectorStatusResponse represents single Kafka Connector API status response
	KafkaConnectorStatusResponse struct {
		APIResponse
		Status KafkaConnectorStatus `json:"status"`
	}

	// KafkaConnectorStatus represents the status of a kafka connector
	KafkaConnectorStatus struct {
		State string                     `json:"state"`
		Tasks []KafkaConnectorTaskStatus `json:"tasks"`
	}

	// KafkaConnectorTaskStatus represents the status of a kafka connector task
	KafkaConnectorTaskStatus struct {
		Id    int    `json:"id"`
		State string `json:"state"`
		Trace string `json:"trace"`
	}
)

// Create creates Kafka Connector attached to Kafka or Kafka Connector service based on configuration
func (h *KafkaConnectorsHandler) Create(project, service string, c KafkaConnectorConfig) error {
	path := buildPath("project", project, "service", service, "connectors")
	bts, err := h.client.doPostRequest(path, c)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes Kafka Connector by name
func (h *KafkaConnectorsHandler) Delete(project, service, name string) error {
	path := buildPath("project", project, "service", service, "connectors", name)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List lists all available Kafka Connectors for a service
func (h *KafkaConnectorsHandler) List(project, service string) (*KafkaConnectorsResponse, error) {
	path := buildPath("project", project, "service", service, "connectors")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp KafkaConnectorsResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	return &rsp, nil
}

// GetByName gets a KafkaConnector by name
func (h *KafkaConnectorsHandler) GetByName(project, service, name string) (*KafkaConnector, error) {
	path := buildPath("project", project, "service", service, "connectors")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp KafkaConnectorsResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	for i := range rsp.Connectors {
		con := rsp.Connectors[i]
		if con.Name == name {
			return &con, nil
		}
	}

	// TODO: This is a hack. We pretend that this was an API call all along, even though this is only convenience
	// It is acceptable because all other functions here have the contract that we return a non nil result if the
	// error is nil. So for the sake of API consistency we pretend that the remote API returned this error.
	return nil, Error{
		Message: fmt.Sprintf("no kafka connector with name '%s' found in project '%s' for service '%s'", name, project, service),
		Status:  http.StatusNotFound,
	}
}

// Get the status of a single Kafka Connector by name
func (h *KafkaConnectorsHandler) Status(project, service, name string) (*KafkaConnectorStatusResponse, error) {
	path := buildPath("project", project, "service", service, "connectors", name, "status")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp KafkaConnectorStatusResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}

// Update updates a Kafka Connector configuration by Connector Name
func (h *KafkaConnectorsHandler) Update(project, service, name string, c KafkaConnectorConfig) (*KafkaConnectorResponse, error) {
	path := buildPath("project", project, "service", service, "connectors", name)
	bts, err := h.client.doPutRequest(path, c)
	if err != nil {
		return nil, err
	}

	var rsp KafkaConnectorResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}
	return &rsp, nil
}
