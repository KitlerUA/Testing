package main

import (
	"fmt"
)

//START OMIT
func FreeSpace(usage, quota int64) int64 {
	return quota - usage
}

func main() {
	used := int64(123)
	quota := int64(1000)
	fmt.Println(FreeSpace(used, quota))
}

//END OMIT
