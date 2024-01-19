package main

import (
	"fmt"
	"time"
)

// Создаем свою функцию, которая принимает time.Duration, который прнимает, сколько спать будем
func sleep(dur time.Duration) {
	//Используем time.After, которая возвращает канал, который отправляет текущее время после истечения указанного времени в dur
	//<- используем для ожидания сигнала от канала.
	<-time.After(dur)
}

func tsak25() {
	fmt.Println("Щас усну")
	sleep(2 * time.Second)
	fmt.Println("Поспал аж 2 секунды")
}
