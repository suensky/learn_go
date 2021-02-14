package main

import (
	"fmt"
	"strconv"
)

func plus(a int, b int) int {
	return a + b
}

func plusAll(a, b, c int, d string) int {
	dInt, _ := strconv.Atoi(d)
	return a + b + c + dInt
}

func plusOne(a, b int) (int, int) {
	return a + 1, b + 1
}

func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func functions() {
	fmt.Println("--------- function -------------")
	var result int = plusAll(1, 2, 3, "4")
	fmt.Println("plusAll(1,2,3,\"4\") => ", result)

	a, b := plusOne(10, 11)
	fmt.Println("a= ", a, " b = ", b)

	fmt.Println("sum(1, 2) = ", sum(1, 2))
	nums := []int{1, 2, 3, 4}
	fmt.Println("sum of ", nums, " is ", sum(nums...))
}
