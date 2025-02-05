package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEncoded(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []int
	}{
		{"Case 1: LLRR=", "LLRR=", []int{2, 1, 0, 1, 2, 2}},
		{"Case 2: ==RLL", "==RLL", []int{0, 0, 0, 2, 1, 0}},
		{"Case 3: =LLRR", "=LLRR", []int{2, 2, 1, 0, 1, 2}},
		{"Case 4: RRL=R", "RRL=R", []int{0, 1, 2, 0, 0, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := encoded(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("\n---\nTest Failed: %s\nInput:    %v\nExpected: %v\nGot:      %v\n---",
					tc.name, tc.input, tc.expected, result)
			} else {
				fmt.Printf("\nTest Passed: %s\nInput:    %v\nOutput:   %v\n", tc.name, tc.input, result)
			}
		})
	}
}
