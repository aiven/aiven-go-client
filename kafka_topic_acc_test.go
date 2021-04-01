package aiven

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var _ = Describe("Kafka Topic", func() {
	var (
		projectName string
		project     *Project
		err         error
	)

	Context("Kafka Topic CRUD", func() {
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
			serviceName = "test-acc-kafka-topic-sr-" + strconv.Itoa(rand.Int())
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
			errC            error
			topicName       string
			segmentJitterMs int64
		)
		segmentJitterMs = 10
		topicName = "test1"

		It("create kafka topic", func() {
			time.Sleep(10 * time.Second)
			errC = client.KafkaTopics.Create(projectName, serviceName, CreateKafkaTopicRequest{
				TopicName: topicName,
				Config: KafkaTopicConfig{
					CleanupPolicy:   "compact",
					SegmentJitterMs: &segmentJitterMs,
				},
				Tags: []KafkaTopicTag{
					{
						Key:   "tag1-key",
						Value: "tag1-value",
					},
				},
			})

			Eventually(func() string {
				topic, _ := client.KafkaTopics.Get(projectName, serviceName, topicName)

				if topic != nil {
					return topic.State
				}

				return ""
			}, 25*time.Minute, 1*time.Minute).Should(Equal("ACTIVE"))
		})

		It("should not error kafka topic with config", func() {
			Expect(errC).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			t, errT := client.KafkaTopics.Get(projectName, serviceName, topicName)
			Expect(errT).NotTo(HaveOccurred())

			if t != nil {
				Expect(t.Config.CleanupPolicy.Value).NotTo(BeEmpty())
				Expect(t.Config.SegmentJitterMs.Value).To(Equal(segmentJitterMs))
				Expect(t.Tags).NotTo(BeEmpty())
				Expect(t.Tags[0].Key).To(Equal("tag1-key"))
				Expect(t.Tags[0].Value).To(Equal("tag1-value"))
			}
		})

		It("should update topic config", func() {
			var uncleanLeaderElectionEnable = true

			errU := client.KafkaTopics.Update(projectName, serviceName, topicName, UpdateKafkaTopicRequest{
				Config: KafkaTopicConfig{
					UncleanLeaderElectionEnable: &uncleanLeaderElectionEnable,
				},
				Tags: []KafkaTopicTag{
					{
						Key:   "tag1-key",
						Value: "tag1-value",
					},
					{
						Key:   "tag2-key",
						Value: "tag2-value",
					},
				},
			})
			Expect(errU).NotTo(HaveOccurred())

			t2, errG := client.KafkaTopics.Get(projectName, serviceName, topicName)
			Expect(errG).NotTo(HaveOccurred())
			Expect(t2).NotTo(BeNil())

			if t2 != nil {
				Expect(t2.Config.UncleanLeaderElectionEnable.Value).Should(Equal(true))
				Expect(t2.Tags).ShouldNot(BeEmpty())
				Expect(len(t2.Tags)).To(Equal(2))
			}
		})

		It("list v2", func() {
			list, errV2 := client.KafkaTopics.V2List(projectName, serviceName, []string{topicName})
			Expect(errV2).NotTo(HaveOccurred())

			Expect(len(list)).Should(Equal(1))
			Expect(list[0].Config.CleanupPolicy.Value).NotTo(BeEmpty())
			Expect(list[0].Config.SegmentJitterMs.Value).To(Equal(segmentJitterMs))
			Expect(list[0].Tags).NotTo(BeEmpty())
			Expect(len(list[0].Tags)).To(Equal(2))
		})

		It("delete Kafka Topic and Kafka service", func() {
			if errD := client.KafkaTopics.Delete(projectName, serviceName, topicName); errD != nil {
				Fail("cannot delete kafka topic:" + errD.Error())
			}

			if errD := client.Services.Delete(projectName, serviceName); errD != nil {
				Fail("cannot delete service:" + errD.Error())
			}
		})
	})
})
