package main

import "fmt"

func ranges() {
	fmt.Println("--------- range -------------")
	nums := []int{2, 3, 4}
	for i, val := range nums {
		fmt.Println("i: ", i, " => ", val)
	}

	m := map[string]string{"apple": "APLE", "Microsoft": "MSFT"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "golang" {
		fmt.Printf("%d -> %c\n", i, c)
	}
}
