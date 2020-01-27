package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupAccountTeamProjectsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Account Team Projects test case")

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

		if r.URL.Path == "/account/a28707e316df/team/at28707ea77e2/projects" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountTeamProjectsResponse{
				Projects: []AccountTeamProject{
					{
						ProjectName: "test-pr",
						TeamType:    "admin",
					},
				},
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/account/a28707e316df/team/at28707ea77e2/project/test-pr" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(struct {
				APIResponse
				Message string `json:"message"`
			}{
				Message: "associated",
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

	}))

	apiurl = ts.URL

	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown Accounts Team Projects test case")
	}
}

func TestAccountTeamProjectsHandler_List(t *testing.T) {
	c, tearDown := setupAccountTeamProjectsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountTeamProjectsResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
			},
			&AccountTeamProjectsResponse{
				APIResponse: APIResponse{},
				Projects: []AccountTeamProject{
					{
						ProjectName: "test-pr",
						TeamType:    "admin",
					},
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
			},
			nil,
			true,
		},
		{
			"empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamProjectsHandler{
				client: tt.fields.client,
			}
			got, err := h.List(tt.args.accountId, tt.args.teamId)
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

func TestAccountTeamProjectsHandler_Create(t *testing.T) {
	c, tearDown := setupAccountTeamProjectsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
		p         AccountTeamProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "admin",
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "admin",
				},
			},
			true,
		},
		{
			"empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "admin",
				},
			},
			true,
		},
		{
			"empty-project-name",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "",
					TeamType:    "admin",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamProjectsHandler{
				client: tt.fields.client,
			}
			if err := h.Create(tt.args.accountId, tt.args.teamId, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountTeamProjectsHandler_Update(t *testing.T) {
	c, tearDown := setupAccountTeamProjectsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
		p         AccountTeamProject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "developer",
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "admin",
				},
			},
			true,
		},
		{
			"empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "",
				p: AccountTeamProject{
					ProjectName: "test-pr",
					TeamType:    "admin",
				},
			},
			true,
		},
		{
			"empty-project-name",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				p: AccountTeamProject{
					ProjectName: "",
					TeamType:    "admin",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamProjectsHandler{
				client: tt.fields.client,
			}
			if err := h.Update(tt.args.accountId, tt.args.teamId, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountTeamProjectsHandler_Delete(t *testing.T) {
	c, tearDown := setupAccountTeamProjectsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId   string
		teamId      string
		projectName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId:   "a28707e316df",
				teamId:      "at28707ea77e2",
				projectName: "test-pr",
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId:   "",
				teamId:      "at28707ea77e2",
				projectName: "test-pr",
			},
			true,
		},
		{
			"empty-team-id",
			fields{client: c},
			args{
				accountId:   "a28707e316df",
				teamId:      "",
				projectName: "test-pr",
			},
			true,
		},
		{
			"empty-project-name",
			fields{client: c},
			args{
				accountId:   "a28707e316df",
				teamId:      "at28707ea77e2",
				projectName: "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamProjectsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.accountId, tt.args.teamId, tt.args.projectName); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
