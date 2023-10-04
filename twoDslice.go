//create multiD slice

package main

import "fmt"

func main() {
	x := 2
	y := 3
	s := make([][]int, x)
	fmt.Println(s)
	for i, _ := range s {
		s[i] = make([]int, y)
		fmt.Println(s[i])
	}
	fmt.Println(s)
}
