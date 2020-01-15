package main

import (
	client "github.com/aiven/aiven-go-client"
	"log"
	"os"
)

func main() {
	// create new user client
	c, err := client.NewUserClient(
		os.Getenv("AIVEN_USERNAME"),
		os.Getenv("AIVEN_PASSWORD"), "aiven-go-client-test/"+client.Version())
	if err != nil {
		log.Fatalf("user authentication error: %s", err)
	}

	// create account
	acc, err := c.Accounts.Create(client.Account{
		Name: "test-acc1@aiven.io",
	})
	if err != nil {
		log.Fatalf("cannot create account err: %s", err)
	}
	log.Printf("account created %v", acc)

	// get account by id
	accG, err := c.Accounts.Get(acc.Account.Id)
	if err != nil {
		log.Fatalf("cannot get account err: %s", err)
	}
	log.Printf("account get %v", accG)

	// update account
	accU, err := c.Accounts.Update(accG.Account.Id, client.Account{
		Name: "test-acc1+update@aiven.io",
	})
	if err != nil {
		log.Fatalf("cannot create account err: %s", err)
	}
	log.Printf("account update %v", accU)

	// delete account
	err = c.Accounts.Delete(accG.Account.Id)
	if err != nil {
		log.Fatalf("cannot delete account err: %s", err)
	}
	log.Printf("accont with id %s was deleted", accG.Account.Id)
}
