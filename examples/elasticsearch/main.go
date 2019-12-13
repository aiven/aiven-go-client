package main

import (
	client "github.com/aiven/aiven-go-client"
	"log"
	"os"
)

func main() {
	// Create new user client
	c, err := client.NewUserClient(
		os.Getenv("AIVEN_USERNAME"),
		os.Getenv("AIVEN_PASSWORD"), "aiven-go-client-test/"+client.Version())
	if err != nil {
		log.Fatalf("user authentication error: %s", err)
	}

	// Create new project
	project, err := c.Projects.Create(client.CreateProjectRequest{
		CardID:  os.Getenv("AIVEN_CARD_ID"),
		Cloud:   "google-europe-west1",
		Project: "testproject1",
	})
	if err != nil {
		log.Fatalf("project creation error: %s", err)
	}

	// Create new Elasticsearch service inside the project
	userConfig := make(map[string]interface{})
	userConfig["elasticsearch_version"] = "7"
	service, err := c.Services.Create(project.Name, client.CreateServiceRequest{
		Cloud:                 "google-europe-west1",
		GroupName:             "default",
		MaintenanceWindow:     nil,
		Plan:                  "startup-4",
		ProjectVPCID:          nil,
		ServiceName:           "my-test-elasticsearch",
		ServiceType:           "elasticsearch",
		TerminationProtection: false,
		UserConfig:            userConfig,
		ServiceIntegrations:   nil,
	})
	if err != nil {
		log.Fatalf("cannot create new Elasticsearch service, error: %s", err)
	}

	// Create new Elasticsearch user
	user, err := c.ServiceUsers.Create(project.Name, service.Name, client.CreateServiceUserRequest{Username: "es_test_user1"})
	if err != nil {
		log.Fatalf("cannot create new Elasticsearch user, error: %s", err)
	}

	// Create new Elasticsearch ACLs for a user
	var rules []client.ElasticsearchACLRule
	rules = append(rules, client.ElasticsearchACLRule{
		Index:      "_all",
		Permission: "admin",
	})

	// Add Elasticsearch ACL
	err = c.ElasticsearchACLs.Add(project.Name, service.Name, client.ElasticSearchACL{
		Rules:    rules,
		Username: user.Username,
	})
	if err != nil {
		log.Fatalf("cannot create new Elasticsearch ACLs, error: %s", err)
	}

	// List Elasticsearch ACLs
	esACLs, err := c.ElasticsearchACLs.List(project.Name, service.Name)
	if err != nil {
		log.Fatalf("cannot get an Elasticsearch ACLs list, error: %s", err)
	}

	log.Printf("Elastic search ACLs : %+v", esACLs)

	// Delete Elasticsearch ACL
	err = c.ElasticsearchACLs.Delete(project.Name, service.Name, client.ElasticSearchACL{
		Rules:    rules,
		Username: user.Username,
	})
	if err != nil {
		log.Fatalf("cannot delete Elasticsearch ACL, error: %s", err)
	}
}
