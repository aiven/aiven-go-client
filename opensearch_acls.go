package aiven

import "context"

type (
	// OpenSearchACLsHandler Aiven go-client handler for OpenSearch ACLs
	OpenSearchACLsHandler struct {
		client *Client
	}

	// OpenSearchACLRequest Aiven API request
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/opensearch/acl
	OpenSearchACLRequest struct {
		OpenSearchACLConfig OpenSearchACLConfig `json:"opensearch_acl_config"`
	}

	// OpenSearchACLResponse Aiven API response
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/opensearch/acl
	OpenSearchACLResponse struct {
		APIResponse
		OpenSearchACLConfig OpenSearchACLConfig `json:"opensearch_acl_config"`
	}

	// OpenSearchACLConfig represents a configuration for OpenSearch ACLs
	OpenSearchACLConfig struct {
		ACLs        []OpenSearchACL `json:"acls"`
		Enabled     bool            `json:"enabled"`
		ExtendedAcl bool            `json:"extendedAcl"`
	}

	// OpenSearchACL represents a OpenSearch ACLs entry
	OpenSearchACL struct {
		Rules    []OpenSearchACLRule `json:"rules"`
		Username string              `json:"username"`
	}

	// OpenSearchACLRule represents a OpenSearch ACLs Rule entry
	OpenSearchACLRule struct {
		Index      string `json:"index"`
		Permission string `json:"permission"`
	}
)

// Update updates OpenSearch ACL config
func (h *OpenSearchACLsHandler) Update(ctx context.Context, project, service string, req OpenSearchACLRequest) (*OpenSearchACLResponse, error) {
	path := buildPath("project", project, "service", service, "opensearch", "acl")
	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r OpenSearchACLResponse
	return &r, checkAPIResponse(bts, &r)
}

// Get gets all existing OpenSearch ACLs config
func (h *OpenSearchACLsHandler) Get(ctx context.Context, project, service string) (*OpenSearchACLResponse, error) {
	path := buildPath("project", project, "service", service, "opensearch", "acl")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r OpenSearchACLResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete removes the specified ACL from the existing OpenSearch ACLs config.
func (conf *OpenSearchACLConfig) Delete(ctx context.Context, acl OpenSearchACL) *OpenSearchACLConfig {
	newACLs := []OpenSearchACL{} // Create a new slice to hold the updated list of ACLs.

	// Iterate over each existing ACL entry.
	for _, existingAcl := range conf.ACLs {
		// If the ACL usernames match, we'll potentially modify the rules.
		if acl.Username == existingAcl.Username {
			newRules := []OpenSearchACLRule{} // Create a new slice to hold the updated list of rules.

			// Check each existing rule against the rules in the ACL to be deleted.
			for _, existingRule := range existingAcl.Rules {
				match := false // Flag to track if the existing rule matches any rule in the ACL to be deleted.
				for _, ruleToDelete := range acl.Rules {
					if existingRule.Permission == ruleToDelete.Permission && existingRule.Index == ruleToDelete.Index {
						match = true // The existing rule matches a rule in the ACL to be deleted.
						break
					}
				}
				// If the existing rule doesn't match any rule in the ACL to be deleted, add it to the new list.
				if !match {
					newRules = append(newRules, existingRule)
				}
			}

			// If there are remaining rules after deletion, add the modified ACL to the new list.
			if len(newRules) > 0 {
				existingAcl.Rules = newRules
				newACLs = append(newACLs, existingAcl)
			}
		} else {
			// If the usernames don't match, directly add the existing ACL to the new list.
			newACLs = append(newACLs, existingAcl)
		}
	}

	// Replace the original list of ACLs with the updated list.
	conf.ACLs = newACLs
	return conf
}

// Add appends new ACL to the existing OpenSearch ACLs config.
func (conf *OpenSearchACLConfig) Add(acl OpenSearchACL) *OpenSearchACLConfig {
	var userIndex int
	userExists := false

	// Iterate over the existing ACLs to identify duplicates and determine user existence.
	for p, existingAcl := range conf.ACLs {
		if acl.Username == existingAcl.Username {
			userExists = true
			userIndex = p

			// Filter out any rules in the ACL to add that already exist for the user.
			remainingRules := []OpenSearchACLRule{}
			for _, rule := range acl.Rules {
				exists := false
				for _, existingRule := range existingAcl.Rules {
					if rule.Permission == existingRule.Permission && rule.Index == existingRule.Index {
						exists = true
						break
					}
				}
				if !exists {
					remainingRules = append(remainingRules, rule)
				}
			}
			acl.Rules = remainingRules
		}
	}

	// If no rules remain for the user, return the existing configuration.
	if len(acl.Rules) == 0 {
		return conf
	}

	// Add the new or updated ACL to the config.
	if userExists {
		conf.ACLs[userIndex].Rules = append(conf.ACLs[userIndex].Rules, acl.Rules...)
	} else {
		conf.ACLs = append(conf.ACLs, acl)
	}

	return conf
}
