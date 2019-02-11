package main

import "testing"

//START OMIT
func TestFree_Table(t *testing.T) {

	testCases := []struct {
		usage    int64
		quota    int64
		expected int64
	}{
		{5, 7, 2},
		{0, 0, 0},
		{12, 7, 0},
	}

	for _, test := range testCases {
		result := FreeSpace(test.usage, test.quota)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}

//END OMIT

// go test -timeout 120s github.com/KitlerUA/Testing/simple -run ^TestFree_Table$ -count=1
