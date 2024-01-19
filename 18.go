package main

import (
	"fmt"
	"sync"
)

// Создаем структуру Counter с полями val и мутексом для синхронизации
type Counter struct {
	val int
	mut sync.Mutex
}

// Создаем Inc(), который увеличивает значени счетчика
func (c *Counter) Inc() {
	//Блокируем доступ для других горутин
	c.mut.Lock()
	c.val++
	//Разблокируем
	c.mut.Unlock()
}

// Создаем Val(), который возрващает значение счетчика
func (c *Counter) Val() int {
	c.mut.Lock()
	//Отложенная разблокировка
	defer c.mut.Unlock()
	return c.val
}

func task18() {
	var wg sync.WaitGroup
	counter := Counter{}
	//Создаем 500 горутин
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Итог:", counter.Val())
}
