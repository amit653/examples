package main

import "fmt"

func main() {
	var f1, f2 func()
	var data []int
	for i := 0; i < 4; i++ {
		// closure that captures 'data' from the enclosing scope
		fmt.Println("insert", i)
		closure := func() {
			fmt.Println("--Append an element to 'data'")
			data = append(data, i)
		}
		fmt.Println("-Invoke the closure")
		closure()
	}
	fmt.Println(data)
	f1 = func() {
		fmt.Println("Function f1 is called")
	}
	f2 = func() {
		fmt.Println("Function f2 is called")
		f1()
	}
	// Invoking f1
	f2()

	/////////// stateless anonymous func which returns closure func
	anonymous := func() {
		fmt.Println("anonymous Function ")
	}
	counter := func() func() int {
		fmt.Println("Outside")
		cnt := 0
		return func() int {
			fmt.Println("Inside")
			cnt++
			return cnt
		}
	}() // prints Println("Outside")
	anonymous()
	fmt.Println(counter())
	fmt.Println(counter())

}
