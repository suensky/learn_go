package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	fmt.Println("--------- Closure -------------")
	nextInt := intSeq()

	// i is captured.
	for i := 0; i < 5; i++ {
		fmt.Print(nextInt(), " ")
	}

	fmt.Println()
	newNextInt := intSeq()
	fmt.Println(newNextInt())
}
