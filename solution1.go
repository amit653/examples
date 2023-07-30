package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		//fmt.Println(d, len(d))
		//fmt.Println(b, reflect.TypeOf(b), reflect.TypeOf(b).Kind())
			str_slc := []string{str}  // will convert string to slice of length 1
			fmt.Println(str_slc, reflect.TypeOf(str_slc), reflect.TypeOf(str_slc).Kind())
			//c := strings.Join(str_slc, "-")
			fmt.Println(d)
	*/
	str := "reverse"
	str_slc := strings.Split(str, "") // returns []string
	// OR  problem can be string or below slice of characters
	//str_slc := []string{"r", "e", "v", "e", "r", "s", "e"} // count elements in slice
	fmt.Println("\n Problem 1 - count number of elements in slice and print unique elements  ", str_slc, len(str_slc))
	dc := make(map[string]int)
	empty_slc := []string{}
	for _, key := range str_slc {
		//fmt.Println(key)
		_, ok := dc[key]
		if ok {
			dc[key] += 1
		} else {
			dc[key] = 1
			empty_slc = append(empty_slc, key) // extra code to remove duplicates and enter only unique elements in empty slice
		}

	}
	fmt.Println("count of elements in slice are ", dc)
	fmt.Println("Unique elements in slice are ", empty_slc)
	fmt.Println("-------------------------------------------------\n")

	fmt.Println("Problem 2- Reverse the  slice of string ", str_slc)
	//// Reverse a string using empty slice
	slc := []string{}
	//read from last and populate in new slice
	for j := len(str_slc) - 1; j >= 0; j-- {

		slc = append(slc, str_slc[j])

	}
	fmt.Println(slc, "Reverse of string ->", strings.Join(slc, ""))
}
