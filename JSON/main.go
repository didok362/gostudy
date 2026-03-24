package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/k0kubun/pp"
)

type user struct {
	Name      string  `json:"name"`
	Adress    string  `json:"adress"`
	Age       int     `json:"age"`
	IsMarried bool    `json:"isMarried"`
	Height    float32 `json:"height"`
}

func handlerdef(w http.ResponseWriter, r *http.Request) {
	var User user
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	json.Unmarshal(body, &User)
	pp.Println(User)
}

func handleruser(w http.ResponseWriter, r *http.Request) {
	User := user{
		Name:      "shk",
		Adress:    "penbiso!",
		Age:       67,
		IsMarried: true,
		Height:    200,
	}
	output, err := json.Marshal(User)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = w.Write(output)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func main() {
	http.HandleFunc("/", handlerdef)
	http.HandleFunc("/user", handleruser)
	http.ListenAndServe(":8088", nil)
}
