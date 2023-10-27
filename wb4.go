package main

import (
	"fmt"
	"github.com/rs/xid"
	"sync"
)

func main() {
	var wg sync.WaitGroup //делайм вет группу чтобы определять количество воркеров
	var amountWorkers int //переменная для хранения количества воркеров

	fmt.Println("Enter amount workers...")
	fmt.Scanln(&amountWorkers) //ввод

	wg.Add(amountWorkers) //добавляем количество

	//цикл для запуска воркеров
	for i := 0; i < amountWorkers; i++ {
		//запускаем воркера в горутине
		go func(i int) {
			for {
				id := GenerateId()                   // генерируем id
				fmt.Printf("Worker %d: %s\n", i, id) // выводим информацию о воркере и id
			}
			wg.Done()
		}(i)
	}

	wg.Wait() // бесконечно ждем
}

func GenerateId() string {
	return xid.New().String()
}
