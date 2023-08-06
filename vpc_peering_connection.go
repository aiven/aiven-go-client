package aiven

import (
	"context"
	"encoding/json"
)

type (
	// VPCPeeringConnectionsHandler is the client that interacts with the VPC
	// Peering Connections API on Aiven.
	VPCPeeringConnectionsHandler struct {
		client *Client
	}

	// CreateVPCPeeringConnectionRequest holds the parameters to create a new
	// peering connection.
	CreateVPCPeeringConnectionRequest struct {
		PeerCloudAccount     string   `json:"peer_cloud_account"`
		PeerVPC              string   `json:"peer_vpc"`
		PeerRegion           *string  `json:"peer_region,omitempty"`
		PeerAzureAppId       string   `json:"peer_azure_app_id,omitempty"`
		PeerAzureTenantId    string   `json:"peer_azure_tenant_id,omitempty"`
		PeerResourceGroup    string   `json:"peer_resource_group,omitempty"`
		UserPeerNetworkCIDRs []string `json:"user_peer_network_cidrs,omitempty"`
	}
)

// Create the given VPC on Aiven.
// when CreateVPCPeeringConnectionRequest.PeerRegion == nil the PeerVPC must be in the
// same region as the project VPC (vpcID)
func (h *VPCPeeringConnectionsHandler) Create(
	ctx context.Context,
	project string,
	vpcID string,
	req CreateVPCPeeringConnectionRequest,
) (*VPCPeeringConnection, error) {
	path := buildPath("project", project, "vpcs", vpcID, "peering-connections")
	rsp, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var response *VPCPeeringConnection
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	// if region was not set in the request omit it from the response
	if req.PeerRegion == nil {
		response.PeerRegion = nil
	}

	return response, nil
}

// GetVPCPeering Connection from Aiven.
// if peerRegion == nil the peered VPC is assumed to be in the same region as the project VPC (vpcID)
func (h *VPCPeeringConnectionsHandler) GetVPCPeering(
	ctx context.Context,
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
	peerRegion *string,
) (*VPCPeeringConnection, error) {
	// There's no API call for getting individual peering connection. Get the VPC
	// info and filter from there
	vpc, err := h.client.VPCs.Get(ctx, project, vpcID)
	if err != nil {
		return nil, err
	}

	for _, pc := range vpc.PeeringConnections {
		if (peerRegion == nil || eqStrPointers(pc.PeerRegion, peerRegion)) && pc.PeerCloudAccount == peerCloudAccount && pc.PeerVPC == peerVPC {
			return pc, nil
		}
	}

	err = Error{Message: "Peering connection not found", Status: 404}
	return nil, err
}

// GetVPCPeeringWithResourceGroup retrieves a VPC peering connection
func (h *VPCPeeringConnectionsHandler) GetVPCPeeringWithResourceGroup(
	ctx context.Context,
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
	peerRegion *string,
	peerResourceGroup *string,
) (*VPCPeeringConnection, error) {
	// There's no API call for getting individual peering connection. Get the VPC
	// info and filter from there
	vpc, err := h.client.VPCs.Get(ctx, project, vpcID)
	if err != nil {
		return nil, err
	}

	for _, pc := range vpc.PeeringConnections {
		found := pc.PeerCloudAccount == peerCloudAccount &&
			pc.PeerVPC == peerVPC &&
			// Not given or equal
			(peerRegion == nil || eqStrPointers(pc.PeerRegion, peerRegion)) &&
			(peerResourceGroup == nil || eqStrPointers(pc.PeerResourceGroup, peerResourceGroup))

		if found {
			return pc, nil
		}
	}

	err = Error{Message: "Peering connection not found", Status: 404}
	return nil, err
}

// Get a VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Get(
	ctx context.Context,
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
) (*VPCPeeringConnection, error) {
	return h.GetVPCPeering(ctx, project, vpcID, peerCloudAccount, peerVPC, nil)
}

// DeleteVPCPeering Connection from Aiven.
// If peerRegion == nil the peering VPC must be in the same region as project VPC (vpcID)
func (h *VPCPeeringConnectionsHandler) DeleteVPCPeering(ctx context.Context, project, vpcID, peerCloudAccount, peerVPC string, peerRegion *string) error {
	pathElements := []string{"project", project, "vpcs", vpcID, "peering-connections", "peer-accounts", peerCloudAccount, "peer-vpcs", peerVPC}
	if peerRegion != nil {
		pathElements = append(pathElements, "peer-regions", *peerRegion)
	}

	bts, err := h.client.doDeleteRequest(ctx, buildPath(pathElements...), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// DeleteVPCPeeringWithResourceGroup deletes a VPC peering connection
func (h *VPCPeeringConnectionsHandler) DeleteVPCPeeringWithResourceGroup(ctx context.Context, project, vpcID, peerCloudAccount, peerVPC, peerResourceGroup string, peerRegion *string) error {
	pathElements := []string{"project", project,
		"vpcs", vpcID,
		"peering-connections", "peer-accounts", peerCloudAccount,
		"peer-resource-groups", peerResourceGroup,
		"peer-vpcs", peerVPC,
	}
	if peerRegion != nil {
		pathElements = append(pathElements, "peer-regions", *peerRegion)
	}

	bts, err := h.client.doDeleteRequest(ctx, buildPath(pathElements...), nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)

}

// Delete the given VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Delete(ctx context.Context, project, vpcID, peerCloudAccount, peerVPC string) error {
	return h.DeleteVPCPeering(ctx, project, vpcID, peerCloudAccount, peerVPC, nil)
}

// List all VPC peering connections for a given VPC.
func (h *VPCPeeringConnectionsHandler) List(ctx context.Context, project, vpcID string) ([]*VPCPeeringConnection, error) {
	vpc, err := h.client.VPCs.Get(ctx, project, vpcID)
	if err != nil {
		return nil, err
	}

	return vpc.PeeringConnections, nil
}

func eqStrPointers(a, b *string) bool {
	if a != nil && b != nil {
		return *a == *b
	}
	// one or both of them nil
	return a == b
}
