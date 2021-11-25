// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

type (
	// Service represents the Service model on Aiven.
	Service struct {
		ACL                   []*KafkaACL            `json:"acl"`
		Backups               []*Backup              `json:"backups"`
		CloudName             string                 `json:"cloud_name"`
		ConnectionPools       []*ConnectionPool      `json:"connection_pools"`
		CreateTime            string                 `json:"create_time"`
		UpdateTime            string                 `json:"update_time"`
		GroupList             []string               `json:"group_list"`
		NodeCount             int                    `json:"node_count"`
		Plan                  string                 `json:"plan"`
		Name                  string                 `json:"service_name"`
		Type                  string                 `json:"service_type"`
		ProjectVPCID          *string                `json:"project_vpc_id"`
		URI                   string                 `json:"service_uri"`
		URIParams             map[string]string      `json:"service_uri_params"`
		State                 string                 `json:"state"`
		Metadata              interface{}            `json:"metadata"`
		Users                 []*ServiceUser         `json:"users"`
		UserConfig            map[string]interface{} `json:"user_config"`
		ConnectionInfo        ConnectionInfo         `json:"connection_info"`
		TerminationProtection bool                   `json:"termination_protection"`
		MaintenanceWindow     MaintenanceWindow      `json:"maintenance"`
		Integrations          []*ServiceIntegration  `json:"service_integrations"`
		Components            []*ServiceComponents   `json:"components"`
		Powered               bool                   `json:"powered"`
		NodeStates            []*NodeState           `json:"node_states"`
		DiskSpaceMB           int                    `json:"disk_space_mb"`
	}

	// NodeState represents the Node State model on Aiven
	NodeState struct {
		Name            string           `json:"name"`
		ProgressUpdates []ProgressUpdate `json:"progress_updates"`
		Role            string           `json:"role"`
		State           string           `json:"state"`
	}

	// ProgressUpdate state represents the Progress Update model on Aiven
	ProgressUpdate struct {
		Completed bool   `json:"completed"`
		Current   int    `json:"current"`
		Max       int    `json:"max"`
		Min       int    `json:"min"`
		Phase     string `json:"phase"`
		Unit      string `json:"unit"`
	}

	// ServiceComponents represents Service Components which may contain
	// information regarding service components Dynamic/Public DNS records
	ServiceComponents struct {
		Component                 string `json:"component"`
		Host                      string `json:"host"`
		Port                      int    `json:"port"`
		Route                     string `json:"route"`
		Usage                     string `json:"usage"`
		Ssl                       *bool  `json:"ssl"`
		KafkaAuthenticationMethod string `json:"kafka_authentication_method"`
	}

	// Backup represents an individual backup of service data on Aiven
	Backup struct {
		BackupTime string `json:"backup_time"`
		DataSize   int    `json:"data_size"`
	}

	// ConnectionInfo represents the Service Connection information on Aiven.
	ConnectionInfo struct {
		CassandraHosts []string `json:"cassandra"`

		ElasticsearchURIs     []string `json:"elasticsearch"`
		ElasticsearchUsername string   `json:"elasticsearch_username"`
		ElasticsearchPassword string   `json:"elasticsearch_password"`
		KibanaURI             string   `json:"kibana_uri"`

		OpensearchURIs          []string `json:"opensearch"`
		OpensearchDashboardsURI string   `json:"opensearch_dashboards_uri"`
		OpensearchPassword      string   `json:"opensearch_password"`
		OpensearchUsername      string   `json:"opensearch_username"`

		GrafanaURIs []string `json:"grafana"`

		InfluxDBURIs         []string `json:"influxdb"`
		InfluxDBDatabaseName string   `json:"influxdb_dbname"`
		InfluxDBUsername     string   `json:"influxdb_username"`
		InfluxDBPassword     string   `json:"influxdb_password"`

		KafkaHosts        []string `json:"kafka"`
		KafkaAccessCert   string   `json:"kafka_access_cert"`
		KafkaAccessKey    string   `json:"kafka_access_key"`
		KafkaConnectURI   string   `json:"kafka_connect_uri"`
		KafkaRestURI      string   `json:"kafka_rest_uri"`
		SchemaRegistryURI string   `json:"schema_registry_uri"`

		PostgresParams      []PostgresParams `json:"pg_params"`
		PostgresReplicaURI  string           `json:"pg_replica_uri"`
		PostgresStandbyURIs []string         `json:"pg_standby"`
		PostgresURIs        []string         `json:"pg"`

		RedisPassword  string   `json:"redis_password"`
		RedisSlaveURIs []string `json:"redis_slave"`
		RedisURIs      []string `json:"redis"`

		FlinkHostPorts []string `json:"flink"`
	}

	// PostgresParams represents individual parameters for a PostgreSQL ConnectionInfo
	PostgresParams struct {
		DatabaseName string `json:"dbname"`
		Host         string `string:"host"`
		Password     string `string:"password"`
		Port         string `string:"port"`
		SSLMode      string `string:"sslmode"`
		User         string `string:"user"`
	}

	// KafkaACL represents a Kafka ACL entry on Aiven.
	KafkaACL struct {
		ID         string `json:"id"`
		Permission string `json:"permission"`
		Topic      string `json:"topic"`
		Username   string `json:"username"`
	}

	// ConnectionPool represents a PostgreSQL PGBouncer connection pool on Aiven
	ConnectionPool struct {
		ConnectionURI string `json:"connection_uri"`
		Database      string `json:"database"`
		PoolMode      string `json:"pool_mode"`
		PoolName      string `json:"pool_name"`
		PoolSize      int    `json:"pool_size"`
		Username      string `json:"username"`
	}

	// MaintenanceWindow during which maintenance operations should take place
	MaintenanceWindow struct {
		DayOfWeek string `json:"dow"`
		TimeOfDay string `json:"time"`
	}

	// ServicesHandler is the client that interacts with the Service API
	// endpoints on Aiven.
	ServicesHandler struct {
		client *Client
	}

	// CreateServiceRequest are the parameters to create a Service.
	CreateServiceRequest struct {
		Cloud                 string                  `json:"cloud,omitempty"`
		GroupName             string                  `json:"group_name,omitempty"`
		MaintenanceWindow     *MaintenanceWindow      `json:"maintenance,omitempty"`
		Plan                  string                  `json:"plan,omitempty"`
		ProjectVPCID          *string                 `json:"project_vpc_id"`
		ServiceName           string                  `json:"service_name"`
		ServiceType           string                  `json:"service_type"`
		TerminationProtection bool                    `json:"termination_protection"`
		UserConfig            map[string]interface{}  `json:"user_config,omitempty"`
		ServiceIntegrations   []NewServiceIntegration `json:"service_integrations"`
		DiskSpaceMB           int                     `json:"disk_space_mb,omitempty"`
	}

	// UpdateServiceRequest are the parameters to update a Service.
	UpdateServiceRequest struct {
		Cloud                 string                 `json:"cloud,omitempty"`
		GroupName             string                 `json:"group_name,omitempty"`
		MaintenanceWindow     *MaintenanceWindow     `json:"maintenance,omitempty"`
		Plan                  string                 `json:"plan,omitempty"`
		ProjectVPCID          *string                `json:"project_vpc_id"`
		Powered               bool                   `json:"powered"`
		TerminationProtection bool                   `json:"termination_protection"`
		UserConfig            map[string]interface{} `json:"user_config,omitempty"`
		DiskSpaceMB           int                    `json:"disk_space_mb,omitempty"`
	}

	// ServiceResponse represents the response from Aiven after interacting with
	// the Service API.
	ServiceResponse struct {
		APIResponse
		Service *Service `json:"service"`
	}

	// ServiceListResponse represents the response from Aiven for listing
	// services.
	ServiceListResponse struct {
		APIResponse
		Services []*Service `json:"services"`
	}
)

// Hostname provides host name for the service. This method is provided for backwards
// compatibility, typically it is easier to just get the value from URIParams directly.
func (s *Service) Hostname() (string, error) {
	return s.URIParams["host"], nil
}

// Port provides port for the service. This method is provided for backwards
// compatibility, typically it is easier to just get the value from URIParams directly.
func (s *Service) Port() (string, error) {
	return s.URIParams["port"], nil
}

// Create creates the given Service on Aiven.
func (h *ServicesHandler) Create(project string, req CreateServiceRequest) (*Service, error) {
	path := buildPath("project", project, "service")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Get gets a specific service from Aiven.
func (h *ServicesHandler) Get(project, service string) (*Service, error) {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Update will update the given service with the given parameters.
func (h *ServicesHandler) Update(project, service string, req UpdateServiceRequest) (*Service, error) {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Delete will delete the given service from Aiven.
func (h *ServicesHandler) Delete(project, service string) error {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List will fetch all services for a given project.
func (h *ServicesHandler) List(project string) ([]*Service, error) {
	path := buildPath("project", project, "service")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Services, errR
}
