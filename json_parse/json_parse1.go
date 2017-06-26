package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type SimpleObject struct {
	Type string
	Value string
}

type SelfObject   SimpleObject
type ParentObject SimpleObject
type TaskObject   SimpleObject
type EntityObject SimpleObject
type AlarmObject  SimpleObject

type DeclaredAlarm struct {
	Key                string
	Entity             EntityObject
	Alarm              AlarmObject
	OverallStatus      string
	Time               string
	Acknowledged       bool
	AcknowledgedByUser string
	AcknowledgedTime   string
	EventKey           uint
}

type RuntimeObject struct {
	ConnectionState string
	PowerState      string
}

type SummaryObject struct {
	Vm SimpleObject
	Runtime RuntimeObject
	Guest map[string]string
}

type VirtualMachine struct {
	Self               SelfObject
	Value              string
	AvailableField     string
	Parent             ParentObject
	CustomValue        string
	OverallStatus      string
	ConfigStatus       string
	ConfigIssue        string
	EffectiveRole      []int
	Permission         string
	Name               string
	DisabledMethod     []string
	RecentTask         []TaskObject
	DeclaredAlarmState []DeclaredAlarm
	Capability         map[string]bool
	Summary            SummaryObject
}

type VMInfo struct {
	VirtualMachines []VirtualMachine
}

func main() {
	json_stream, err := ioutil.ReadFile("./jmh-test-govc.powerOn.json")
	if err != nil {
		fmt.Println(err)
	}

	var json_unmarshalled VMInfo

	err = json.Unmarshal(json_stream, &json_unmarshalled)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("json_unmarshalled is of type: %T\n", json_unmarshalled)
	fmt.Printf("%v\n", json_unmarshalled)

	fmt.Printf("There are %d VMs reporting info.\n", len(json_unmarshalled.VirtualMachines))
	fmt.Println()
	fmt.Println(json_unmarshalled.VirtualMachines[0].Summary.Runtime.PowerState)
}
