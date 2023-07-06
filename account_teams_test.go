package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupAccountsTeamsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Account Teams test case")

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
		} else if r.URL.Path == "/account/a28707e316df/teams" {
			// get a list of account teams
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(AccountTeamsResponse{
					APIResponse: APIResponse{},
					Teams: []AccountTeam{
						{
							AccountId:  "a28707e316df",
							Name:       "Account Owners",
							Id:         "at28707ea77e2",
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

			//create mew account team
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(AccountTeamResponse{
					APIResponse: APIResponse{},
					Team: AccountTeam{
						AccountId:  "a28707e316df",
						Name:       "test-team-1",
						Id:         "at28761bc6348",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
					},
				})

				if err != nil {
					t.Error(err)
				}
				return
			}
		} else if r.URL.Path == "/account/a28707e316df/team/at28707ea77e2" {
			//update account team
			if r.Method == "PUT" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(AccountTeamResponse{
					APIResponse: APIResponse{},
					Team: AccountTeam{
						AccountId:  "a28707e316df",
						Name:       "new team name",
						Id:         "at28707ea77e2",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
					},
				})

				if err != nil {
					t.Error(err)
				}
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountTeamResponse{
				APIResponse: APIResponse{},
				Team: AccountTeam{
					AccountId:  "a28707e316df",
					Name:       "Account Owners",
					Id:         "at28707ea77e2",
					CreateTime: getTime(t),
					UpdateTime: getTime(t),
				},
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
		t.Log("teardown Account Teams test case")
		ts.Close()
	}
}

func TestAccountsTeamsHandler_List(t *testing.T) {
	c, tearDown := setupAccountsTeamsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountTeamsResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{accountId: "a28707e316df"},
			&AccountTeamsResponse{
				APIResponse: APIResponse{},
				Teams: []AccountTeam{
					{
						AccountId:  "a28707e316df",
						Name:       "Account Owners",
						Id:         "at28707ea77e2",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
					},
				},
			},
			false,
		},
		{
			"error",
			fields{client: c},
			args{accountId: ""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamsHandler{
				client: tt.fields.client,
			}
			got, err := h.List(tt.args.accountId)
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

func TestAccountsTeamsHandler_Get(t *testing.T) {
	c, tearDown := setupAccountsTeamsTestCase(t)
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
		want    *AccountTeamResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
			},
			&AccountTeamResponse{
				APIResponse: APIResponse{},
				Team: AccountTeam{
					AccountId:  "a28707e316df",
					Name:       "Account Owners",
					Id:         "at28707ea77e2",
					CreateTime: getTime(t),
					UpdateTime: getTime(t),
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
			h := AccountTeamsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.accountId, tt.args.teamId)
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

func TestAccountsTeamsHandler_Create(t *testing.T) {
	c, tearDown := setupAccountsTeamsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		team      AccountTeam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountTeamResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				team: AccountTeam{
					Name: "test-team-1",
				},
			},
			&AccountTeamResponse{
				APIResponse: APIResponse{},
				Team: AccountTeam{
					AccountId:  "a28707e316df",
					Name:       "test-team-1",
					Id:         "at28761bc6348",
					CreateTime: getTime(t),
					UpdateTime: getTime(t),
				},
			},
			false,
		},
		{
			"error-empty-id",
			fields{client: c},
			args{
				accountId: "",
				team: AccountTeam{
					Name: "test-team-1",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamsHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.accountId, tt.args.team)
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

func TestAccountsTeamsHandler_Delete(t *testing.T) {
	c, tearDown := setupAccountsTeamsTestCase(t)
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
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
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
			true,
		},
		{
			"error-empty-team-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.accountId, tt.args.teamId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountTeamsHandler_Update(t *testing.T) {
	c, tearDown := setupAccountsTeamsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		teamId    string
		team      AccountTeam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountTeamResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				teamId:    "at28707ea77e2",
				team: AccountTeam{
					Name: "new team name",
				},
			},
			&AccountTeamResponse{
				APIResponse: APIResponse{},
				Team: AccountTeam{
					AccountId:  "a28707e316df",
					Name:       "new team name",
					Id:         "at28707ea77e2",
					CreateTime: getTime(t),
					UpdateTime: getTime(t),
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
				team: AccountTeam{
					Name: "new team name",
				},
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
				team: AccountTeam{
					Name: "new team name",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountTeamsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.accountId, tt.args.teamId, tt.args.team)
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
