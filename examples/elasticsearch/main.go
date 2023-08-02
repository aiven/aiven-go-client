package main

import (
	"context"
	"log"
	"os"

	client "github.com/aiven/aiven-go-client"
)

func main() {
	ctx := context.Background()

	// Create new user client
	c, err := client.NewUserClient(
		os.Getenv("AIVEN_USERNAME"),
		os.Getenv("AIVEN_PASSWORD"), "aiven-go-client-test/"+client.Version())
	if err != nil {
		log.Fatalf("user authentication error: %s", err)
	}

	// Create new project
	project, err := c.Projects.Create(ctx, client.CreateProjectRequest{
		CardID:  client.ToStringPointer(os.Getenv("AIVEN_CARD_ID")),
		Cloud:   client.ToStringPointer("google-europe-west1"),
		Project: "testproject1",
	})
	if err != nil {
		log.Fatalf("project creation error: %s", err)
	}

	// Create new Elasticsearch service inside the project
	userConfig := make(map[string]interface{})
	userConfig["elasticsearch_version"] = "7"
	service, err := c.Services.Create(ctx, project.Name, client.CreateServiceRequest{
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
	user, err := c.ServiceUsers.Create(ctx, project.Name, service.Name, client.CreateServiceUserRequest{Username: "es_test_user1"})
	if err != nil {
		log.Fatalf("cannot create new Elasticsearch user, error: %s", err)
	}

	// List Elasticsearch ACLs
	esACLs, err := c.ElasticsearchACLs.Get(ctx, project.Name, service.Name)
	if err != nil {
		log.Fatalf("cannot get an Elasticsearch ACLs list, error: %s", err)
	}

	log.Printf("Elastic search ACLs : %+v", esACLs)

	// Create new Elasticsearch ACLs for a user
	var rules []client.ElasticsearchACLRule
	rules = append(rules, client.ElasticsearchACLRule{
		Index:      "_all",
		Permission: "admin",
	})

	esACLs.ElasticSearchACLConfig.Add(client.ElasticSearchACL{
		Rules:    rules,
		Username: user.Username,
	})

	_, err = c.ElasticsearchACLs.Update(ctx, project.Name, service.Name, client.ElasticsearchACLRequest{
		ElasticSearchACLConfig: esACLs.ElasticSearchACLConfig})
	if err != nil {
		log.Fatalf("cannot update Elasticsearch ACLs, error: %s", err)
	}
}
