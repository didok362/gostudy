package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	var output string
	switch code {
	case "200":
		w.WriteHeader(200)
		output = "StatusOK"
	case "400":
		w.WriteHeader(400)
		output = "StatusBadRequest"
	default:
		w.WriteHeader(404)
		output = "StatusNotFound"
	}
	_, err := w.Write([]byte(output))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Корректная работа")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8082", nil)
}
