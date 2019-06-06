// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
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
		PeerCloudAccount string  `json:"peer_cloud_account"`
		PeerVPC          string  `json:"peer_vpc"`
		PeerRegion       *string `json:"peer_region,omitempty"`
	}
)

// Create the given VPC on Aiven.
// when CreateVPCPeeringConnectionRequest.PeerRegion == nil the PeerVPC must be in the
// same region as the project VPC (vpcID)
func (h *VPCPeeringConnectionsHandler) Create(
	project string,
	vpcID string,
	req CreateVPCPeeringConnectionRequest,
) (*VPCPeeringConnection, error) {
	path := buildPath("project", project, "vpcs", vpcID, "peering-connections")
	rsp, err := h.client.doPostRequest(path, req)
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
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
	peerRegion *string,
) (*VPCPeeringConnection, error) {
	// There's no API call for getting individual peering connection. Get the VPC
	// info and filter from there
	vpc, err := h.client.VPCs.Get(project, vpcID)
	if err != nil {
		return nil, err
	}

	for _, pc := range vpc.PeeringConnections {
		if (peerRegion == nil || pc.PeerRegion == nil || *pc.PeerRegion == *peerRegion) && pc.PeerCloudAccount == peerCloudAccount && pc.PeerVPC == peerVPC {
			return pc, nil
		}
	}

	err = Error{Message: "Peering connection not found", Status: 404}
	return nil, err
}

// Get a VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Get(
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
) (*VPCPeeringConnection, error) {
	return h.GetVPCPeering(project, vpcID, peerCloudAccount, peerVPC, nil)
}

// DeleteVPCPeering Connection from Aiven.
// If peerRegion == nil the peering VPC must be in the same region as project VPC (vpcID)
func (h *VPCPeeringConnectionsHandler) DeleteVPCPeering(project, vpcID, peerCloudAccount, peerVPC string, peerRegion *string) error {
	pathElements := []string{"project", project, "vpcs", vpcID, "peering-connections", "peer-accounts", peerCloudAccount, "peer-vpcs", peerVPC}
	if peerRegion != nil {
		pathElements = append(pathElements, "peer-regions", *peerRegion)
	}

	bts, err := h.client.doDeleteRequest(buildPath(pathElements...), nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// Delete the given VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Delete(project, vpcID, peerCloudAccount, peerVPC string) error {
	return h.DeleteVPCPeering(project, vpcID, peerCloudAccount, peerVPC, nil)
}

// List all VPC peering connections for a given VPC.
func (h *VPCPeeringConnectionsHandler) List(project, vpcID string) ([]*VPCPeeringConnection, error) {
	vpc, err := h.client.VPCs.Get(project, vpcID)
	if err != nil {
		return nil, err
	}

	return vpc.PeeringConnections, nil
}
