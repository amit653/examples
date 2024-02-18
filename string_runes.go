package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	st := "hey¶"

	for i, rn := range st { /// range of string returns rune rn
		//fmt.Printf("%T", j)
		fmt.Println(i)
		fmt.Printf("%c", rn)
	}
	fmt.Println()
	///////

	//st := "hey¶"
	var size int
	slc := []byte("a§•°")
	/*
				for i, _ := range string(slc) { /// range of string returns rune j
					//fmt.Printf("%T", slc)
					//fmt.Printf("%c", slc)
					if i > 0 {
		size = i /// find the byte size of 1st rune
						break
					}
					//fmt.Println(i)
				} */
	// OR
	_, size = utf8.DecodeRune(slc) //returns 1st rune and its size in bytes
	fmt.Printf("%s %d\n", slc[:size], size)
	caps := bytes.ToUpper(slc[:size])
	fmt.Printf("%s\n", caps)
	copy(slc[:size], caps)
	fmt.Printf("%s=%[1]x\n", slc)

	/////////////// opposite word
	word := os.Args[1:][0] /// pass the word as argument

	english := []string{"good", "weak"}
	opposite := []string{"bad", "strong"}
	for i, w := range english {

		if word == w {
			fmt.Printf("%q\n", opposite[i])
			//return  will end the program and break will terminate the loop
		}

	}

	//var dct map[string]string
	dct := map[string]string{} //map literal creates an empty map
	dct["good"] = "bad"

	if _, ok := dct[word]; ok {

		fmt.Printf("dict value=%#v\n", dct[word])
		fmt.Printf("dict value=%s\n", dct[word])
	}

}
