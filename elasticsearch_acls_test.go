package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupElasticsearchACLsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup ElasticsearchACLs test case")

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

		if r.URL.Path == "/project/test-pr/service/test-sr/elasticsearch/acl" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(ElasticSearchACLResponse{
				ElasticSearchACLConfig: ElasticSearchACLConfig{
					ACLs: []ElasticSearchACL{
						{
							Rules: []ElasticsearchACLRule{{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							}},
							Username: "test-user",
						},
					},
					Enabled:     true,
					ExtendedAcl: false,
				}})

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
		t.Log("teardown ElasticsearchACLs test case")
	}
}

func TestElasticSearchACLsHandler_Create(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		req     ElasticsearchACLRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ElasticSearchACLResponse
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: ElasticsearchACLRequest{
					ElasticSearchACLConfig: ElasticSearchACLConfig{
						ACLs: []ElasticSearchACL{
							{
								Rules: []ElasticsearchACLRule{{
									Index:      "_all",
									Permission: ElasticsearchAclPermissionAdmin,
								}},
								Username: "test-user",
							},
						},
						Enabled:     true,
						ExtendedAcl: false,
					},
				},
			},
			&ElasticSearchACLResponse{
				ElasticSearchACLConfig: ElasticSearchACLConfig{
					ACLs: []ElasticSearchACL{
						{
							Rules: []ElasticsearchACLRule{{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							}},
							Username: "test-user",
						},
					},
					Enabled:     true,
					ExtendedAcl: false,
				},
			},
			false,
		},
		{
			"empty-username",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: ElasticsearchACLRequest{
					ElasticSearchACLConfig: ElasticSearchACLConfig{
						ACLs: []ElasticSearchACL{
							{
								Rules: []ElasticsearchACLRule{{
									Index:      "_all",
									Permission: ElasticsearchAclPermissionAdmin,
								}},
								Username: "",
							},
						},
						Enabled:     true,
						ExtendedAcl: false,
					},
				},
			},
			nil,
			true,
		},
		{
			"empty-rules",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: ElasticsearchACLRequest{
					ElasticSearchACLConfig: ElasticSearchACLConfig{
						ACLs: []ElasticSearchACL{
							{
								Rules:    []ElasticsearchACLRule{},
								Username: "test-user",
							},
						},
						Enabled:     true,
						ExtendedAcl: false,
					},
				},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
				client: tt.fields.client,
			}
			got, err := h.Create(tt.args.project, tt.args.service, tt.args.req)
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

