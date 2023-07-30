package main

import (
	"fmt"
	"sort"
	"strings"
)

func Anagram() {
	var (
		dummyDict = []string{"cat", "from", "time", "tac", "DIANI", "mate", "MURDER", "item", "tame", "act", "form", "off", "test", "INDIA", "REDRUM"}
	)
	//fmt.Println(dummyDict)
	dct := make(map[string][]string)
	//slc := []string{}
	for _, word := range dummyDict {
		//fmt.Println(i, word)
		key := sortStr(word)
		//fmt.Println(key, dct[key]) // dct[key] returns []string
		dct[key] = append(dct[key], word)
		//1st val dct[act]=append([],cat)   -> dct[act]=[cat]
		//2nd val dct[fmor]=append([],from) -> dct[fmor]=[from]
		//3rd val dct[eimt]=append([],time) -> dct[eimt]=[time]
		//4th val dct[act]=append([cat],tac) -> dct[act]=[cat,tac]    -- appended word in existing value [cat]for dct[act]
	}
	fmt.Println("Given words in list \n", dummyDict)
	fmt.Println("\nAnagram words from the given list are \n", dct)
	for _, words := range dct {
		fmt.Println(words)
	}

}
func sortStr(k string) string {

	slc := strings.Split(k, "") //[]string
	sort.Strings(slc)
	return strings.Join(slc, "")

}
func main() {

	Anagram()

}
