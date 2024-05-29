package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// Парсинг аргументов командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Использование: go run main.go [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Установка соединения с заданным таймаутом
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Printf("Не удалось подключиться к %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Подключено к %s\n", address)

	// Создание каналов для остановки горутин
	done := make(chan struct{})

	// Чтение из сокета и вывод в STDOUT
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Ошибка чтения из соединения: %v\n", err)
		}
		close(done)
	}()

	// Чтение из STDIN и запись в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintln(conn, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Ошибка чтения из STDIN: %v\n", err)
		}
		conn.Close()
		close(done)
	}()

	// Ожидание завершения одной из горутин
	<-done
	fmt.Println("Соединение закрыто")
}
