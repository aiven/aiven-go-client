package aiven

import (
	"context"
	"math/rand"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Accounts", func() {
	ctx := context.Background()
	Context("Check all possible accounts interactions", func() {
		// Account
		var (
			accountName string
			account     *AccountResponse
			err         error
		)

		It("Account creation should not error", func() {
			accountName = "test-acc-account-" + strconv.Itoa(rand.Int())
			account, err = client.Accounts.Create(ctx, Account{Name: accountName})
			Expect(err).NotTo(HaveOccurred())
		})

		It("Account creation should populate fields properly", func() {
			Expect(account.Account.Id).NotTo(BeEmpty())
			Expect(account.Account.Name).To(Equal(accountName))
			Expect(account.Account.CreateTime).NotTo(BeNil())
			Expect(account.Account.UpdateTime).NotTo(BeNil())
			Expect(account.Account.IsAccountOwner).To(Equal(true))
			Expect(account.Account.PrimaryBillingGroupId).NotTo(BeNil())
			Expect(account.APIResponse.Message).To(BeEmpty())
			Expect(account.APIResponse.Errors).To(BeEmpty())
		})

		// AccountTeams
		var (
			teamName string
			team     *AccountTeamResponse
			errT     error
		)

		It("Create account team", func() {
			if account == nil {
				Fail("account is nil")
			}
			teamName = "test-acc-team" + strconv.Itoa(rand.Int())
			team, errT = client.AccountTeams.Create(ctx, account.Account.Id, AccountTeam{
				Name: teamName,
			})

			Expect(errT).NotTo(HaveOccurred())

			Expect(team.Team.Id).NotTo(BeEmpty())
			Expect(team.Team.Name).To(Equal(teamName))
			Expect(team.Team.UpdateTime).NotTo(BeNil())
			Expect(team.Team.CreateTime).NotTo(BeNil())
			Expect(team.APIResponse.Message).To(BeEmpty())
			Expect(team.APIResponse.Errors).To(BeEmpty())
		})

		// Account members
		var (
			memberEmail string
			errM        error
		)

		It("should not error", func() {
			memberEmail = "ivan.savciuc+" + strconv.Itoa(rand.Int()) + "@aiven.io"
			errM = client.AccountTeamMembers.Invite(ctx, account.Account.Id, team.Team.Id, memberEmail)

			Expect(memberEmail).NotTo(BeEmpty())
			Expect(errM).NotTo(HaveOccurred())
		})

		It("should send invite", func() {
			invites, errI := client.AccountTeamInvites.List(ctx, account.Account.Id, team.Team.Id)
			Expect(errI).NotTo(HaveOccurred())

			var found bool
			for _, invite := range invites.Invites {
				if invite.UserEmail == memberEmail {
					found = true
				}
			}

			Expect(found).To(Equal(true), "cannot find invitation for newly created member")

			if errD := client.AccountTeamInvites.Delete(ctx, account.Account.Id, team.Team.Id, memberEmail); errD != nil {
				Fail("cannot delete an invitation :" + errD.Error())
			}
		})

		// Account project
		var (
			projectName string
			projectType = "admin"
		)

		It("Account projects creation ", func() {
			projectName = "test-acc-pr" + strconv.Itoa(rand.Int())

			By("Create project")
			if _, errP := client.Projects.Create(ctx, CreateProjectRequest{
				Project:   projectName,
				AccountId: ToStringPointer(account.Account.Id),
				Tags:      map[string]string{},
			}); errP != nil {
				Fail("cannot create project :" + errP.Error())
			}

			By("Create account team project association")
			if errTP := client.AccountTeamProjects.Create(
				ctx,
				account.Account.Id,
				team.Team.Id,
				AccountTeamProject{ProjectName: projectName, TeamType: "developer"}); errTP != nil {
				Fail("cannot create account team project association:" + errTP.Error())
			}

			By("Update account team project")
			if errTPu := client.AccountTeamProjects.Update(
				ctx,
				account.Account.Id,
				team.Team.Id,
				AccountTeamProject{ProjectName: projectName, TeamType: projectType}); errTPu != nil {
				Fail("cannot update account team project:" + errTPu.Error())
			}
		})

		It("Account team project association should be in the list", func() {
			projects, errL := client.AccountTeamProjects.List(ctx, account.Account.Id, team.Team.Id)
			Expect(errL).NotTo(HaveOccurred())

			if projects != nil {
				var found bool
				for _, p := range projects.Projects {
					if p.ProjectName == projectName {
						Expect(p.TeamType).To(Equal(projectType))
						found = true
					}
				}

				Expect(found).To(Equal(true), "cannot find project in the account team projects list")
			}
		})

		It("AccountAuthentications check authentication methods", func() {
			list, errL := client.AccountAuthentications.List(ctx, account.Account.Id)
			Expect(errL).NotTo(HaveOccurred())
			Expect(list.AuthenticationMethods).NotTo(BeEmpty(), "default auth methods should be created automatically")
		})

		It("AccountAuthentications  add new one", func() {
			resp, errL := client.AccountAuthentications.Create(ctx, account.Account.Id, AccountAuthenticationMethodCreate{
				AuthenticationMethodName: "test-auth",
				AuthenticationMethodType: "saml",
			})
			Expect(errL).NotTo(HaveOccurred())
			Expect(resp.AuthenticationMethod.AuthenticationMethodID).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.SAMLMetadataURL).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.SAMLAcsURL).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.AccountID).To(Equal(account.Account.Id))
			Expect(resp.APIResponse.Message).To(BeEmpty())
			Expect(resp.APIResponse.Errors).To(BeEmpty())
		})

		It("remove all the stuff", func() {
			if errD := client.AccountTeamProjects.Delete(ctx, account.Account.Id, team.Team.Id, projectName); errD != nil {
				Fail("cannot delete account team project association :" + errD.Error())
			}

			if errD := client.Projects.Delete(ctx, projectName); errD != nil {
				Fail("cannot delete project :" + errD.Error())
			}

			if list, errL := client.AccountTeamMembers.List(ctx, account.Account.Id, team.Team.Id); errL != nil {
				Fail("cannot get a list of account team members:" + errL.Error())
			} else {
				for _, m := range list.Members {
					if errD := client.AccountTeamMembers.Delete(ctx, account.Account.Id, team.Team.Id, m.UserEmail); errD != nil {
						Fail("cannot delete account team member:" + errD.Error())
					}
				}
			}
			if errD := client.AccountTeams.Delete(ctx, account.Account.Id, team.Team.Id); errD != nil {
				Fail("cannot delete account team :" + errD.Error())
			}

			if errD := client.Accounts.Delete(ctx, account.Account.Id); errD != nil {
				Fail("cannot delete account :" + errD.Error())
			}
		})
	})
})
