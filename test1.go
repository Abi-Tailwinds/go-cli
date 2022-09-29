package main

import (

	"bytes"
	//"strings"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"os/exec"
)
var result string
var vpcname, group, imageid, instancetype, cmd1, keyname = "vpc-9df038e0", "any-grp", "ami-f0e7d19a", "t2.micro", "sg-0a94428cbd8cc8bfa", "awskey1"
func createAWSVM() {
	aws := "aws"
	ec2 := "ec2"
	authorize := "authorize-security-group-ingress"
	groupid := "--group-id"
	protocol := "--protocol"
	tcp := "tcp"
	port := "--port"
	port1 := "3389"
	port2 := "22"
	cidr := "--cidr"
	cidr1 := "0.0.0.0/24"
	createtags := "create-tags"
	resources := "--resources"
        tags := "--tags"
	keyn := "Key=Mykey,Value=MyClivm"
	//createsecuritygroup := "create-security-group"
	//vpcname_ := "--vpc-name"
	//group_ := "--group-name"
	//description := "awsvm"
	runinstances := "run-instances"
	instancetype_ := "--instance-type"
	imageid_ := "--image-id"
	securitygroupid := "--security-group-id"
	//createkeypair := "create-key-pair"
	keyname_ := "--key-name"
	//query := "--query"
	//output := "--output"
	//string1 := "'KeyMaterial'"
	//text := "text"
	fmt.Println("Creating directory")
	if err := os.MkdirAll("/home/Abi-dev/program/.md/vm", 0777); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created directory")
	fmt.Println("Creating key-pair values..")
	cmd := exec.Command("bash","-c","aws ec2 create-key-pair --key-name awskey1 --query 'KeyMaterial' --output text")
	log.Printf("%v",cmd)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		log.Printf("STDERR of install  command : %v\n", stderr.String())
	}
	log.Printf("Output of install command : %v\n", string(out))
	f, err := os.Create("/home/Abi-dev/program/.md/vm/llss.pem")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    _, err2 := f.WriteString(string(out))

    if err2 != nil {
        log.Fatal(err2)
    }

    fmt.Println("done")
	fmt.Println("Started creating sg")
	cmd1 := exec.Command("bash","-c","aws ec2 create-security-group --group-name isec1 --description 'mg' --vpc-id vpc-9df038e0")
	cmd1.Stderr = &stderr
	outp, err := cmd1.Output()
	if err != nil {
		log.Printf("STDERR of sg command : %v\n", stderr.String())
	}
	myvar := string(outp)
        var m map[string]interface{}
	err = json.Unmarshal(outp, &m)
	      var s = m["GroupId"].(string)
	   fmt.Println("%s",s)
           fmt.Println("Decoded:", m)
	log.Printf("Output of sg command : %s\n", myvar)
     fmt.Println("Created security grp")
     fmt.Println("Assigning port access for the security grp created ")
     cmd3 := exec.Command(aws,ec2,authorize,groupid,s,protocol,tcp,port,port1,cidr,cidr1)
     log.Print("command: >> %v\n",cmd3)
     cmd3.Stderr = &stderr
     output1, err := cmd3.Output()
     if err != nil {
	     log.Printf("Stderr of sg authorize command: %v\n",stderr.String())
	}
	fmt.Println("Op is %v\n",string(output1))
      fmt.Println("Authorize port 22 access")
      cmd4 := exec.Command(aws,ec2,authorize,groupid,s,protocol,tcp,port,port2,cidr,cidr1)
      log.Print("command: >> %v\n",cmd4)
     cmd4.Stderr = &stderr
     output2, err := cmd4.Output()
     if err != nil {
             log.Printf("Stderr of sg authorize command: %v\n",stderr.String())
        }
        fmt.Println("Op is %v\n",string(output2))
     cmd2 := exec.Command(aws,ec2,runinstances,imageid_,imageid,instancetype_,instancetype,keyname_,"awskey1",securitygroupid,s)
     log.Printf("install vm :CMD>>:%v\n", cmd2)
	cmd2.Stderr = &stderr
	outpu, err := cmd2.Output()
	if err != nil {
		log.Printf("STDERR of vm command : %v\n", stderr.String())
	}
        var m1 map[string]interface{}
        err = json.Unmarshal(outpu, &m1)
         var s1 = m1["Instances"]
           fmt.Println("%s",s1)
	   result = fmt.Sprintf("%s", m1["Instances"].([]interface{})[0].(map[string]interface{})["InstanceId"])
           fmt.Println("Decoded:", result)
	log.Printf("Output of vm install command : %v\n", string(outpu))
	fmt.Println("Giving name for vm")
	cmd6 := exec.Command(aws,ec2,createtags,resources,result,tags,keyn)
	log.Printf("install vm :CMD>>:%v\n", cmd6)
        cmd6.Stderr = &stderr
        output3, err := cmd6.Output()
        if err != nil {
                log.Printf("STDERR of vm command : %v\n", stderr.String())
        }
	fmt.Println("tags created output",output3)

}
func deleteAWSVM(result string) {
	aws := "aws"
	ec2 := "ec2"
	delete := "delete-instances"
	instanceid := "--instance-id"
	cmd5 := exec.Command(aws,ec2,delete,instanceid,result)
	log.Printf("install vm :CMD>>:%v\n", cmd5)
        cmd5.Stderr = &stderr
        output7, err := cmd5.Output()
        if err != nil {
                log.Printf("STDERR of vm command : %v\n", stderr.String())
        }

}
func main() {
	createAWSVM()
	deleteAWSVM(result)
}
