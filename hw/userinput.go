package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	database := make(map[string]int)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Ввеедите текст: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			continue
		}
		text := scanner.Text()
		input := strings.Fields(text)
		if len(input) != 3 {
			fmt.Printf("\nInvalid input!\n\nUse:\nдобавить key value\nудалить key value\n\n")
			continue
		}
		command := input[0]
		key := input[1]
		value, err := strconv.Atoi(input[2])
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}
		fmt.Printf("Команда: %s\nКлюч: %s\nЗначение: %d\n", command, key, value)
		if strings.ToLower(command) == "добавить" {
			database[key] += value
			fmt.Println("")
			pp.Println("db: ", database)
			fmt.Println("")
		} else if (strings.ToLower(command)) == "удалить" {
			if database[key] > value {
				database[key] -= value
			} else if database[key] <= value {
				database[key] = 0
			}
			fmt.Println("")
			pp.Println("db: ", database)
			fmt.Println("")
		} else if (strings.ToLower(command)) == "получить" {
			fmt.Printf("\nКлючь: %s, Значение: %d\n\n", key, database[key])
		} else {
			fmt.Printf("Invalid input!\n\nUse:\nдобавить key value\nудалить key value\n\n")
		}

	}
}
