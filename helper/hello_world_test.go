package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Gema",
			request:  "Gema",
			expected: "Hello Gema",
		},
		{
			name:     "Akbar",
			request:  "Akbar",
			expected: "Hello Akbar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Gema")
	if result != "Hello Gema" {
		panic("Result is not Hello Gema")
	}
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Gema")
	assert.Equal(t, "Hello Gema", result)
}
