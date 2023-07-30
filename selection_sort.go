package main

import (
	"fmt"
)

func main() {
//Selection sort in golang
	//arr_slice := []string{"S", "O", "R", "T", "E", "X"}
	//fmt.Println(strings.Join(arr_slice, ""))
	//varr := "SORTEX" //"1àha漢字Pépy5" // "SORT"
	//arr := []byte(varr)

	arr := []int{20, 23, 12, 22, 11}
	//arr := []int{64, 25, 12, 22, 11}
	size := len(arr)

	fmt.Println(arr)

	for i := 0; i < size-1; i++ {
		min_max := i // assuming min/max elelement is at i position
		for j := i + 1; j < size; j++ {
			//fmt.Printf("for  i-%dth position scanning linearly to find min element position\n ", i)

			if arr[j] < arr[min_max] { // for Descending  if arr[j] > arr[min_max]
				min_max = j // // assign min elelement to j position
			}

			// swap
			arr[i], arr[min_max] = arr[min_max], arr[i]
		}

	}
	fmt.Println(arr)
}
