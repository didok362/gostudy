package main

import (
	"fmt"
	"sync"
	"time"
)

type Storage struct {
	db map[string]string
	mu sync.Mutex
}

func (s *Storage) Set(key string, value string) {
	s.mu.Lock()
	s.db[key] = value
	s.mu.Unlock()
}

func (s *Storage) Get(key string) string {
	s.mu.Lock()
	value := s.db[key]
	s.mu.Unlock()
	return value
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	delete(s.db, key)
	s.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	starttime := time.Now()
	s := &Storage{
		db: make(map[string]string),
	}
	for i := 0; i < 1000; i++ {
		wg.Add(3)
		go func() {
			defer wg.Done()
			s.Set("Govno", "Macha")
		}()
		go func() {
			defer wg.Done()
			fmt.Println(s.Get("Govno"))
		}()
		go func() {
			defer wg.Done()
			s.Delete("Govno")
		}()
	}
	wg.Wait()
	fmt.Printf("Time passed: %v\n", time.Since(starttime))

}
