package main

import "fmt"

// Создаем структуру Human
type Human struct {
	Name string
	Age  int
}

// Создаем метод MyName() стурктуры Human
func (h Human) MyName() {
	fmt.Println("Меня зовут", h.Name)
}

// Создаем структуру Action, которая встраивает Human
type Action struct {
	Human
	Action string
}

// Создаем новый экземпляр Action
func NewAction(name string, age int, action string) Action {
	return Action{Human: Human{name, age}, Action: action}
}

// Создаем метод Now() структуры Action
func (a Action) Now() {
	fmt.Println(a.Name, "cейчас", a.Action)
}

func task1() {
	fmt.Println("Начало работы")
	//Создаем экземпляр Action
	act := NewAction("Петя", 25, "идет домой")
	//Вызываем метод MyName(), унаследованный от Human
	act.MyName()
	//Вызываем метод Now(), опредленный в Action
	act.Now()
}
