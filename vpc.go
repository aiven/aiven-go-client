package aiven

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

type (
	// VPCPeeringConnection holds parameters associated with a VPC peering connection
	VPCPeeringConnection struct {
		CreateTime               *time.Time              `json:"create_time"`
		PeerCloudAccount         string                  `json:"peer_cloud_account"`
		PeerVPC                  string                  `json:"peer_vpc"`
		PeerRegion               *string                 `json:"peer_region"`
		State                    string                  `json:"state"`
		UpdateTime               *time.Time              `json:"update_time"`
		StateInfo                *map[string]interface{} `json:"state_info"`
		PeerAzureAppId           string                  `json:"peer_azure_app_id,omitempty"`
		PeerAzureTenantId        string                  `json:"peer_azure_tenant_id,omitempty"`
		UserPeerNetworkCIDRs     []string                `json:"user_peer_network_cidrs,omitempty"`
		VPCPeeringConnectionType string                  `json:"vpc_peering_connection_type,omitempty"`
		PeerResourceGroup        *string                 `json:"peer_resource_group,omitempty"`
	}

	// VPC holds parameters associated with a Virtual Private Cloud
	VPC struct {
		CloudName          string                  `json:"cloud_name"`
		CreateTime         *time.Time              `json:"create_time"`
		NetworkCIDR        string                  `json:"network_cidr"`
		ProjectVPCID       string                  `json:"project_vpc_id"`
		State              string                  `json:"state"`
		UpdateTime         *time.Time              `json:"update_time"`
		PeeringConnections []*VPCPeeringConnection `json:"peering_connections"`
	}

	// VPCsHandler is the client that interacts with the VPCs API on Aiven.
	VPCsHandler struct {
		client *Client
	}

	// CreateVPCRequest holds the parameters to create a new VPC.
	CreateVPCRequest struct {
		CloudName          string                  `json:"cloud_name"`
		NetworkCIDR        string                  `json:"network_cidr"`
		PeeringConnections []*VPCPeeringConnection `json:"peering_connections"`
	}

	// VPCListResponse represents the response from Aiven for listing VPCs.
	VPCListResponse struct {
		APIResponse
		VPCs []*VPC `json:"vpcs"`
	}
)

// Create the given VPC on Aiven.
func (h *VPCsHandler) Create(ctx context.Context, project string, req CreateVPCRequest) (*VPC, error) {
	path := buildPath("project", project, "vpcs")
	if req.PeeringConnections == nil {
		req.PeeringConnections = []*VPCPeeringConnection{}
	}
	rsp, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	return parseVPCResponse(rsp)
}

// Get a specific VPC from Aiven.
func (h *VPCsHandler) Get(ctx context.Context, project, vpcID string) (*VPC, error) {
	path := buildPath("project", project, "vpcs", vpcID)
	rsp, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	return parseVPCResponse(rsp)
}

// Delete the given VPC from Aiven.
func (h *VPCsHandler) Delete(ctx context.Context, project, vpcID string) error {
	path := buildPath("project", project, "vpcs", vpcID)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// List all VPCs for a given project.
func (h *VPCsHandler) List(ctx context.Context, project string) ([]*VPC, error) {
	path := buildPath("project", project, "vpcs")
	rsp, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var response *VPCListResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.VPCs, nil
}

func parseVPCResponse(rsp []byte) (*VPC, error) {
	var response *VPC
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	return response, nil
}
