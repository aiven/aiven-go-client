package aiven

import (
	"testing"
)

func TestClient_Init(t *testing.T) {
	var c Client = Client{}
	c.Init()
}
