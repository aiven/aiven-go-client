// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

import (
	"encoding/json"
	"errors"
)

type (
	// KafkaTopic represents a Kafka Topic on Aiven.
	KafkaTopic struct {
		CleanupPolicy         string       `json:"cleanup_policy"`
		MinimumInSyncReplicas int          `json:"min_insync_replicas"`
		Partitions            []*Partition `json:"partitions"`
		Replication           int          `json:"replication"`
		RetentionBytes        int          `json:"retention_bytes"`
		RetentionHours        int          `json:"retention_hours"`
		State                 string       `json:"state"`
		TopicName             string       `json:"topic_name"`
	}

	// KafkaListTopic represents kafka list topic model on Aiven.
	KafkaListTopic struct {
		CleanupPolicy         string `json:"cleanup_policy"`
		MinimumInSyncReplicas int    `json:"min_insync_replicas"`
		Partitions            int    `json:"partitions"`
		Replication           int    `json:"replication"`
		RetentionBytes        int    `json:"retention_bytes"`
		RetentionHours        int    `json:"retention_hours"`
		State                 string `json:"state"`
		TopicName             string `json:"topic_name"`
	}

	// Partition represents a Kafka partition.
	Partition struct {
		ConsumerGroups []*ConsumerGroup `json:"consumer_groups"`
		EarliestOffset int              `json:"earliest_offset"`
		ISR            int              `json:"isr"`
		LatestOffset   int              `json:"latest_offset"`
		Partition      int              `json:"partition"`
		Size           int              `json:"size"`
	}

	// ConsumerGroup is the group used in partitions.
	ConsumerGroup struct {
		GroupName string `json:"group_name"`
		Offset    int    `json:"offset"`
	}

	// KafkaTopicsHandler is the client which interacts with the kafka endpoints
	// on Aiven.
	KafkaTopicsHandler struct {
		client *Client
	}

	// CreateKafkaTopicRequest are the parameters used to create a kafka topic.
	CreateKafkaTopicRequest struct {
		CleanupPolicy         *string `json:"cleanup_policy,omitempty"`
		MinimumInSyncReplicas *int    `json:"min_insync_replicas,omitempty"`
		Partitions            *int    `json:"partitions,omitempty"`
		Replication           *int    `json:"replication,omitempty"`
		RetentionBytes        *int    `json:"retention_bytes,omitempty"`
		RetentionHours        *int    `json:"retention_hours,omitempty"`
		TopicName             string  `json:"topic_name"`
	}

	// UpdateKafkaTopicRequest are the parameters used to update a kafka topic.
	UpdateKafkaTopicRequest struct {
		MinimumInSyncReplicas *int `json:"min_insync_replicas,omitempty"`
		Partitions            *int `json:"partitions,omitempty"`
		Replication           *int `json:"replication,omitempty"`
		RetentionBytes        *int `json:"retention_bytes,omitempty"`
		RetentionHours        *int `json:"retention_hours,omitempty"`
	}

	// KafkaTopicResponse is the response for Kafka Topic requests.
	KafkaTopicResponse struct {
		APIResponse
		Topic *KafkaTopic `json:"topic"`
	}

	// KafkaTopicsResponse is the response for listing kafka topics.
	KafkaTopicsResponse struct {
		APIResponse
		Topics []*KafkaListTopic `json:"topics"`
	}
)

// Create creats a specific kafka topic.
func (h *KafkaTopicsHandler) Create(project, service string, req CreateKafkaTopicRequest) error {
	path := buildPath("project", project, "service", service, "topic")
	bts, err := h.client.doPostRequest(path, req)
	if err != nil {
		return err
	}

	var rsp *APIResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return err
	}

	if rsp == nil {
		return ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return errors.New(rsp.Message)
	}

	return nil
}

// Get gets a specific kafka topic.
func (h *KafkaTopicsHandler) Get(project, service, topic string) (*KafkaTopic, error) {
	path := buildPath("project", project, "service", service, "topic", topic)
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var response *KafkaTopicResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Topic, nil
}

// List lists all the kafka topics.
func (h *KafkaTopicsHandler) List(project, service string) ([]*KafkaListTopic, error) {
	path := buildPath("project", project, "service", service, "topic")
	rsp, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var response *KafkaTopicsResponse
	if err := json.Unmarshal(rsp, &response); err != nil {
		return nil, err
	}

	if len(response.Errors) != 0 {
		return nil, errors.New(response.Message)
	}

	return response.Topics, nil
}

// Update updates a specific topic with the given parameters.
func (h *KafkaTopicsHandler) Update(project, service, topic string, req UpdateKafkaTopicRequest) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return err
	}

	var rsp *APIResponse
	if err := json.Unmarshal(bts, &rsp); err != nil {
		return err
	}

	if rsp == nil {
		return ErrNoResponseData
	}

	if rsp.Errors != nil && len(rsp.Errors) != 0 {
		return errors.New(rsp.Message)
	}

	return nil
}

// Delete deletes a specific kafka topic.
func (h *KafkaTopicsHandler) Delete(project, service, topic string) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return handleDeleteResponse(bts)
}
