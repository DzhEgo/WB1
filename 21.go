package main

import "fmt"

type Adaptee struct{}

type Adapter struct {
	adaptee *Adaptee
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

type Target interface {
	Request() string
}

func ClientCode(target Target) {
	fmt.Println("Client: Я могу работать с целевыми объектами:", target.Request())
}

func task21() {
	target := NewAdapter(&Adaptee{})

	ClientCode(target)
}
