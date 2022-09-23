package main

import (
	"fmt"
)

type z interface {
	func1()
	func2()
}
type a struct {
	b
}
type b struct {
}

func (a *a) func1() {
	fmt.Println("func1")
}
func (a *b) func2() {
	fmt.Println("func2")
}
func main() {
	asd := &a{}
	var y z = asd
	y.func1()
	asd.func1()
	asd.func2()

}
