package main

import (
	"fmt"
	"time"
)

func forLoop() {
	fmt.Println("--------- for loop -------------")
	i := 0
	for i < 10 {
		fmt.Printf("%v ", i)
		i = i + 1
	}
	fmt.Println()

	for j := 0; j < 10; j++ {
		fmt.Printf("%d ", j)
	}
	fmt.Println()

	for {
		fmt.Println("break")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

func ifElse() {
	fmt.Println("--------- if-else -------------")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is Odd")
	}

	if n := 9; n < 0 {
		fmt.Println(n, " is negative")
	} else if n < 10 {
		fmt.Println(n, " has 1 digit")
	} else {
		fmt.Println(n, " has multiple digits")
	}
}

func switchStatement() {
	fmt.Println("--------- switch -------------")
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Print("One")
	case 2:
		fmt.Print("Two")
	case 3:
		fmt.Print("Three")
	}
	fmt.Println()

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's AM")
	default:
		fmt.Println("It's PM")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case float32, float64:
			fmt.Println("I'm a float32")
		default:
			fmt.Printf("I don't know the type: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(20)
	whatAmI(20.0)
	whatAmI("hello")
}
