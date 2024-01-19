package main

import (
	"fmt"
	"strings"
)

// Создаем функцию для реверс слов
func reverse(str string) string {
	//Разбиваем строку на слова
	word := strings.Fields(str)
	//i - 0 индекс, j - последний
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		//Делаем перестановку
		word[i], word[j] = word[j], word[i]
	}
	//Соединяем слова в строку и возвращаем строку
	return strings.Join(word, " ")
}

func task20() {
	str := "snow dog sun"

	fmt.Println("Начальная строка:", str)
	fmt.Println("Реверс:", reverse(str))
}
