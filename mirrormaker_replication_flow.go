// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

type (
	// MirrorMakerReplicationFlowHandler is the client which interacts with the
	// Kafka MirrorMaker 2 ReplicationFlows endpoints on Aiven.
	MirrorMakerReplicationFlowHandler struct {
		client *Client
	}

	// ReplicationFlow a replication flow entity
	ReplicationFlow struct {
		Enabled                         bool     `json:"enabled"`
		OffsetSyncsTopicLocation        string   `json:"offset_syncs_topic_location,omitempty"`
		SourceCluster                   string   `json:"source_cluster,omitempty"`
		TargetCluster                   string   `json:"target_cluster,omitempty"`
		ReplicationPolicyClass          string   `json:"replication_policy_class"`
		SyncGroupOffsetsEnabled         bool     `json:"sync_group_offsets_enabled"`
		SyncGroupOffsetsIntervalSeconds int      `json:"sync_group_offsets_interval_seconds"`
		EmitHeartbeatsEnabled           bool     `json:"emit_heartbeats_enabled"`
		Topics                          []string `json:"topics"`
		TopicsBlacklist                 []string `json:"topics.blacklist"`
	}

	// MirrorMakerReplicationFlowRequest request used to create a Kafka MirrorMaker 2
	// ReplicationFlows entry.
	MirrorMakerReplicationFlowRequest struct {
		ReplicationFlow
	}

	// MirrorMakerReplicationFlowsResponse represents the response from Aiven after
	// interacting with the Kafka MirrorMaker 2 API.
	MirrorMakerReplicationFlowsResponse struct {
		APIResponse
		ReplicationFlows []ReplicationFlow `json:"replication_flows"`
	}

	// MirrorMakerReplicationFlowResponse represents the response from Aiven after
	// interacting with the Kafka MirrorMaker 2 API.
	MirrorMakerReplicationFlowResponse struct {
		APIResponse
		ReplicationFlow ReplicationFlow `json:"replication_flow"`
	}
)

// Create creates new Kafka MirrorMaker 2 Replication Flows entry.
func (h *MirrorMakerReplicationFlowHandler) Create(project, service string, req MirrorMakerReplicationFlowRequest) error {
	path := buildPath("project", project, "service", service, "mirrormaker", "replication-flows")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Update updates new Kafka MirrorMaker 2 Replication Flows entry.
func (h *MirrorMakerReplicationFlowHandler) Update(project, service, sourceCluster, targetCluster string, req MirrorMakerReplicationFlowRequest) (*MirrorMakerReplicationFlowResponse, error) {
	path := buildPath("project", project, "service", service, "mirrormaker", "replication-flows", sourceCluster, targetCluster)

	// unset source and destination clusters fields
	req.SourceCluster = ""
	req.TargetCluster = ""

	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return nil, err
	}

	var rsp MirrorMakerReplicationFlowResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// List gets a Kafka MirrorMaker 2 Replication Flows.
func (h *MirrorMakerReplicationFlowHandler) List(project, service string) (*MirrorMakerReplicationFlowsResponse, error) {
	path := buildPath("project", project, "service", service, "mirrormaker", "replication-flows")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp MirrorMakerReplicationFlowsResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Get gets a Kafka MirrorMaker 2 Replication Flows.
func (h *MirrorMakerReplicationFlowHandler) Get(project, service, sourceCluster, targetCluster string) (*MirrorMakerReplicationFlowResponse, error) {
	path := buildPath("project", project, "service", service, "mirrormaker", "replication-flows", sourceCluster, targetCluster)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var rsp MirrorMakerReplicationFlowResponse
	if errR := checkAPIResponse(bts, &rsp); errR != nil {
		return nil, errR
	}

	return &rsp, nil
}

// Delete deletes a Kafka MirrorMaker 2 Replication Flows entry.
func (h *MirrorMakerReplicationFlowHandler) Delete(project, service, sourceCluster, targetCluster string) error {
	path := buildPath("project", project, "service", service, "mirrormaker", "replication-flows", sourceCluster, targetCluster)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
