package aiven

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strconv"
)

var _ = Describe("BillingGroup", func() {
	var (
		billingGName string
		billingG     *BillingGroup
		err          error
	)

	BeforeEach(func() {
		billingGName = "test-acc-bg-" + strconv.Itoa(rand.Int())
		billingG, err = client.BillingGroup.Create(BillingGroupRequest{
			BillingGroupName: billingGName,
			Company:          ToStringPointer("testC1"),
			AddressLines:     []string{"NYC Some Street 123 A"},
			CountryCode:      ToStringPointer("US"),
			City:             ToStringPointer("NY"),
			ZipCode:          ToStringPointer("101778"),
			BillingCurrency:  ToStringPointer("EUR"),
		})
	})

	Context("Billing group tests", func() {
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			Expect(billingG).NotTo(BeNil())

			if billingG != nil {
				Expect(*billingG.Company).NotTo(BeEmpty())
				Expect(billingG.AddressLines).NotTo(BeEmpty())
				Expect(*billingG.CountryCode).NotTo(BeEmpty())
				Expect(*billingG.City).NotTo(BeEmpty())
				Expect(*billingG.ZipCode).NotTo(BeEmpty())
				Expect(*billingG.BillingCurrency).NotTo(BeEmpty())
				Expect(billingG.BillingGroupName).NotTo(BeEmpty())
			}
		})

		It("update and get checks", func() {
			_, err = client.BillingGroup.Update(billingG.Id, BillingGroupRequest{
				BillingExtraText: ToStringPointer("some text ..."),
			})
			Expect(err).NotTo(HaveOccurred())

			billingG, err = client.BillingGroup.Get(billingG.Id)
			Expect(err).NotTo(HaveOccurred())

			if billingG != nil {
				Expect(*billingG.Company).NotTo(BeEmpty())
				Expect(billingG.AddressLines).NotTo(BeEmpty())
				Expect(*billingG.CountryCode).NotTo(BeEmpty())
				Expect(*billingG.City).NotTo(BeEmpty())
				Expect(*billingG.ZipCode).NotTo(BeEmpty())
				Expect(*billingG.BillingCurrency).NotTo(BeEmpty())
				Expect(billingG.BillingGroupName).NotTo(BeEmpty())
				Expect(*billingG.BillingExtraText).NotTo(BeEmpty())
			}
		})

		It("assign a project", func() {
			projectName := "test-acc-pr-" + strconv.Itoa(rand.Int())
			_, err = client.Projects.Create(CreateProjectRequest{
				Project: projectName,
				VatID:   ToStringPointer(""),
			})
			Expect(err).NotTo(HaveOccurred())

			err = client.BillingGroup.AssignProjects(billingG.Id, []string{projectName})
			Expect(err).NotTo(HaveOccurred())

			projects, err := client.BillingGroup.GetProjects(billingG.Id)
			Expect(projects).NotTo(BeEmpty())

			err = client.Projects.Delete(projectName)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	AfterEach(func() {
		err = client.BillingGroup.Delete(billingG.Id)
		if err != nil {
			Fail("cannot delete billing group : " + err.Error())
		}
	})
})
