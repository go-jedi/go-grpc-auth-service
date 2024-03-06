package bcrypt

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	t.Parallel()
	// Arrange
	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "OK",
			input:    "test",
			expected: "$2a$14$ux1dOZJriumbEIheHryf1eA6qwa0qG2j3bnjU0i2g3KUbD5QqEUr.",
			err:      nil,
		},
		{
			name:     "ERROR (NO EQUAL)",
			input:    "test",
			expected: "686a7172686a7177313234363137616a6668616a739f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f121312",
		},
	}
	//	 Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := HashPassword(test.input)

			// Assert
			switch test.name {
			case "OK":
				require.Equal(t, test.err, err)
				require.Equal(t, strings.Contains(test.expected, "$2a$14$"), strings.Contains(result, "$2a$14$"), fmt.Sprintf("Incorrect result. Expect %s, got %s",
					test.expected,
					result,
				))
			case "ERROR (NO EQUAL)":
				require.Equal(t, test.err, err)
				require.NotEqual(t, test.expected, result)
			}
		})
	}
}
