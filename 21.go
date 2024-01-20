package main

import "fmt"

// Создаем интерфейс Adaptee, который несовместим с существующим клиентским кодом
type Adaptee struct{}

// Создаем Адаптер, который делает интерфейс Adaptee совместимым с Target
type Adapter struct {
	adaptee *Adaptee
}

// Создаем интерфейс Target, который использует клиент
type Target interface {
	Request() string
}

func (a *Adaptee) SpecRequest() string {
	return "Конкретный запрос"
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{adaptee: adaptee}
}

func (a *Adapter) Request() string {
	return a.adaptee.SpecRequest()
}

// Создаем клиентский код, который поддерживает все классы, которые следуют Target
func ClientCode(target Target) {
	fmt.Println("Client: Я могу работать с целевыми объектами:", target.Request())
}

func task21() {
	target := NewAdapter(&Adaptee{})

	ClientCode(target)
}
