package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val
		}
		return items
	}
	environment := getenvironment(os.Environ(), func(item string) (key, val string) {
		splits := strings.SplitN(item, "=", 2)
		key = splits[0]
		val = splits[1]
		return
	})
	fmt.Println("USER environment variable has value of", environment["USER"])
}
