package main

import (
	"fmt"
	"sync"
)

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup   // вэйт группа для определения количества воркеров
	var resultMu sync.Mutex // мьютекс чтобы избежать гонки за данными
	var result []int        // результирующий слайс
	wg.Add(len(arr))
	for _, i := range arr {
		go Qua(i, &result, &resultMu, &wg) // передаем адреса переменных
	}
	wg.Wait() // ждем заверщения всех воркеров
	fmt.Println(result)
}

// функция для расчета квадратов
func Qua(val int, result *[]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()                          // лочим
	*result = append(*result, val*val) //записываем данные в память
	mu.Unlock()                        // анлочим
	wg.Done()                          // передаем информацию о выполнении
}
