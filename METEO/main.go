package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type data struct {
	value int
	x     int
	y     int
	z     int
}

func pressuresensor(ctx context.Context, ch chan data, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Датчик давления номре %d завершил дейсвтие\n", i)
			return
		default:
			randnum := rand.Intn(50) + 50
			fmt.Printf("Датчик давления номре %d работает показатель: %d Па\n", i, randnum)
			newData := data{
				value: randnum,
				x:     i * 1000,
				y:     i * 1500,
				z:     i * 500,
			}
			ch <- newData
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func humiditysensor(ctx context.Context, ch chan data, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Датчик влажность номре %d завершил дейсвтие\n", i)
			return
		default:
			randnum := rand.Intn(50) + 50
			fmt.Printf("Датчик влажность номре %d работает показатель: %d Па\n", i, randnum)
			newData := data{
				value: randnum,
				x:     i * 1000,
				y:     i * 1500,
				z:     i * 500,
			}
			ch <- newData
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func seismicsensor(ctx context.Context, ch chan data, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Датчик сейсмической актинвости номре %d завершил дейсвтие\n", i)
			return
		default:
			randnum := rand.Intn(50) + 50
			fmt.Printf("Датчик сейсмической актинвости номре %d рабоатет", i)
			newData := data{
				value: randnum,
				x:     i * 1000,
				y:     i * 1500,
				z:     i * 500,
			}
			ch <- newData
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	mainCtx, cancelAll := context.WithCancel(context.Background())
	pressureCtx, cancelPressure := context.WithCancel(mainCtx)
	humidityCtx, cancelHumidity := context.WithCancel(mainCtx)
	seismicCtx, cancelSeismic := context.WithCancel(mainCtx)
	pressureChan := make(chan data)
	humidityChan := make(chan data)
	seismicChan := make(chan data)

	wg.Add(1)
	go func() {
		defer wg.Done()
		pressuresensor(pressureCtx, pressureChan, 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		humiditysensor(humidityCtx, humidityChan, 1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		seismicsensor(seismicCtx, seismicChan, 1)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		cancelPressure()
		cancelHumidity()
		cancelSeismic()
		cancelAll()
	}()

	for {
		select {
		case <-mainCtx.Done():
			wg.Wait()
			return
		case d := <-pressureChan:
			fmt.Println("Давление:", d.value)
		case d := <-humidityChan:
			fmt.Println("Влажность:", d.value)
		case d := <-seismicChan:
			fmt.Println("Сейсмика:", d.value)
		}
	}
}
