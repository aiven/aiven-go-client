// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"fmt"
)

type (
	// KafkaSchemaRegistryACLHandler is the client which interacts with the Kafka Schema Registry ACL endpoints
	// on Aiven.
	KafkaSchemaRegistryACLHandler struct {
		client *Client
	}

	// CreateKafkaSchemaRegistryACLRequest are the parameters used to create a Kafka Schema Registry ACL entry.
	CreateKafkaSchemaRegistryACLRequest struct {
		Permission string `json:"permission"`
		Resource   string `json:"resource"`
		Username   string `json:"username"`
	}

	// KafkaSchemaRegistryACLResponse represents the response from Aiven after interacting with
	// the Kafka Schema Registry ACL API.
	KafkaSchemaRegistryACLResponse struct {
		APIResponse
		ACL []*KafkaSchemaRegistryACL `json:"acl"`
	}
)

// Create creates new Kafka Schema Registry ACL entry.
func (h *KafkaSchemaRegistryACLHandler) Create(project, service string, req CreateKafkaSchemaRegistryACLRequest) (*KafkaSchemaRegistryACL, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema-registry", "acl")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var rsp KafkaSchemaRegistryACLResponse
	if err := checkAPIResponse(bts, &rsp); err != nil {
		return nil, err
	}

	// The server doesn't return the Schema Registry ACL we created but list of all ACLs currently
	// defined. Need to find the correct one manually. There could be multiple ACLs
	// with same attributes. Assume the one that was created is the last one matching.
	var foundACL *KafkaSchemaRegistryACL
	for _, acl := range rsp.ACL {
		if acl.Permission == req.Permission && acl.Resource == req.Resource && acl.Username == req.Username {
			foundACL = acl
		}
	}

	if foundACL == nil {
		return nil, fmt.Errorf("created ACL not found from response ACL list")
	}

	return foundACL, nil
}

// Get gets a specific Kafka Schema Registry ACL.
func (h *KafkaSchemaRegistryACLHandler) Get(project, serviceName, aclID string) (*KafkaSchemaRegistryACL, error) {
	// There's no API for getting individual ACL entry. List instead and filter from there
	acls, err := h.List(project, serviceName)
	if err != nil {
		return nil, err
	}

	for _, acl := range acls {
		if acl.ID == aclID {
			return acl, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Schema Registry ACL with ID %v not found", aclID), Status: 404}
	return nil, err
}

// List lists all the Kafka Schema Registry ACL entries.
func (h *KafkaSchemaRegistryACLHandler) List(project, serviceName string) ([]*KafkaSchemaRegistryACL, error) {
	// Get Kafka Schema Registry ACL entries from service info, as in Kafka ACLs.
	service, err := h.client.Services.Get(project, serviceName)
	if err != nil {
		return nil, err
	}

	return service.SchemaRegistryACL, nil
}

// Delete deletes a specific Kafka Schema Registry ACL entry.
func (h *KafkaSchemaRegistryACLHandler) Delete(project, serviceName, aclID string) error {
	path := buildPath("project", project, "service", serviceName, "kafka", "schema-registry", "acl", aclID)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
