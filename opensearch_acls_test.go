package aiven

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupOpenSearchACLsTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup OpenSearchACLs test case")

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

		if r.URL.Path == "/project/test-pr/service/test-sr/opensearch/acl" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(OpenSearchACLResponse{
				OpenSearchACLConfig: OpenSearchACLConfig{
					ACLs: []OpenSearchACL{
						{
							Rules: []OpenSearchACLRule{{
								Index:      "_all",
								Permission: "admin",
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

	apiUrl = ts.URL
	c, err := NewUserClient(UserName, UserPassword, "aiven-go-client-test/"+Version())
	if err != nil {
		t.Fatalf("user authentication error: %s", err)
	}

	return c, func(t *testing.T) {
		t.Log("teardown OpenSearchACLs test case")
		ts.Close()
	}
}

func TestOpenSearchACLsHandler_Update(t *testing.T) {
	c, tearDown := setupOpenSearchACLsTestCase(t)
	defer tearDown(t)

	ctx := context.Background()

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		req     OpenSearchACLRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *OpenSearchACLResponse
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: OpenSearchACLRequest{
					OpenSearchACLConfig: OpenSearchACLConfig{
						ACLs: []OpenSearchACL{
							{
								Rules: []OpenSearchACLRule{{
									Index:      "_all",
									Permission: "admin",
								}},
								Username: "test-user",
							},
						},
						Enabled:     true,
						ExtendedAcl: false,
					},
				},
			},
			&OpenSearchACLResponse{
				OpenSearchACLConfig: OpenSearchACLConfig{
					ACLs: []OpenSearchACL{
						{
							Rules: []OpenSearchACLRule{{
								Index:      "_all",
								Permission: "admin",
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
			h := &OpenSearchACLsHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(ctx, tt.args.project, tt.args.service, tt.args.req)
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

func TestOpenSearchACLsHandler_Get(t *testing.T) {
	c, tearDown := setupOpenSearchACLsTestCase(t)
	defer tearDown(t)

	ctx := context.Background()

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
		want    *OpenSearchACLResponse
		wantErr bool
	}{
		{
			"correct",
			fields{client: c},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&OpenSearchACLResponse{
				OpenSearchACLConfig: OpenSearchACLConfig{
					ACLs: []OpenSearchACL{
						{
							Rules: []OpenSearchACLRule{{
								Index:      "_all",
								Permission: "admin",
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
			h := &OpenSearchACLsHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(ctx, tt.args.project, tt.args.service)
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

func TestOpenSearchACLConfig_Add(t *testing.T) {
	type fields struct {
		ACLs        []OpenSearchACL
		Enabled     bool
		ExtendedAcl bool
	}
	type args struct {
		acl OpenSearchACL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OpenSearchACLConfig
	}{
		{
			"add-multiple",
			fields{
				ACLs: []OpenSearchACL{
					{
						Username: "test-user",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_rw",
								Permission: "write",
							},
						}},
					{
						Username: "test-user2",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_all",
								Permission: "admin",
							},
						}},
				},
			},
			args{acl: OpenSearchACL{
				Rules: []OpenSearchACLRule{
					{
						Index:      "_all",
						Permission: "admin",
					},
					{
						Index:      "_test",
						Permission: "write",
					},
				},
				Username: "test-user",
			}},
			&OpenSearchACLConfig{
				ACLs: []OpenSearchACL{
					{
						Username: "test-user",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_rw",
								Permission: "write",
							},
							{
								Index:      "_all",
								Permission: "admin",
							},
							{
								Index:      "_test",
								Permission: "write",
							},
						}},
					{
						Username: "test-user2",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_all",
								Permission: "admin",
							},
						}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := OpenSearchACLConfig{
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

func TestOpenSearchACLConfig_Delete(t *testing.T) {
	ctx := context.Background()
	type fields struct {
		ACLs        []OpenSearchACL
		Enabled     bool
		ExtendedAcl bool
	}
	type args struct {
		acl OpenSearchACL
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *OpenSearchACLConfig
	}{
		{
			"multiple",
			fields{
				ACLs: []OpenSearchACL{
					{
						Username: "test-user",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_all",
								Permission: "admin",
							},
							{
								Index:      "_rw",
								Permission: "readwrite",
							},
							{
								Index:      "_test",
								Permission: "write",
							},
						}},
					{
						Username: "test-user2",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_all",
								Permission: "admin",
							},
						}},
				},
				Enabled:     false,
				ExtendedAcl: false,
			},
			args{acl: OpenSearchACL{
				Username: "test-user",
				Rules: []OpenSearchACLRule{
					{
						Index:      "_all",
						Permission: "admin",
					},
					{
						Index:      "_rw",
						Permission: "readwrite",
					},
				}}},
			&OpenSearchACLConfig{
				ACLs: []OpenSearchACL{
					{
						Username: "test-user",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_test",
								Permission: "write",
							},
						}},
					{
						Username: "test-user2",
						Rules: []OpenSearchACLRule{
							{
								Index:      "_all",
								Permission: "admin",
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
			conf := OpenSearchACLConfig{
				ACLs:        tt.fields.ACLs,
				Enabled:     tt.fields.Enabled,
				ExtendedAcl: tt.fields.ExtendedAcl,
			}
			if got := conf.Delete(ctx, tt.args.acl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}
