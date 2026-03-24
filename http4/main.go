package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		fmt.Println(key, value)
	}
	name := r.Header.Get("My-Name")
	_, err := w.Write([]byte("Привет, " + name + "!"))
	if err != nil {
		fmt.Println("Somthing went wrong")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8083", nil)
}
