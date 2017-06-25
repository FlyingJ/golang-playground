package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	json_stream, err := ioutil.ReadFile("./jmh-test-govc.powerOn.json")
	if err != nil {
		fmt.Println(err)
	}

	var json_unmarshalled interface{}

	err = json.Unmarshal(json_stream, &json_unmarshalled)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", json_unmarshalled)
	fmt.Printf("%v", json_unmarshalled["VirtualMachines"])
}
