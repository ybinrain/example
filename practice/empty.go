package main

import (
	"fmt"
)

type Tester struct {
	s interface {
		String() string
	}
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func Print(v interface{}) {
	fmt.Printf("%T: %v\n", v, v)
}

func main() {
	t := Tester{&User{1, "Ybin"}}
	fmt.Println(t.s.String())
}
