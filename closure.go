package main

import "fmt"

func square() func(n int) int {
	fmt.Println("outside square")
	return func(n int) int {
		fmt.Println("inside square")
		return n * n
	}
}
func multiply(n1 int) func(n2 int) int {
	fmt.Println("outside multiple")
	return func(n2 int) int {
		fmt.Println("inside multiply")
		return n1 * n2
	}
}
func sum(f func(n int) int) {
	fmt.Println("inside sum")
	fmt.Println(f(8))

}
func main() {
	fmt.Println("square of number\n")
	sq := square()     // returns func(n int)  mapped to its alias sq
	fmt.Println(sq(3)) //sq calls return func(n int) int
	fmt.Println("multiplication of 2 numbers\n")
	mlt := multiply(2)
	fmt.Println(mlt(5))

	fn := func(n int) int {
		fmt.Println("inside main fn")
		return n + n
	}
	sum(fn) //It calls the sum function by passing the function value as an argument.Directly calls the function value by providing a value as an argument.
}
