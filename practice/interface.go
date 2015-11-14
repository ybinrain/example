package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("User %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func main() {

	// u := User{2, "Hrr"}
	var u User = User{2, "Hrr"}
	var i interface{} = u
	u.id = 3
	u.name = "Jack"

	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))

	var t Printer = &User{1, "Ybin"}
	t.Print()
}
