package main

import (
	"fmt"
	"strconv"
)

func slices() {
	fmt.Println("--------- slice -------------")
	a := make([]string, 3)
	fmt.Println("Initialized slices: ", a)

	i := 0
	for ; i < len(a); i++ {
		a[i] = strconv.Itoa(i)
	}
	fmt.Println("a: ", a)

	a = append(a, strconv.Itoa(i))
	i++
	a = append(a, strconv.Itoa(i), strconv.Itoa(i+1), strconv.Itoa(i+2))
	fmt.Println("a: ", a)

	s := make([]string, len(a))
	copy(s, a)

	l := s[2:5]
	fmt.Println("s[2:5]= ", l)

	l = s[:5]
	fmt.Println("s[:5]= ", l)

	l = s[2:]
	fmt.Println("s[2:] = ", l)

	t := []string{"g"}
	fmt.Println("Initialized: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i * j
		}
	}
	fmt.Println("twoD: ", twoD)
}
