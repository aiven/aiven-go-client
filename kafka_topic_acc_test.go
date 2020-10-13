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
				Expect(project.AccountId).To(BeEmpty())
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
			segmentJitterMs int
		)
		segmentJitterMs = 10
		topicName = "test1"

		It("create kafka topic", func() {
			time.Sleep(10 * time.Second)
			errC = client.KafkaTopics.Create(projectName, serviceName, CreateKafkaTopicRequest{
				CleanupPolicy:         nil,
				MinimumInSyncReplicas: nil,
				Partitions:            nil,
				Replication:           nil,
				RetentionBytes:        nil,
				RetentionHours:        nil,
				TopicName:             topicName,
				Config: KafkaTopicConfig{
					CleanupPolicy:                   "compact",
					CompressionType:                 "",
					DeleteRetentionMs:               nil,
					FileDeleteDelayMs:               nil,
					FlushMessages:                   nil,
					FlushMs:                         nil,
					IndexIntervalBytes:              nil,
					MaxCompactionLagMs:              nil,
					MaxMessageBytes:                 nil,
					MessageDownconversionEnable:     nil,
					MessageFormatVersion:            "",
					MessageTimestampDifferenceMaxMs: nil,
					MessageTimestampType:            "",
					MinCleanableDirtyRatio:          nil,
					MinCompactionLagMs:              nil,
					MinInsyncReplicas:               nil,
					Preallocate:                     nil,
					RetentionBytes:                  nil,
					RetentionMs:                     nil,
					SegmentBytes:                    nil,
					SegmentIndexBytes:               nil,
					SegmentJitterMs:                 &segmentJitterMs,
					SegmentMs:                       nil,
					UncleanLeaderElectionEnable:     nil,
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
			}
		})

		It("should update topic config", func() {
			var uncleanLeaderElectionEnable = true

			errU := client.KafkaTopics.Update(projectName, serviceName, topicName, UpdateKafkaTopicRequest{Config: KafkaTopicConfig{
				UncleanLeaderElectionEnable: &uncleanLeaderElectionEnable,
			}})
			Expect(errU).NotTo(HaveOccurred())

			t2, errG := client.KafkaTopics.Get(projectName, serviceName, topicName)
			Expect(errG).NotTo(HaveOccurred())
			Expect(t2).NotTo(BeNil())

			if t2 != nil {
				Expect(t2.Config.UncleanLeaderElectionEnable.Value).Should(Equal(true))
			}
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
