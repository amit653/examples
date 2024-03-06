package main

import "fmt"

func test(fn func() int) int {
	fmt.Println("outside main func")
	return 5
}
func main() {
	add := func(v int) int {

		return v + v
	}(7)
  fmt.Println(add)/// addition of number 7+7

	multiply := func(v int) int {
		return v * v
	}
	m := multiply(7) ////// square  of number 7
	fmt.Println(m)


	c := test(func() int {
		fmt.Println("inside main func")
		return 57

	})
	fmt.Println(c)
}
