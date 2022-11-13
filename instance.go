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
	cmd1 := exec.Command("bash","-c","pulumi login --local")
        log.Printf("%v",cmd1)
        var stderr bytes.Buffer
        cmd1.Stderr = &stderr
        out1, err1 := cmd1.Output()
        if err1 != nil {
                log.Printf("STDERR of pulumi login command : %v\n", stderr.String())
        }
        log.Printf("Output of pulumi login command : %v\n", string(out1))
		cmd5 := exec.Command("pulumi","stack","init","dev1","--cwd","/home/Abi-dev/pulumi/create-delete-aws-vm")
	log.Printf("%v", cmd5)
	cmd5.Stderr = &stderr
	out5, err5 := cmd5.Output()
	if err5 != nil {
		log.Printf("STEDERR OF pulumi stack init command: %v/n", stderr.String())
	}
	log.Printf("%v", string(out5))
	cmd2 := exec.Command("pulumi","stack","select","dev1","--cwd","/home/Abi-dev/pulumi/create-delete-aws-vm")
        log.Printf("%v",cmd2)
        //var stderr bytes.Buffer
        cmd2.Stderr = &stderr
        out2, err2 := cmd2.Output()
        if err2 != nil {
                log.Printf("STDERR of pulumi stack select command : %v\n", stderr.String())
        }
        log.Printf("Output of pulumi stack select command : %v\n", string(out2))
		cmd6 := exec.Command("pulumi","plugin","install","resource","aws","5.13.0","--cwd","/home/Abi-dev/pulumi/create-delete-aws-vm")
        log.Printf("%v",cmd6)
        cmd6.Stderr = &stderr
        out6, err6 := cmd6.Output()
        if err6 != nil {
                log.Printf("STDERR of pulumi plugin  select command : %v\n", stderr.String())
        }
        log.Printf("Output of pulumi plugin  select command : %v\n", string(out6))
	cmd7 := exec.Command("pulumi","config","set","aws:region","us-east-1","--cwd","/home/Abi-dev/pulumi/create-delete-aws-vm")
        log.Printf("%v",cmd7)
        cmd6.Stderr = &stderr
        out7, err7 := cmd7.Output()
        if err7 != nil {
                log.Printf("STDERR of pulumi plugin  select command : %v\n", stderr.String())
        }
        log.Printf("Output of pulumi plugin  select command : %v\n", string(out7)) 
	cmd := exec.Command("pulumi", "up", "--cwd", "/home/Abi-dev/pulumi/create-delete-aws-vm", "--config=vmname=awsclivm", "--config=subnetid=subnet-9cf64cbd", "--config=vpcid=vpc-9df038e0", "--config=keyname=awskey4", "-y")
	fmt.Println(cmd)
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
