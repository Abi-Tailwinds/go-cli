package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func vmcreate() (string, string) {
	cmd := exec.Command("pulumi", "up", "--cwd", "/home/Abi-dev/pulumi/create-delete-aws-vm", "--config=vmname=awsclivm", "--config=subnetid=subnet-9cf64cbd", "--config=vpcid=vpc-9df038e0", "--config=keyname=awskey4", "-y")
	fmt.Println(cmd)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Printf("STDERR of keypair command : %v\n", stderr.String())
	}
	log.Printf("Output of keypair command : %v\n", string(out))
	var s = string(out)
	var str1 = "awsclivm" + ":ElasticIP:"
	var IP string
	var ID string
	v := strings.Fields(s)
	fmt.Println(v)
	for i := 0; i < len(v); i++ {
		match1, err1 := regexp.MatchString(str1, v[i])
		if err1 != nil {
			log.Println(err1)
		}
		if match1 {
			IP = v[i+1]
			log.Printf("%s", v[i+1])
			log.Printf("%s", IP)
			ID = v[i+3]
			log.Printf("%s", v[i+3])
			log.Printf("%s", ID)
		}

	}
	if IP[0] == '"' {
		IP = IP[1:]
	}
	if i := len(IP) - 1; IP[i] == '"' {
		IP = IP[:i]
	}
	log.Println(IP)
	if ID[0] == '"' {
		ID = ID[1:]
	}
	if i := len(ID) - 1; ID[i] == '"' {
		ID = ID[:i]
	}
	log.Println(ID)
	return IP, ID
}
func main() {
	var ID string
	_, ID = vmcreate()
	cmd := exec.Command("aws", "ec2", "describe-instance-status", "--instance-id", ID)
	fmt.Println(cmd)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Printf("STDERR of keypair command : %v\n", stderr.String())
	}
	log.Printf("Output of keypair command : %v\n", string(out))
	var m map[string]interface{}
	err = json.Unmarshal(out, &m)
	//	result := m["InstanceStatuses"].(map[string]interface{})["SystemStatus"]
	result := string(out)
	fmt.Println(result)
	var str1 = "Status"
	var IP string
	v := strings.Fields(result)
	fmt.Println(v)
	for i := 0; i < len(v); i++ {
		match1, err1 := regexp.MatchString(str1, v[i])
		if err1 != nil {
			log.Println(err)
		}
		if match1 {
			IP = v[i]
			if IP == "passed" {
				log.Println("passed")
				break
			}
		}
	}

}
