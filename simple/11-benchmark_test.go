package main

import (
	"log"
	"os"
	"testing"
)

//START OMIT

var f, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0775)
var l = log.New(f, "", 0)

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l.Println(".")
	}
}

//END OMIT

// go test -benchmem -run=^$ github.com/KitlerUA/Testing/simple -bench ^BenchmarkPrint$ -count=1
