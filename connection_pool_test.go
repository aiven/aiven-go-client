package aiven_test

import (
	"testing"

	"github.com/jelmersnoeck/aiven/internal/test_helpers"
)

func TestConnectionPool(t *testing.T) {
	pn := test_helpers.ProjectName("connection-pool")
	cl := test_helpers.Client()
	_, err := test_helpers.NewProject(cl, pn)
	if err != nil {
		t.Errorf("Error creating project: %s", err)
		return
	}
	defer func() {
		if err := cl.Projects.Delete(pn); err != nil {
			t.Errorf("Error deleting project: %s", err)
		}
	}()
}
