package aiven

import (
	"errors"
	"regexp"
	"strconv"
)

type (
	// KafkaSchemaHandler is the client which interacts with the Kafka Schema endpoints on Aiven.
	KafkaSchemaHandler struct {
		client *Client
	}

	// KafkaSchemaConfig represents Aiven Kafka Schema Configuration options
	KafkaSchemaConfig struct {
		CompatibilityLevel string `json:"compatibilityLevel"`
	}

	// KafkaSchemaConfigResponse represents the response from Aiven Kafka Schema Configuration endpoint
	KafkaSchemaConfigResponse struct {
		APIResponse
		KafkaSchemaConfig
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

	// KafkaSchemaSubjectVersion represents clist of versions
	KafkaSchemaSubjectVersion struct {
		Versions []int `json:"versions"`
	}

	// KafkaSchemaSubjectVersionResponse represents the response from Aiven Kafka Schema Subject versions endpoint
	KafkaSchemaSubjectVersionResponse struct {
		APIResponse
		KafkaSchemaSubjectVersion
	}

	// KafkaSchemaSubject Kafka SchemaS Subject representation
	KafkaSchemaSubject struct {
		Schema string `json:"schema"`
	}

	// KafkaSchemaSubjectResponse Kafka Schemas Subject API endpoint response representation
	KafkaSchemaSubjectResponse struct {
		APIResponse
		Id int `json:"id"`
	}

	// KafkaSchemaValidateResponse Kafka Schemas Subject validation API endpoint response representation
	KafkaSchemaValidateResponse struct {
		APIResponse
		IsCompatible bool `json:"is_compatible"`
	}
)

func NewKafkaSchema(s string) KafkaSchemaSubject {
	space := regexp.MustCompile(`\s+`)

	return KafkaSchemaSubject{
		Schema: string(space.ReplaceAllString(s, "")),
	}
}

// UpdateConfig updates new Kafka Schema config entry
func (h *KafkaSchemaHandler) UpdateConfig(project, service string, c KafkaSchemaConfig) (*KafkaSchemaConfigResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config")
	bts, err := h.client.doPutRequest(path, c)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// GetConfig gets a Kafka Schema configuration
func (h *KafkaSchemaHandler) GetConfig(project, service string) (*KafkaSchemaConfigResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "config")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaConfigResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// GetSubjects gets a list of Kafka Schema Subjects configuration
func (h *KafkaSchemaHandler) GetSubjects(project, service string) (*KafkaSchemaSubjectsResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectsResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// GetSubjectVersions gets a Kafka Schema Subject versions
func (h *KafkaSchemaHandler) GetSubjectVersions(project, service, name string) (*KafkaSchemaSubjectVersionResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectVersionResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// DeleteSubjectVersions deletes a Kafka Schema Subject versions
func (h *KafkaSchemaHandler) DeleteSubjectVersions(project, service, name string, version int) error {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions", strconv.Itoa(version))
	bts, err := h.client.doDeleteRequest(path, nil)
	if err != nil {
		return err
	}

	return checkAPIResponse(bts, nil)
}

// GetSubject gets a Kafka Schema Subject
func (h *KafkaSchemaHandler) GetSubject(project, service, name string, version int) (*KafkaSchemaSubjectResponse, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions", strconv.Itoa(version))
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}

// ValidateSchema validates Kafka Schema
func (h *KafkaSchemaHandler) ValidateSchema(
	project, service, name string,
	version int,
	subject KafkaSchemaSubject) (bool, error) {
	path := buildPath("project", project, "service", service, "kafka", "schema", "compatibility", "subjects", name, "versions", strconv.Itoa(version))

	bts, err := h.client.doPostRequest(path, subject)
	if err != nil {
		return false, err
	}

	var r KafkaSchemaValidateResponse
	errR := checkAPIResponse(bts, &r)

	return r.IsCompatible, errR
}

// AddSubject adds a new kafka Schema
func (h *KafkaSchemaHandler) AddSubject(project, service, name string, subject KafkaSchemaSubject) (*KafkaSchemaSubjectResponse, error) {
	vR, err := h.GetSubjectVersions(project, service, name)
	if err != nil && err.(Error).Status != 404 {
		return nil, err
	}

	// Get latest version
	if vR != nil {
		if len(vR.Versions) != 0 {
			hVersion := 0
			for _, v := range vR.Versions {
				if hVersion < v {
					hVersion = v
				}
			}

			// Validate Kafka schema against the latest existing version
			isValid, err := h.ValidateSchema(project, service, name, hVersion, subject)
			if err != nil {
				return nil, err
			}

			if !isValid {
				return nil, errors.New("kafka schema is not compatible with version :" + strconv.Itoa(hVersion))
			}
		}
	}

	println("hi again")

	path := buildPath("project", project, "service", service, "kafka", "schema", "subjects", name, "versions")
	println(path)
	bts, err := h.client.doPostRequest(path, subject)
	if err != nil {
		return nil, err
	}

	var r KafkaSchemaSubjectResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
