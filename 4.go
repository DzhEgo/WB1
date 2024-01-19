package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Создаем воркер для чтения из канала и вывода
func worker(id int, wg *sync.WaitGroup, ch <-chan string) {
	defer wg.Done()

	for data := range ch {
		fmt.Printf("Воркер %d: %s\n", id, data)
	}
}

func task4() {
	//Кол-во воркеров
	num := 3
	//Создание канала для передачи данных от главного потока
	dataChan := make(chan string)
	//С помощью WaitGroup ожидаем завершения всех воркеров
	var wg sync.WaitGroup

	//Запускаем воркеры
	for i := 1; i <= num; i++ {
		wg.Add(1)
		go worker(i, &wg, dataChan)
	}

	//Создаем горутину для постоянной записи в канал
	go func() {
		count := 1
		for {
			dataChan <- fmt.Sprintf("Данные %d", count)
			count++
		}
	}()

	//Создаем канал, который будет отслеживать сигналы для завершения программы
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	//Ожидаем сигнал завершения
	select {
	case <-stopCh:
		//Закрываем канал для завершения работы воркеров
		close(dataChan)
		//Ожидаем завершения всех воркеров
		wg.Wait()
		fmt.Println("Программа завершена!")

	}
}
