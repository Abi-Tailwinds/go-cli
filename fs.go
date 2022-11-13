package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.command("java", "jar", "url", "-s", "build", "jobname", "-s")
	fmt.Println(cmd)
}