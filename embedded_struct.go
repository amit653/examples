
package main

import "fmt"

type car struct {
	make  string
	model string
}
type truck struct {
	car
	bedsize int
}

func main() {
	fmt.Println("Hello, 世界")
	lane := truck{
		bedsize: 10,
		car: car{
			make:  "wr",
			model: "SDD",
		},
	}

	fmt.Println(lane.bedsize)
	fmt.Println(lane.make, lane.car.make) //// embedded fields promoted to the top-level  instead of lane.car.make
}
