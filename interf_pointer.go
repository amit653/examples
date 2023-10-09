package main

import (
	"fmt"
)

type Isayhi interface {
	Say()
}
type Sayhi struct{}

func (s *Sayhi) Say() {
	fmt.Println("Hi!")
}
func main() {
	var i Isayhi
	fmt.Println(i == nil)

	var sayerImplementation *Sayhi
	i = sayerImplementation
	fmt.Println(i)
	i.Say()

	//i = &Sayhi{}
	//fmt.Println(i)
	//i.Say()

}
