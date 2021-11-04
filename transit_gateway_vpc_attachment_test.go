package aiven

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupTransitGatewayVPCAttachmentTestCase(t *testing.T) (*Client, func(t *testing.T)) {
	t.Log("setup Service test case")

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

		vpc := VPC{
			CloudName:    "aws-eu-west-1",
			NetworkCIDR:  "172.16.0.0/24",
			ProjectVPCID: "test-vpc-id",
			State:        "ACTIVE",
		}

		t.Log(r.URL.Path)

		if r.URL.Path == "/project/test-pr/vpcs/test-vpc-id/user-peer-network-cidrs" {
			w.Header().Set("Content-Type", "application/json")

			if r.Method == http.MethodPut {
				var req TransitGatewayVPCAttachmentRequest

				err := json.NewDecoder(r.Body).Decode(&req)
				if err != nil {
					t.Error(err)
				}
				if req.Add == nil {
					errorResponse(t, w, http.StatusBadRequest, "Invalid input for add: None is not of type 'array'")
					return
				}
				if req.Delete == nil {
					errorResponse(t, w, http.StatusBadRequest, "Invalid input for delete: None is not of type 'array'")
					return
				}
				if len(req.Add) == 0 && len(req.Delete) == 0 {
					errorResponse(t, w, http.StatusBadRequest, "Both add and delete cidr sets cannot be empty")
					return
				}
				for _, a := range req.Add {
					if a.PeerResourceGroup != nil && *a.PeerResourceGroup == "" {
						errorResponse(t, w, http.StatusBadRequest, "peer_resource_group must always be None for cloud aws-eu-west-1")
						return
					}
				}
			}

			w.WriteHeader(http.StatusOK)

			err := json.NewEncoder(w).Encode(vpc)

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
		t.Log("teardown TransitGatewayVPCAttachment test case")
	}
}

func errorResponse(t *testing.T, w http.ResponseWriter, statusCode int, msg string) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(Error{
		Status:  statusCode,
		Message: msg,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestTransitGatewayVPCAttachmentHandler_Update(t *testing.T) {
	c, _ := setupTransitGatewayVPCAttachmentTestCase(t)

	type fields struct {
		client *Client
	}
	type args struct {
		project      string
		projectVPCId string
		req          TransitGatewayVPCAttachmentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *VPC
		wantErr bool
	}{
		{
			"bad-add",
			fields{client: c},
			args{
				"test-pr",
				"test-vpc-id",
				TransitGatewayVPCAttachmentRequest{
					Add:    nil,
					Delete: []string{},
				},
			},
			nil,
			true,
		},
		{
			"bad-del",
			fields{client: c},
			args{
				"test-pr",
				"test-vpc-id",
				TransitGatewayVPCAttachmentRequest{
					Add:    []TransitGatewayVPCAttachment{},
					Delete: nil,
				},
			},
			nil,
			true,
		},
		{
			"empty-add-del",
			fields{client: c},
			args{
				"test-pr",
				"test-vpc-id",
				TransitGatewayVPCAttachmentRequest{
					Add:    []TransitGatewayVPCAttachment{},
					Delete: []string{},
				},
			},
			nil,
			true,
		},
		{
			"empty-peer-resource-group",
			fields{client: c},
			args{
				"test-pr",
				"test-vpc-id",
				TransitGatewayVPCAttachmentRequest{
					Add: []TransitGatewayVPCAttachment{
						{
							CIDR:              "10.0.0.0/24",
							PeerCloudAccount:  "111222333444",
							PeerVPC:           "tgw-EXAMPLE",
							PeerResourceGroup: ToStringPointer(""),
						},
					},
					Delete: []string{},
				},
			},
			nil,
			true,
		},
		{
			"normal",
			fields{client: c},
			args{
				"test-pr",
				"test-vpc-id",
				TransitGatewayVPCAttachmentRequest{
					Add: []TransitGatewayVPCAttachment{
						{
							CIDR:              "10.0.0.0/24",
							PeerCloudAccount:  "111222333444",
							PeerVPC:           "tgw-EXAMPLE",
							PeerResourceGroup: nil,
						},
					},
					Delete: []string{},
				},
			},
			&VPC{
				CloudName:    "aws-eu-west-1",
				NetworkCIDR:  "172.16.0.0/24",
				ProjectVPCID: "test-vpc-id",
				State:        "ACTIVE",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &TransitGatewayVPCAttachmentHandler{
				client: tt.fields.client,
			}
			got, err := h.Update(tt.args.project, tt.args.projectVPCId, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
