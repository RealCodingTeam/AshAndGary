package problems

import "testing"

func Test_parse(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{
			name:           "Test for 2",
			input:          2,
			expectedOutput: 2,
		},
		{
			name:           "Test for 12",
			input:          12,
			expectedOutput: 13,
		},
		{
			name:           "Test for 24",
			input:          24,
			expectedOutput: 27,
		},
		{
			name:           "Test for 1000",
			input:          1000,
			expectedOutput: 1370,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parse(test.input)
			if result != test.expectedOutput {
				t.Errorf("Expected %d got %d", test.expectedOutput, result)
			}
		})
	}
}
