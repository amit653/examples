// declare and pass function as variables

package main
import "fmt"

func domath(passfun func(int, int) (string, int)) {
	result := passfun
	fmt.Println(result(7, 8))
	//OR
	//result := passfun(2, 5)
	//fmt.Println(result)

	// OR
	//fmt.Println(passfun(7, 8))

}

func add(a, b int) (string, int) {

	return "Add", a + b

}
func multiply(a, b int) (string, int) {

	return "multiply func", a * b

}

func main() {

	domath(add)
	domath(multiply)
	fmt.Println("-----------")
	//OR
	var passfun func(int, int) (string, int) // declare variable of type func, similar declaration as done in domath function
	passfun = add
	fmt.Println(passfun(7, 8))
	passfun = multiply
	fmt.Println(passfun(7, 8))

	fmt.Println(">>>>>>>>>>>>>>>>")
	slc_fun := []func(int, int) (string, int){}
	slc_fun = append(slc_fun, add, multiply)
	for _, math_fun := range slc_fun {
		fmt.Println(math_fun(7, 8))

	}
}
