package aiven

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Projects", func() {
	var (
		projectName  string
		billingGName string
		project      *Project
		billingG     *BillingGroup
		err          error
	)

	ctx := context.Background()

	BeforeEach(func() {
		billingGName = "test-acc-bg-" + generateRandomID()
		billingG, err = client.BillingGroup.Create(ctx, BillingGroupRequest{
			BillingGroupName: billingGName,
			Company:          ToStringPointer("testC1"),
			AddressLines:     []string{"NYC Some Street 123 A"},
			CountryCode:      ToStringPointer("US"),
			City:             ToStringPointer("NY"),
			ZipCode:          ToStringPointer("101778"),
			BillingCurrency:  ToStringPointer("EUR"),
		})

		projectName = "test-acc-pr" + generateRandomID()
		project, err = client.Projects.Create(ctx, CreateProjectRequest{
			Project:                      projectName,
			BillingCurrency:              "EUR",
			TechnicalEmails:              ContactEmailFromStringSlice([]string{"test@example.com"}),
			UseSourceProjectBillingGroup: false,
			BillingGroupId:               billingG.Id,
			Tags:                         map[string]string{},
		})
	})

	Context("Create new project", func() {
		It("should not error", func() {
			if !IsAlreadyExists(err) {
				Expect(err).NotTo(HaveOccurred())
			}
		})

		It("should populate fields properly", func() {
			project, err = client.Projects.Get(ctx, projectName)
			Expect(err).NotTo(HaveOccurred())
			Expect(project).NotTo(BeNil())

			if project != nil {
				Expect(project.Name).NotTo(BeEmpty())
				Expect(project.AccountId).To(BeEmpty())
				Expect(project.BillingCurrency).NotTo(BeEmpty())
				Expect(project.GetTechnicalEmailsAsStringSlice()).NotTo(BeEmpty())
				Expect(project.BillingGroupId).NotTo(BeEmpty())
				Expect(project.BillingGroupName).Should(Equal(billingGName))
			}
		})

		It("update project name", func() {
			project, err = client.Projects.Update(ctx, projectName, UpdateProjectRequest{
				Name: projectName + "-new",
				Tags: map[string]string{},
			})

			if err == nil {
				projectName = projectName + "-new"
			}

			Expect(err).NotTo(HaveOccurred())
			Expect(project).NotTo(BeNil())
		})
	})

	Context("Get project event logs", func() {
		It("should be event logs", func() {
			events, logErr := client.Projects.GetEventLog(ctx, projectName)
			Expect(logErr).To(BeNil())
			Expect(events).ToNot(BeNil())
			Expect(events).Should(Not(BeEmpty()))
			for _, event := range events {
				Expect(event).NotTo(BeNil())
			}
		})
	})

	Context("Get service types", func() {
		It("check returned service types", func() {
			serviceTypes, err := client.Projects.ServiceTypes(ctx, projectName)
			Expect(err).To(BeNil())
			Expect(serviceTypes).ToNot(BeNil())
			Expect(serviceTypes).Should(Not(BeEmpty()))
			for _, serviceType := range serviceTypes {
				Expect(serviceType).NotTo(BeNil())
			}
		})
	})

	Context("Get integration types", func() {
		It("check returned integration types", func() {
			integrationTypes, err := client.Projects.IntegrationTypes(ctx, projectName)
			Expect(err).To(BeNil())
			Expect(integrationTypes).ToNot(BeNil())
			Expect(integrationTypes).Should(Not(BeEmpty()))
			for _, integrationType := range integrationTypes {
				Expect(integrationType).NotTo(BeNil())
			}
		})
	})

	Context("Get integration endpoint types", func() {
		It("check returned integration endpoint types", func() {
			endpointTypes, err := client.Projects.IntegrationEndpointTypes(ctx, projectName)
			Expect(err).To(BeNil())
			Expect(endpointTypes).ToNot(BeNil())
			Expect(endpointTypes).Should(Not(BeEmpty()))
			for _, endpointType := range endpointTypes {
				Expect(endpointType).NotTo(BeNil())
			}
		})
	})

	AfterEach(func() {
		err = client.Projects.Delete(ctx, projectName)
		if err != nil {
			Fail("cannot delete project : " + err.Error())
		}

		err = client.BillingGroup.Delete(ctx, billingG.Id)
		if err != nil {
			Fail("cannot delete billing group : " + err.Error())
		}
	})
})
