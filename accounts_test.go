package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func setupAccountsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Accounts test case")

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

		if r.URL.Path == "/account" {
			// get a list of account
			if r.Method == "GET" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(AccountsResponse{
					APIResponse: APIResponse{},
					Accounts: []Account{
						{
							Id:             "a28707e316df",
							Name:           "test@aiven.fi",
							OwnerTeamId:    "at28707cc9aa3",
							CreateTime:     getTime(t),
							UpdateTime:     getTime(t),
							BillingEnabled: false,
							TenantId:       "aiven",
						},
					},
				})

				if err != nil {
					t.Error(err)
				}
				return
			}

			// create mew account
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(getAccountResponse(t))

				if err != nil {
					t.Error(err)
				}
				return
			}
		}

		// get account by id
		if r.URL.Path == "/account/a28707e316df" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(getAccountResponse(t))

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
		t.Log("teardown Accounts test case")
		ts.Close()
	}
}

func getTime(t *testing.T) *time.Time {
	var err error
	tt, err := time.Parse(time.RFC3339, "2020-01-15T10:35:33Z")
	if err != nil {
		t.Error(err)
	}

	return &tt
}

func TestAccountsHandler_List(t *testing.T) {
	c, tearDown := setupAccountsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *AccountsResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			&AccountsResponse{
				APIResponse: APIResponse{},
				Accounts: []Account{
					{
						Id:             "a28707e316df",
						Name:           "test@aiven.fi",
						OwnerTeamId:    "at28707cc9aa3",
						CreateTime:     getTime(t),
						UpdateTime:     getTime(t),
						BillingEnabled: false,
						TenantId:       "aiven",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountsHandler{
				client: tt.fields.client,
			}
			got, err := h.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountsHandler_Get(t *testing.T) {
	c, tearDown := setupAccountsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{id: "a28707e316df"},
			getAccountResponse(t),
			false,
		},
		{
			"error-empty-id",
			fields{client: c},
			args{id: ""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.id)
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

func getAccountResponse(t *testing.T) *AccountResponse {
	return &AccountResponse{
		APIResponse: APIResponse{},
		Account: Account{
			Id:             "a28707e316df",
			Name:           "test@aiven.fi",
			CreateTime:     getTime(t),
			UpdateTime:     getTime(t),
			BillingEnabled: false,
			TenantId:       "aiven",
		},
	}
}

func TestAccountsHandler_Create(t *testing.T) {
	c, tearDown := setupAccountsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		account Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountResponse
		wantErr bool
	}{
		{
			"",
			fields{client: c},
			args{account: Account{
				Name: "test@aiven.fi",
			}},
			getAccountResponse(t),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountsHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.account)
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

func TestAccountsHandler_Update(t *testing.T) {
	c, tearDown := setupAccountsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		id      string
		account Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				id: "a28707e316df",
				account: Account{
					Name: "test@aiven.fi",
				},
			},
			getAccountResponse(t),
			false,
		},
		{
			"error-empty-id",
			fields{client: c},
			args{
				id: "",
				account: Account{
					Name: "test@aiven.fi",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.id, tt.args.account)
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

func TestAccountsHandler_Delete(t *testing.T) {
	c, tearDown := setupAccountsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		id string
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
			args{id: "a28707e316df"},
			false,
		},
		{
			"error-empty-id",
			fields{client: c},
			args{id: ""},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
