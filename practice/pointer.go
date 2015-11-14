package main

import (
	"fmt"
)

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	x, y := 1, 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
