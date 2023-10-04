package main

//Modifiy string values
import "fmt"

func main() {

	str := "hello"

	byte := []byte(str)
	byte[0] = 'H' // "H" will  be invalid , assignment needs a byte value
	byte[4] = 'O' 
	fmt.Println(string(byte))

}
