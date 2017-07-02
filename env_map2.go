package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("USER environment variable has value of", os.Getenv("USER"))
}
