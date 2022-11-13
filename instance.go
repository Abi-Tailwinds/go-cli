package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"encoding/json"
	"bytes"
)

func main() {
	cmd := exec.Command("bash", "-c", "aws ec2 describe-instance-status --instance-id i-06b1ab9249f49f036")
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
