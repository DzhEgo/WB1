package main

import (
	"fmt"
	"time"
)

// Создаем функцию для отправки данных в канал
func send(ch chan<- int) {
	for i := 1; ; i++ {
		//Отправляем значения в канал
		ch <- i
		//Задрежка
		time.Sleep(time.Second)
	}
}

func task5() {
	//Кол-во времени работы (секунды)
	num := 3
	//Создаем канал для передачи данных
	dataChan := make(chan int)
	//Запускаем горутину для отправки
	go send(dataChan)
	//Запускаем таймер, который по истечению N секунд завершит программу
	timer := time.NewTimer(time.Duration(num) * time.Second)
	defer timer.Stop()

	//Ожидаем данные из канала или истечение таймера
	for {
		select {
		case data := <-dataChan:
			fmt.Printf("Полученные данные: %d\n", data)
		case <-timer.C:
			fmt.Println("Программа завершена!")
			return
		}
	}
}
