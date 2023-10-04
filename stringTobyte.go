package main

//Modifiy string values
import "fmt"

func main() {

	str := "hello"

	byte := []byte(str)
	byte[0] = 'H' // "H" will  be invalid , assignment needs a byte value
	byte[4] = 'O'
	fmt.Println(string(byte))

	for _, r := range str { // range iterates over rune in string
		fmt.Print(string(r))
		fmt.Println()
	}

//convert a string to rune slice to access by index
	r := []rune(str)
	for _, i := range r {
		fmt.Print(string(i))
		fmt.Println()
	}

}
