package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/k0kubun/pp"
)

type task struct {
	Headline     string
	Body         string
	Creationtime time.Time
	IsDone       bool
	DoneTime     time.Time
}

func main() {
	tasks := []task{}
	events := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ввеедите текст: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			continue
		}
		text := scanner.Text()
		events = append(events, text)
		input := strings.Fields(text)
		command := input[0]
		switch command {
		case "help":
			fmt.Println("help/add/list/del/done/events/exit")
		case "add":
			newTask := task{
				Headline:     input[1],
				Body:         strings.Join(input[2:], " "),
				Creationtime: time.Now(),
			}
			tasks = append(tasks, newTask)
			fmt.Printf("Задача %s добвалена\n", input[1])
		case "list":
			pp.Println(tasks)
		case "del":
			IsFound := false
			if len(input) < 2 || len(input) > 3 {
				fmt.Println("Invalind input")
			}
			for i := 0; i < len(tasks); i++ {
				if tasks[i].Headline == input[1] {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Задача %s удалена\n", input[1])
					IsFound = true
					break
				}
			}
			if IsFound == false {
				fmt.Printf("Задача %s не найдена\n", input[1])
			}
		case "done":
			IsFound := false
			if len(input) < 2 || len(input) > 3 {
				fmt.Println("Invalind input")
			}
			for i := 0; i < len(tasks); i++ {
				if tasks[i].Headline == input[1] {
					tasks[i].DoneTime = time.Now()
					tasks[i].IsDone = true
					fmt.Printf("Задача %s помечана как выполененая\n", input[1])
					IsFound = true
					break
				}
			}
			if IsFound == false {
				fmt.Printf("Задача %s не найдена\n", input[1])
			}
		case "events":
			pp.Println("История:", events)
		case "exit":
			return
		}
	}
}
