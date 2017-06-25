package main

import (
        "bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "test.file")

        var stdout, stderr bytes.Buffer
        cmd.Stdout = &stdout
        cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
                fmt.Println("inside error")
                fmt.Printf("STDERR: %s\n", stderr.String())
		log.Fatal(err)
	} else {
		fmt.Println("no error encountered")
		fmt.Printf("STDOUT: %s", stdout.String())
	}
}
