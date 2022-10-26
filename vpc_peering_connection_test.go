package aiven

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesteqStrPointers(t *testing.T) {
	foo := "foo"
	bar := "bar"
	cases := []struct {
		a, b     *string
		expected bool
	}{
		{
			a:        &foo,
			b:        &foo,
			expected: true,
		},
		{
			a:        nil,
			b:        nil,
			expected: true,
		},
		{
			a:        &foo,
			b:        &bar,
			expected: false,
		},
		{
			a:        &bar,
			b:        &foo,
			expected: false,
		},
		{
			a:        &foo,
			b:        nil,
			expected: false,
		},
		{
			a:        nil,
			b:        &foo,
			expected: false,
		},
	}

	for i, o := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			assert.Equal(t, eqStrPointers(o.a, o.b), o.expected)
		})
	}
}
