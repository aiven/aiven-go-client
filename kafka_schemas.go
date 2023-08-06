package aiven

import (
	"context"
	"errors"
	"strconv"
)

type (
	// KafkaSubjectSchemasHandler is the client which interacts with the Kafka Schema endpoints on Aiven
	KafkaSubjectSchemasHandler struct {
		client *Client
	}

	// KafkaGlobalSchemaConfigHandler is the client which interacts with the Kafka Schema endpoints on Aiven
	KafkaGlobalSchemaConfigHandler struct {
		client *Client
	}

	// KafkaSchemaConfig represents Aiven Kafka Schema Configuration options
	KafkaSchemaConfig struct {
		CompatibilityLevel string `json:"compatibility"`
	}

	// KafkaSchemaConfigUpdateResponse represents the PUT method response from Aiven Kafka Schema Configuration endpoint
	KafkaSchemaConfigUpdateResponse struct {
		APIResponse
		KafkaSchemaConfig
	}

	// KafkaSchemaConfigResponse represents the response from Aiven Kafka Schema Configuration endpoint
	KafkaSchemaConfigResponse struct {
		APIResponse
		CompatibilityLevel string `json:"compatibilityLevel"`
	}

	// KafkaSchemaSubjects represents a list of Aiven Kafka Schema subjects
	KafkaSchemaSubjects struct {
		Subjects []string `json:"subjects"`
	}

	// KafkaSchemaSubjectsResponse represents the response from Aiven Kafka Schema Subjects endpoint
	KafkaSchemaSubjectsResponse struct {
		APIResponse
		KafkaSchemaSubjects
	}

	// KafkaSchemaSubjectVersions represents a list of versions
	KafkaSchemaSubjectVersions struct {
		Versions []int `json:"versions"`
	}

	// KafkaSchemaSubjectVersionsResponse represents the response from Aiven Kafka Schema Subject versions endpoint
	KafkaSchemaSubjectVersionsResponse struct {
		APIResponse
		KafkaSchemaSubjectVersions
	}

	// KafkaSchemaSubject Kafka SchemaS Subject representation
	KafkaSchemaSubject struct {
		Schema     string `json:"schema"`
		SchemaType string `json:"schemaType,omitempty"`
	}

	// KafkaSchemaSubjectResponse Kafka Schemas Subject API endpoint response representation
	KafkaSchemaSubjectResponse struct {
		APIResponse
		Id int `json:"id"`
	}

	// KafkaSchemaSubjectVersion Kafka Schema Subject Version representation
	KafkaSchemaSubjectVersion struct {
		Id         int    `json:"id"`
		Schema     string `json:"schema"`
		Subject    string `json:"subject"`
		Version    int    `json:"version"`
		SchemaType string `json:"schemaType"`
	}

	// KafkaSchemaSubjectVersionResponse Kafka Schemas Subject Version API endpoint response representation
	KafkaSchemaSubjectVersionResponse struct {
		APIResponse
		Version KafkaSchemaSubjectVersion `json:"version"`
	}

	// KafkaSchemaValidateResponse Kafka Schemas Subject validation API endpoint response representation
	KafkaSchemaValidateResponse struct {
		APIResponse
		IsCompatible bool `json:"is_compatible"`
	}
)

// Update updates new Kafka Schema config entry
func (h *KafkaGlobalSchemaConfigHandler) Update(ctx context.Context, project, service string, c KafkaSchemaConfig) (*KafkaSchemaConfigUpdateResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config")
	bts, err := h.client.doPutRequest(ctx, path, c)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigUpdateResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Get gets a Kafka Schema configuration
func (h *KafkaGlobalSchemaConfigHandler) Get(ctx context.Context, project, service string) (*KafkaSchemaConfigResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// List gets a list of Kafka Schema Subjects configuration
func (h *KafkaSubjectSchemasHandler) List(ctx context.Context, project, service string) (*KafkaSchemaSubjectsResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectsResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// GetVersions gets a Kafka Schema Subject versions
func (h *KafkaSubjectSchemasHandler) GetVersions(ctx context.Context, project, service, name string) (*KafkaSchemaSubjectVersionsResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions")
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectVersionsResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Delete delete a Kafka Schema Subject versions, of versions parameter is empty it delete all existing versions
func (h *KafkaSubjectSchemasHandler) Delete(ctx context.Context, project, service, name string, versions ...int) error {
	if len(versions) == 0 {
		path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name)
		bts, err := h.client.doDeleteRequest(ctx, path, nil)
		if err != nil {
			return err
		}

		if errR := checkAPIResponse(bts, nil); errR != nil {
			return errR
		}
	}

	for _, version := range versions {
		path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions", strconv.Itoa(version))
		bts, err := h.client.doDeleteRequest(ctx, path, nil)
		if err != nil {
			return err
		}

		if errR := checkAPIResponse(bts, nil); errR != nil {
			return errR
		}
	}

	return nil
}

// Get gets a Kafka Schema Subject
func (h *KafkaSubjectSchemasHandler) Get(ctx context.Context, project, service, name string, version int) (*KafkaSchemaSubjectVersionResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions", strconv.Itoa(version))
	bts, err := h.client.doGetRequest(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectVersionResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// Validate validates Kafka Schema
func (h *KafkaSubjectSchemasHandler) Validate(
	ctx context.Context,
	project, service, name string,
	version int,
	subject KafkaSchemaSubject) (bool, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "compatibility", "subjects", name, "versions", strconv.Itoa(version))

	bts, err := h.client.doPostRequest(ctx, path, subject)
	if err != nil {
		return false, err
	}

	var r KafkaSchemaValidateResponse
	errR := checkAPIResponse(bts, &r)

	return r.IsCompatible, errR
}

// Add adds a new kafka Schema
func (h *KafkaSubjectSchemasHandler) Add(ctx context.Context, project, service, name string, subject KafkaSchemaSubject) (*KafkaSchemaSubjectResponse, error) {
	vR, err := h.GetVersions(ctx, project, service, name)
	if err != nil && !IsNotFound(err) {
		return nil, err
	}

	// GetVersions latest version
	if vR != nil {
		if len(vR.Versions) != 0 {
			hVersion := 0
			for _, v := range vR.Versions {
				if hVersion < v {
					hVersion = v
				}
			}

			// Validate Kafka schema against the latest existing version
			isValid, err := h.Validate(ctx, project, service, name, hVersion, subject)
			if err != nil {
				return nil, err
			}

			if !isValid {
				return nil, errors.New("kafka schema is not compatible with version :" + strconv.Itoa(hVersion))
			}
		}
	}

	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions")
	bts, err := h.client.doPostRequest(ctx, path, subject)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// UpdateConfiguration updates configuration for Schema Registry subject
func (h *KafkaSubjectSchemasHandler) UpdateConfiguration(ctx context.Context, project, service, subjectName, compatibility string) (
	*KafkaSchemaConfigUpdateResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config", subjectName)

	bts, err := h.client.doPutRequest(ctx, path, KafkaSchemaConfig{
		CompatibilityLevel: compatibility,
	})
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigUpdateResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

func (h *KafkaSubjectSchemasHandler) GetConfiguration(ctx context.Context, project, service, subjectName string) (
	*KafkaSchemaConfigResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config", subjectName)

	bts, err := h.client.doGetRequest(ctx, path+"?global_default_fallback=false", nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
