package main

import (
	"flag"
	"testing"
)

//START OMIT
var integrations = flag.Bool("integration", false, "run integration tests")

func TestFree_Skip(t *testing.T) {

	if *integrations {
		t.Skip("not integration test")
	}

	testCases := []struct {
		name     string
		usage    int64
		quota    int64
		expected int64
	}{
		{"positive result", 5, 7, 2},
		{"all zeros", 0, 0, 0},
		{"usage>quota", 12, 7, 0},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result := FreeSpace(test.usage, test.quota)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

//END OMIT

// go test -timeout 120s github.com/KitlerUA/Testing/simple -run ^TestFree_Skip$ -count=1
