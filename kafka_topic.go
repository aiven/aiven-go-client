package aiven

import "context"

type (
	// KafkaTopicConfig represents a Kafka Topic Config on Aiven.
	KafkaTopicConfig struct {
		CleanupPolicy                   string   `json:"cleanup_policy,omitempty"`
		CompressionType                 string   `json:"compression_type,omitempty"`
		DeleteRetentionMs               *int64   `json:"delete_retention_ms,omitempty"`
		FileDeleteDelayMs               *int64   `json:"file_delete_delay_ms,omitempty"`
		FlushMessages                   *int64   `json:"flush_messages,omitempty"`
		FlushMs                         *int64   `json:"flush_ms,omitempty"`
		IndexIntervalBytes              *int64   `json:"index_interval_bytes,omitempty"`
		MaxCompactionLagMs              *int64   `json:"max_compaction_lag_ms,omitempty"`
		MaxMessageBytes                 *int64   `json:"max_message_bytes,omitempty"`
		MessageDownconversionEnable     *bool    `json:"message_downconversion_enable,omitempty"`
		MessageFormatVersion            string   `json:"message_format_version,omitempty"`
		MessageTimestampDifferenceMaxMs *int64   `json:"message_timestamp_difference_max_ms,omitempty"`
		MessageTimestampType            string   `json:"message_timestamp_type,omitempty"`
		MinCleanableDirtyRatio          *float64 `json:"min_cleanable_dirty_ratio,omitempty"`
		MinCompactionLagMs              *int64   `json:"min_compaction_lag_ms,omitempty"`
		MinInsyncReplicas               *int64   `json:"min_insync_replicas,omitempty"`
		Preallocate                     *bool    `json:"preallocate,omitempty"`
		RetentionBytes                  *int64   `json:"retention_bytes,omitempty"`
		RetentionMs                     *int64   `json:"retention_ms,omitempty"`
		SegmentBytes                    *int64   `json:"segment_bytes,omitempty"`
		SegmentIndexBytes               *int64   `json:"segment_index_bytes,omitempty"`
		SegmentJitterMs                 *int64   `json:"segment_jitter_ms,omitempty"`
		SegmentMs                       *int64   `json:"segment_ms,omitempty"`
		UncleanLeaderElectionEnable     *bool    `json:"unclean_leader_election_enable,omitempty"`
		RemoteStorageEnable             *bool    `json:"remote_storage_enable,omitempty"`
		//LocalRetentionBytes             *int64   `json:"local_retention_bytes,omitempty"`
		//LocalRetentionMs                *int64   `json:"local_retention_ms,omitempty"`
	}

	// KafkaTopicConfigResponse represents a Kafka Topic Config on Aiven.
	KafkaTopicConfigResponse struct {
		CleanupPolicy                   KafkaTopicConfigResponseString `json:"cleanup_policy,omitempty"`
		CompressionType                 KafkaTopicConfigResponseString `json:"compression_type,omitempty"`
		DeleteRetentionMs               KafkaTopicConfigResponseInt    `json:"delete_retention_ms,omitempty"`
		FileDeleteDelayMs               KafkaTopicConfigResponseInt    `json:"file_delete_delay_ms,omitempty"`
		FlushMessages                   KafkaTopicConfigResponseInt    `json:"flush_messages,omitempty"`
		FlushMs                         KafkaTopicConfigResponseInt    `json:"flush_ms,omitempty"`
		IndexIntervalBytes              KafkaTopicConfigResponseInt    `json:"index_interval_bytes,omitempty"`
		MaxCompactionLagMs              KafkaTopicConfigResponseInt    `json:"max_compaction_lag_ms,omitempty"`
		MaxMessageBytes                 KafkaTopicConfigResponseInt    `json:"max_message_bytes,omitempty"`
		MessageDownconversionEnable     KafkaTopicConfigResponseBool   `json:"message_downconversion_enable,omitempty"`
		MessageFormatVersion            KafkaTopicConfigResponseString `json:"message_format_version,omitempty"`
		MessageTimestampDifferenceMaxMs KafkaTopicConfigResponseInt    `json:"message_timestamp_difference_max_ms,omitempty"`
		MessageTimestampType            KafkaTopicConfigResponseString `json:"message_timestamp_type,omitempty"`
		MinCleanableDirtyRatio          KafkaTopicConfigResponseFloat  `json:"min_cleanable_dirty_ratio,omitempty"`
		MinCompactionLagMs              KafkaTopicConfigResponseInt    `json:"min_compaction_lag_ms,omitempty"`
		MinInsyncReplicas               KafkaTopicConfigResponseInt    `json:"min_insync_replicas,omitempty"`
		Preallocate                     KafkaTopicConfigResponseBool   `json:"preallocate,omitempty"`
		RetentionBytes                  KafkaTopicConfigResponseInt    `json:"retention_bytes,omitempty"`
		RetentionMs                     KafkaTopicConfigResponseInt    `json:"retention_ms,omitempty"`
		SegmentBytes                    KafkaTopicConfigResponseInt    `json:"segment_bytes,omitempty"`
		SegmentIndexBytes               KafkaTopicConfigResponseInt    `json:"segment_index_bytes,omitempty"`
		SegmentJitterMs                 KafkaTopicConfigResponseInt    `json:"segment_jitter_ms,omitempty"`
		SegmentMs                       KafkaTopicConfigResponseInt    `json:"segment_ms,omitempty"`
		UncleanLeaderElectionEnable     KafkaTopicConfigResponseBool   `json:"unclean_leader_election_enable,omitempty"`
		RemoteStorageEnable             KafkaTopicConfigResponseBool   `json:"remote_storage_enable,omitempty"`
		//LocalRetentionBytes             KafkaTopicConfigResponseInt    `json:"local_retention_bytes,omitempty"`
		//LocalRetentionMs                KafkaTopicConfigResponseInt    `json:"local_retention_ms,omitempty"`
		Tags []KafkaTopicTag `json:"tags,omitempty"`
	}

	KafkaTopicTag struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	KafkaTopicConfigResponseString struct {
		Source   string `json:"source"`
		Value    string `json:"value"`
		Synonyms []struct {
			Source string `json:"source"`
			Value  string `json:"value"`
			Name   string `json:"name"`
		} `json:"synonyms"`
	}

	KafkaTopicConfigResponseInt struct {
		Source   string `json:"source"`
		Value    int64  `json:"value"`
		Synonyms []struct {
			Source string `json:"source"`
			Value  int64  `json:"value"`
			Name   string `json:"name"`
		} `json:"synonyms"`
	}

	KafkaTopicConfigResponseBool struct {
		Source   string `json:"source"`
		Value    bool   `json:"value"`
		Synonyms []struct {
			Source string `json:"source"`
			Value  bool   `json:"value"`
			Name   string `json:"name"`
		} `json:"synonyms"`
	}

	KafkaTopicConfigResponseFloat struct {
		Source   string  `json:"source"`
		Value    float64 `json:"value"`
		Synonyms []struct {
			Source string  `json:"source"`
			Value  float64 `json:"value"`
			Name   string  `json:"name"`
		} `json:"synonyms"`
	}

	// KafkaTopic represents a Kafka Topic on Aiven.
	KafkaTopic struct {
		CleanupPolicy         string                   `json:"cleanup_policy"`
		MinimumInSyncReplicas int                      `json:"min_insync_replicas"`
		Partitions            []*Partition             `json:"partitions"`
		Replication           int                      `json:"replication"`
		RetentionBytes        int                      `json:"retention_bytes"`
		RetentionHours        *int                     `json:"retention_hours,omitempty"`
		State                 string                   `json:"state"`
		TopicName             string                   `json:"topic_name"`
		Config                KafkaTopicConfigResponse `json:"config"`
		Tags                  []KafkaTopicTag          `json:"tags,omitempty"`
	}

	// KafkaListTopic represents kafka list topic model on Aiven.
	KafkaListTopic struct {
		CleanupPolicy         string `json:"cleanup_policy"`
		MinimumInSyncReplicas int    `json:"min_insync_replicas"`
		Partitions            int    `json:"partitions"`
		Replication           int    `json:"replication"`
		RetentionBytes        int    `json:"retention_bytes"`
		RetentionHours        *int64 `json:"retention_hours,omitempty"`
		State                 string `json:"state"`
		TopicName             string `json:"topic_name"`
	}

	// Partition represents a Kafka partition.
	Partition struct {
		ConsumerGroups []*ConsumerGroup `json:"consumer_groups"`
		EarliestOffset int64            `json:"earliest_offset"`
		ISR            int              `json:"isr"`
		LatestOffset   int64            `json:"latest_offset"`
		Partition      int              `json:"partition"`
		Size           int64            `json:"size"`
	}

	// ConsumerGroup is the group used in partitions.
	ConsumerGroup struct {
		GroupName string `json:"group_name"`
		Offset    int64  `json:"offset"`
	}

	// KafkaTopicsHandler is the client which interacts with the kafka endpoints
	// on Aiven.
	KafkaTopicsHandler struct {
		client *Client
	}

	// CreateKafkaTopicRequest are the parameters used to create a kafka topic.
	CreateKafkaTopicRequest struct {
		CleanupPolicy         *string          `json:"cleanup_policy,omitempty"`
		MinimumInSyncReplicas *int             `json:"min_insync_replicas,omitempty"`
		Partitions            *int             `json:"partitions,omitempty"`
		Replication           *int             `json:"replication,omitempty"`
		RetentionBytes        *int             `json:"retention_bytes,omitempty"`
		RetentionHours        *int             `json:"retention_hours,omitempty"`
		TopicName             string           `json:"topic_name"`
		Config                KafkaTopicConfig `json:"config"`
		Tags                  []KafkaTopicTag  `json:"tags,omitempty"`
	}

	// UpdateKafkaTopicRequest are the parameters used to update a kafka topic.
	UpdateKafkaTopicRequest struct {
		MinimumInSyncReplicas *int             `json:"min_insync_replicas,omitempty"`
		Partitions            *int             `json:"partitions,omitempty"`
		Replication           *int             `json:"replication,omitempty"`
		RetentionBytes        *int             `json:"retention_bytes,omitempty"`
		RetentionHours        *int             `json:"retention_hours,omitempty"`
		Config                KafkaTopicConfig `json:"config"`
		Tags                  []KafkaTopicTag  `json:"tags,omitempty"`
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

	// KafkaV2TopicsResponse is the response for listing kafka topics specific for API V2 endpoint.
	KafkaV2TopicsResponse struct {
		APIResponse
		Topics []*KafkaTopic `json:"topics"`
	}
)

// Create creats a specific kafka topic.
func (h *KafkaTopicsHandler) Create(ctx context.Context, project, service string, req CreateKafkaTopicRequest) error {
	path := buildPath("project", project, "service", service, "topic")
	bts, err := h.client.doPostRequest(ctx, path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Get gets a specific kafka topic.
func (h *KafkaTopicsHandler) Get(ctx context.Context, project, service, topic string) (*KafkaTopic, error) {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaTopicResponse
	errR := checkAPIResponse(bts, &r)

	return r.Topic, errR
}

// List lists all the kafka topics.
func (h *KafkaTopicsHandler) List(ctx context.Context, project, service string) ([]*KafkaListTopic, error) {
	path := buildPath("project", project, "service", service, "topic")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaTopicsResponse
	errR := checkAPIResponse(bts, &r)

	return r.Topics, errR
}

// Update updates a specific topic with the given parameters.
func (h *KafkaTopicsHandler) Update(ctx context.Context, project, service, topic string, req UpdateKafkaTopicRequest) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doPutRequest(ctx, path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes a specific kafka topic.
func (h *KafkaTopicsHandler) Delete(ctx context.Context, project, service, topic string) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doDeleteRequest(ctx, path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// V2List lists selected kafka topics using v2 API endpoint.
func (h *KafkaTopicsHandler) V2List(ctx context.Context, project, service string, topics []string) ([]*KafkaTopic, error) {
	type v2ListRequest struct {
		TopicNames []string `json:"topic_names"`
	}

	req := v2ListRequest{TopicNames: topics}

	path := buildPath("project", project, "service", service, "topic")
	bts, err := h.client.doV2PostRequest(ctx, path, req)
	if err != nil {
		return nil, err
	}

	var r KafkaV2TopicsResponse
	errR := checkAPIResponse(bts, &r)

	return r.Topics, errR
}
