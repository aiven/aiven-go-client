package aiven

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"strconv"
)

var _ = Describe("Accounts", func() {
	var (
		accountName string
		account     *AccountResponse
		err         error
	)

	BeforeEach(func() {
		accountName = "test-acc-account" + strconv.Itoa(rand.Int())
		account, err = client.Accounts.Create(Account{Name: accountName})
	})

	Context("Create new account", func() {
		It("should not error", func() {
			Expect(err).NotTo(HaveOccurred())
		})

		It("should populate fields properly", func() {
			Expect(account.Account.Id).NotTo(BeEmpty())
			Expect(account.Account.Name).To(Equal(accountName))
			Expect(account.Account.CreateTime).NotTo(BeNil())
			Expect(account.Account.UpdateTime).NotTo(BeNil())
			Expect(account.APIResponse.Message).To(BeEmpty())
			Expect(account.APIResponse.Errors).To(BeEmpty())
		})
	})

	Describe("AccountTeams", func() {
		var (
			teamName string
			team     *AccountTeamResponse
			errT     error
		)

		BeforeEach(func() {
			if account == nil {
				Fail("account is nil")
			}
			teamName = "test-acc-team" + strconv.Itoa(rand.Int())
			team, errT = client.AccountTeams.Create(account.Account.Id, AccountTeam{
				Name: teamName,
			})
		})

		Context("Create account team", func() {
			It("should not error", func() {
				Expect(errT).NotTo(HaveOccurred())
			})

			It("should populate fields properly", func() {
				Expect(team.Team.Id).NotTo(BeEmpty())
				Expect(team.Team.Name).To(Equal(teamName))
				Expect(team.Team.UpdateTime).NotTo(BeNil())
				Expect(team.Team.CreateTime).NotTo(BeNil())
				Expect(team.APIResponse.Message).To(BeEmpty())
				Expect(team.APIResponse.Errors).To(BeEmpty())
			})

			var (
				memberEmail string
				errM        error
			)

			Context("accountTeams", func() {

			})
			JustBeforeEach(func() {
				memberEmail = "savciuci+" + strconv.Itoa(rand.Int()) + "@aiven.io"
				errM = client.AccountTeamMembers.Invite(account.Account.Id, team.Team.Id, memberEmail)
			})
			Context("Invite account team member", func() {
				It("should not error", func() {
					Expect(memberEmail).NotTo(BeEmpty())
					Expect(errM).NotTo(HaveOccurred())
				})

				It("should be in the team members list", func() {

				})

				It("should send invite", func() {
					invites, errI := client.AccountTeamInvites.List(account.Account.Id, team.Team.Id)
					Expect(errI).NotTo(HaveOccurred())

					var found bool
					for _, invite := range invites.Invites {
						if invite.UserEmail == memberEmail {
							found = true
						}
					}

					Expect(found).To(Equal(true), "cannot find invitation for newly created member")

					if errD := client.AccountTeamInvites.Delete(account.Account.Id, team.Team.Id, memberEmail); errD != nil {
						Fail("cannot delete an invitation :" + errD.Error())
					}
				})
			})

			AfterEach(func() {
				if list, errL := client.AccountTeamMembers.List(account.Account.Id, team.Team.Id); errL != nil {
					Fail("cannot get a list of account team members:" + errL.Error())
				} else {
					for _, m := range list.Members {
						if m.UserEmail == memberEmail {
							if errD := client.AccountTeamMembers.Delete(account.Account.Id, team.Team.Id, memberEmail); errD != nil {
								Fail("cannot delete account team member:" + errD.Error())
							}
						}
					}
				}
			})

			Context("AccountTeamProjects", func() {
				var (
					projectName string
					projectType string = "admin"
				)

				BeforeEach(func() {
					projectName = "test-acc-pr" + strconv.Itoa(rand.Int())

					By("Create project")
					if _, errP := client.Projects.Create(CreateProjectRequest{
						Project:   projectName,
						AccountId: account.Account.Id,
					}); errP != nil {
						Fail("cannot create project :" + errP.Error())
					}

					By("Create account team project association")
					if errTP := client.AccountTeamProjects.Create(
						account.Account.Id,
						team.Team.Id,
						AccountTeamProject{ProjectName: projectName, TeamType: projectType}); errTP != nil {
						Fail("cannot create account team project association:" + errTP.Error())
					}
				})
				Context("Account team project association", func() {
					It("should be in the list", func() {
						projects, errL := client.AccountTeamProjects.List(account.Account.Id, team.Team.Id)
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
				})
				AfterEach(func() {
					if errD := client.AccountTeamProjects.Delete(account.Account.Id, team.Team.Id, projectName); errD != nil {
						Fail("cannot delete account team project association :" + errD.Error())
					}

					if errD := client.Projects.Delete(projectName); errD != nil {
						Fail("cannot delete project :" + errD.Error())
					}
				})

			})
		})

		AfterEach(func() {
			if errD := client.AccountTeams.Delete(account.Account.Id, team.Team.Id); errD != nil {
				Fail("cannot delete account team :" + errD.Error())
			}
		})

	})

	Context("AccountAuthentications", func() {
		It("check authentication methods", func() {
			list, errL := client.AccountAuthentications.List(account.Account.Id)
			Expect(errL).NotTo(HaveOccurred())
			Expect(list.AuthenticationMethods).NotTo(BeEmpty(), "default auth methods should be created automatically")
		})

		It("add new one", func() {
			resp, errL := client.AccountAuthentications.Create(account.Account.Id, AccountAuthenticationMethod{
				Name: "test-auth",
				Type: "saml",
			})
			Expect(errL).NotTo(HaveOccurred())
			Expect(resp.AuthenticationMethod.Id).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.SAMLMetadataUrl).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.SAMLAcsUrl).NotTo(BeEmpty())
			Expect(resp.AuthenticationMethod.AccountId).To(Equal(account.Account.Id))
			Expect(resp.APIResponse.Message).To(BeEmpty())
			Expect(resp.APIResponse.Errors).To(BeEmpty())
		})
	})

	AfterEach(func() {
		if errD := client.Accounts.Delete(account.Account.Id); errD != nil {
			Fail("cannot delete account :" + errD.Error())
		}
	})
})
