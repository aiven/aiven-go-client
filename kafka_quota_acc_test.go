package aiven

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafka Quota", func() {
	var (
		projectName string
		project     *Project
		err         error
	)

	Context("Kafka Quota CRUD", func() {
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
			serviceName = "test-acc-kafka-quota-sr-" + strconv.Itoa(rand.Int())
			service, errS = client.Services.Create(projectName, CreateServiceRequest{
				Cloud:        "google-europe-west1",
				Plan:         "business-4",
				ProjectVPCID: nil,
				ServiceName:  serviceName,
				ServiceType:  "kafka",
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

		// kafka topic
		var (
			errC          error
			quotaUserName string
		)
		quotaUserName = "test1"

		It("create kafka quota", func() {
			time.Sleep(10 * time.Second)
			errC = client.KafkaQuotas.Create(projectName, serviceName, KafkaQuota{
				User:              quotaUserName,
				ClientId:          "",
				ConsumerByteRate:  42,
				ProducerByteRate:  43,
				RequestPercentage: 44,
			})

			Eventually(func() string {
				quotas, _ := client.KafkaQuotas.List(projectName, serviceName)

				if quotas != nil {
					for _, quota := range quotas {
						if quota.User == quotaUserName {
							return quota.User
						}
					}
					return ""
				}

				return ""
			}, 25*time.Minute, 1*time.Minute).Should(Equal("test1"))
		})

		It("should not error kafka quota with config", func() {
			Expect(errC).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			quotas, errQ := client.KafkaQuotas.List(projectName, serviceName)
			var foundQuota *KafkaQuota
			if quotas != nil {
				for _, quota := range quotas {
					if quota.User == quotaUserName {
						foundQuota = quota
					}
				}
			}
			Expect(errQ).NotTo(HaveOccurred())

			if foundQuota != nil {
				Expect(foundQuota.User).To(Equal("test1"))
				Expect(foundQuota.ClientId).To(BeEmpty())
				Expect(foundQuota.RequestPercentage).To(Equal(44))
				Expect(foundQuota.ConsumerByteRate).To(Equal(42))
				Expect(foundQuota.ProducerByteRate).To(Equal(43))
			}
		})

		It("should update topic config", func() {

			errU := client.KafkaQuotas.Update(projectName, serviceName, KafkaQuota{
				User:              quotaUserName,
				ClientId:          "",
				ConsumerByteRate:  45,
				ProducerByteRate:  46,
				RequestPercentage: 47,
			})
			Expect(errU).NotTo(HaveOccurred())

			quotas, errQ := client.KafkaQuotas.List(projectName, serviceName)
			var foundQuota *KafkaQuota
			if quotas != nil {
				for _, quota := range quotas {
					if quota.User == quotaUserName {
						foundQuota = quota
					}
				}
			}
			Expect(errQ).NotTo(HaveOccurred())
			Expect(foundQuota).NotTo(BeNil())

			if foundQuota != nil {
				Expect(foundQuota.RequestPercentage).To(Equal(47))
				Expect(foundQuota.ConsumerByteRate).To(Equal(45))
				Expect(foundQuota.ProducerByteRate).To(Equal(46))
			}
		})

		It("delete Kafka Topic and Kafka service", func() {
			if errD := client.KafkaQuotas.Delete(projectName, serviceName, DeleteKafkaQuotaRequest{User: quotaUserName, ClientId: ""}); errD != nil {
				Fail("cannot delete kafka quota" + errD.Error())
			}

			if errD := client.Services.Delete(projectName, serviceName); errD != nil {
				Fail("cannot delete service:" + errD.Error())
			}
		})
	})
})
