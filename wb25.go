package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	endTime := time.Now().Unix() + int64(seconds) // определяем время выхода из сна
	fmt.Println(endTime)
	// бесконечный цикл пока время не подойдет к нужному
	for {
		if time.Now().Unix() >= endTime {
			break
		}
	}
	fmt.Println("Конец сна")
}

func main() {
	fmt.Println("Спим")
	sleepDuration := 5 // пауза од 5 секунд
	sleep(sleepDuration)
	fmt.Println("Завершаем работу")
}
