package helper

import "testing"

func TestTrimSequenceNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Standard numeric suffix removal
		{
			name:     "simple name with sequence number",
			input:    "Laptop - 5",
			expected: "Laptop",
		},
		{
			name:     "only last separator removed when multiple dashes",
			input:    "Device - Name - 3",
			expected: "Device - Name",
		},

		// No separator present
		{
			name:     "no separator returns original",
			input:    "NoSequence",
			expected: "NoSequence",
		},

		// Suffix is not a number
		{
			name:     "non-numeric suffix returns original",
			input:    "Name - Location",
			expected: "Name - Location",
		},

		// Zero is valid sequence number
		{
			name:     "zero is a valid sequence number",
			input:    "Name - 0",
			expected: "Name",
		},

		// Only last numeric suffix is removed
		{
			name:     "only last numeric suffix removed",
			input:    "Name - 123 - 456",
			expected: "Name - 123",
		},

		// Empty string
		{
			name:     "empty string returns empty string",
			input:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimSequenceNumber(tt.input)
			if result != tt.expected {
				t.Errorf("TrimSequenceNumber(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
