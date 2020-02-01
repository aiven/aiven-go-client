// Copyright (c) 2020 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupAccountsTeamMembersTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Account Team Members test case")

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

		if r.URL.Path == "/account/a28707e316df/team/at28707ea77e2/members" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountTeamMembersResponse{
				Members: []AccountTeamMember{
					{
						TeamId:     "at28707ea77e2",
						TeamName:   "Account Owners",
						RealName:   "Test User",
						UserId:     "u286c52034d3",
						UserEmail:  "test@example.com",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
					},
				},
			})

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/account/a28707e316df/team/at28707ea77e2/member/u286c52034d3" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountTeamMembersResponse{})

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
		t.Log("teardown Accounts Team Members test case")
		ts.Close()
	}
}

func TestAccountTeamMembersHandler_List(t *testing.T) {
	c, tearDown := setupAccountsTeamMembersTestCase(t)
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
		want    *AccountTeamMembersResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
			},
			&AccountTeamMembersResponse{
				Members: []AccountTeamMember{
					{
						TeamId:     "at28707ea77e2",
						TeamName:   "Account Owners",
						RealName:   "Test User",
						UserId:     "u286c52034d3",
						UserEmail:  "test@example.com",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
					},
				},
			},
			false,
		},
		{
			"error-empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
			},
			nil,
			true,
		},
		{
			"error-empty-team-id",
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
			h := AccountTeamMembersHandler{
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

func TestAccountTeamMembersHandler_Delete(t *testing.T) {
	c, tearDown := setupAccountsTeamMembersTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
		userId    string
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
				userId:    "u286c52034d3",
			},
			false,
		},
		{
			"error-empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
				userId:    "u286c52034d3",
			},
			true,
		},
		{
			"error-empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "",
				userId:    "u286c52034d3",
			},
			true,
		},
		{
			"error-empty-user-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				userId:    "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamMembersHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.accountId, tt.args.teamId, tt.args.userId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountTeamMembersHandler_Invite(t *testing.T) {
	c, tearDown := setupAccountsTeamMembersTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
		email     string
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
				email:     "test@example.com",
			},
			false,
		},
		{
			"error-empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				teamId:    "at28707ea77e2",
				email:     "test@example.com",
			},
			true,
		},
		{
			"error-empty-team-id",
			fields{client: c},
			args{
				accountId: "at28707ea77e2",
				teamId:    "",
				email:     "test@example.com",
			},
			true,
		},
		{
			"error-empty-email",
			fields{client: c},
			args{
				accountId: "at28707ea77e2",
				teamId:    "at28707ea77e2",
				email:     "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamMembersHandler{
				client: tt.fields.client,
			}
			if err := h.Invite(tt.args.accountId, tt.args.teamId, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
