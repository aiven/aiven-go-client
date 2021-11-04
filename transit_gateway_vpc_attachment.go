package aiven

type (
	// TransitGatewayVPCAttachmentHandler is the client that interacts with the
	// Transit Gateway VPC Attachment API on Aiven.
	TransitGatewayVPCAttachmentHandler struct {
		client *Client
	}

	// TransitGatewayVPCAttachmentRequest holds the parameters to create a new
	// or update an existing Transit Gateway VPC Attachment.
	TransitGatewayVPCAttachmentRequest struct {
		Add    []TransitGatewayVPCAttachment `json:"add"`
		Delete []string                      `json:"delete"`
	}

	// TransitGatewayVPCAttachment represents Transit Gateway VPC Attachment
	TransitGatewayVPCAttachment struct {
		CIDR              string  `json:"cidr"`
		PeerCloudAccount  string  `json:"peer_cloud_account"`
		PeerResourceGroup *string `json:"peer_resource_group"`
		PeerVPC           string  `json:"peer_vpc"`
	}
)

// Update updates user-defined peer network CIDRs for a project VPC
func (h *TransitGatewayVPCAttachmentHandler) Update(
	project, projectVPCId string,
	req TransitGatewayVPCAttachmentRequest,
) (*VPC, error) {
	path := buildPath("project", project, "vpcs", projectVPCId, "user-peer-network-cidrs")
	rsp, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	return parseVPCResponse(rsp)
}
