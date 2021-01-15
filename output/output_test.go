package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePhone(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected bool
	}{
		"Show value": {"Ola!", true},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Show(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
