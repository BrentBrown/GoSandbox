package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex = sync.Mutex{}

//go run -race main.go

func main() {
	wg.Add(2)
	//go incrementWithRace("Foo:")
	//go incrementWithRace("Bar:")
	go incrementWithMutex("Foo:")
	go incrementWithMutex("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementWithRace(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func incrementWithMutex(s string) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
		mutex.Unlock()
	}
	wg.Done()
}
