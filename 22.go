package main

import (
	"fmt"
	"math/big"
)

func task22() {
	//Используем пакет big для работы с большими числами
	a, b := big.NewInt(0), big.NewInt(0)
	//Числа
	a.SetString("1048677", 10)
	b.SetString("1048678", 10)
	//Сложение
	sum := big.NewInt(0).Add(a, b)
	//Вычитание
	diff := big.NewInt(0).Sub(a, b)
	//Умножение
	pro := big.NewInt(0).Mul(a, b)
	//Деление
	quo := big.NewInt(0).Div(a, b)

	fmt.Println("Сложение:", sum)
	fmt.Println("Вычитание:", diff)
	fmt.Println("Умножение:", pro)
	fmt.Println("Деление:", quo)
}
