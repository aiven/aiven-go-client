package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupMirrormakerReplicationFlowTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Kafka Mirrormaker 2 Replication Flow test case")

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

		if r.URL.Path == "/project/test-pr/service/test-sr/mirrormaker/replication-flows" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`
				{
					"errors": [
					],
					"message": "Completed",
					"replication_flows": [{
						"enabled": true,
						"offset_syncs_topic_location": "source",
						"source_cluster": "kafka-sc",
						"target_cluster": "kafka-tc",
						"topics": [
								".*"
						],
						"topics.blacklist": [
								".*[\\-\\.]internal", 
								".*\\.replica", 
								"__.*"
						]
					}]
				}
			`))
			if err != nil {
				t.Error(err)
			}
			return
		}

		if r.URL.Path == "/project/test-pr/service/test-sr/mirrormaker/replication-flows/kafka-sc/kafka-tc" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`
				{
					"replication_flow": {
						"enabled": true,
						"offset_syncs_topic_location": "source",
						"source_cluster": "kafka-sc",
						"target_cluster": "kafka-tc",
						"topics": [
								".*"
						],
						"topics.blacklist": [
								".*[\\-\\.]internal", 
								".*\\.replica", 
								"__.*"
						]
					}
				}
			`))
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
		t.Log("teardown Kafka Mirrormaker 2 Replication Flow test case")
		ts.Close()
	}
}

func TestMirrorMakerReplicationFlowHandler_Create(t *testing.T) {
	c, tearDown := setupMirrormakerReplicationFlowTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project string
		service string
		req     MirrorMakerReplicationFlowRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"basic",
			fields{c},
			args{
				project: "test-pr",
				service: "test-sr",
				req: MirrorMakerReplicationFlowRequest{
					ReplicationFlow{
						Enabled:       true,
						SourceCluster: "kafka-sc",
						TargetCluster: "kafka-tc",
						Topics: []string{
							".*",
						},
						TopicsBlacklist: []string{
							".*[\\-\\.]internal",
							".*\\.replica",
							"__.*",
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MirrorMakerReplicationFlowHandler{
				client: tt.fields.client,
			}
			err := h.Create(tt.args.project, tt.args.service, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMirrorMakerReplicationFlowHandler_Update(t *testing.T) {
	c, tearDown := setupMirrormakerReplicationFlowTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project       string
		service       string
		sourceCluster string
		targetCluster string
		req           MirrorMakerReplicationFlowRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MirrorMakerReplicationFlowResponse
		wantErr bool
	}{
		{
			"basic",
			fields{
				client: c,
			},
			args{
				project:       "test-pr",
				service:       "test-sr",
				sourceCluster: "kafka-sc",
				targetCluster: "kafka-tc",
				req: MirrorMakerReplicationFlowRequest{
					ReplicationFlow{
						Enabled:       true,
						SourceCluster: "kafka-sc",
						TargetCluster: "kafka-tc",
						Topics: []string{
							".*",
						},
						TopicsBlacklist: []string{
							".*[\\-\\.]internal",
							".*\\.replica",
							"__.*",
						},
					},
				},
			},
			&MirrorMakerReplicationFlowResponse{
				ReplicationFlow: ReplicationFlow{
					Enabled:                  true,
					OffsetSyncsTopicLocation: "source",
					SourceCluster:            "kafka-sc",
					TargetCluster:            "kafka-tc",
					Topics: []string{
						".*",
					},
					TopicsBlacklist: []string{
						".*[\\-\\.]internal",
						".*\\.replica",
						"__.*",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MirrorMakerReplicationFlowHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.project, tt.args.service, tt.args.sourceCluster, tt.args.targetCluster, tt.args.req)
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

func TestMirrorMakerReplicationFlowHandler_List(t *testing.T) {
	c, tearDown := setupMirrormakerReplicationFlowTestCase(t)
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
		want    *MirrorMakerReplicationFlowsResponse
		wantErr bool
	}{
		{
			"basic",
			fields{
				client: c,
			},
			args{
				project: "test-pr",
				service: "test-sr",
			},
			&MirrorMakerReplicationFlowsResponse{
				APIResponse: APIResponse{
					Message: "Completed",
					Errors:  []Error{},
				},
				ReplicationFlows: []ReplicationFlow{
					{
						Enabled:                  true,
						OffsetSyncsTopicLocation: "source",
						SourceCluster:            "kafka-sc",
						TargetCluster:            "kafka-tc",
						Topics: []string{
							".*",
						},
						TopicsBlacklist: []string{
							".*[\\-\\.]internal",
							".*\\.replica",
							"__.*",
						},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MirrorMakerReplicationFlowHandler{
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

func TestMirrorMakerReplicationFlowHandler_Delete(t *testing.T) {
	c, tearDown := setupMirrormakerReplicationFlowTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project       string
		service       string
		sourceCluster string
		targetCluster string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"basic",
			fields{
				client: c,
			},
			args{
				project:       "test-pr",
				service:       "test-sr",
				sourceCluster: "kafka-sc",
				targetCluster: "kafka-tc",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MirrorMakerReplicationFlowHandler{
				client: tt.fields.client,
			}
			if err := h.Delete(tt.args.project, tt.args.service, tt.args.sourceCluster, tt.args.targetCluster); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMirrorMakerReplicationFlowHandler_Get(t *testing.T) {
	c, tearDown := setupMirrormakerReplicationFlowTestCase(t)
	defer tearDown(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project       string
		service       string
		sourceCluster string
		targetCluster string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *MirrorMakerReplicationFlowResponse
		wantErr bool
	}{
		{
			"basic",
			fields{c},
			args{
				project:       "test-pr",
				service:       "test-sr",
				sourceCluster: "kafka-sc",
				targetCluster: "kafka-tc",
			},
			&MirrorMakerReplicationFlowResponse{
				ReplicationFlow: ReplicationFlow{
					Enabled:                  true,
					OffsetSyncsTopicLocation: "source",
					SourceCluster:            "kafka-sc",
					TargetCluster:            "kafka-tc",
					Topics: []string{
						".*",
					},
					TopicsBlacklist: []string{
						".*[\\-\\.]internal",
						".*\\.replica",
						"__.*",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &MirrorMakerReplicationFlowHandler{
				client: tt.fields.client,
			}
			got, err := h.Get(tt.args.project, tt.args.service, tt.args.sourceCluster, tt.args.targetCluster)
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
