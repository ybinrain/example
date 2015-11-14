package main

import (
	"fmt"
)

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEventGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func sum(args []int) int {
	t := 0
	for _, v := range args {
		t += v
	}
	return t
}

func half(n int) (float64, bool) {
	b := false
	if n%2 == 0 {
		b = true
	}
	return float64(n) / 2.0, b
}

func nlist(a int, b ...int) int {
	return len(b)
}

func fib(b int) int {
	if b == 0 || b == 1 {
		return b
	}
	return fib(b-1) + fib(b-2)
}

func A() {
	panic("A")
}

func B() {
	defer func() {
		str := recover()
		fmt.Printf("%vAAA\n", str)
	}()
	A()
}

func average(args ...float64) float64 {
	var total float64 = 0
	for _, v := range args {
		total += v
	}
	return total / float64(len(args))
}

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEvent := makeEventGenerator()
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())
	fmt.Println(nextEvent())

	fmt.Println(factorial(20))

	tt := []int{10, 20, -1}
	fmt.Println(sum(tt))

	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))

	fmt.Println(nlist(1, 2, 3, 4))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	B()
	//var x1 [5]int
	x1 := [5]int{4: 100}
	x1[4] = 100
	fmt.Println(x1)

	x2 := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}
	x3 := x2[:]
	fmt.Println(average(x3...))

	x4 := make([]float64, 5, 10)
	x4[4] = 100

	fmt.Printf("%T, %p, %p\n", x2, &x2, &x2[0])
	fmt.Printf("%T, %p, %p\n", x4, x4, &x4[0])
	fmt.Println(x4)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	fmt.Printf("%p, %p, %p\n", &slice1, slice1, slice2)

	// var x5 map[string]int
	x5 := make(map[string]int)
	x5["key"] = 10
	fmt.Println(x5)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	if v, ok := elements["H"]; ok {
		fmt.Println(v)
	}

	x6 := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var min int = x6[0]
	for i := 1; i < len(x6); i++ {
		if min > x6[i] {
			min = x6[i]
		}
	}
	fmt.Println(min)
	for _, v := range x6 {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)

	//panic("PANIC")
}
