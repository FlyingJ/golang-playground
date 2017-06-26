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
	vminfo_json, err := ioutil.ReadFile("./jmh-test-govc.powerOn.json")
	if err != nil {
		fmt.Println(err)
	}

	var VMInfoResult VMInfo

	err = json.Unmarshal(vminfo_json, &VMInfoResult)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("json_unmarshalled is of type: %T\n", VMInfoResult)
	fmt.Printf("%v\n", VMInfoResult)

	fmt.Printf("There are %d VMs reporting info.\n", len(VMInfoResult.VirtualMachines))
	fmt.Println()
	// vmPowerState := VMInfoResult.VirtualMachines[0].Summary.Runtime.PowerState
	vmToolsStatus := VMInfoResult.VirtualMachines[0].Summary.Guest["ToolsStatus"]
	vmHostName := VMInfoResult.VirtualMachines[0].Summary.Guest["HostName"]
	vmIpAddress := VMInfoResult.VirtualMachines[0].Summary.Guest["IpAddress"]
	fmt.Println(vmToolsStatus)
	fmt.Println(vmHostName)
	if ok := IsRunning(&VMInfoResult); ok {
		fmt.Printf("%s is running\n", VMInfoResult.VirtualMachines[0].Name)
	} else {
		fmt.Printf("%s is not running\n", VMInfoResult.VirtualMachines[0].Name)
	}
	fmt.Println(vmIpAddress)
}

func IsRunning(vminfo *VMInfo) bool {
	if vminfo.VirtualMachines[0].Summary.Runtime.PowerState == "poweredOn" {
		return true
	}
	return false
}
