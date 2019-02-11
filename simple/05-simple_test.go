package main

import "testing"

//START OMIT
func TestFree(t *testing.T) {
	result := FreeSpace(10, 20)
	if result != 10 {
		t.Errorf("Expected %d, got %d", -10, result)
		// t.Logf(""Expected %d, got %d", 10, result")
		// t.Fail()
	}
	if t.Failed() {
		return
	}
	t.Log("we are in the end of test")
}

func TestFree_FailNow(t *testing.T) {
	result := FreeSpace(10, 20)
	if result != -10 {
		t.Fatalf("Expected %d, got %d", -10, result)
		// t.Logf(""Expected %d, got %d", 10, result")
		// t.FailNow()
	}
	t.Log("we are in the end of test")
}

//END OMIT

// go test  github.com/KitlerUA/Testing/simple -run ^TestFree$ -count=1
// go test -timeout 120s github.com/KitlerUA/Testing/simple -run ^TestFree_FailNow$ -count=1
