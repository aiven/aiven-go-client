package aiven

import (
	"context"
	"time"
)

type (
	// Service represents the Service model on Aiven.
	Service struct {
		ACL                   []*KafkaACL               `json:"acl"`
		SchemaRegistryACL     []*KafkaSchemaRegistryACL `json:"schema_registry_acl"`
		Backups               []*Backup                 `json:"backups"`
		CloudName             string                    `json:"cloud_name"`
		ConnectionPools       []*ConnectionPool         `json:"connection_pools"`
		CreateTime            *time.Time                `json:"create_time"`
		UpdateTime            *time.Time                `json:"update_time"`
		GroupList             []string                  `json:"group_list"`
		NodeCount             int                       `json:"node_count"`
		Plan                  string                    `json:"plan"`
		Name                  string                    `json:"service_name"`
		Type                  string                    `json:"service_type"`
		ProjectVPCID          *string                   `json:"project_vpc_id"`
		URI                   string                    `json:"service_uri"`
		URIParams             map[string]string         `json:"service_uri_params"`
		State                 string                    `json:"state"`
		Metadata              interface{}               `json:"metadata"`
		Users                 []*ServiceUser            `json:"users"`
		UserConfig            map[string]interface{}    `json:"user_config"`
		ConnectionInfo        ConnectionInfo            `json:"connection_info"`
		TerminationProtection bool                      `json:"termination_protection"`
		MaintenanceWindow     MaintenanceWindow         `json:"maintenance"`
		Integrations          []*ServiceIntegration     `json:"service_integrations"`
		Components            []*ServiceComponents      `json:"components"`
		Powered               bool                      `json:"powered"`
		NodeStates            []*NodeState              `json:"node_states"`
		DiskSpaceMB           int                       `json:"disk_space_mb"`
		Features              ServiceFeatures           `json:"features"`
		TechnicalEmails       []ContactEmail            `json:"tech_emails"`
	}

	ServiceFeatures struct {
		EnhancedLogging                bool `json:"enhanced_logging"`
		ImprovedTopicManagement        bool `json:"improved_topic_management"`
		ServiceIntegrations            bool `json:"service_integrations"`
		KafkaConnectServiceIntegration bool `json:"kafka_connect_service_integration"`
		PGAllowReplication             bool `json:"pg_allow_replication"`
		Letsencrypt                    bool `json:"letsencrypt"`
		IndexPatterns                  bool `json:"index_patterns"`
		Karapace                       bool `json:"karapace"`
		KarapaceRest                   bool `json:"karapace_rest"`
		KarapaceJSONSchema             bool `json:"karapace_json_schema"`
		KafkaConfigBackupsEnabled      bool `json:"kafka_config_backups_enabled"`
		TopicManagement                bool `json:"topic_management"`
		KafkaStrictAccessCertChecks    bool `json:"kafka_strict_access_cert_checks"`
		KafkaTopicInfo                 bool `json:"kafka_topic_info"`
		KafkaConnect                   bool `json:"kafka_connect"`
		KafkaMirrormaker               bool `json:"kafka_mirrormaker"`
		KafkaRest                      bool `json:"kafka_rest"`
		SchemaRegistry                 bool `json:"schema_registry"`
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
		BackupTime        *time.Time                `json:"backup_time"`
		BackupName        string                    `json:"backup_name"`
		DataSize          int                       `json:"data_size"`
		StorageLocation   string                    `json:"storage_location"`
		AdditionalRegions []*BackupAdditionalRegion `json:"additional_regions"`
	}

	// BackupAdditionalRegion represents a remote region where the backup is synchronized
	BackupAdditionalRegion struct {
		Cloud  string `json:"cloud"`
		Region string `json:"region"`
	}

	// ConnectionInfo represents the Service Connection information on Aiven.
	ConnectionInfo struct {
		// Common
		Direct      []string `json:"direct"`
		LetsEncrypt bool     `json:"letsencrypt"`

		// Kafka
		KafkaHosts        []string `json:"kafka"` // TODO: Rename to KafkaURIs in the next major version.
		KafkaAccessCert   string   `json:"kafka_access_cert"`
		KafkaAccessKey    string   `json:"kafka_access_key"`
		KafkaConnectURI   string   `json:"kafka_connect_uri"`
		KafkaRestURI      string   `json:"kafka_rest_uri"`
		SchemaRegistryURI string   `json:"schema_registry_uri"`

		// PostgreSQL
		PostgresURIs        []string         `json:"pg"`
		PostgresBouncer     string           `json:"pg_bouncer"`
		PostgresParams      []PostgresParams `json:"pg_params"`
		PostgresReplicaURI  string           `json:"pg_replica_uri"`
		PostgresStandbyURIs []string         `json:"pg_standby"`
		PostgresSyncingURIs []string         `json:"pg_syncing"`

		// Thanos
		ThanosURIs                      []string `json:"thanos"`
		QueryFrontendURI                string   `json:"query_frontend_uri"`
		QueryURI                        string   `json:"query_uri"`
		ReceiverIngestingRemoteWriteURI string   `json:"receiver_ingesting_remote_write_uri"`
		ReceiverRemoteWriteURI          string   `json:"receiver_remote_write_uri"`
		StoreURI                        string   `json:"store_uri"`

		// MySQL
		MySQLURIs        []string      `json:"mysql"`
		MySQLParams      []MySQLParams `json:"mysql_params"`
		MySQLReplicaURI  string        `json:"mysql_replica_uri"`
		MySQLStandbyURIs []string      `json:"mysql_standby"`

		// ElasticSearch
		ElasticsearchURIs     []string `json:"elasticsearch"`
		ElasticsearchPassword string   `json:"elasticsearch_password"`
		ElasticsearchUsername string   `json:"elasticsearch_username"`
		KibanaURI             string   `json:"kibana_uri"` // This field is available in OpenSearch as well.

		// OpenSearch
		// TODO: Rename Opensearch to OpenSearch in the next major version.
		OpensearchURIs          []string `json:"opensearch"`
		OpensearchDashboardsURI string   `json:"opensearch_dashboards_uri"`
		OpensearchUsername      string   `json:"opensearch_username"`
		OpensearchPassword      string   `json:"opensearch_password"`

		// Cassandra
		CassandraHosts []string `json:"cassandra"` // TODO: Rename to CassandraURIs in the next major version.

		// Redis and Dragonfly
		RedisURIs       []string `json:"redis"`
		RedisSlaveURIs  []string `json:"redis_slave"`
		RedisReplicaURI string   `json:"redis_replica_uri"`
		RedisPassword   string   `json:"redis_password"`

		// InfluxDB
		InfluxDBURIs         []string `json:"influxdb"`
		InfluxDBUsername     string   `json:"influxdb_username"`
		InfluxDBPassword     string   `json:"influxdb_password"`
		InfluxDBDatabaseName string   `json:"influxdb_dbname"`

		// Grafana
		GrafanaURIs []string `json:"grafana"`

		// M3DB
		M3DBURIs                 []string `json:"m3db"`
		HTTPClusterURI           string   `json:"http_cluster_uri"`
		HTTPNodeURI              string   `json:"http_node_uri"`
		InfluxDBURI              string   `json:"influxdb_uri"`
		PrometheusRemoteReadURI  string   `json:"prometheus_remote_read_uri"`
		PrometheusRemoteWriteURI string   `json:"prometheus_remote_write_uri"`

		// M3 Aggregator
		M3AggregatorURIs  []string `json:"m3aggregator"`
		AggregatorHTTPURI string   `json:"aggregator_http_uri"`

		// ClickHouse
		ClickHouseURIs []string `json:"clickhouse"`

		// Flink
		FlinkHostPorts []string `json:"flink"` // TODO: Rename to FlinkURIs in the next major version.
	}

	// PostgresParams represents individual parameters for a PostgreSQL ConnectionInfo
	PostgresParams struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		SSLMode      string `json:"sslmode"`
		User         string `json:"user"`
		Password     string `json:"password"`
		DatabaseName string `json:"dbname"`
	}

	// MySQLParams represents individual parameters for a MySQL ConnectionInfo
	MySQLParams struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		SSLMode      string `json:"ssl-mode"`
		User         string `json:"user"`
		Password     string `json:"password"`
		DatabaseName string `json:"dbname"`
	}

	// KafkaACL represents a Kafka ACL entry on Aiven.
	KafkaACL struct {
		ID         string `json:"id"`
		Permission string `json:"permission"`
		Topic      string `json:"topic"`
		Username   string `json:"username"`
	}

	// KafkaSchemaRegistryACL represents a Kafka Schema Registry ACL entry on Aiven.
	KafkaSchemaRegistryACL struct {
		ID         string `json:"id"`
		Permission string `json:"permission"`
		Resource   string `json:"resource"`
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
		DayOfWeek string               `json:"dow"`
		TimeOfDay string               `json:"time"`
		Updates   []*MaintenanceUpdate `json:"updates,omitempty"`
	}

	// MaintenanceUpdate represents a maintenance needing to be applied on the service.
	MaintenanceUpdate struct {
		Deadline    *string `json:"deadline,omitempty"`
		Description string  `json:"description,omitempty"`
		StartAfter  string  `json:"start_after,omitempty"`
		StartAt     *string `json:"start_at,omitempty"`
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
		StaticIPs             []string                `json:"static_ips,omitempty"`
		TechnicalEmails       *[]ContactEmail         `json:"tech_emails,omitempty"`
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
		Karapace              *bool                  `json:"karapace,omitempty"`
		TechnicalEmails       *[]ContactEmail        `json:"tech_emails,omitempty"`
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
func (h *ServicesHandler) Create(ctx context.Context, project string, req CreateServiceRequest) (*Service, error) {
	path := buildPath("project", project, "service")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Get gets a specific service from Aiven.
func (h *ServicesHandler) Get(ctx context.Context, project, service string) (*Service, error) {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Update will update the given service with the given parameters.
func (h *ServicesHandler) Update(ctx context.Context, project, service string, req UpdateServiceRequest) (*Service, error) {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(bts, &r)

	return r.Service, errR
}

// Delete will delete the given service from Aiven.
func (h *ServicesHandler) Delete(ctx context.Context, project, service string) error {
	path := buildPath("project", project, "service", service)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List will fetch all services for a given project.
func (h *ServicesHandler) List(ctx context.Context, project string) ([]*Service, error) {
	path := buildPath("project", project, "service")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceListResponse
	errR := checkAPIResponse(bts, &r)

	return r.Services, errR
}
