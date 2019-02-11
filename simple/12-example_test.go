package main

import (
	"fmt"
)

//START OMIT
func Haha() {
	fmt.Println("hahaha")
}

func ExampleTest() {
	Haha()
	Haha()
	// Output: hahaha
	// hahaha
}

//END OMIT

// go test -timeout 120s github.com/KitlerUA/Testing/simple -run ^TestFree_Skip$ -count=1
