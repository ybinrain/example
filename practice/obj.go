package main

import (
	"fmt"
)

type X struct{}

func (self *X) test() {
	fmt.Println("X.test")
}

type Data struct {
	x int
}

func (self Data) ValueTest() { // func ValueTest(self Data);
	self.x = 10
	fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() { // func PointerTest(self *Data);
	self.x = 20
	fmt.Printf("Pointer: %p\n", self)
}

func (self *Data) Printx() {
	fmt.Println(self.x)
}

type User struct {
	id   int
	name string
}

func (self *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", self, self)
}

type Manager struct {
	User
	title string
}

func (self *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

type Circle struct {
	x float64
	y float64
	r float64
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {

	var c Circle
	fmt.Printf("%T, %v\n", c, c)

	o := new(Circle)
	fmt.Printf("%T, %v\n", o, o)

	a := new(Android)
	a.Name = "ybin"
	a.Talk()
}
