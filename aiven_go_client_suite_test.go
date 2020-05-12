package aiven

import (
	"os"
	"testing"

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

	url := os.Getenv("AIVEN_WEB_URL")
	if url == "" {
		Fail("environment variable `AIVEN_WEB_URL` is not set")
	}
	apiurl = url + "/v1"

	token := os.Getenv("AIVEN_TOKEN")
	if token == "" {
		Fail("cannot create Aiven API client, `AIVEN_TOKEN` is required")
	}

	cardId := os.Getenv("AIVEN_CARD_ID")
	if cardId == "" {
		Fail("cannot create Aiven API client, `AIVEN_CARD_ID` is required")
	}

	client, err = NewTokenClient(
		token,
		"aiven-go-client-test/")

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
