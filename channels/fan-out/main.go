/* NOTE
- Buffered Channel = Send can be completed without "recv" => Reduce the latency between "send" & "recv"
- Buffered Channel => Synchronization Problem, same goroutine could try to read the same line
- So we still see latency on send side due to the fact that multiple sends are happening at the same time
- Good for CLI Tool, Cron Job, lambda function
- Need to be careful if you want to use this pattern for web service
- Multiple send might put alot of load on the system very quickly
- Goroutines are running in parallel & concurrently (out-of-order execution)
*/

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fanOut()
}

// fanOut: In this pattern, the parent goroutine creates 2000 child goroutines
// and waits for them to signal their results.
func fanOut() {
	children := 2000
	ch := make(chan string, children) // Buffered Channel, Signal with String Data

	// Send
	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	// Recv
	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
