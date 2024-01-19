package main

import (
	"fmt"
	"sync"
)

func task9() {
	//Создаем канал передачи чисел и удоенных чисел
	numCh := make(chan int)
	doublCh := make(chan int)
	var wg sync.WaitGroup

	//Запускаем горутину для чтения данных из масива и записи в канал чисел
	go func() {
		//Закрываем канал, после передачи чисел
		defer close(numCh)
		numbers := []int{1, 2, 3, 4, 5}
		for _, num := range numbers {
			numCh <- num
		}
	}()

	//Запускаем горутину для удвоения чисел и записи в другой канал
	go func() {
		//Закрываем канал, после передачи чисел
		defer close(doublCh)
		for num := range numCh {
			doublCh <- num * 2
		}
	}()
	//Счетчик. Ожидаем завершения обеих горутин
	wg.Add(2)
	//Запускаем горутину для вывода значений из второго канала
	go func() {
		defer wg.Done()
		for doublNum := range doublCh {
			fmt.Println(doublNum)
		}
	}()
	//Ожидаем завершения горутин
	wg.Wait()
}
