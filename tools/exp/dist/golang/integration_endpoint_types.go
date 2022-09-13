// Code generated by gojson. DO NOT EDIT.

package golang

import (
	"bytes"

	"github.com/aiven/aiven-go-client/tools/exp/dist"
	"gopkg.in/yaml.v3"
)

// IntegrationEndpointTypes is a generated type that represents integration_endpoint_types.yml.
type IntegrationEndpointTypes struct {
	Datadog struct {
		Properties struct {
			DatadogAPIKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
				UserError string `yaml:"user_error"`
			} `yaml:"datadog_api_key"`
			DatadogTags struct {
				Example []struct {
					Comment string `yaml:"comment"`
					Tag     string `yaml:"tag"`
				} `yaml:"example"`
				Items struct {
					Properties struct {
						Comment struct {
							Example   string `yaml:"example"`
							MaxLength int    `yaml:"max_length"`
							Title     string `yaml:"title"`
							Type      string `yaml:"type"`
						} `yaml:"comment"`
						Tag struct {
							Description string `yaml:"description"`
							Example     string `yaml:"example"`
							MaxLength   int    `yaml:"max_length"`
							MinLength   int    `yaml:"min_length"`
							Pattern     string `yaml:"pattern"`
							Title       string `yaml:"title"`
							Type        string `yaml:"type"`
							UserError   string `yaml:"user_error"`
						} `yaml:"tag"`
					} `yaml:"properties"`
					Title string `yaml:"title"`
					Type  string `yaml:"type"`
				} `yaml:"items"`
				MaxItems int    `yaml:"max_items"`
				Title    string `yaml:"title"`
				Type     string `yaml:"type"`
			} `yaml:"datadog_tags"`
			DisableConsumerStats struct {
				Example bool   `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"disable_consumer_stats"`
			KafkaConsumerCheckInstances struct {
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"kafka_consumer_check_instances"`
			KafkaConsumerStatsTimeout struct {
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"kafka_consumer_stats_timeout"`
			MaxPartitionContexts struct {
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"max_partition_contexts"`
			Site struct {
				Enum []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"site"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"datadog"`
	ExternalAwsCloudwatchLogs struct {
		Properties struct {
			AccessKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"access_key"`
			LogGroupName struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"log_group_name"`
			Region struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"region"`
			SecretKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"secret_key"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_aws_cloudwatch_logs"`
	ExternalAwsCloudwatchMetrics struct {
		Properties struct {
			AccessKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"access_key"`
			Namespace struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"namespace"`
			Region struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"region"`
			SecretKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"secret_key"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_aws_cloudwatch_metrics"`
	ExternalElasticsearchLogs struct {
		Properties struct {
			Ca struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ca"`
			IndexDaysMax struct {
				Default string `yaml:"default"`
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"index_days_max"`
			IndexPrefix struct {
				Default   string `yaml:"default"`
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
				UserError string `yaml:"user_error"`
			} `yaml:"index_prefix"`
			Timeout struct {
				Default string `yaml:"default"`
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"timeout"`
			URL struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"url"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_elasticsearch_logs"`
	ExternalGoogleCloudLogging struct {
		Properties struct {
			LogID struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"log_id"`
			ProjectID struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"project_id"`
			ServiceAccountCredentials struct {
				Description string `yaml:"description"`
				Example     string `yaml:"example"`
				MaxLength   int    `yaml:"max_length"`
				Title       string `yaml:"title"`
				Type        string `yaml:"type"`
			} `yaml:"service_account_credentials"`
		} `yaml:"properties"`
		Title string `yaml:"title"`
		Type  string `yaml:"type"`
	} `yaml:"external_google_cloud_logging"`
	ExternalKafka struct {
		Properties struct {
			BootstrapServers struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"bootstrap_servers"`
			SaslMechanism struct {
				Enum []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string   `yaml:"example"`
				Title   string   `yaml:"title"`
				Type    []string `yaml:"type"`
			} `yaml:"sasl_mechanism"`
			SaslPlainPassword struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				MinLength int      `yaml:"min_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"sasl_plain_password"`
			SaslPlainUsername struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				MinLength int      `yaml:"min_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"sasl_plain_username"`
			SecurityProtocol struct {
				Enum []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"security_protocol"`
			SslCaCert struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ssl_ca_cert"`
			SslClientCert struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ssl_client_cert"`
			SslClientKey struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ssl_client_key"`
			SslEndpointIdentificationAlgorithm struct {
				Enum []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string   `yaml:"example"`
				Title   string   `yaml:"title"`
				Type    []string `yaml:"type"`
			} `yaml:"ssl_endpoint_identification_algorithm"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_kafka"`
	ExternalOpensearchLogs struct {
		Properties struct {
			Ca struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ca"`
			IndexDaysMax struct {
				Default string `yaml:"default"`
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"index_days_max"`
			IndexPrefix struct {
				Default   string `yaml:"default"`
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
				UserError string `yaml:"user_error"`
			} `yaml:"index_prefix"`
			Timeout struct {
				Default string `yaml:"default"`
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"timeout"`
			URL struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"url"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_opensearch_logs"`
	ExternalPostgresql struct {
		Properties struct {
			Host struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"host"`
			Password struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"password"`
			Port struct {
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"port"`
			SslMode struct {
				Default string `yaml:"default"`
				Enum    []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"ssl_mode"`
			SslRootCert struct {
				Default   string `yaml:"default"`
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"ssl_root_cert"`
			Username struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"username"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_postgresql"`
	ExternalSchemaRegistry struct {
		Properties struct {
			Authentication struct {
				Enum []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"authentication"`
			BasicAuthPassword struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"basic_auth_password"`
			BasicAuthUsername struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"basic_auth_username"`
			URL struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"url"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"external_schema_registry"`
	Jolokia struct {
		Properties struct {
			BasicAuthPassword struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"basic_auth_password"`
			BasicAuthUsername struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
				UserError string `yaml:"user_error"`
			} `yaml:"basic_auth_username"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"jolokia"`
	Prometheus struct {
		Properties struct {
			BasicAuthPassword struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"basic_auth_password"`
			BasicAuthUsername struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Pattern   string `yaml:"pattern"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
				UserError string `yaml:"user_error"`
			} `yaml:"basic_auth_username"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"prometheus"`
	Rsyslog struct {
		Properties struct {
			Ca struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"ca"`
			Cert struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"cert"`
			Format struct {
				Default string `yaml:"default"`
				Enum    []struct {
					Value string `yaml:"value"`
				} `yaml:"enum"`
				Example string `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"format"`
			Key struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"key"`
			Logline struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"logline"`
			Port struct {
				Default string `yaml:"default"`
				Example string `yaml:"example"`
				Maximum int    `yaml:"maximum"`
				Minimum int    `yaml:"minimum"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"port"`
			Sd struct {
				Example   string   `yaml:"example"`
				MaxLength int      `yaml:"max_length"`
				Title     string   `yaml:"title"`
				Type      []string `yaml:"type"`
			} `yaml:"sd"`
			Server struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"server"`
			TLS struct {
				Default bool   `yaml:"default"`
				Example bool   `yaml:"example"`
				Title   string `yaml:"title"`
				Type    string `yaml:"type"`
			} `yaml:"tls"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"rsyslog"`
	Signalfx struct {
		Properties struct {
			EnabledMetrics struct {
				Default []string `yaml:"default"`
				Example []string `yaml:"example"`
				Items   struct {
					MaxLength int    `yaml:"max_length"`
					Title     string `yaml:"title"`
					Type      string `yaml:"type"`
				} `yaml:"items"`
				MaxItems int    `yaml:"max_items"`
				Title    string `yaml:"title"`
				Type     string `yaml:"type"`
			} `yaml:"enabled_metrics"`
			SignalfxAPIKey struct {
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"signalfx_api_key"`
			SignalfxRealm struct {
				Default   string `yaml:"default"`
				Example   string `yaml:"example"`
				MaxLength int    `yaml:"max_length"`
				MinLength int    `yaml:"min_length"`
				Title     string `yaml:"title"`
				Type      string `yaml:"type"`
			} `yaml:"signalfx_realm"`
		} `yaml:"properties"`
		Type string `yaml:"type"`
	} `yaml:"signalfx"`
}

// FilledIntegrationEndpointTypes is a function that returns IntegrationEndpointTypes filled with generated data.
func FilledIntegrationEndpointTypes() *IntegrationEndpointTypes {
	d := yaml.NewDecoder(bytes.NewReader(dist.IntegrationEndpointTypes))

	r := &IntegrationEndpointTypes{}

	if err := d.Decode(r); err != nil {
		panic(err)
	}

	return r
}