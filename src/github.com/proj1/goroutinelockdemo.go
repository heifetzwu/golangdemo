package main

import (
	"fmt"
	"sync"
	"time"
)

func withdraw() {
	mu.Lock()
	balance := money
	time.Sleep(0000 * time.Millisecond)
	balance -= 1000
	money = balance
	mu.Unlock()
	fmt.Println("After withdrawing $1000, balace: ", money)
	
	wg.Done()
}

var wg sync.WaitGroup
var mu sync.Mutex
var money int = 1500

func goroutinelockdemo() {
	fmt.Println("3~~~We have $1500")
	wg.Add(3)
	go withdraw() // first withdraw
	go withdraw() // second withdraw
	go withdraw() // second withdraw
	wg.Wait()
}