package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		_, err := w.Write([]byte("Incorrect method type"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		w.WriteHeader(200)
		fmt.Println("Успешно обарбоатн")
		_, err := w.Write([]byte("You used right metod!"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
