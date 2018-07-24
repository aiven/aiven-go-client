package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type (
	// VpcPeeringHandler is the client which interacts with the Projects VPC Peering endpoint
	// on Aiven
	VpcPeeringHandler struct {
		client *Client
	}
	// PeeringConnection represents the Peering Connection model on Aiven.
	PeeringConnection struct {
		CreateTime       string    `json:"create_time"`
		PeerCloudAccount string    `json:"peer_cloud_account"`
		PeerVpc          string    `json:"peer_vpc"`
		State            string    `json:"state"`
		StateInfo        StateInfo `json:"state_info"`
		UpdateTime       string    `json:"update_time"`
	}

	// CreatePeeringConnection are the parameters for creating a VPC peering connection
	CreatePeeringConnection struct {
		PeerCloudAccount string `json:"peer_cloud_account"`
		PeerVpc          string `json:"peer_vpc"`
	}

	// VpcPeeringResponse is the response from Aiven for Creating a Peering Connection
	VpcPeeringResponse struct {
		APIResponse
		PeeringConnection
	}
)

// Create the specified VPC Peering Connection
func (h *VpcPeeringHandler) Create(project, vpc string, req CreatePeeringConnection) (*PeeringConnection, error) {
	url := fmt.Sprintf("/project/%s/vpcs/%s/peering-connections", project, vpc)
	bts, err := h.client.doPostRequest(url, req)
	if err != nil {
		return nil, err
	}

	if bts == nil {
		return nil, ErrNoResponseData
	}

	var rsp *VpcPeeringResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return nil, err
	}

	if rsp == nil {
		return nil, ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return &rsp.PeeringConnection, nil
}

// Delete the specified VPC Peering Connection
func (h *VpcPeeringHandler) Delete(project, aivenVpc, peerAccountID, peerVpc string) error {
	url := fmt.Sprintf("/project/%s/vpcs/%s/peering-connections/peer-accounts/%s/peer-vpcs/%s", project, aivenVpc, peerAccountID, peerVpc)
	bts, err := h.client.doDeleteRequest(url, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}

// Get gets the specified VPC
func (h *VpcPeeringHandler) Get(project, aivenVpc, peerAccountID, peerVpc string) (*PeeringConnection, error) {
	log.Printf(fmt.Sprintf("Getting information for Project `%s` Aiven Vpc `%s` Peer Account `%s` Peer VPC `%s`", project, aivenVpc, peerAccountID, peerVpc))

	// Aiven does not provide a GET call for VPC Peering Connections, so we use the VPC handler
	vpc, err := h.client.Vpcs.Get(project, aivenVpc)
	if err != nil {
		return nil, err
	}

	var peeringConnection *PeeringConnection
	for _, pc := range vpc.PeeringConnections {
		if pc.PeerCloudAccount == peerAccountID && pc.PeerVpc == peerVpc {
			peeringConnection = &pc
			break
		}
	}
	if peeringConnection == nil {
		return nil, ErrNoResponseData
	}
	return peeringConnection, nil
}

// List all VPCs for a project
func (h *VpcPeeringHandler) List(project, aivenVpc, peerAccountID string) ([]*PeeringConnection, error) {
	// Aiven does not provide a GET call for VPC Peering Connections, so we use the VPC handler
	vpc, err := h.client.Vpcs.Get(project, aivenVpc)
	if err != nil {
		return nil, err
	}

	var connections []*PeeringConnection
	for _, pc := range vpc.PeeringConnections {
		if pc.PeerCloudAccount == peerAccountID {
			connections = append(connections, &pc)
		}
	}
	if len(connections) == 0 {
		return nil, ErrNoResponseData
	}

	return connections, nil
}
