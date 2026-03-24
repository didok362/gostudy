package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func handelr(w http.ResponseWriter, r *http.Request) {
	strings := []string{"hello world!", "DINIS", "123", "kakashki"}
	randnum := rand.Intn(len(strings))
	str := strings[randnum]
	_, err := w.Write([]byte(str))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Корректная работа")
	}
}

func main() {
	http.HandleFunc("/", handelr)
	http.ListenAndServe(":8080", nil)

}
