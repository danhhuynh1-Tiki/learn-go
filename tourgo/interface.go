package main

import "fmt"

type Animal interface {
	SayHello() string
}

type Dog struct {
	voice string
}
type Cat struct {
	voice string
}

func (d Dog) SayHello() string {
	return "Dog : " + d.voice
}

func (c Cat) SayHello() string {
	return "Cat : " + c.voice
}
func showInfor(a Animal) {
	fmt.Println(a)
}
func main() {
	c := Cat{voice: "Meow"}
	d := Dog{voice: "Gau"}

	animals := []Animal{c, d}

	for _, v := range animals {
		// fmt.Println("Hi" + v.SayHello())
		showInfor(v)
	}

	// fmt.Println(c.SayHello())
	// fmt.Println(d.SayHello())

}
