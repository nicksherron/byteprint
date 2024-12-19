package main

import (
	"fmt"
	"testing"
)

func TestGetByteArrayFromArg(t *testing.T) {
	var tests = []struct {
		name     string
		input    []string
		expected string
		wantErr  bool
	}{
		{
			name:     "basic conversion",
			input:    []string{"102, 111, 111, 109, 97, 110, 99, 104, 117"},
			expected: "foomanchu",
			wantErr:  false,
		},
		{
			name:     "longer string",
			input:    []string{"102, 111, 111, 109, 97, 110, 99, 104, 117, 32, 105, 115, 32, 97, 32, 103, 111, 111, 100, 32, 98, 111, 121"},
			expected: "foomanchu is a good boy",
			wantErr:  false,
		},
		{
			name:     "invalid input",
			input:    []string{"invalid"},
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := convertByteArrayFromArg(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertByteArrayFromArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && output != tt.expected {
				t.Errorf("convertByteArrayFromArg() = %v, want %v", output, tt.expected)
			}
		})
	}
}

func ExampleconvertByteArrayFromArg() {
	result, err := convertByteArrayFromArg([]string{"102, 111, 111, 109, 97, 110, 99, 104, 117"})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(result)
	// Output: foomanchu
}
