package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestProcessMessage(t *testing.T) {
	tests := []struct {
		messages []string
		expect   []string
	}{
		{
			messages: []string{"Sunlit", "Man"},
			expect:   []string{"Man-processed", "Sunlit-processed"},
		},
		{
			messages: []string{"Nomad do you copy?"},
			expect:   []string{"Nomad do you copy?-processed"},
		},
		{
			messages: []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"},
			expect:   []string{"Taldain-processed", "Roshar-processed", "Sel-processed", "Nalthis-processed", "Scadriel-processed"},
		},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{},
				expect:   []string{},
			},
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{"Scadriel"},
				expect:   []string{"Scadriel-processed"},
			},
		)
	}

	for _, tc := range tests {
		fail := false
		result := processMessages(tc.messages)

		if len(result) != len(tc.expect) {
			fail = true
		}
		if slices.Equal(result, tc.expect) {
			fail = true
		}
		if fail {
			t.Errorf("\n========\nFAIL\nInputs: %v\nExpected: %v\nActual: %v\n========\n", tc.messages, tc.expect, result)
		}
		fmt.Printf("\n========\nPASS\nInputs: %v\nExpected: %v\nActual: %v\n========\n", tc.messages, tc.expect, result)
	}
}

var withSubmit = true
