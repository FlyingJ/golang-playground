package main

import (
	"fmt"
)

func main() {
	var stringSlice []string
	stringSlice = append(stringSlice, "foo")
	stringSlice = append(stringSlice, "bar", "bazz")

	fmt.Println(stringSlice)
}
