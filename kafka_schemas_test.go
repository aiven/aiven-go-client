package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewKafkaSchema(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want KafkaSchemaSubject
	}{
		{
			"simple",
			args{s: `
				test
			`},
			KafkaSchemaSubject{Schema: "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKafkaSchema(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKafkaSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setupKafkaSchemasTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Kafka Schemas test case")

	const (
		UserName     = "test@aiven.io"
		UserPassword = "testabcd"
		AccessToken  = "some-random-token"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// auth
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

		// config
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/config" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(KafkaSchemaConfigResponse{
				APIResponse{},
				KafkaSchemaConfig{CompatibilityLevel: "FULL"},
			})

			if err != nil {
				t.Error(err)
			}

			return
		}

		// config
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/config" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(KafkaSchemaConfigResponse{
				APIResponse{},
				KafkaSchemaConfig{CompatibilityLevel: "FULL"},
			})

			if err != nil {
				t.Error(err)
			}

			return
		}

		// subjects
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/subjects" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(KafkaSchemaSubjectsResponse{
				APIResponse{},
				KafkaSchemaSubjects{Subjects: []string{"testSb1", "testSb2"}},
			})

			if err != nil {
				t.Error(err)
			}

			return
		}

		// add subject no versions
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/subjects/test-schema-no-versions/versions" {
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)

				return
			}

			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err := json.NewEncoder(w).Encode(KafkaSchemaSubjectResponse{
					APIResponse{},
					1,
				})

				if err != nil {
					t.Error(err)
				}

				return
			}
		}

		// add subject has versions
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/subjects/test-schema/versions" {
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err := json.NewEncoder(w).Encode(KafkaSchemaSubjectVersionResponse{
					APIResponse{},
					KafkaSchemaSubjectVersion{Versions: []int{1, 2, 3, 4}},
				})

				if err != nil {
					t.Error(err)
				}

				return
			}

			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err := json.NewEncoder(w).Encode(KafkaSchemaSubjectResponse{
					APIResponse{},
					5,
				})

				if err != nil {
					t.Error(err)
				}

				return
			}
		}

		// validate against version 4
		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/compatibility/subjects/test-schema/versions/4" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(KafkaSchemaValidateResponse{
				APIResponse{},
				true,
			})

			if err != nil {
				t.Error(err)
			}

			return
		}

		if r.URL.Path == "/project/test-pr/service/test-sr/kafka/schema/subjects/test-schema/versions/5" {
			if r.Method == "DELETE" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err := json.NewEncoder(w).Encode(APIResponse{})

				if err != nil {
					t.Error(err)
				}

				return
			}

			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err := json.NewEncoder(w).Encode(KafkaSchemaSubjectResponse{
					APIResponse: APIResponse{},
					Id:          5,
				})

				if err != nil {
					t.Error(err)
				}

				return
			}
		}

	}))

	apiurl = ts.URL

	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown ElasticsearchACLs test case")
	}
}

func TestKafkaSchemaHandler_UpdateConfig(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		c       KafkaSchemaConfig
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaSchemaConfigResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				c: KafkaSchemaConfig{
					CompatibilityLevel: "FULL",
				},
			},
			&KafkaSchemaConfigResponse{
				APIResponse: APIResponse{},
				KafkaSchemaConfig: KafkaSchemaConfig{
					CompatibilityLevel: "FULL",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.UpdateConfig(tt.args.project, tt.args.service, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_AddSubject(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
	defer tearDown(t)

	schema := `
		{
				"doc": "example",
				"fields": [{
					"default": 5,
					"doc": "my test number",
					"name": "test",
					"namespace": "test",
					"type": "int"
				}],
				"name": "example",
				"namespace": "example",
				"type": "record"
			}`

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
		subject KafkaSchemaSubject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaSchemaSubjectResponse
		wantErr bool
	}{
		{
			"no-versions",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-schema-no-versions",
				subject: NewKafkaSchema(schema),
			},
			&KafkaSchemaSubjectResponse{
				Id: 1,
			},
			false,
		},
		{
			"has-versions",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-schema",
				subject: NewKafkaSchema(schema),
			},
			&KafkaSchemaSubjectResponse{
				Id: 5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.AddSubject(tt.args.project, tt.args.service, tt.args.name, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddSubject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_DeleteSubjectVersions(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
		version int
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
				name:    "test-schema",
				version: 5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			if err := h.DeleteSubjectVersions(tt.args.project, tt.args.service, tt.args.name, tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("DeleteSubjectVersions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKafkaSchemaHandler_GetConfig(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
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
		want    *KafkaSchemaConfigResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&KafkaSchemaConfigResponse{
				APIResponse: APIResponse{},
				KafkaSchemaConfig: KafkaSchemaConfig{
					CompatibilityLevel: "FULL",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.GetConfig(tt.args.project, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_GetSubject(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
		version int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *KafkaSchemaSubjectResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-schema",
				version: 5,
			},
			&KafkaSchemaSubjectResponse{
				APIResponse: APIResponse{},
				Id:          5,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.GetSubject(tt.args.project, tt.args.service, tt.args.name, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_GetSubjectVersions(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
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
		want    *KafkaSchemaSubjectVersionResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-schema",
			},
			&KafkaSchemaSubjectVersionResponse{
				APIResponse: APIResponse{},
				KafkaSchemaSubjectVersion: KafkaSchemaSubjectVersion{
					Versions: []int{1, 2, 3, 4},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.GetSubjectVersions(tt.args.project, tt.args.service, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubjectVersions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubjectVersions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_GetSubjects(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
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
		want    *KafkaSchemaSubjectsResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&KafkaSchemaSubjectsResponse{
				APIResponse: APIResponse{},
				KafkaSchemaSubjects: KafkaSchemaSubjects{
					Subjects: []string{"testSb1", "testSb2"},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.GetSubjects(tt.args.project, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSubjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKafkaSchemaHandler_ValidateSchema(t *testing.T) {
	c, tearDown := setupKafkaSchemasTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		name    string
		version int
		subject KafkaSchemaSubject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				name:    "test-schema",
				version: 4,
				subject: NewKafkaSchema(`
				{
					"doc": "example",
					"fields": [{
						"default": 5,
						"doc": "my test number",
						"name": "test",
						"namespace": "test",
						"type": "int"
					}],
					"name": "example",
					"namespace": "example",
					"type": "record"
				}`),
			},
			true,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &KafkaSchemaHandler{
				client: tt.fields.client,
			}
			got, err := h.ValidateSchema(tt.args.project, tt.args.service, tt.args.name, tt.args.version, tt.args.subject)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateSchema() got = %v, want %v", got, tt.want)
			}
		})
	}
}
