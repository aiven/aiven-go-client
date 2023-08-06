package aiven

import (
	"context"
	"fmt"
)

type (
	// KafkaACLHandler is the client which interacts with the Kafka ACL endpoints
	// on Aiven.
	KafkaACLHandler struct {
		client *Client
	}

	// CreateKafkaACLRequest are the parameters used to create a Kafka ACL entry.
	CreateKafkaACLRequest struct {
		Permission string `json:"permission"`
		Topic      string `json:"topic"`
		Username   string `json:"username"`
	}

	// KafkaACLResponse represents the response from Aiven after interacting with
	// the Kafka ACL API.
	KafkaACLResponse struct {
		APIResponse
		ACL []*KafkaACL `json:"acl"`
	}
)

// Create creates new Kafka ACL entry.
func (h *KafkaACLHandler) Create(ctx context.Context, project, service string, req CreateKafkaACLRequest) (*KafkaACL, error) {
	path := buildPath("project", project, "service", service, "acl")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var rsp KafkaACLResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	// The server doesn't return the ACL we created but list of all ACLs currently
	// defined. Need to find the correct one manually. There could be multiple ACLs
	// with same attributes. Assume the one that was created is the last one matching.
	var foundACL *KafkaACL
	for _, acl := range rsp.ACL {
		if acl.Permission == req.Permission && acl.Topic == req.Topic && acl.Username == req.Username {
			foundACL = acl
		}
	}

	if foundACL == nil {
		return nil, fmt.Errorf("created ACL not found from response ACL list")
	}

	return foundACL, nil
}

// Get gets a specific Kafka ACL.
func (h *KafkaACLHandler) Get(ctx context.Context, project, serviceName, aclID string) (*KafkaACL, error) {
	// There's no API for getting individual ACL entry. List instead and filter from there
	acls, err := h.List(ctx, project, serviceName)
	if err != nil {
		return nil, err
	}

	for _, acl := range acls {
		if acl.ID == aclID {
			return acl, nil
		}
	}

	err = Error{Message: fmt.Sprintf("ACL with ID %v not found", aclID), Status: 404}
	return nil, err
}

// List lists all the Kafka ACL entries.
func (h *KafkaACLHandler) List(ctx context.Context, project, serviceName string) ([]*KafkaACL, error) {
	// There's no API for listing Kafka ACL entries. Need to get them from
	// service info instead
	service, err := h.client.Services.Get(ctx, project, serviceName)
	if err != nil {
		return nil, err
	}

	return service.ACL, nil
}

// Delete deletes a specific Kafka ACL entry.
func (h *KafkaACLHandler) Delete(ctx context.Context, project, serviceName, aclID string) error {
	path := buildPath("project", project, "service", serviceName, "acl", aclID)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
