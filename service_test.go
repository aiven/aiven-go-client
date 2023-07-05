package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupServiceTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Service test case")

	const (
		UserName     = "test@aiven.io"
		UserPassword = "testabcd"
		AccessToken  = "some-random-token"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		service := ServiceResponse{
			Service: &Service{
				CloudName: "google-europe-west1",
				NodeCount: 1,
				Plan:      "hobbyist",
				Name:      "test-service",
				Type:      "kafka",
				NodeStates: []*NodeState{
					{
						Name:            "test-service-1",
						Role:            "master",
						State:           "running",
						ProgressUpdates: []ProgressUpdate{},
					},
				},
			},
		}

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
		} else if r.URL.Path == "/project/test-pr/service" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			err := json.NewEncoder(w).Encode(service)

			if err != nil {
				t.Error(err)
			}
			return
		} else if r.URL.Path == "/project/test-pr/service/test-sr" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			err := json.NewEncoder(w).Encode(service)

			if err != nil {
				t.Error(err)
			}
			return
		} else if r.URL.Path == "/project/test-pr-list/service" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(ServiceListResponse{
				APIResponse: APIResponse{},
				Services:    nil,
			})

			if err != nil {
				t.Error(err)
			}
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))

	apiUrl = ts.URL

	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown ElasticsearchACLs test case")
	}
}

func TestServicesHandler_Create(t *testing.T) {
	c, _ := setupServiceTestCase(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		req     CreateServiceRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Service
		wantErr bool
	}{
		{
			"error-expected",
			fields{client: c},
			args{
				project: "test-pr-wrong",
				req: CreateServiceRequest{
					Cloud:       "google-europe-west1",
					GroupName:   "default",
					Plan:        "hobbyist",
					ServiceName: "test-service",
					ServiceType: "kafka",
				},
			},
			nil,
			true,
		},
		{
			"normal",
			fields{client: c},
			args{
				project: "test-pr",
				req: CreateServiceRequest{
					Cloud:       "google-europe-west1",
					GroupName:   "default",
					Plan:        "hobbyist",
					ServiceName: "test-service",
					ServiceType: "kafka",
				},
			},
			&Service{
				CloudName: "google-europe-west1",
				NodeCount: 1,
				Plan:      "hobbyist",
				Name:      "test-service",
				Type:      "kafka",
				NodeStates: []*NodeState{
					{
						Name:            "test-service-1",
						Role:            "master",
						State:           "running",
						ProgressUpdates: []ProgressUpdate{},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ServicesHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.project, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServicesHandler_Get(t *testing.T) {
	c, _ := setupServiceTestCase(t)

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
		want    *Service
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{client: c},
			args:   args{project: "test-pr", service: "test-sr"},
			want: &Service{
				CloudName: "google-europe-west1",
				NodeCount: 1,
				Plan:      "hobbyist",
				Name:      "test-service",
				Type:      "kafka",
				NodeStates: []*NodeState{
					{
						Name:            "test-service-1",
						Role:            "master",
						State:           "running",
						ProgressUpdates: []ProgressUpdate{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ServicesHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.project, tt.args.service)
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

func TestServicesHandler_Update(t *testing.T) {
	c, _ := setupServiceTestCase(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		req     UpdateServiceRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Service
		wantErr bool
	}{
		{
			name:   "normal",
			fields: fields{client: c},
			args: args{project: "test-pr", service: "test-sr", req: UpdateServiceRequest{
				Cloud:     "google-europe-west1",
				GroupName: "default",
				Plan:      "hobbyist",
			}},
			want: &Service{
				CloudName: "google-europe-west1",
				NodeCount: 1,
				Plan:      "hobbyist",
				Name:      "test-service",
				Type:      "kafka",
				NodeStates: []*NodeState{
					&NodeState{
						Name:            "test-service-1",
						Role:            "master",
						State:           "running",
						ProgressUpdates: []ProgressUpdate{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ServicesHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.project, tt.args.service, tt.args.req)
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

func TestServicesHandler_Delete(t *testing.T) {
	c, _ := setupServiceTestCase(t)

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
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{client: c},
			args: args{
				project: "test-pr",
				service: "test-sr",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ServicesHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.project, tt.args.service); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServicesHandler_List(t *testing.T) {
	c, _ := setupServiceTestCase(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*Service
		wantErr bool
	}{
		{
			name:    "",
			fields:  fields{client: c},
			args:    args{project: "test-pr-list"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ServicesHandler{
				client: tt.fields.client,
			}
			got, err := h.List(tt.args.project)
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
