package aiven

import "context"

type (
	// ElasticSearchACLsHandler Aiven go-client handler for Elastisearch ACLs
	ElasticSearchACLsHandler struct {
		client *Client
	}

	// ElasticsearchACLRequest Aiven API request
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/elasticsearch/acl
	ElasticsearchACLRequest struct {
		ElasticSearchACLConfig ElasticSearchACLConfig `json:"elasticsearch_acl_config"`
	}

	// ElasticSearchACLResponse Aiven API response
	// https://api.aiven.io/v1/project/<project>/service/<service_name>/elasticsearch/acl
	ElasticSearchACLResponse struct {
		APIResponse
		ElasticSearchACLConfig ElasticSearchACLConfig `json:"elasticsearch_acl_config"`
	}

	// ElasticSearchACLConfig represents a configuration for Elasticsearch ACLs
	ElasticSearchACLConfig struct {
		ACLs        []ElasticSearchACL `json:"acls"`
		Enabled     bool               `json:"enabled"`
		ExtendedAcl bool               `json:"extendedAcl"`
	}

	// ElasticSearchACL represents a ElasticSearch ACLs entry
	ElasticSearchACL struct {
		Rules    []ElasticsearchACLRule `json:"rules"`
		Username string                 `json:"username"`
	}

	// ElasticsearchACLRule represents a ElasticSearch ACLs Rule entry
	ElasticsearchACLRule struct {
		Index      string `json:"index"`
		Permission string `json:"permission"`
	}
)

// Update updates Elasticsearch ACL config
//
// Deprecated: Use OpenSearchACLsHandler.Update instead.
func (h *ElasticSearchACLsHandler) Update(ctx context.Context, project, service string, req ElasticsearchACLRequest) (*ElasticSearchACLResponse, error) {
	path := buildPath("project", project, "service", service, "elasticsearch", "acl")
	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r ElasticSearchACLResponse
	return &r, checkAPIResponse(bts, &r)
}

// Get gets all existing Elasticsearch ACLs config
//
// Deprecated: Use OpenSearchACLsHandler.Get instead.
func (h *ElasticSearchACLsHandler) Get(ctx context.Context, project, service string) (*ElasticSearchACLResponse, error) {
	path := buildPath("project", project, "service", service, "elasticsearch", "acl")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r ElasticSearchACLResponse
	return &r, checkAPIResponse(bts, &r)
}

// Delete removes the specified ACL from the existing ElasticSearch ACLs config.
//
// Deprecated: Use OpenSearchACLConfig.Delete instead.
func (conf *ElasticSearchACLConfig) Delete(ctx context.Context, acl ElasticSearchACL) *ElasticSearchACLConfig {
	newACLs := []ElasticSearchACL{} // Create a new slice to hold the updated list of ACLs.

	// Iterate over each existing ACL entry.
	for _, existingAcl := range conf.ACLs {
		// If the ACL usernames match, we'll potentially modify the rules.
		if acl.Username == existingAcl.Username {
			newRules := []ElasticsearchACLRule{} // Create a new slice to hold the updated list of rules.

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

// Add appends new ACL to the existing ElasticSearch ACLs config.
//
// Deprecated: Use OpenSearchACLConfig.Add instead.
func (conf *ElasticSearchACLConfig) Add(acl ElasticSearchACL) *ElasticSearchACLConfig {
	var userIndex int
	userExists := false

	// Iterate over the existing ACLs to identify duplicates and determine user existence.
	for p, existingAcl := range conf.ACLs {
		if acl.Username == existingAcl.Username {
			userExists = true
			userIndex = p

			// Filter out any rules in the ACL to add that already exist for the user.
			remainingRules := []ElasticsearchACLRule{}
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
