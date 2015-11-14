package main

import (
	"fmt"
	"math/cmplx"
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	//t      uint64     = 18446744073709551616
)

func main() {
	const fm = "%T(%v)\n"
	fmt.Printf(fm, ToBe, ToBe)
	fmt.Printf(fm, MaxInt, MaxInt)
	fmt.Printf(fm, z, z)
	//fmt.Printf(f, t, t)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	var it int = 10
	var fl float64 = float64(it)

	fmt.Printf("%v\n", fl)

	const test = 1
	fmt.Printf("%v\n", test)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}
