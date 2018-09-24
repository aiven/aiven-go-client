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
		PeerCloudAccount string `json:"peer_cloud_account"`
		PeerVPC          string `json:"peer_vpc"`
	}
)

// Create the given VPC on Aiven.
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

	return response, nil
}

// Get a specific VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Get(
	project string,
	vpcID string,
	peerCloudAccount string,
	peerVPC string,
) (*VPCPeeringConnection, error) {
	// There's no API call for getting individual peering connection. Get the VPC
	// info and filter from there
	vpc, err := h.client.VPCs.Get(project, vpcID)
	if err != nil {
		return nil, err
	}

	for _, pc := range vpc.PeeringConnections {
		if pc.PeerCloudAccount == peerCloudAccount && pc.PeerVPC == peerVPC {
			return pc, nil
		}
	}

	err = Error{Message: "Peering connection not found", Status: 404}
	return nil, err
}

// Delete the given VPC Peering Connection from Aiven.
func (h *VPCPeeringConnectionsHandler) Delete(project, vpcID, peerCloudAccount, peerVPC string) error {
	path := buildPath(
		"project",
		project,
		"vpcs",
		vpcID,
		"peering-connections",
		"peer-accounts",
		peerCloudAccount,
		"peer-vpcs",
		peerVPC,
	)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// List all VPC peering connections for a given VPC.
func (h *VPCPeeringConnectionsHandler) List(project, vpcID string) ([]*VPCPeeringConnection, error) {
	vpc, err := h.client.VPCs.Get(project, vpcID)
	if err != nil {
		return nil, err
	}

	return vpc.PeeringConnections, nil
}