func TestElasticSearchACLsHandler_Update(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		req     ElasticsearchACLRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ElasticSearchACLResponse
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: ElasticsearchACLRequest{
					ElasticSearchACLConfig: ElasticSearchACLConfig{
						ACLs: []ElasticSearchACL{
							{
								Rules: []ElasticsearchACLRule{{
									Index:      "_all",
									Permission: ElasticsearchAclPermissionAdmin,
								}},
								Username: "test-user",
							},
						},
						Enabled:     true,
						ExtendedAcl: false,
					},
				},
			},
			&ElasticSearchACLResponse{
				ElasticSearchACLConfig: ElasticSearchACLConfig{
					ACLs: []ElasticSearchACL{
						{
							Rules: []ElasticsearchACLRule{{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							}},
							Username: "test-user",
						},
					},
					Enabled:     true,
					ExtendedAcl: false,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
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

func TestElasticSearchACLsHandler_response(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		response *ElasticSearchACLResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{response: &ElasticSearchACLResponse{
				APIResponse:            APIResponse{},
				ElasticSearchACLConfig: ElasticSearchACLConfig{},
			}},
			false,
		},
		{
			"error-in-response",
			fields{client: c},
			args{response: &ElasticSearchACLResponse{
				APIResponse: APIResponse{
					Errors: []Error{
						{
							Message:  "test-er-message",
							MoreInfo: "more_info",
							Status:   500,
						},
					},
				},
				ElasticSearchACLConfig: ElasticSearchACLConfig{},
			}},
			true,
		},
		{
			"empty",
			fields{client: c},
			args{response: nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
				client: tt.fields.client,
			}

			b, _ := json.Marshal(tt.args.response)
			_, err := h.response(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("response() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

}

func TestElasticSearchACLsHandler_List(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
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
		want    *ElasticSearchACLResponse
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&ElasticSearchACLResponse{
				ElasticSearchACLConfig: ElasticSearchACLConfig{
					ACLs: []ElasticSearchACL{
						{
							Rules: []ElasticsearchACLRule{{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							}},
							Username: "test-user",
						},
					},
					Enabled:     true,
					ExtendedAcl: false,
				},
				APIResponse: APIResponse{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
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

func TestElasticSearchACLsHandler_Add(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		acl     ElasticSearchACL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"add-empty",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				acl:     ElasticSearchACL{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
				client: tt.fields.client,
			}
			if err := h.Add(tt.args.project, tt.args.service, tt.args.acl); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestElasticSearchACLsHandler_Delete(t *testing.T) {
	c, tearDown := setupElasticsearchACLsTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		acl     ElasticSearchACL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"empty",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				acl:     ElasticSearchACL{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ElasticSearchACLsHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.project, tt.args.service, tt.args.acl); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestElasticSearchACLConfig_Check(t *testing.T) {
	type fields struct {
		ACLs        []ElasticSearchACL
		Enabled     bool
		ExtendedAcl bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"correct",
			fields{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionReadAndWrite,
							},
						},
					},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
			false,
		},
		{
			"correct",
			fields{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: "random-text...",
							},
						},
					},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ElasticSearchACLConfig{
				ACLs:        tt.fields.ACLs,
				Enabled:     tt.fields.Enabled,
				ExtendedAcl: tt.fields.ExtendedAcl,
			}
			if err := conf.Check(); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestElasticSearchACLConfig_Add(t *testing.T) {
	type fields struct {
		ACLs        []ElasticSearchACL
		Enabled     bool
		ExtendedAcl bool
	}
	type args struct {
		acl ElasticSearchACL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ElasticSearchACLConfig
	}{
		{
			"add-multiple",
			fields{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_rw",
								Permission: ElasticsearchAclPermissionReadAndWrite,
							},
						}},
					{
						Username: "test-user2",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
						}},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
			args{acl: ElasticSearchACL{
				Rules: []ElasticsearchACLRule{
					{
						Index:      "_all",
						Permission: ElasticsearchAclPermissionAdmin,
					},
					{
						Index:      "_test",
						Permission: ElasticsearchAclPermissionWrite,
					},
				},
				Username: "test-user",
			}},
			ElasticSearchACLConfig{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_rw",
								Permission: ElasticsearchAclPermissionReadAndWrite,
							},
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
							{
								Index:      "_test",
								Permission: ElasticsearchAclPermissionWrite,
							},
						}},
					{
						Username: "test-user2",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
						}},
				},
				Enabled:     true,
				ExtendedAcl: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ElasticSearchACLConfig{
				ACLs:        tt.fields.ACLs,
				Enabled:     tt.fields.Enabled,
				ExtendedAcl: tt.fields.ExtendedAcl,
			}
			if got := conf.Add(tt.args.acl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestElasticSearchACLConfig_Delete(t *testing.T) {
	type fields struct {
		ACLs        []ElasticSearchACL
		Enabled     bool
		ExtendedAcl bool
	}
	type args struct {
		acl ElasticSearchACL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ElasticSearchACLConfig
	}{
		{
			"multiple",
			fields{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
							{
								Index:      "_rw",
								Permission: ElasticsearchAclPermissionReadAndWrite,
							},
							{
								Index:      "_test",
								Permission: ElasticsearchAclPermissionWrite,
							},
						}},
					{
						Username: "test-user2",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
						}},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
			args{acl: ElasticSearchACL{
				Username: "test-user",
				Rules: []ElasticsearchACLRule{
					{
						Index:      "_all",
						Permission: ElasticsearchAclPermissionAdmin,
					},
					{
						Index:      "_rw",
						Permission: ElasticsearchAclPermissionReadAndWrite,
					},
				}}},
			ElasticSearchACLConfig{
				ACLs: []ElasticSearchACL{
					{
						Username: "test-user",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_test",
								Permission: ElasticsearchAclPermissionWrite,
							},
						}},
					{
						Username: "test-user2",
						Rules: []ElasticsearchACLRule{
							{
								Index:      "_all",
								Permission: ElasticsearchAclPermissionAdmin,
							},
						}},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := ElasticSearchACLConfig{
				ACLs:        tt.fields.ACLs,
				Enabled:     tt.fields.Enabled,
				ExtendedAcl: tt.fields.ExtendedAcl,
			}
			if got := conf.Delete(tt.args.acl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
