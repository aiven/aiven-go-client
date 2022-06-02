package aiven

import (
	"math/rand"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BillingGroup", func() {
	var (
		billingGName string
		billingG     *BillingGroup
		copiedBG     *BillingGroup
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

		It("check empty assigned projects list", func() {
			projects, err := client.BillingGroup.GetProjects(billingG.Id)
			Expect(err).NotTo(HaveOccurred())
			Expect(projects).To(BeEmpty())
		})

		It("assign a project", func() {
			projectName := "test-acc-pr-" + strconv.Itoa(rand.Int())
			_, err = client.Projects.Create(CreateProjectRequest{
				Project: projectName,
				VatID:   ToStringPointer(""),
				Tags:    map[string]string{},
			})
			Expect(err).NotTo(HaveOccurred())

			err = client.BillingGroup.AssignProjects(billingG.Id, []string{projectName})
			Expect(err).NotTo(HaveOccurred())

			projects, errG := client.BillingGroup.GetProjects(billingG.Id)
			Expect(errG).NotTo(HaveOccurred())
			Expect(projects).NotTo(BeEmpty())

			err = client.Projects.Delete(projectName)
			Expect(err).NotTo(HaveOccurred())
		})

		It("list all billing groups", func() {
			list, err := client.BillingGroup.ListAll()
			Expect(err).NotTo(HaveOccurred())
			Expect(list).NotTo(BeEmpty())
		})

		It("create billing group by copy from an existing one", func() {
			billingGName = "copy-from-" + billingG.BillingGroupName
			copiedBG, err = client.BillingGroup.Create(BillingGroupRequest{
				BillingGroupName:     billingGName,
				CopyFromBillingGroup: ToStringPointer(billingG.Id),
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(copiedBG.Id).NotTo(BeNil())
			Expect(copiedBG.Id).NotTo(Equal(billingG.Id))
			Expect(copiedBG.Company).To(Equal(billingG.Company))
			Expect(copiedBG.AddressLines).To(Equal(billingG.AddressLines))
			Expect(copiedBG.CountryCode).To(Equal(billingG.CountryCode))
			Expect(copiedBG.City).To(Equal(billingG.City))
			Expect(copiedBG.ZipCode).To(Equal(billingG.ZipCode))
			Expect(copiedBG.BillingCurrency).To(Equal(billingG.BillingCurrency))
		})
	})

	AfterEach(func() {
		err = client.BillingGroup.Delete(billingG.Id)
		if err != nil {
			Fail("cannot delete billing group : " + err.Error())
		}
	})
})
