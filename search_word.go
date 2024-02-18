package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	/*
		////////////// print unique word count as  map from input file
		//str := "the quick brown fox jumps over the lazy dog"
		// words := strings.Split(str, " ")
		wd := bufio.NewScanner(os.Stdin)
		wd.Split(bufio.ScanWords)
		dct := map[string]int{}
		for wd.Scan() {

			//fmt.Println(wd.Text())
			key := strings.ToLower(wd.Text())
			fmt.Println(key)

			_, ok := dct[key]
			if ok {
				dct[key] += 1
			} else {
				dct[key] = 1
			}

		}
		fmt.Println(dct)
	*/
	/////////////////////////////////////////////////////////////////////////////////////////////////////////
	///  // search word amit in test.txt
	//1) go run string.go amit<test.txt
	//2) curl -s https://python.langchain.com/docs/modules/model_io/output_parsers/types/json|go run string.go json
	//3) go run string.go day
	//lovely day
	//day
	//file contains search word "day"
	rx := regexp.MustCompile(`[^a-z]+`) // search any character with a-z
	search := os.Args[1:][0]

	wd1 := bufio.NewScanner(os.Stdin)
	wd1.Split(bufio.ScanWords)
	word := map[string]bool{}
	for wd1.Scan() {
		k := strings.ToLower(wd1.Text())
		k = string(rx.ReplaceAllString(k, "")) // replace special chars with blank

		word[k] = true
	}
	for word := range word {
		fmt.Println(word)

	}
	fmt.Println(search) // search word
	if word[search] {
		fmt.Printf("file contains search word %q\n", search)
		return
	}
	fmt.Printf("file does not contains search word %q\n", search)
}

/*
	//// no of lines in file
	//os.Stdin.Close()

	scan := bufio.NewScanner(os.Stdin)
	var lines int
	if err := scan.Err(); err != nil {
		fmt.Println("error")
	}
	for scan.Scan() {
		lines++
		fmt.Println(scan.Text())
	}
	fmt.Println("no of lines in input file", lines)
*/
