package main

import "fmt"

func concat(st ...string) string {
	final := ""
	// st is just a slice of strings
	fmt.Println(st)
	for _, j := range st {
		final = final + " " + j

	}
	return final
}
func main() {
	slc := []string{"hi", "amit"}
	fmt.Println(slc)
	res := concat("hi", "amit")
	fmt.Println(res)
	//func append(slice []Type, elems ...Type) []Type  , append() is variadic
	slc = append(slc, "Hi", "Abhi")
	slc = append(slc, "Adi")
	fmt.Println(slc)
}
