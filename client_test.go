package aiven

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func TestCheckRetry(t *testing.T) {
	type option struct {
		name     string // test name
		code     int    // response status code
		expected bool   // expected result
		body     string
		method   string
	}

	opts := []*option{
		{
			name:     "Retries 408",
			code:     http.StatusRequestTimeout,
			expected: true,
		},
		{
			name:     "Retries 429",
			code:     http.StatusTooManyRequests,
			expected: true,
		},
		{
			name:     "Retries 404",
			code:     http.StatusNotFound,
			body:     "Service 'test-foo' does not exist",
			method:   "POST",
			expected: true,
		},
		{
			name:     "Does not retry 404 with different error message",
			code:     http.StatusNotFound,
			body:     "User 'test-foo' does not exist",
			method:   "POST",
			expected: false,
		},
		{
			name:     "Retries deletion 417",
			code:     http.StatusExpectationFailed,
			method:   "DELETE",
			expected: true,
		},
	}

	doRetry := []int{500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}
	for _, code := range doRetry {
		opts = append(opts, &option{
			name:     fmt.Sprintf("Tests code %d, should retry", code),
			code:     code,
			expected: true,
		})
	}

	// Except 408 and 429
	doNotRetry := []int{
		100, 101, 102, 103,
		200, 201, 202, 203, 204, 205, 206, 207, 208, 226,
		300, 301, 302, 303, 304, 305, 306, 307, 308,
		400, 401, 402, 403, 404, 405, 406, 407, 409, 410, 411, 412, 413, 414,
		415, 416, 417, 418, 421, 422, 423, 424, 425, 426, 428, 431, 451,
	}

	for _, code := range doNotRetry {
		opts = append(opts, &option{
			name:     fmt.Sprintf("Tests code %d, should not retry", code),
			code:     code,
			expected: false,
		})
	}

	ctx := context.Background()
	for _, opt := range opts {
		t.Run(opt.name, func(t *testing.T) {
			rsp := &http.Response{
				Status:     fmt.Sprintf("%d", opt.code),
				StatusCode: opt.code,
				Body:       io.NopCloser(strings.NewReader(opt.body)),
				Request:    &http.Request{Method: opt.method},
			}
			retry, err := checkRetry(ctx, rsp, nil)
			assert.Equal(t, opt.expected, retry)
			assert.Nil(t, err)
		})
	}
}

func TestIsServiceLagError(t *testing.T) {
	cases := []struct {
		body     string
		method   string
		expected bool
	}{
		{
			body:     "Service lol does not exist",
			method:   "POST",
			expected: true,
		},
		{
			// Invalid body
			body:     "User lol does not exist",
			method:   "POST",
			expected: false,
		},
		{
			// Invalid method
			body:     "Service lol does not exist",
			method:   "GET",
			expected: false,
		},
	}

	for i, opt := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, opt.expected, isServiceLagError(opt.method, opt.body))
		})
	}
}

func TestIsUserError(t *testing.T) {
	cases := []struct {
		body     string
		method   string
		expected bool
	}{
		{
			body:     "User avnadmin with component main not found",
			method:   "",
			expected: true,
		},
		{
			body:     "User root with component main not found",
			method:   "",
			expected: true,
		},
		{
			// Invalid user
			body:     "User lol with component main not found",
			method:   "",
			expected: false,
		},
		{
			// Invalid body
			body:     "User root with component",
			method:   "",
			expected: false,
		},
	}

	for i, opt := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, opt.expected, isUserError(opt.method, opt.body))
		})
	}
}
