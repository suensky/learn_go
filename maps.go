package main

import "fmt"

func maps() {
	fmt.Println("--------- map -------------")
	m := make(map[string]int)
	fmt.Println("m: ", m)

	m["k1"] = 1
	m["k2"] = 2
	fmt.Println("m: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len(m): ", len(m))

	delete(m, "k2")
	fmt.Println("m: ", m)

	val, present := m["k2"]
	fmt.Println("val: ", val, "\npresent: ", present)

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("m: ", m)
}
