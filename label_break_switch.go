package main

import (
	"fmt"
	"time"
)

func main() {

loopi: //Labeled break and continue
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(x, y)
			//break loopi
			continue loopi
		}
	}
// //
loop:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("timeout reached")
			break loop // will terminate from for loop
			//break // will terminate from select
		}
	}
	fmt.Println("the end")
////
myswitch:
	switch {
	case true:
		for {
			fmt.Println("switch")
			break myswitch 

		}
	}
	fmt.Println("switch end")
}
