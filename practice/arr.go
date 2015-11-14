package main

import (
	"fmt"
)

func test(x *[2]int) {
	fmt.Printf("%p\n", x)
	x[1] = 100
}

func main() {
	b := [...]int{1, 2, 3, 4}
	for i, j := range b {
		fmt.Println(i)
		fmt.Println(j)
	}

	c := [...]int{10: 1, 15: 5}
	fmt.Println(c)

	d := [...]struct {
		name string
		age  int8
	}{
		{"ybin", 10},
		{"hrr", 12},
	}
	fmt.Println(d)

	e := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(e)

	f := [2]int{}
	fmt.Printf("%p\n", &f)
	test(&f)
	fmt.Println(f)
	fmt.Println(len(f), cap(f))

	g := [...]int{0, 1, 2, 3, 4, 5, 6}
	s := g[1:4:5]

	s[0] += 100
	fmt.Println(g)
	fmt.Println(s)

	s1 := []int{0, 1, 2, 3, 8: 200}
	//s1 := []int{0, 1, 2, 3, 8}
	fmt.Printf("%T\t%v\n", s1, s1[8])
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8)
	fmt.Printf("%T\n", s2)
	fmt.Println(s2, len(s2), cap(s2))

	d1 := [5]struct {
		x int
	}{}
	s3 := d1[:]

	d1[1].x = 10
	s3[2].x = 20

	fmt.Println(d1, s3)
	fmt.Printf("%p\t%p\t%p\n", &d1, &s3, &s3[0])

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	ss1 := data[:3]
	ss2 := append(ss1, 100, 200)
	fmt.Println(data, ss1, ss2)

	z := make([]int, 0, 1)
	x := cap(z)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > x {
			fmt.Printf("cap: %d -> %d\n", x, n)
			x = n
		}
	}

	m := map[string]int{
		"a": 1,
	}

	if v, ok := m["a"]; ok {
		fmt.Println(v)
	}

	fmt.Println(m["c"])

	m["b"] = 2
	delete(m, "c")
	println(len(m))
	for k, v := range m {
		println(k, v)
	}

}
