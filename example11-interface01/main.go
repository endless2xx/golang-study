package main

import "fmt"

type Car interface {
	GetName() string
	Run()
	Didi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) Didi() {
	fmt.Println("Didi")
}

func main() {
	var car Car
	fmt.Println(car)

	var bmw BMW

	car = &bmw
	fmt.Println(car)
	bmw.Name = "宝马"
	car.Run()
}
