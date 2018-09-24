// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"fmt"
)

type (
	// ConnectionPoolsHandler is the client which interacts with the connection pool endpoints
	// on Aiven.
	ConnectionPoolsHandler struct {
		client *Client
	}

	// CreateConnectionPoolRequest are the parameters used to create a connection pool entry.
	CreateConnectionPoolRequest struct {
		Database string `json:"database"`
		PoolMode string `json:"pool_mode"`
		PoolName string `json:"pool_name"`
		PoolSize int    `json:"pool_size"`
		Username string `json:"username"`
	}

	// UpdateConnectionPoolRequest are the parameters used to update a connection pool entry.
	UpdateConnectionPoolRequest struct {
		Database string `json:"database"`
		PoolMode string `json:"pool_mode"`
		PoolSize int    `json:"pool_size"`
		Username string `json:"username"`
	}
)

// Create new connection pool entry.
func (h *ConnectionPoolsHandler) Create(
	project string,
	serviceName string,
	req CreateConnectionPoolRequest,
) (*ConnectionPool, error) {
	path := buildPath("project", project, "service", serviceName, "connection_pool")
	_, err := h.client.doPostRequest(path, req)
	if err != nil {
		return nil, err
	}

	// Server doesn't return the connection pool we created, need to fetch it separately.
	return h.Get(project, serviceName, req.PoolName)
}

// Get a specific connection pool.
func (h *ConnectionPoolsHandler) Get(project, serviceName, poolName string) (*ConnectionPool, error) {
	// There's no API for getting individual connection pool entry. List instead and filter from there
	pools, err := h.List(project, serviceName)
	if err != nil {
		return nil, err
	}

	for _, pool := range pools {
		if pool.PoolName == poolName {
			return pool, nil
		}
	}

	err = Error{Message: fmt.Sprintf("Connection pool with name %v not found", poolName), Status: 404}
	return nil, err
}

// List all the connection pool entries for given service.
func (h *ConnectionPoolsHandler) List(project, serviceName string) ([]*ConnectionPool, error) {
	// There's no API for listing connection pool entries. Need to get them from
	// service info instead
	service, err := h.client.Services.Get(project, serviceName)
	if err != nil {
		return nil, err
	}

	return service.ConnectionPools, nil
}

// Update a specific connection pool with the given parameters.
func (h *ConnectionPoolsHandler) Update(
	project string,
	serviceName string,
	poolName string,
	req UpdateConnectionPoolRequest,
) (*ConnectionPool, error) {
	path := buildPath("project", project, "service", serviceName, "connection_pool", poolName)
	_, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	// Server doesn't return the connection pool we updated, need to fetch it separately.
	return h.Get(project, serviceName, poolName)
}

// Delete a specific connection pool entry.
func (h *ConnectionPoolsHandler) Delete(project, serviceName, poolName string) error {
	path := buildPath("project", project, "service", serviceName, "connection_pool", poolName)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
