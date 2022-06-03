package aiven

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafka", func() {
	var (
		projectName string
		project     *Project
		err         error
	)

	Context("Kafka Schemas CRUD", func() {
		It("should not error", func() {
			projectName = os.Getenv("AIVEN_PROJECT_NAME")
			project, err = client.Projects.Get(projectName)

			Expect(err).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			Expect(project).NotTo(BeNil())

			if project != nil {
				Expect(project.Name).NotTo(BeEmpty())
			}
		})

		// kafka service
		var (
			serviceName string
			service     *Service
			errS        error
		)

		It("creating service", func() {
			serviceName = "test-acc-kafka-sc-" + strconv.Itoa(rand.Int())
			service, errS = client.Services.Create(projectName, CreateServiceRequest{
				Cloud:        "google-europe-west1",
				Plan:         "business-4",
				ProjectVPCID: nil,
				ServiceName:  serviceName,
				ServiceType:  "kafka",
				UserConfig: map[string]interface{}{
					"schema_registry": true,
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
				Expect(service.Type).Should(Equal("kafka"))

				Eventually(func() string {
					service, _ = client.Services.Get(projectName, serviceName)
					return service.State
				}, 25*time.Minute, 1*time.Minute).Should(Equal("RUNNING"))
			}
		})

		// kafka schema
		var (
			errR        error
			errC        error
			subjectName string
		)

		It("create kafka schema subject", func() {
			time.Sleep(25 * time.Second)
			_, errC = client.KafkaGlobalSchemaConfig.Update(projectName, serviceName, KafkaSchemaConfig{
				CompatibilityLevel: "BACKWARD",
			})

			time.Sleep(25 * time.Second)
			subjectName = "test-subj"
			_, errR = client.KafkaSubjectSchemas.Add(projectName, serviceName, subjectName, KafkaSchemaSubject{
				Schema: `{
					"doc": "example",
					"fields": [{
						"default": 5,
						"doc": "my test number",
						"name": "test",
						"namespace": "test",
						"type": "int"
					}],
					"name": "example",
					"namespace": "example",
					"type": "record"
				}`})
		})

		It("should not error global config", func() {
			Expect(errC).NotTo(HaveOccurred())
		})

		It("should not error subject", func() {
			Expect(errR).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			s, errG := client.KafkaSubjectSchemas.Get(projectName, serviceName, subjectName, 1)
			Expect(errG).NotTo(HaveOccurred())
			Expect(s).NotTo(BeNil())

			if s != nil {
				Expect(s.Version.Schema).NotTo(BeEmpty())
				Expect(s.Version.Subject).NotTo(BeEmpty())
				Expect(s.Version.Version).To(Equal(1))
			}
		})

		It("should update configuration", func() {
			_, err := client.KafkaSubjectSchemas.GetConfiguration(projectName, serviceName, subjectName)
			Expect(err).To(HaveOccurred())
			Expect(IsNotFound(err)).To(Equal(true))

			s, errU := client.KafkaSubjectSchemas.UpdateConfiguration(projectName, serviceName, subjectName, "FORWARD")
			Expect(errU).NotTo(HaveOccurred())
			Expect(s).NotTo(BeNil())

			if s != nil {
				Expect(s.CompatibilityLevel).Should(Equal("FORWARD"))
			}

			s2, errG := client.KafkaSubjectSchemas.GetConfiguration(projectName, serviceName, subjectName)
			Expect(errG).NotTo(HaveOccurred())
			Expect(s2).NotTo(BeNil())

			if s2 != nil {
				Expect(s2.CompatibilityLevel).Should(Equal("FORWARD"))
			}
		})

		It("delete Kafka Schema subject and Kafka service", func() {
			if errD := client.KafkaSubjectSchemas.Delete(projectName, serviceName, subjectName); errD != nil {
				Fail("cannot delete kafka schema subject:" + errD.Error())
			}

			if errD := client.Services.Delete(projectName, serviceName); errD != nil {
				Fail("cannot delete service:" + errD.Error())
			}
		})
	})
})
