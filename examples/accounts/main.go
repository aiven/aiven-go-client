package main

import (
	"context"
	"log"
	"os"

	client "github.com/aiven/aiven-go-client/v2"
)

func main() {
	ctx := context.Background()

	// create new user client
	c, err := client.NewUserClient(
		os.Getenv("AIVEN_USERNAME"),
		os.Getenv("AIVEN_PASSWORD"), "aiven-go-client-test/"+client.Version())
	if err != nil {
		log.Fatalf("user authentication error: %s", err)
	}

	// create account
	acc, err := c.Accounts.Create(ctx, client.Account{
		Name: "test-acc1@aiven.io",
	})
	if err != nil {
		log.Fatalf("cannot create account err: %s", err)
	}
	log.Printf("account created %v", acc)

	// get account by id
	accG, err := c.Accounts.Get(ctx, acc.Account.Id)
	if err != nil {
		log.Fatalf("cannot get account err: %s", err)
	}
	log.Printf("account get %v", accG)

	// update account
	accU, err := c.Accounts.Update(ctx, accG.Account.Id, client.Account{
		Name: "test-acc1+update@aiven.io",
	})
	if err != nil {
		log.Fatalf("cannot update account err: %s", err)
	}
	log.Printf("account update %v", accU)

	// create a team
	team, err := c.AccountTeams.Create(ctx, accU.Account.Id, client.AccountTeam{
		Name: "test-team1",
	})
	if err != nil {
		log.Fatalf("cannot create account team err: %s", err)
	}

	teamU, err := c.AccountTeams.Update(ctx, accU.Account.Id, team.Team.Id, client.AccountTeam{
		Name: "test-team2",
	})
	if err != nil {
		log.Fatalf("cannot update account team err: %s", err)
	}
	log.Printf("account team %v", teamU)

	// delete account
	err = c.Accounts.Delete(ctx, accG.Account.Id)
	if err != nil {
		log.Fatalf("cannot delete account err: %s", err)
	}
	log.Printf("accont with id %s was deleted", accG.Account.Id)
}
