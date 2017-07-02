package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func sliceToMap(newSlice []string) map[string]string {
	newMap := make(map[string]string)

	for _, newString := range newSlice {
		splitString := strings.SplitN(newString, "=", 2)
		k := splitString[0]
		v := splitString[1]
		newMap[k] = v
	}

	return newMap
}

func mapToSlice(newMap map[string]string) []string {
	newSlice := make([]string, len(newMap))

	for k, v := range newMap {
		newSlice = append(newSlice, strings.Join([]string{k, v}, "="))
	}

	return newSlice
}

func NewGovcCmd(path string, config map[string]string) (*exec.Cmd, error) {
	cmd := exec.Command(path)

	envMap := sliceToMap(cmd.Env)
	envMap["GOVC_INSECURE"] = "true"

	govcUrl := fmt.Sprintf("https://%s:%s@%s/sdk", config["VC_USER"], config["VC_PASSWORD"], config["VC_HOST"])
	envMap["GOVC_URL"] = govcUrl

	cmd.Env = mapToSlice(envMap)

	return cmd, nil
}

func GovcAbout(govcPath string, govcConfig map[string]string) error {
	govcCmd, err := NewGovcCmd(govcPath, govcConfig)

	if err != nil {
		log.Fatal(err)
		return err
	}

	var stdout, stderr bytes.Buffer
	govcCmd.Stdout = &stdout
	govcCmd.Stderr = &stderr

	govcCmd.Args = append(govcCmd.Args, "about")

	err = govcCmd.Run()
	if err != nil {
		log.Println("GovcAbout: govc command run error")
		log.Printf("STDERR:\n%s\n", stderr.String())
		log.Fatal(err)
		return err
	} else {
		fmt.Println("GovcAbout: no error encountered")
		fmt.Printf("STDOUT:\n%s\n", stdout.String())
	}

	return nil
}

func main() {
	govcPath := "govc"

	govcConfig := make(map[string]string)
	govcConfigFields := []string{"VC_USER", "VC_PASSWORD", "VC_HOST"}
	for _, govcConfigField := range govcConfigFields {
		govcConfig[govcConfigField] = os.Getenv(govcConfigField)
	}

	err := GovcAbout(govcPath, govcConfig)
	if err != nil {
		log.Println("main: GovcAbout error")
		log.Fatal(err)
	}

}
