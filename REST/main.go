package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type book struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Value    int       `json:"value"`
	IsReaded bool      `json:"isReaded"`
	AddedAt  time.Time `json:"addedAt"`
	ReadAt   time.Time `json:"readAt"`
}

var books = []book{}

func bookshandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	splitted := strings.Split(path, "/")
	var textid string

	if len(splitted) > 2 {
		textid = splitted[2]
	}

	switch method {
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		newBook := book{
			AddedAt: time.Now(),
			ID:      rand.IntN(9999),
		}
		json.Unmarshal(body, &newBook)
		books = append(books, newBook)
	case "PATCH":
		id, err := strconv.Atoi(textid)
		if err != nil {
			fmt.Println(err.Error())
		}
		for i := 0; i < len(books); i++ {
			if books[i].ID == id {
				books[i].IsReaded = true
			}
		}
	case "GET":
		if textid == "" {
			filtered := []book{}
			author := r.URL.Query().Get("author")
			isReaded := r.URL.Query().Get("isReaded")
			for _, b := range books {
				if author != "" && b.Author != author {
					continue
				}
				if isReaded != "" && b.IsReaded != (isReaded == "true") {
					continue
				}
				filtered = append(filtered, b)
			}
			databasae, _ := json.Marshal(filtered)
			_, err := w.Write(databasae)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			id, err := strconv.Atoi(textid)
			if err != nil {
				fmt.Println(err.Error())
			}
			for i := 0; i < len(books); i++ {
				if books[i].ID == id {
					output, err := json.Marshal(books[i])
					if err != nil {
						fmt.Println(err.Error())
					}
					_, err = w.Write(output)
					if err != nil {
						fmt.Println(err.Error())
					}
					return
				}
			}
			http.Error(w, "Book not found", http.StatusNotFound)
		}
	case "DELETE":
		if textid == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(textid)
		if err != nil {
			fmt.Println(err.Error())
		}
		for i := 0; i < len(books); i++ {
			if books[i].ID == id {
				books = append(books[:i], books[i+1:]...)
				return
			}
		}
		http.Error(w, "Book not found", http.StatusNotFound)
	}

}

func main() {
	http.HandleFunc("/books/", bookshandler)
	http.ListenAndServe(":8088", nil)
}
