package main

import (
	"log"
	//"strings"
	"os/exec"
	//"text/template"
	"os"
	"bytes"
	"fmt"

)

func main() {
	keyname := "awskey4"
	g := fmt.Sprintf("%s", keyname )
	fmt.Printf("%v",g)
	/*aws := "aws"
	ec2 := "ec2"
	create := "create-key-pair"
	keyname_ := "--key-name"
	query := "--query"
	keymaterial :=  `'KeyMaterial'`
	output := "--output"
	text := "text"*/
		cmd := exec.Command("aws", "ec2", "create-key-pair", "--key-name", keyname, "--query", "KeyMaterial", "--output", "text")
        //cmd := exec.Command(aws,ec2,create,keyname_,g,query,keymaterial,output,text)
        //cmd := exec.Command("bash","-c","aws ec2 create-key-pair --key-name "%s", --query 'KeyMaterial' --output text",g)
        log.Printf("%v",cmd)
        var stderr bytes.Buffer
        cmd.Stderr = &stderr
        out, err := cmd.Output()
        if err != nil {
                log.Printf("STDERR of keypair command : %v\n", stderr.String())
        }
        log.Printf("Output of keypair command : %v\n", string(out))
	filename := keyname+".pem"
	f, err := os.Create(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    _, err2 := f.WriteString(string(out))
    if err2 != nil {
        log.Fatal(err2)
    }
	chmod := "chmod"
	num := "600"
	cmd = exec.Command(chmod,num,filename)
	log.Printf("%v",cmd)
	cmd.Stderr = &stderr
	out1, err := cmd.Output()
	if err != nil {
			log.Printf("STDERR of keypair command : %v\n", stderr.String())
	}
	log.Printf("Output of keypair chmod command : %v\n", string(out1))
	cmd3 := exec.Command("ssh-keygen", "-y", "-f", "/home/Abi-dev/go-cli/awskey4.pem")
        //cmd := exec.Command(aws,ec2,create,keyname_,g,query,keymaterial,output,text)
        //cmd := exec.Command("bash","-c","aws ec2 create-key-pair --key-name "%s", --query 'KeyMaterial' --output text",g)
        log.Printf("%v",cmd3)
       // var stderr bytes.Buffer
        cmd3.Stderr = &stderr
        out3, err3 := cmd3.Output()
        if err3 != nil {
                log.Printf("STDERR of keypair command : %v\n", stderr.String())
        }
        log.Printf("Output of keypair command : %v\n", string(out3))

}
