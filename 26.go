package main

import (
	"fmt"
	"strings"
)

// Фкнциия проверки на уникальность
func Uniq(str string) bool {
	Chmap := make(map[rune]bool)
	//Проверяем все сивмолы, в независимости от регистра
	for _, char := range strings.ToLower(str) {
		if _, exists := Chmap[char]; exists {
			//Если нашли повторяющийся символ
			return false
		}
		Chmap[char] = true
	}
	//Если все символы уникальны
	return true
}

func task26() {
	fmt.Println("abcd:", Uniq("abcd"))
	fmt.Println("abCdefAaf:", Uniq("abCdefAaf"))
	fmt.Println("aabcd", Uniq("aabcd"))
}
