package aiven_test

import (
	"testing"

	"github.com/jelmersnoeck/aiven"
	"github.com/jelmersnoeck/aiven/internal/test_helpers"
)

func TestService(t *testing.T) {
	pn := test_helpers.ProjectName("service")
	cl := test_helpers.Client()
	_, err := test_helpers.NewProject(cl, pn)
	if err != nil {
		t.Errorf("Error creating project: %s", err)
		return
	}
	defer func() {
		if err := cl.Projects.Delete(pn); err != nil {
			t.Errorf("Error deleting project: %s", err)
		}
	}()

	t.Run("integration test", func(t *testing.T) {
		t.Run("with all required params", func(t *testing.T) {
			var service *aiven.Service
			var err error
			serviceName := test_helpers.ProjectName("successful-pg")

			t.Run("it should create the service without errors", func(t *testing.T) {
				service, err = cl.Services.Create(pn, aiven.CreateServiceRequest{
					Plan:        test_helpers.ServicePlan,
					ServiceName: serviceName,
					ServiceType: "pg",
				})

				if err != nil {
					t.Errorf("Expected error to be nil, got %s", err)
				}

				if service == nil {
					t.Errorf("Expected service to be created", err)
				}
			})

			t.Run("it should get the service without errors", func(t *testing.T) {
				service, err = cl.Services.Get(pn, serviceName)

				if err != nil {
					t.Errorf("Expected error to be nil, got %s", err)
				}

				if service == nil {
					t.Errorf("Expected service to be fetched", err)
				}
			})

			t.Run("it should update the service without errors", func(t *testing.T) {
			})

			t.Run("it should delete the service without errors", func(t *testing.T) {
				if err = cl.Services.Delete(pn, serviceName); err != nil {
					t.Errorf("Expected service to be nil, got %s", err)
				}
			})
		})
	})
}
