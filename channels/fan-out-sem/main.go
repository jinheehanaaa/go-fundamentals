/*
- We can control # of goroutines that are executing (ex: 8 out of 2000) at any given time
- Buffered Channel = Send can be completed without rect => Reduce the latency between send & recv
- Semaphore Channel = Buffered Channel with Running State
- Let scheduler choose goroutine state from runnable to running state
- Only 8 goroutines can do the work inside sending block b/c of hardware spec
*/

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Advanced patterns
	fanOutSem()

}

// fanOutSem: In this pattern, a semaphore is added to the fan out pattern
// to restrict the number of child goroutines that can be schedule to run.
func fanOutSem() {
	children := 2000
	ch := make(chan string, children) // Buffered Channel, 1 slot for each goroutine

	g := runtime.GOMAXPROCS(0) // Semaphore Channel
	sem := make(chan bool, g)

	// send
	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true // Signal for sem
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			<-sem // pull value out of semaphor channel, Remove
		}(c)
	}

	// recv
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
