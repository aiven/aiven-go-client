package aiven

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service Task", func() {
	var (
		projectName string
	)

	Context("PG Service Task", func() {
		It("should not error", func() {
			projectName = os.Getenv("AIVEN_PROJECT_NAME")
			_, err := client.Projects.Get(projectName)

			Expect(err).NotTo(HaveOccurred())
		})

		// PG service
		var (
			serviceName string
			service     *Service
			errS        error
		)

		It("creating service", func() {
			serviceName = "test-acc-pg-sr-" + strconv.Itoa(rand.Int())
			service, errS = client.Services.Create(projectName, CreateServiceRequest{
				Cloud:        "google-europe-west1",
				Plan:         "business-4",
				ProjectVPCID: nil,
				ServiceName:  serviceName,
				ServiceType:  "pg",
				UserConfig: map[string]interface{}{
					"pg_version": "11",
				},
			})
		})

		It("should not error", func() {
			Expect(errS).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			Expect(service).NotTo(BeNil())

			if service != nil {
				Expect(service.Name).NotTo(BeEmpty())
				Expect(service.Plan).NotTo(BeEmpty())
				Expect(service.Type).Should(Equal("pg"))

				Eventually(func() string {
					service, _ = client.Services.Get(projectName, serviceName)
					return service.State
				}, 25*time.Minute, 1*time.Minute).Should(Equal("RUNNING"))
			}
		})

		// service task
		It("create pg service task", func() {
			t, errT := client.ServiceTask.Create(projectName, serviceName, ServiceTaskRequest{
				TargetVersion: "12",
				TaskType:      "upgrade_check",
			})

			Expect(errT).NotTo(HaveOccurred())
			Expect(t.Task.CreateTime).NotTo(BeEmpty())
			Expect(t.Task.Id).NotTo(BeEmpty())
			Expect(t.Task.Result).NotTo(BeEmpty())
			Expect(t.Task.SourcePgVersion).NotTo(BeEmpty())
			Expect(t.Task.TargetPgVersion).NotTo(BeEmpty())
			Expect(t.Task.TaskType).NotTo(BeEmpty())

			Eventually(func() *bool {
				t, errT = client.ServiceTask.Get(projectName, serviceName, t.Task.Id)
				return t.Task.Success
			}, 5*time.Minute, 15*time.Second).Should(Not(BeNil()))

			t, errT = client.ServiceTask.Get(projectName, serviceName, t.Task.Id)
			Expect(errT).NotTo(HaveOccurred())
			Expect(*t.Task.Success).To(BeTrue())
		})

		It("delete PG service", func() {
			if errD := client.Services.Delete(projectName, serviceName); errD != nil {
				Fail("cannot delete service:" + errD.Error())
			}
		})
	})
})
