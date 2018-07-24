package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type (
	// VpcHandler is the client which interacts with the Projects VPC endpoint
	// on Aiven
	VpcHandler struct {
		client *Client
	}

	// Vpc represents the Vpc model on Aiven.
	Vpc struct {
		CloudName          string              `json:"cloud_name"`
		CreateTime         string              `json:"create_time"`
		NetworkCidr        string              `json:"network_cidr"`
		PeeringConnections []PeeringConnection `json:"peering_connections,omitempty"`
		ProjectVpcID       string              `json:"project_vpc_id"`
		State              string              `json:"state"`
		UpdateTime         string              `json:"update_time"`
	}

	// StateInfo is the response from Aiven for the Peering endpoint state
	StateInfo struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	}

	// CreateVpcRequest are the parameters for creating a VPC
	CreateVpcRequest struct {
		CloudName          string                    `json:"cloud_name,omitempty"`
		NetworkCidr        string                    `json:"network_cidr"`
		PeeringConnections []CreatePeeringConnection `json:"peering_connections,omitempty"`
	}

	// VpcResponse is the response from Aiven for the Vpc Endpoints
	VpcResponse struct {
		APIResponse
		Vpc
	}

	// VpcListResponse is the response from Aiven for listing VPCs
	VpcListResponse struct {
		APIResponse
		Vpcs []*Vpc `json:"vpcs"`
	}
)

// Create creates a new VPC
func (h *VpcHandler) Create(project string, req CreateVpcRequest) (*Vpc, error) {
	rsp, err := h.client.doPostRequest(fmt.Sprintf("/project/%s/vpcs", project), req)
	if err != nil {
		return nil, err
	}
	return parseVpcResponse(rsp)
}

// Get gets the specified VPC
func (h *VpcHandler) Get(project, vpc string) (*Vpc, error) {
	log.Printf(fmt.Sprintf("Getting information for Project `%s` Vpc `%s`", project, vpc))

	rsp, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/vpcs/%s", project, vpc), nil)
	if err != nil {
		return nil, err
	}

	return parseVpcResponse(rsp)
}

// List all VPCs for a project
func (h *VpcHandler) List(project string) ([]*Vpc, error) {
	bts, err := h.client.doGetRequest(fmt.Sprintf("/project/%s/vpcs", project), nil)
	if err != nil {
		return nil, err
	}

	var rsp *VpcListResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return nil, err
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return rsp.Vpcs, nil
}

func parseVpcResponse(bts []byte) (*Vpc, error) {
	if bts == nil {
		return nil, ErrNoResponseData
	}

	var rsp *VpcResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return nil, err
	}

	if rsp == nil {
		return nil, ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return &rsp.Vpc, nil
}
