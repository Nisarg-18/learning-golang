package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var scores = []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("1")
		mut.Lock()
		scores = append(scores, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("2")
		mut.Lock()
		scores = append(scores, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		fmt.Println("3")
		mut.Lock()
		scores = append(scores, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex){
		mut.RLock()
		fmt.Println(scores)
		mut.RUnlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
	fmt.Println(scores)

}