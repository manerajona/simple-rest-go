package main

import "fmt"

type hotdog int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenceToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, `says "Good morning, Miss. Moneypenny."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says "My name is`, sa.lname+`,`, sa.fname, sa.lname, `."`)
}

type human interface {
	speak()
}

func saySommething(h human) {
	h.speak()
}

func main() {
	// Variables
	x1 := 7
	fmt.Println(x1)

	var x2 int
	fmt.Println(x2) // is zero

	var x3 hotdog
	x3 = 8
	fmt.Println(x3)

	// structs
	jona := person{
		"Jonathan",
		"Manera",
	}
	fmt.Println(jona)

	jamesBond := secretAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(jamesBond)

	// Array
	xi := []int{1, 2, 3, 45, 342342}
	fmt.Println(xi)

	// Map
	m := map[string]int{
		"Todd": 45,
		"Jona": 30,
	}
	fmt.Println(m)

	// function
	jona.speak()
	jamesBond.speak()
	jamesBond.person.speak()

	// interfaces
	saySommething(jona)
	saySommething(jamesBond)
}
