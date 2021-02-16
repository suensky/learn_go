package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(ival *int) {
	*ival = 0
}

func pointers() {
	fmt.Println("--------- Pointer -------------")
	i := 100

	zeroVal(i)
	fmt.Println("i: ", i)

	zeroPtr(&i)
	fmt.Println("i: ", i)

	fmt.Println("&i: ", &i)
}
