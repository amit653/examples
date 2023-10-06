package main

import (
	"fmt"
)

// func timeNow() (t string) {
func timeNow() {
	for i := 0; i < 2; i++ {
		defer func() { // only applicable for function , not for code block
			//t = "Current time is: " + t
			fmt.Println("inside defer", i)
		}()
		fmt.Println("outside defer", i)
		//return time.Now().Format(time.Stamp)

	}
	fmt.Println("exiting timenow func")
}

func main() {
	//fmt.Println(timeNow())
	timeNow()
}
