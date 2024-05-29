package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Выводим приглашение к вводу
		fmt.Print("$ ")

		// Читаем введенную команду
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		// Удаляем символ новой строки из введенной строки
		input = strings.TrimSpace(input)

		// Разбиваем введенную строку на части по пробелам
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		// Первое слово - команда, остальные - аргументы
		command := parts[0]
		args := parts[1:]

		// Выполняем команду
		switch command {
		case "cd":
			err := changeDirectory(args)
			if err != nil {
				fmt.Println("Ошибка смены директории:", err)
			}
		case "pwd":
			printWorkingDirectory()
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "kill":
			err := killProcess(args)
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			listProcesses()
		case "exit":
			fmt.Println("Выход из шелла.")
			return
		default:
			cmd := exec.Command(command, args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка выполнения команды:", err)
			}
		}
	}
}

// Функция для смены текущей директории
func changeDirectory(args []string) error {
	if len(args) == 0 {
		return os.Chdir(getHomeDir())
	}
	return os.Chdir(args[0])
}

// Функция для вывода текущей директории
func printWorkingDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка получения текущей директории:", err)
		return
	}
	fmt.Println(dir)
}

// Функция для получения домашней директории пользователя
func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Ошибка получения домашней директории:", err)
	}
	return home
}

// Функция для завершения процесса по PID
func killProcess(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("не указан PID процесса")
	}
	pidStr := args[0]                // Получаем строку с PID
	pid, err := strconv.Atoi(pidStr) // Преобразуем строку в int
	if err != nil {
		return fmt.Errorf("некорректный PID: %v", err)
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	return process.Signal(syscall.SIGKILL)
}

// Функция для вывода списка запущенных процессов
func listProcesses() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Ошибка при выполнении команды ps:", err)
	}
}
