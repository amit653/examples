package main

import (
	"fmt"
)

func main() {

	slc := []string{"cat", "dog", "bat", "dog", "cat", "dog", "cat", "dog"} //o/p-> [[cat:[2]]  [bat:[1]]  ]
	fmt.Println(slc)
	dct1 := make(map[string][]int) // Populate slice []int vas cat:[1 2] and then replace 1 with 2
	dct := make(map[string]int) //Populate empty map with values
	
	for _, key := range slc {
		fmt.Println(key)
		if _, ok := dct[key]; ok {
			dct[key]++
			if dct[key] > 1 { // if [cat:[1] int already present then remove it and update it with 2 at below append ]  cat:[2]]
				//replace last inserted element from slice using  dct1[cat]=[1]  remove it []
				dct1[key] = dct1[key][1:1]
				
			}
			dct1[key] = append(dct1[key], dct[key])
			
		} else {
			dct[key] = 1
			
			dct1[key] = append(dct1[key], dct[key])
		}

	}
	fmt.Println(dct1)

	/*
		//dc := []string{"cat", "time", "cat", "time"} // o/p [cat:[cat]]
		//dc := map[int]int{10: 2, 20: 4}
		//https://www.dotnetperls.com/convert-map-slice-go
		dc := map[int]int{}
		fmt.Println(dc)
		fruit := map[string]float64{
			"pears":  1.75,
			"cherry": 3.22,
			"apple":  1.89,
		}go run ma
		//Ordering map keys
		var fruits []string
		for val := range fruit {
			fruits = append(fruits, val)
			//fmt.Println(dish)
		}
		sort.Strings(fruits)
		fmt.Println("Arrange dict in  alphabetical order:")
		//fmt.Println(fruits)
		for _, key := range fruits {
			fmt.Println(key, fruit[key])
		}
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		// // Convert map to slice of key-value pairs.
		m := map[string]string{
			"java": "coffee",
			"go":   "verb",
			"ruby": "gemstone",
		}
		
		val := [][]string{} //slice of 2-element/nested slice string arrays
		for k, v := range m {
			//fmt.Println([]string{k,v})
			val = append(val, []string{k,v}) // [ [],[]...]
		}
		fmt.Println(val, "-", val[0])
	*/
}
