package main

import (
	client "github.com/aiven/aiven-go-client"
	"log"
	"os"
	"time"
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
		CardID:  client.ToStringPointer(os.Getenv("AIVEN_CARD_ID")),
		Cloud:   client.ToStringPointer("google-europe-west1"),
		Project: "test-kafka-con1",
	})
	if err != nil {
		log.Fatalf("project creation error: %s", err)
	}

	// Create new Kafka service inside the project
	userConfig := make(map[string]interface{})
	userConfig["kafka_version"] = "2.4"
	userConfig["kafka_connect"] = true

	kService, err := c.Services.Create(project.Name, client.CreateServiceRequest{
		Cloud:                 "google-europe-west1",
		GroupName:             "default",
		MaintenanceWindow:     nil,
		Plan:                  "business-4",
		ProjectVPCID:          nil,
		ServiceName:           "my-test-kafka",
		ServiceType:           "kafka",
		TerminationProtection: false,
		UserConfig:            userConfig,
		ServiceIntegrations:   nil,
	})
	if err != nil {
		log.Fatalf("cannot create new Kafka service, error: %s", err)
	}

	// Create new Elasticsearch service inside the project
	userConfig = make(map[string]interface{})
	userConfig["elasticsearch_version"] = "7"

	esService, err := c.Services.Create(project.Name, client.CreateServiceRequest{
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

	for {
		err = c.KafkaConnectors.Create(project.Name, kService.Name, client.KafkaConnectorConfig{
			"topics":              "TestT1",
			"connection.username": esService.URIParams["user"],
			"name":                "es-connector",
			"connection.password": esService.URIParams["password"],
			"connector.class":     "io.aiven.connect.elasticsearch.ElasticsearchSinkConnector",
			"type.name":           "es-connector",
			"connection.url":      " https://" + esService.URIParams["host"] + ":" + esService.URIParams["port"],
		})

		if err != nil {
			if err.(client.Error).Status == 501 {
				log.Println("Kafka service is not fully up and running, wait 10 seconds and try again ...")
				time.Sleep(10 * time.Second)
				continue
			}

			if err.(client.Error).Status == 503 {
				log.Println("Kafka service is not available, try again in 10 seconds ...")
				time.Sleep(10 * time.Second)
				continue
			}

			if client.IsNotFound(err) {
				log.Println("Kafka service is not found by some reason, try again in 10 seconds ...")
				time.Sleep(10 * time.Second)
				continue
			}

			log.Fatalf("cannot create new Kafka connector, error: %s", err)
		}

		break
	}

	listCon, err := c.KafkaConnectors.List(project.Name, kService.Name)
	if err != nil {
		log.Fatalf("cannot get a Kafka Connectors list, error: %s", err)
	}

	for _, c := range listCon.Connectors {
		log.Printf("Kafka connector: %v", c)
	}
}
