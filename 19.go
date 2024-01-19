package main

import (
	"fmt"
)

// Создаем функцию для реверса строки
func reverseLet(str string) string {
	//Преобразуем строки в слайс рун
	let := []rune(str)
	//i - 0 индекс, j - последний
	for i, j := 0, len(let)-1; i < j; i, j = i+1, j-1 {
		//Делаем перестановку
		let[i], let[j] = let[j], let[i]
	}
	//Преобразуем обратно в строку
	return string(let)
}

func task19() {
	str := "главрыба"

	fmt.Println("Начальное слово:", str)
	fmt.Println("Реверс:", reverseLet(str))
}
