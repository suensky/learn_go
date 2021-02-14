package main

import "fmt"

func values() {
	fmt.Println("--------- values -------------")
	fmt.Println("go" + "lang")
	fmt.Println("1+1=", 1+1)
	fmt.Println("7/3 = ", 7/3)
	fmt.Println("7.0/3 = ", 7.0/3)
	fmt.Println("7/3.0 = ", 7/3.0)
}

func variables() {
	fmt.Println("--------- variables -------------")
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var intI int
	var float32I float32
	var float64I float64
	var stringI string
	var booleanI bool
	fmt.Println(intI, float32I, float64I, stringI, booleanI)

	f := "apple"
	fmt.Println(f)
}
