package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupProjectTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Projects test case")

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

		if r.URL.Path == "/project/test-pr" {
			// UPDATE
			if r.Method == "PUT" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(ProjectResponse{
					APIResponse: APIResponse{},
					Project: &Project{
						AvailableCredits: "0.00",
						BillingAddress:   "",
						BillingEmails:    nil,
						BillingExtraText: "",
						Card:             Card{},
						Country:          "",
						CountryCode:      "",
						DefaultCloud:     "google-europe-east1",
						EstimatedBalance: "0.00",
						PaymentMethod:    "card",
						Name:             "test-pr",
						TechnicalEmails:  nil,
						VatID:            "",
						AccountId:        "account-id",
					},
				})

				if err != nil {
					t.Error(err)
				}
				return
			}

			// GET
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(ProjectResponse{
				APIResponse: APIResponse{},
				Project:     testGetProject(),
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/project" {
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(ProjectResponse{
					APIResponse: APIResponse{},
					Project:     testGetProject(),
				})

				if err != nil {
					t.Error(err)
				}
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(ProjectListResponse{
				APIResponse: APIResponse{},
				Projects: []*Project{
					testGetProject(),
				},
			})

			if err != nil {
				t.Error(err)
			}
			return
		}
	}))

	apiUrl = ts.URL

	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown Projects test case")
	}
}

func TestProjectsHandler_Create(t *testing.T) {
	c, tearDown := setupProjectTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		req CreateProjectRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{req: CreateProjectRequest{
				Project:   "test-pr",
				AccountId: ToStringPointer("account-id"),
			}},
			testGetProject(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ProjectsHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.req)
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

func TestProjectsHandler_Get(t *testing.T) {
	c, tearDown := setupProjectTestCase(t)
	defer tearDown(t)

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
		want    *Project
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{project: "test-pr"},
			testGetProject(),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ProjectsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.project)
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

func testGetProject() *Project {
	return &Project{
		AvailableCredits: "0.00",
		BillingAddress:   "",
		BillingEmails:    nil,
		BillingExtraText: "",
		Card:             Card{},
		Country:          "",
		CountryCode:      "",
		DefaultCloud:     "google-europe-west3",
		EstimatedBalance: "0.00",
		PaymentMethod:    "card",
		Name:             "test-pr",
		TechnicalEmails:  nil,
		VatID:            "",
		AccountId:        "account-id",
	}
}

func TestProjectsHandler_Update(t *testing.T) {
	c, tearDown := setupProjectTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		req     UpdateProjectRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Project
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{
				project: "test-pr",
				req: UpdateProjectRequest{
					Cloud: ToStringPointer("google-europe-east1"),
				},
			},
			&Project{
				AvailableCredits: "0.00",
				BillingAddress:   "",
				BillingEmails:    nil,
				BillingExtraText: "",
				Card:             Card{},
				Country:          "",
				CountryCode:      "",
				DefaultCloud:     "google-europe-east1",
				EstimatedBalance: "0.00",
				PaymentMethod:    "card",
				Name:             "test-pr",
				TechnicalEmails:  nil,
				VatID:            "",
				AccountId:        "account-id",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ProjectsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.project, tt.args.req)
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

func TestProjectsHandler_Delete(t *testing.T) {
	c, tearDown := setupProjectTestCase(t)
	defer tearDown(t)

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
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{project: "test-pr"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ProjectsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.project); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProjectsHandler_List(t *testing.T) {
	c, tearDown := setupProjectTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Project
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			[]*Project{
				testGetProject(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ProjectsHandler{
				client: tt.fields.client,
			}
			got, err := h.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("1 %T %v", got, got)
			t.Logf("2 %T %v", tt.want, tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
