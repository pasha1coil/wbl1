package main

import "fmt"

type Human struct{}

// метод структуры Human
func (c *Human) Print() {
	fmt.Println("Human")
}

type Action struct {
	Human // встраиваем структуру human в action, action наследует метод human
}

// метод структуры Action
func (p *Action) Print() {
	fmt.Println("Action")
}

func main() {
	var x Action //определяем тип переменной x как Action

	x.Print()       //вызываем метод Action
	x.Human.Print() //вызываем метод Human
}
