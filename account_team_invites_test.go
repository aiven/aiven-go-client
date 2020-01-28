package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupAccountTeamInvitesTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Account Team Invites test case")

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

		if r.URL.Path == "/account/a28707e316df/team/b28707e316df/invites" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountTeamInvitesResponse{
				APIResponse: APIResponse{},
				Invites: []AccountTeamInvite{
					{
						AccountId:          "a28707e316df",
						AccountName:        "test2@aiven.fi",
						InvitedByUserEmail: "test1@aiven.fi",
						TeamId:             "b28707e316df",
						TeamName:           "Account Owners",
						UserEmail:          "test_invite_sent_to@aiven.fi",
						CreateTime:         getTime(t),
					},
				},
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
		t.Log("teardown Accounts Team Invites test case")
	}
}

func TestAccountTeamInvitesHandler_List(t *testing.T) {
	c, tearDown := setupAccountTeamInvitesTestCase(t)
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
		want    *AccountTeamInvitesResponse
		wantErr bool
	}{
		{
			"basic",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "b28707e316df",
			},
			&AccountTeamInvitesResponse{
				APIResponse: APIResponse{},
				Invites: []AccountTeamInvite{
					{
						AccountId:          "a28707e316df",
						AccountName:        "test2@aiven.fi",
						InvitedByUserEmail: "test1@aiven.fi",
						TeamId:             "b28707e316df",
						TeamName:           "Account Owners",
						UserEmail:          "test_invite_sent_to@aiven.fi",
						CreateTime:         getTime(t),
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
				teamId:    "b28707e316df",
			},
			nil,
			true,
		},
		{
			"empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316dfs",
				teamId:    "",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamInvitesHandler{
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
