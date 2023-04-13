package aiven

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const (
	UpdateOperationResetCredentials = "reset-credentials"
	UpdateOperationSetAccessControl = "set-access-control"
)

type (
	// ServiceUser is the representation of a Service User in the Aiven API.
	ServiceUser struct {
		Username                    string        `json:"username"`
		Password                    string        `json:"password"`
		Type                        string        `json:"type"`
		AccessCert                  string        `json:"access_cert"`
		AccessKey                   string        `json:"access_key"`
		AccessCertNotValidAfterTime string        `json:"access_cert_not_valid_after_time"`
		AccessControl               AccessControl `json:"access_control,omitempty"`
	}

	AccessControl struct {
		M3Group                  *string  `json:"m3_group"`
		RedisACLCategories       []string `json:"redis_acl_categories"`
		RedisACLCommands         []string `json:"redis_acl_commands"`
		RedisACLKeys             []string `json:"redis_acl_keys"`
		RedisACLChannels         []string `json:"redis_acl_channels"`
		PostgresAllowReplication *bool    `json:"pg_allow_replication"`
	}

	// ServiceUsersHandler is the client that interacts with the ServiceUsers
	// endpoints.
	ServiceUsersHandler struct {
		client *Client
	}

	// CreateServiceUserRequest are the parameters required to create a
	// ServiceUser.
	CreateServiceUserRequest struct {
		Username       string         `json:"username"`
		Authentication *string        `json:"authentication,omitempty"`
		AccessControl  *AccessControl `json:"access_control,omitempty"`
	}

	// ModifyServiceUserRequest params required to modify a ServiceUser
	ModifyServiceUserRequest struct {
		Operation      *string        `json:"operation"`
		Authentication *string        `json:"authentication,omitempty"`
		NewPassword    *string        `json:"new_password,omitempty"`
		AccessControl  *AccessControl `json:"access_control,omitempty"`
	}

	// ServiceUserResponse represents the response after creating a ServiceUser.
	ServiceUserResponse struct {
		APIResponse
		User *ServiceUser `json:"user"`
	}
)

// MarshalJSON implements a custom marshalling process for AccessControl where only null fields are omitted
func (ac AccessControl) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})

	fields := reflect.TypeOf(ac)
	values := reflect.ValueOf(ac)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)

		switch value.Kind() {
		case reflect.Pointer, reflect.Slice: // *string, *bool, []string
			if !value.IsNil() {
				jsonName := field.Tag.Get("json")
				out[jsonName] = value.Interface()
			}
		}
	}
	return json.Marshal(out)
}

// Create creates the given User on Aiven.
func (h *ServiceUsersHandler) Create(project, service string, req CreateServiceUserRequest) (*ServiceUser, error) {
	path := buildPath("project", project, "service", service, "user")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	var r ServiceUserResponse
	errR := checkAPIResponse(bts, &r)

	return r.User, errR
}

// List Service Users for given service in Aiven.
func (h *ServiceUsersHandler) List(project, serviceName string) ([]*ServiceUser, error) {
	// Aiven API does not provide list operation for service users, need to get them via service info instead
	service, err := h.client.Services.Get(project, serviceName)
	if err != nil {
		return nil, err
	}

	return service.Users, nil
}

// Get specific Service User in Aiven.
func (h *ServiceUsersHandler) Get(project, serviceName, username string) (*ServiceUser, error) {
	// Aiven API does not provide get operation for service users, need to get them via list instead
	users, err := h.List(project, serviceName)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Service user with username %v not found", username), Status: 404}
	return nil, err
}

// Update modifies the given Service User in Aiven.
func (h *ServiceUsersHandler) Update(project, service, username string, update ModifyServiceUserRequest) (*ServiceUser, error) {
	var DefaultOperation = UpdateOperationResetCredentials
	if update.Operation == nil {
		update.Operation = &DefaultOperation
	}

	if update.AccessControl != nil && *update.Operation != UpdateOperationSetAccessControl {
		return nil, errors.New("wrong operation for updating access control")
	}

	if (update.NewPassword != nil || update.Authentication != nil) && *update.Operation != UpdateOperationResetCredentials {
		return nil, errors.New("wrong operation for updating credentials")
	}
	path := buildPath("project", project, "service", service, "user", username)
	svc, err := h.client.doPutRequest(path, update)
	if err != nil {
		return nil, err
	}

	var r ServiceResponse
	errR := checkAPIResponse(svc, &r)
	if errR == nil {
		for _, user := range r.Service.Users {
			if user.Username == username {
				return user, nil
			}
		}
		return nil, errors.New("user not found")
	}

	return nil, errR
}

// Delete deletes the given Service User in Aiven.
func (h *ServiceUsersHandler) Delete(project, service, user string) error {
	path := buildPath("project", project, "service", service, "user", user)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
