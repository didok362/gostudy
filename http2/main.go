package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

type data struct {
	Id      int
	Message string
}

// func handler(w http.ResponseWriter, r *http.Request) {
// }

func main() {
	db := []data{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ID := rand.Intn(9999)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		Message := string(body)
		splited := strings.Fields(Message)
		if len(splited) == 0 {
			fmt.Println("Пустое сообщение")
			return
		}
		if splited[0] != "del" {
			newData := data{
				Id:      ID,
				Message: Message,
			}
			db = append(db, newData)
			fmt.Println("Получено:", Message)
			pp.Println(db)
		} else {
			splitedId, err := strconv.Atoi(splited[1])
			if err != nil {
				fmt.Println(err.Error())
			}
			for i := 0; i < len(db); i++ {
				if db[i].Id == splitedId {
					db = append(db[:i], db[i+1:]...)
					fmt.Printf("%d удалено\n", splitedId)
					pp.Println(db)
					return
				}
			}
			fmt.Printf("%d ненайдено\n", splitedId)
		}
	})
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		messeges, err := json.Marshal(db)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			_, err := w.Write(messeges)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("Корректная работа")
			}
		}
	})
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		for i := 0; i < len(db); i++ {
			Id, err := strconv.Atoi(string(body))
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			if db[i].Id == Id {
				_, err := w.Write([]byte(db[i].Message))
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Корректная работа")
				}
				return
			}
		}
		http.Error(w, "сообщение не найдено", 404)
	})
	http.ListenAndServe(":8085", nil)
}
