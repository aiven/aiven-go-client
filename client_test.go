package aiven

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Init(t *testing.T) {
	var c Client = Client{}
	c.Init()
}

func TestClient_Context(t *testing.T) {
	c, err := NewTokenClient("key", "user-agent")
	require.NoError(t, err)

	ctxC := c.WithContext(context.Background())
	assert.Nil(t, c.ctx)
	assert.NotNil(t, ctxC.ctx)
	assert.Equal(t, c.Client, ctxC.Client)
	assert.Equal(t, c.APIKey, ctxC.APIKey)
	assert.Equal(t, c.UserAgent, ctxC.UserAgent)
}
