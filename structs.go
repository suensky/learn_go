package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 30
	return &p
}

func structs() {
	fmt.Println("--------- Struct -------------")
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob"})
	fmt.Println(newPerson("Alan"))

	s := person{"Sean", 20}
	sPtr := &s
	sPtr.age = 30
	fmt.Println(sPtr.name, sPtr.age)
}
