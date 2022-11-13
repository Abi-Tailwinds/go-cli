package main
 
import (
    "fmt"
    "strings"
)
 
// Main function
func main() {
 
    // Creating and initializing the strings
    str1 := "Welcome// to the, online portal, of GeeksforGeeks"
    fmt.Println("String 1: ", str1)
    var a [3]string
    var str1, str2 string
    // Splitting the given strings
    // Using Split() function
    res1 := strings.SplitAfter(str1, "//")
    fmt.Println("Result 1: ", res1)
    for i, k := range res1 {
	    if i==0 {
	       fmt.Println(i, k)
	       str1 = k
	      }
	      else {
		fmt.Println(i, k)
		str2 = k
	      }

     }

}
