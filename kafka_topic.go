// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

type (
	// KafkaTopicConfig represents a Kafka Topic Config on Aiven.
	KafkaTopicConfig struct {
		CleanupPolicy                   string   `json:"cleanup_policy,omitempty"`
		CompressionType                 string   `json:"compression_type,omitempty"`
		DeleteRetentionMs               *int     `json:"delete_retention_ms,omitempty"`
		FileDeleteDelayMs               *int     `json:"file_delete_delay_ms,omitempty"`
		FlushMessages                   *int     `json:"flush_messages,omitempty"`
		FlushMs                         *int     `json:"flush_ms,omitempty"`
		IndexIntervalBytes              *int     `json:"index_interval_bytes,omitempty"`
		MaxCompactionLagMs              *int     `json:"max_compaction_lag_ms,omitempty"`
		MaxMessageBytes                 *int     `json:"max_message_bytes,omitempty"`
		MessageDownconversionEnable     *bool    `json:"message_downconversion_enable,omitempty"`
		MessageFormatVersion            string   `json:"message_format_version,omitempty"`
		MessageTimestampDifferenceMaxMs *int     `json:"message_timestamp_difference_max_ms,omitempty"`
		MessageTimestampType            string   `json:"message_timestamp_type,omitempty"`
		MinCleanableDirtyRatio          *float32 `json:"min_cleanable_dirty_ratio,omitempty"`
		MinCompactionLagMs              *int     `json:"min_compaction_lag_ms,omitempty"`
		MinInsyncReplicas               *int     `json:"min_insync_replicas,omitempty"`
		Preallocate                     *bool    `json:"preallocate,omitempty"`
		RetentionBytes                  *int     `json:"retention_bytes,omitempty"`
		RetentionMs                     *int     `json:"retention_ms,omitempty"`
		SegmentBytes                    *int     `json:"segment_bytes,omitempty"`
		SegmentIndexBytes               *int     `json:"segment_index_bytes,omitempty"`
		SegmentJitterMs                 *int     `json:"segment_jitter_ms,omitempty"`
		SegmentMs                       *int     `json:"segment_ms,omitempty"`
		UncleanLeaderElectionEnable     *bool    `json:"unclean_leader_election_enable,omitempty"`
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
		Value    int    `json:"value"`
		Synonyms []struct {
			Source string `json:"source"`
			Value  int    `json:"value"`
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
		Value    float32 `json:"value"`
		Synonyms []struct {
			Source string  `json:"source"`
			Value  float32 `json:"value"`
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
	}

	// KafkaListTopic represents kafka list topic model on Aiven.
	KafkaListTopic struct {
		CleanupPolicy         string `json:"cleanup_policy"`
		MinimumInSyncReplicas int    `json:"min_insync_replicas"`
		Partitions            int    `json:"partitions"`
		Replication           int    `json:"replication"`
		RetentionBytes        int    `json:"retention_bytes"`
		RetentionHours        *int   `json:"retention_hours,omitempty"`
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
		CleanupPolicy         *string          `json:"cleanup_policy,omitempty"`
		MinimumInSyncReplicas *int             `json:"min_insync_replicas,omitempty"`
		Partitions            *int             `json:"partitions,omitempty"`
		Replication           *int             `json:"replication,omitempty"`
		RetentionBytes        *int             `json:"retention_bytes,omitempty"`
		RetentionHours        *int             `json:"retention_hours,omitempty"`
		TopicName             string           `json:"topic_name"`
		Config                KafkaTopicConfig `json:"config"`
	}

	// UpdateKafkaTopicRequest are the parameters used to update a kafka topic.
	UpdateKafkaTopicRequest struct {
		MinimumInSyncReplicas *int             `json:"min_insync_replicas,omitempty"`
		Partitions            *int             `json:"partitions,omitempty"`
		Replication           *int             `json:"replication,omitempty"`
		RetentionBytes        *int             `json:"retention_bytes,omitempty"`
		RetentionHours        *int             `json:"retention_hours,omitempty"`
		Config                KafkaTopicConfig `json:"config"`
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

	return checkAPIResponse(bts, nil)
}

// Get gets a specific kafka topic.
func (h *KafkaTopicsHandler) Get(project, service, topic string) (*KafkaTopic, error) {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaTopicResponse
	errR := checkAPIResponse(bts, &r)

	return r.Topic, errR
}

// List lists all the kafka topics.
func (h *KafkaTopicsHandler) List(project, service string) ([]*KafkaListTopic, error) {
	path := buildPath("project", project, "service", service, "topic")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaTopicsResponse
	errR := checkAPIResponse(bts, &r)

	return r.Topics, errR
}

// Update updates a specific topic with the given parameters.
func (h *KafkaTopicsHandler) Update(project, service, topic string, req UpdateKafkaTopicRequest) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doPutRequest(path, req)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// Delete deletes a specific kafka topic.
func (h *KafkaTopicsHandler) Delete(project, service, topic string) error {
	path := buildPath("project", project, "service", service, "topic", topic)
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}
