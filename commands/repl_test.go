package commands

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}
	testCases := []testCase{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HELLO  WORLD  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, tc := range testCases {
		result := cleanInput(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("Expected length %d, got %d", len(tc.expected), len(result))
			continue
		}
		for i := range result {
			if result[i] != tc.expected[i] {
				t.Errorf("Expected '%s', got '%s'", tc.expected[i], result[i])
			}
		}
	}
}
