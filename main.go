package main

import "fmt"

type newInt int

type person struct {
	Name string
	Age  int
}

type agent struct {
	person
	license string
}

func (p *person) speak() {
	fmt.Println(p.Name, p.Age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	var t int
	fmt.Printf("%T %d\n", t, t)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	m := map[string]int{}
	m["Lee"] = 34
	m["Kim"] = 45
	fmt.Println(m)

	var ni newInt
	ni = 0
	fmt.Println(ni)

	p1 := person{"John Doe", 10}
	fmt.Println(p1)
	p1.speak()

	p2 := agent{
		person{
			"James Bond",
			36,
		},
		"Fish",
	}
	p2.speak()
	saySomething(&p2)
	return
}
