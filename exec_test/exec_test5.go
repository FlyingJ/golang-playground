package main

import (
    "bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func NewGovcCmd(govcPath string) (*exec.Cmd, error) {
    govcCmd := exec.Command(govcPath)

	return govcCmd
}

type GovcDriver struct {
	govcCmd 
}

func main() {
    fmt.Printf("PATH is %s\n", os.Getenv("PATH"))
    newEnvPathElements := []string{os.Getenv("PATH"), "."}
    os.Setenv("PATH", strings.Join(newEnvPathElements, ":"))
    fmt.Printf("PATH is %s\n", os.Getenv("PATH"))

	cmd := exec.Command("echo_foo.sh")

    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    cmd.Env = append(cmd.Env, "FOO=bar")

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
