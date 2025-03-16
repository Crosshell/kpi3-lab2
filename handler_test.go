package lab2

import (
	"bytes"
	"testing"
)

func TestCompute(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:     "valid input",
			input:    "+ 1 2",  
			expected: "3.000000\n",
			expectErr: false,
		},
		{
			name:      "invalid input",
			input:     "+ 1",  
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := bytes.NewBufferString(tt.input)
			output := new(bytes.Buffer)
			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if output.String() != tt.expected {
				t.Errorf("expected output: %v, got: %v", tt.expected, output.String())
			}
		})
	}
}
