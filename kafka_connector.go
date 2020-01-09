package aiven

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

// Get gets a single Kafka Connector by Connector Name
func (h *KafkaConnectorsHandler) Get(project, service, name string) (*KafkaConnectorResponse, error) {
	path := buildPath("project", project, "service", service, "connectors", name)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp KafkaConnectorResponse
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
