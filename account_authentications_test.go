package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupAccountAuthenticationsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Account Authentications test case")

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

		if r.URL.Path == "/account/a28707e316df/authentication/am28707eb0055" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			var err error
			switch r.Method {
			case "PUT":
				err = json.NewEncoder(w).Encode(AccountAuthenticationResponse{
					APIResponse: APIResponse{},
					AuthenticationMethod: AccountAuthenticationMethod{
						AccountID:  "a28707e316df",
						State:      "active",
						CreateTime: getTime(t),
						UpdateTime: getTime(t),
						DeleteTime: getTime(t),
					},
				})
			default:
				err = json.NewEncoder(w).Encode(AccountAuthenticationResponse{
					APIResponse: APIResponse{},
					AuthenticationMethod: AccountAuthenticationMethod{
						AccountID:                   "a28707e316df",
						AuthenticationMethodEnabled: true,
						AuthenticationMethodName:    "test",
						AuthenticationMethodType:    "saml",
						State:                       "active",
						CreateTime:                  getTime(t),
						UpdateTime:                  getTime(t),
						DeleteTime:                  getTime(t),
					},
				})
			}

			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/account/a28707e316df/authentication" {
			if r.Method == "POST" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(AccountAuthenticationResponse{
					APIResponse: APIResponse{},
					AuthenticationMethod: AccountAuthenticationMethod{
						AccountID:                   "a28707e316df",
						AuthenticationMethodEnabled: true,
						AuthenticationMethodName:    "test",
						AuthenticationMethodType:    "saml",
						State:                       "active",
						CreateTime:                  getTime(t),
						UpdateTime:                  getTime(t),
						DeleteTime:                  getTime(t),
					},
				})

				if err != nil {
					t.Error(err)
				}
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(AccountAuthenticationListResponse{
				AuthenticationMethods: []AccountAuthenticationMethod{
					{
						AccountID:                   "a28707e316df",
						AuthenticationMethodID:      "am28707eb0055",
						AuthenticationMethodEnabled: true,
						AuthenticationMethodName:    "Platform authentication",
						AuthenticationMethodType:    "internal",
						State:                       "active",
						CreateTime:                  getTime(t),
						UpdateTime:                  getTime(t),
						DeleteTime:                  getTime(t),
					},
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
		t.Log("teardown Accounts Authentications test case")
		ts.Close()
	}
}

func TestAccountAuthenticationsHandler_List(t *testing.T) {
	c, tearDown := setupAccountAuthenticationsTestCase(t)
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
		want    *AccountAuthenticationListResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{accountId: "a28707e316df"},
			&AccountAuthenticationListResponse{
				APIResponse: APIResponse{},
				AuthenticationMethods: []AccountAuthenticationMethod{
					{
						AccountID:                   "a28707e316df",
						AuthenticationMethodEnabled: true,
						AuthenticationMethodID:      "am28707eb0055",
						AuthenticationMethodName:    "Platform authentication",
						AuthenticationMethodType:    "internal",
						State:                       "active",
						CreateTime:                  getTime(t),
						UpdateTime:                  getTime(t),
						DeleteTime:                  getTime(t),
					},
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{accountId: ""},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountAuthenticationsHandler{
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

func TestAccountAuthenticationsHandler_Create(t *testing.T) {
	c, tearDown := setupAccountAuthenticationsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId         string
		accountAuthMethId string
		a                 AccountAuthenticationMethodCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountAuthenticationResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId:         "a28707e316df",
				accountAuthMethId: "am28707eb0055",
				a: AccountAuthenticationMethodCreate{
					AuthenticationMethodName: "test1",
					AuthenticationMethodType: "saml",
				},
			},
			&AccountAuthenticationResponse{
				APIResponse: APIResponse{},
				AuthenticationMethod: AccountAuthenticationMethod{
					AccountID:                   "a28707e316df",
					AuthenticationMethodEnabled: true,
					AuthenticationMethodName:    "test",
					AuthenticationMethodType:    "saml",
					State:                       "active",
					CreateTime:                  getTime(t),
					UpdateTime:                  getTime(t),
					DeleteTime:                  getTime(t),
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				a: AccountAuthenticationMethodCreate{
					AuthenticationMethodName: "test1",
					AuthenticationMethodType: "saml",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountAuthenticationsHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.accountId, tt.args.a)
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

func TestAccountAuthenticationsHandler_Update(t *testing.T) {
	c, tearDown := setupAccountAuthenticationsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId         string
		accountAuthMethId string
		a                 AccountAuthenticationMethodUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountAuthenticationResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId:         "a28707e316df",
				accountAuthMethId: "am28707eb0055",
				a: AccountAuthenticationMethodUpdate{
					AuthenticationMethodEnabled: true,
					AuthenticationMethodName:    "test",
				},
			},
			&AccountAuthenticationResponse{
				APIResponse: APIResponse{},
				AuthenticationMethod: AccountAuthenticationMethod{
					AccountID:  "a28707e316df",
					State:      "active",
					CreateTime: getTime(t),
					UpdateTime: getTime(t),
					DeleteTime: getTime(t),
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId:         "",
				accountAuthMethId: "am28707eb0055",
				a: AccountAuthenticationMethodUpdate{
					AuthenticationMethodEnabled: true,
					AuthenticationMethodName:    "test",
				},
			},
			nil,
			true,
		},
		{
			"empty-id",
			fields{client: c},
			args{
				accountId:         "a28707e316df",
				accountAuthMethId: "",
				a: AccountAuthenticationMethodUpdate{
					AuthenticationMethodEnabled: true,
					AuthenticationMethodName:    "test",
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountAuthenticationsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.accountId, tt.args.accountAuthMethId, tt.args.a)
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

func TestAccountAuthenticationsHandler_Delete(t *testing.T) {
	c, tearDown := setupAccountAuthenticationsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		authId    string
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
				authId:    "am28707eb0055",
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				authId:    "am28707eb0055",
			},
			true,
		}, {
			"empty-auth-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				authId:    "",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountAuthenticationsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.accountId, tt.args.authId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccountAuthenticationsHandler_Get(t *testing.T) {
	c, tearDown := setupAccountAuthenticationsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		accountId string
		authId    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AccountAuthenticationResponse
		wantErr bool
	}{
		{
			"normal",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				authId:    "am28707eb0055",
			},
			&AccountAuthenticationResponse{
				APIResponse: APIResponse{},
				AuthenticationMethod: AccountAuthenticationMethod{
					AccountID:                   "a28707e316df",
					AuthenticationMethodEnabled: true,
					AuthenticationMethodName:    "test",
					AuthenticationMethodType:    "saml",
					State:                       "active",
					CreateTime:                  getTime(t),
					UpdateTime:                  getTime(t),
					DeleteTime:                  getTime(t),
				},
			},
			false,
		},
		{
			"empty-account-id",
			fields{client: c},
			args{
				accountId: "",
				authId:    "am28707eb0055",
			},
			nil,
			true,
		}, {
			"empty-auth-id",
			fields{client: c},
			args{
				accountId: "a28707e316df",
				authId:    "",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AccountAuthenticationsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.accountId, tt.args.authId)
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
