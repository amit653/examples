package main

import "fmt"

type game struct {
	title string
	price money
}
type list []*game
type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.5f", m)
}

func (l list) print() { /// Add the method print to type list
	if len(l) == 0 {
		fmt.Println("lenght of game items are 0")
		return

	}
	for _, it := range l {

		fmt.Printf("%10s %s\n", it.title, it.price.string()) ///   price field has money type and money type has string method
	}

}
func main() {
	var (
		tetris    = game{"amit", 3}
		minecraft = game{"abhi", 33}
	)
	var items []*game

	fmt.Println(items, tetris)
	items = append(items, &minecraft, &tetris)
	/*
		for _, it := range items {
			fmt.Println(it.title)
		}
	*/
	my := list(items) // conversion of items type to list type
	//my = nil
	my.print() //list print() method will recieve my variable

}
