package main 
import (
	"fmt"
//	"os"
//	"os/exec"
//	"log"
        "github.com/otiai10/copy"
)

func main() {
	vmname := "myvm"
dirname := "/home/Abi-dev/"+vmname
/*
	fmt.Println("Creating directory")
        if err11 := os.MkdirAll(dirname, 0777); err11 != nil {
                log.Fatal(err11)
        }
	oldDir := "/home/Abi-dev/jenkins"
    cmd := exec.Command("cp", "--recursive", oldDir, dirname)
    cmd.Run() */
    newdir := "/home/Abi-dev/ansible"
    err := Copy(newdir, dirname)
    if err != nil {
	    fmt.Println("error occured")
	}
    fmt.Println("Created directory")

 }
