package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupKafkaConnectorsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Kafka Connectors test case")

	const (
		UserName     = "test@aiven.io"
		UserPassword = "testabcd"
		AccessToken  = "some-random-token"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/userauth" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(authResponse{
				Token: AccessToken,
				State: "active",
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/project/test-pr/service/test-sr/connectors" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(KafkaConnectorsResponse{
				APIResponse: APIResponse{},
				Connectors: []KafkaConnector{
					{
						Name: "es-connector",
						Config: KafkaConnectorConfig{
							"topics":              "TestT1",
							"connection.username": "testUser1",
							"name":                "es-connector",
							"connection.password": "pass",
							"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
							"type.name":           "es-connector",
							"connection.url":      " https://elasticsearchUrl.aive.io:28038",
						},
						Plugin: KafkaConnectorPlugin{},
						Tasks:  []KafkaConnectorTask{},
					},
				},
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/project/test-pr/service/test-sr/connectors/test-kafka-con" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			err := json.NewEncoder(w).Encode(KafkaConnectorResponse{
				APIResponse: APIResponse{},
				Connector: KafkaConnector{
					Name: "test-kafka-con",
					Config: KafkaConnectorConfig{
						"topics":              "TestT1",
						"connection.username": "testUser1",
						"name":                "es-connector",
						"connection.password": "pass",
						"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
						"type.name":           "test-kafka-con",
						"connection.url":      " https://elasticsearchUrl.aive.io:28038",
					},
					Plugin: KafkaConnectorPlugin{},
					Tasks:  []KafkaConnectorTask{},
				},
			})

			if err != nil {
				t.Error(err)
			}
		}

	}))

	apiurl = ts.URL

	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown Kafka Connectors test case")
	}
}

func TestKafkaConnectorsHandler_Create(t *testing.T) {
	c, tearDown := setupKafkaConnectorsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		c       KafkaConnectorConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				c: KafkaConnectorConfig{
					"topics":              "TestT1",
					"connection.username": "testUser1",
					"name":                "es-connector",
					"connection.password": "pass",
					"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
					"type.name":           "es-connector",
					"connection.url":      " https://elasticsearchUrl.aive.io:28038",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaConnectorsHandler{
				client: tt.fields.client,
			}
			if err := h.Create(tt.args.project, tt.args.service, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKafkaConnectorsHandler_Delete(t *testing.T) {
	c, tearDown := setupKafkaConnectorsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-kafka-con",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaConnectorsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.project, tt.args.service, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKafkaConnectorsHandler_List(t *testing.T) {
	c, tearDown := setupKafkaConnectorsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaConnectorsResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&KafkaConnectorsResponse{
				APIResponse: APIResponse{},
				Connectors: []KafkaConnector{
					{
						Name: "es-connector",
						Config: KafkaConnectorConfig{
							"topics":              "TestT1",
							"connection.username": "testUser1",
							"name":                "es-connector",
							"connection.password": "pass",
							"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
							"type.name":           "es-connector",
							"connection.url":      " https://elasticsearchUrl.aive.io:28038",
						},
						Plugin: KafkaConnectorPlugin{},
						Tasks:  []KafkaConnectorTask{},
					}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaConnectorsHandler{
				client: tt.fields.client,
			}
			got, err := h.List(tt.args.project, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaConnectorsHandler_Get(t *testing.T) {
	c, tearDown := setupKafkaConnectorsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaConnectorResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-kafka-con",
			},
			&KafkaConnectorResponse{
				APIResponse: APIResponse{},
				Connector: KafkaConnector{
					Name: "test-kafka-con",
					Config: KafkaConnectorConfig{
						"topics":              "TestT1",
						"connection.username": "testUser1",
						"name":                "es-connector",
						"connection.password": "pass",
						"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
						"type.name":           "test-kafka-con",
						"connection.url":      " https://elasticsearchUrl.aive.io:28038",
					},
					Plugin: KafkaConnectorPlugin{},
					Tasks:  []KafkaConnectorTask{},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaConnectorsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.project, tt.args.service, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaConnectorsHandler_Update(t *testing.T) {
	c, tearDown := setupKafkaConnectorsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
		c       KafkaConnectorConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaConnectorResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-kafka-con",
				c: KafkaConnectorConfig{
					"topics":              "TestT1",
					"connection.username": "testUser1",
					"name":                "es-connector",
					"connection.password": "pass",
					"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
					"type.name":           "test-kafka-con",
					"connection.url":      " https://elasticsearchUrl.aive.io:28038",
				},
			},
			&KafkaConnectorResponse{
				APIResponse: APIResponse{},
				Connector: KafkaConnector{
					Name: "test-kafka-con",
					Config: KafkaConnectorConfig{
						"topics":              "TestT1",
						"connection.username": "testUser1",
						"name":                "es-connector",
						"connection.password": "pass",
						"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
						"type.name":           "test-kafka-con",
						"connection.url":      " https://elasticsearchUrl.aive.io:28038",
					},
					Plugin: KafkaConnectorPlugin{},
					Tasks:  []KafkaConnectorTask{},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaConnectorsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.project, tt.args.service, tt.args.name, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
