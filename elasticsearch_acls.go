package aiven

import (
	"encoding/json"
	"errors"
)

const (
	ElasticsearchAclPermissionRead         = "read"
	ElasticsearchAclPermissionWrite        = "write"
	ElasticsearchAclPermissionAdmin        = "admin"
	ElasticsearchAclPermissionDeny         = "deny"
	ElasticsearchAclPermissionReadAndWrite = "readwrite"
)

// list of supported Aiven API Elasticsearch ACL permissions
var elasticsearchPermissins = []string{
	ElasticsearchAclPermissionRead,
	ElasticsearchAclPermissionWrite,
	ElasticsearchAclPermissionAdmin,
	ElasticsearchAclPermissionDeny,
	ElasticsearchAclPermissionReadAndWrite,
}

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

// Create creates new Elasticsearch ACLs config
func (h *ElasticSearchACLsHandler) Create(project, service string, req ElasticsearchACLRequest) (*ElasticSearchACLResponse, error) {
	err := req.ElasticSearchACLConfig.Check()
	if err != nil {
		return nil, err
	}

	// POST or PUT request
	path := buildPath("project", project, "service", service, "elasticsearch", "acl")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	return h.response(bts)
}

// Update updates Elasticsearch ACL config
func (h *ElasticSearchACLsHandler) Update(project, service string, req ElasticsearchACLRequest) (*ElasticSearchACLResponse, error) {
	err := req.ElasticSearchACLConfig.Check()
	if err != nil {
		return nil, err
	}

	path := buildPath("project", project, "service", service, "elasticsearch", "acl")
	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	return h.response(bts)
}

// List lists all existing Elasticsearch ACLs config
func (h *ElasticSearchACLsHandler) List(project, service string) (*ElasticSearchACLResponse, error) {
	path := buildPath("project", project, "service", service, "elasticsearch", "acl")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	return h.response(bts)
}

// Add function adds one new ACL record to already existing Elasticsearch ACL configuration,
// it first checks if such ES user already exists. And if yes goes through the list of
// ACL rules for such user and check if new rules that user tries to add already exists in
// the configuration or not. We only add new records to the Elasticsearch configuration if case
// there is a difference between actual existing configuration and what user is willing to add.
func (h *ElasticSearchACLsHandler) Add(project, service string, acl ElasticSearchACL) error {
	acls, err := h.List(project, service)
	if err != nil {
		return err
	}

	_, err = h.Update(project, service, ElasticsearchACLRequest{
		ElasticSearchACLConfig: acls.ElasticSearchACLConfig.Add(acl),
	})
	return err
}

// Delete deletes one ACL record (might contain multiple rules) from Elasticsearch ACLs list
func (h *ElasticSearchACLsHandler) Delete(project, service string, acl ElasticSearchACL) error {
	acls, err := h.List(project, service)
	if err != nil {
		return err
	}

	_, err = h.Update(project, service,
		ElasticsearchACLRequest{
			ElasticSearchACLConfig: acls.ElasticSearchACLConfig.Delete(acl),
		},
	)
	return err
}

// response checks if response fom Aiven API contains any errors
func (h *ElasticSearchACLsHandler) response(r []byte) (*ElasticSearchACLResponse, error) {
	var rsp *ElasticSearchACLResponse
	if err := json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	// response cannot be empty
	if rsp == nil {
		return nil, ErrNoResponseData
	}

	// check API response errors
	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return nil, errors.New(rsp.Message)
	}

	return rsp, nil
}

// Delete subtracts ACL from already existing Elasticsearch ACLs config
func (conf ElasticSearchACLConfig) Delete(acl ElasticSearchACL) ElasticSearchACLConfig {
	for p, existingAcl := range conf.ACLs { // subtract ALC from existing ACLs config entry that supposed to be deleted
		if acl.Username == existingAcl.Username {
			for i := range existingAcl.Rules {
				// remove ACL from existing ACLs list
				for _, rule := range acl.Rules {
					if existingAcl.Rules[i].Permission == rule.Permission && existingAcl.Rules[i].Index == rule.Index {
						conf.ACLs[p].Rules = append(conf.ACLs[p].Rules[:i], conf.ACLs[p].Rules[i+1:]...)
					}
				}

				// delete ACL item from ACLs list is there are not rules attached to it
				if len(conf.ACLs[p].Rules) == 0 {
					conf.ACLs = append(conf.ACLs[:p], conf.ACLs[p+1:]...)
				}
			}
		}
	}

	return conf
}

// Check checks Elasticsearch permission and rules
func (conf ElasticSearchACLConfig) Check() error {
	isValidPermission := func(p string) bool {
		for _, permission := range elasticsearchPermissins {
			if p == permission {
				return true
			}
		}

		return false
	}

	for _, acl := range conf.ACLs {
		if acl.Username == "" {
			return errors.New("username cannot be empty")
		}

		if len(acl.Rules) == 0 {
			return errors.New("list of ACL rules cannot be empty")
		}

		// check ACL permission
		for _, rule := range acl.Rules {
			if !isValidPermission(rule.Permission) {
				return errors.New("invalid ACL permission :" + rule.Permission)
			}
		}
	}

	return nil
}

// Add appends new ACL to already existing Elasticsearch ACLs config
func (conf ElasticSearchACLConfig) Add(acl ElasticSearchACL) ElasticSearchACLConfig {
	var userAlreadyExist bool
	var userIndex int

	// check what ACL rules we already have for a user, and if we find that rule already exists,
	// remove it from a rules slice since there is no need of adding duplicates records to the ACL list
	for p, existingAcl := range conf.ACLs {
		if acl.Username == existingAcl.Username { // ACL record for this user already exists
			userAlreadyExist = true
			userIndex = p
			for _, existingRule := range existingAcl.Rules {
				for i, rule := range acl.Rules {
					if existingRule.Permission == rule.Permission && existingRule.Index == rule.Index {
						// remove rule since it already exists for this user
						acl.Rules = append(acl.Rules[:i], acl.Rules[i+1:]...)
					}
				}
			}
		}
	}

	if len(acl.Rules) == 0 {
		return conf // nothing to add to already existing ACL rules list for a user
	}

	// add to existing Elasticsearch ACL config new records
	conf.Enabled = true
	if userAlreadyExist {
		conf.ACLs[userIndex].Rules = append(conf.ACLs[userIndex].Rules, acl.Rules...)
	} else {
		conf.ACLs = append(conf.ACLs, acl)
	}

	return conf
}
