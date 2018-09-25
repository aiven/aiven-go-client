package testhelpers

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aiven/aiven-go-client"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	// ServicePlan is the plan we'll use to test services on
	ServicePlan = "hobbyist"
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
)

// RandStringRunes returns a random string of length n.
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Client returns a new Aiven client which pulls the credentials from the
// environment.
func Client() *aiven.Client {
	cl, err := aiven.NewUserClient(os.Getenv("AIVEN_USERNAME"), os.Getenv("AIVEN_PASSWORD"))
	if err != nil {
		fmt.Println("Client creation")
		panic(err)
	}

	return cl
}

// ProjectName will take a base name and make it unique.
func ProjectName(name string) string {
	return fmt.Sprintf("%s-%s", name, RandStringRunes(10))
}

// NewProject will create a new project based on the given name.
func NewProject(cl *aiven.Client, name string) (*aiven.Project, error) {
	cards, err := cl.CardsHandler.List()
	if err != nil {
		return nil, err
	}
	if len(cards) == 0 {
		return nil, errors.New("No card available")
	}

	return cl.Projects.Create(aiven.CreateProjectRequest{
		CardID:  cards[0].CardID,
		Project: name,
	})
}
