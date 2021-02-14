package main

import "fmt"

func arrays() {
	fmt.Println("--------- switch -------------")
	var a [5]int
	fmt.Println("emp: ", a)

	a[3] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[3])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare: ", b)

	c := [5]int{1, 2, 3}
	fmt.Println("partially declared: ", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("2d: ", twoD)

	fmt.Println("rows: ", len(twoD))
	fmt.Println("cols: ", len(twoD[0]))
}
