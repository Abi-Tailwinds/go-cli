package main
import (
	"fmt"
)
func main(){
//    multiline := "line" + "by line"\n + "and line"\n +"after line"
	str1 := "abi"
	str2 := "naya"
	multiline := "\n"+ str1
	multiline = "\n"+ str2
    fmt.Print(multiline) // New lines as interpreted \n
}
