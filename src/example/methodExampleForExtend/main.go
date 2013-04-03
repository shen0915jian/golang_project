// main
package main

import (
	"fmt"
)

// method 的继承与重写
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//func (h *Human) SayHi() {
//	fmt.Printf("Hi, I am %s you call me on %s\n", h.name, h.phone)
//}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Printf("La la la la ...%s\n", lyrics)
}

//func (e *Employee) SayHi() {
//	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
//		e.company, e.phone)
//}

func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

//interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-yyyy"}, "MIT", 0.00}
	sam := Employee{Human{"Sam", 45, "111-888-xxxx"}, "Golang Inc", 1000}
	paul := Student{Human{"Paul", 26, "111-222-xxx"}, "Harvard", 100}
	tom := Employee{Human{"Sam", 36, "444-222-xxx"}, "Things Ltd.", 5000}

	var i Men

	i = mark
	fmt.Println("This is Mark,a student:")
	i.SayHi()
	i.Sing("November Rain")

	fmt.Println("Let's use a slice of Men and see what happpens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, tom
	for _, value := range x {
		value.SayHi()
	}

	//空 interface
	var a interface{}
	var s int = 5
	a = s
	fmt.Println(a)

}
