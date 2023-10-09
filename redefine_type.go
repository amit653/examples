package main

import (
	"fmt"
)

type t1 struct {
	f1 string
}
type t2 struct { // embedding type
	t1
}
type t3 t1 //redefining type
func (t *t1) timemethod() {
	fmt.Println("inside timemethod")

}
func main() {
	var mt1 t1
	var mt2 t2
	var mt3 t3 // from redefined type
	_ = mt1.f1
	_ = mt2.f1
	_ = mt3.f1
	mt1.timemethod()
	mt2.timemethod()
	//mt3.t1method() // error becuase mt3 is used from redefined type t3

}
