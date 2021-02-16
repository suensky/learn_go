package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect) amplify(scale int) int {
	return r.area() * scale
}

func methods() {
	fmt.Println("--------- Method -------------")
	r := rect{width: 10, height: 20}
	fmt.Println("r's area: ", r.area())
	fmt.Println("r's perim: ", r.perim())
	fmt.Println("r's scale: ", r.amplify(10))

	rPtr := &r
	fmt.Println("r's area: ", rPtr.area())
	fmt.Println("r's perim: ", rPtr.perim())
	fmt.Println("r's scale: ", rPtr.amplify(10))

}
