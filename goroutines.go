// goroutines
package main

import (
	"fmt"
	"sync"
)

/*
goroutines and concurrency are built into the core design of Go.
They're similar to threads but work differently.
Go also gives you full support to sharing memory in your goroutines.
 One goroutine usually uses 4~5 KB of stack memory.
Therefore, it's not hard to run thousands of goroutines on a single computer.
A goroutine is more lightweight, more efficient and more convenient than system threads.
goroutines run on the thread manager at runtime in Go.
 We use the go keyword to create a new goroutine, which is a function at the underlying level
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func sayGoRoutine() {
	//this provides random output, sometimes main ends before this goroutine finishes
	//some synchranization is needed
	go say("World")
	say("hello")
}

func saySynchronized(s string, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
		wg.Done()
	}()
}

func sayGoRoutineSynchronized() {
	var wg sync.WaitGroup
	saySynchronized("world", &wg)
	wg.Wait()
	say("hello")
	wg.Wait()
}

// c <- value sends value to channel
// <-c retrieves value from channel
func sumChannel(s []uint64, c chan uint64) {
	total := uint64(0)
	for _, v := range s {
		total += v
	}
	c <- total
}

var bigNumber uint64 = 100000000

//show some benchmarking here
func sumGoRoutines(s []uint64) uint64 {
	//Default channels are blocking.
	//Channel, way to access shared memory
	//Buffered channel => we are able to send 4 elements without blocking
	c := make(chan uint64, 4)
	l := len(s)
	go sumChannel(s[:l/4], c)
	go sumChannel(s[l/4:l/2], c)
	go sumChannel(s[l/2:3*l/4], c)
	go sumChannel(s[3*l/4:], c)
	sum := <-c + <-c + <-c + <-c
	return sum
}

func sumGoPlain(s []uint64) uint64 {
	sum := uint64(0)
	for _, v := range s {
		sum += v
	}
	return sum
}

func generateLargeSlice(n uint64) (s []uint64) {
	s = make([]uint64, 0, 10)
	for i := uint64(1); i < n; i++ {
		s = append(s, i)
	}
	return
}
