package aiven

import (
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAivenGoClient(t *testing.T) {
	if os.Getenv("AIVEN_ACC") != "" {
		RegisterFailHandler(Fail)
		RunSpecs(t, "AivenGoClient Suite")
	}
}

var (
	client *Client
)

var _ = BeforeSuite(func() {
	var (
		err error
	)

	projectName := os.Getenv("AIVEN_PROJECT_NAME")
	if projectName == "" {
		Fail("cannot create Aiven API client, `AIVEN_PROJECT_NAME` is required")
	}

	client, err = SetupEnvClient("aiven-go-client/test")
	if err != nil {
		Fail("cannot create Aiven API client :" + err.Error())
	}
})

var _ = Describe("Check client", func() {
	Context("Create new client using username and password", func() {
		It("should be valid client", func() {
			Expect(client.Client).NotTo(BeNil())
		})

		It("should have an API token", func() {
			Expect(client.APIKey).NotTo(Equal("some-random-token"))
			Expect(client.APIKey).NotTo(BeEmpty())
		})
	})
})

// generateRandomString generate a random id
func generateRandomID() string {
	var src = rand.NewSource(time.Now().UnixNano())
	return strconv.FormatInt(src.Int63(), 10)
}
