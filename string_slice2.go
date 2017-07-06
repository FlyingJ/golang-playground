package main

import (
	"fmt"
)

func main() {
	datastore := "FORP"
	vmname := "freep"

	var stringSlice []string
	stringSlice = append(stringSlice, "vm.create.disk")
	stringSlice = append(stringSlice, "-ds", datastore)
	stringSlice = append(stringSlice, "-vm", vmname)

	fmt.Println(stringSlice)
}
