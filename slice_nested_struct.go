// Writing slice of Nested struct in different ways.
package main

import "fmt"

type Person struct {
	fname string
	lname string
}
type Emp struct {
	prson Person
	city  string
}

func main() {
	p1 := Person{
		fname: "abhi",
		lname: "g",
	}
	p2 := Person{
		fname: "adi",
		lname: "g",
	}
	people := []Person{}
	people = append(people, p1, p2)
	fmt.Println("slice of People struct", people)

	e1 := Emp{
		prson: Person{"abhi", "g"},
		city:  "test1",
	}
	e2 := Emp{
		prson: Person{"adi", "g"},
		city:  "test2",
	}
	// 1st method  // emp := []Emp{e1, e2} OR

	//2nd Method 
	emp := []Emp{}
	emp = append(emp, e1, e2)
	fmt.Println("slice of Emp struct", emp)

	//3rd Method //declared and defined struct during empl instance creation
	empl := []struct {
		id   int
		prs  Person
		city string
	}{
		{id: 1,
			prs:  Person{"adi", "g"},
			city: "test1",
		},
		{id: 2,
			prs:  Person{"abhi", "g"},
			city: "test2",
		},
	}

	fmt.Println(empl)
}


/*
output
go run .
slice of People struct [{abhi g} {adi g}]
slice of Emp struct [{{abhi g} test1} {{adi g} test2}]
[{1 {adi g} test1} {2 {abhi g} test2}]
*/
