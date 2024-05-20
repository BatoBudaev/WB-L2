package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

func main() {
	//Запрос текущего времени с NTP-сервера
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Printf("Ошибка при получении времени: %v", err)
		os.Exit(1)
	}

	fmt.Println("Время:", time)
}
