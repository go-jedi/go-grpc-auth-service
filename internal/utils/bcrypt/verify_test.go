package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type request struct {
	hashedPassword    string
	candidatePassword string
}

func TestVerifyPassword(t *testing.T) {
	t.Parallel()
	// Arrange

	tests := []struct {
		name     string
		input    request
		expected bool
	}{
		{
			name: "OK",
			input: request{
				hashedPassword:    "$2a$14$ZuAuUPDazXu/BHhbDjX9qeMftgTLPS9FQfYRQxiZD7EZFnVDPKyXC",
				candidatePassword: "test",
			},
			expected: true,
		},
		{
			name: "Error",
			input: request{
				hashedPassword:    "686a7172686a7177313234363137616a6668616a739f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f121312",
				candidatePassword: "1234567890",
			},
			expected: false,
		},
	}
	//	 Act
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := VerifyPassword(
				test.input.hashedPassword,
				test.input.candidatePassword,
			)

			// Assert
			require.Equal(t, test.expected, result)
		})
	}
}
