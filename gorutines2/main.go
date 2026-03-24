package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("peniso!")
		fmt.Println("peniso!")
		fmt.Println("peniso!")
	}()
	go func() {
		fmt.Println("peniso2!")
		fmt.Println("peniso2!")
		fmt.Println("peniso2!")
	}()
	go func() {
		fmt.Println("peniso3!")
		fmt.Println("peniso3!")
		fmt.Println("peniso3!")
	}()
	time.Sleep(500 * time.Microsecond)
}
