/* NOTE:
- Unbuffered Channel => Gaurantee that recv comes first before send (ch <-"paper" is send implementation)
- Goroutine to do some work in multi-threaded environment
- "make" => making channel with open state
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
	// Do not use random number in production code
	rand.Seed(time.Now().UnixNano())
}

func main() {
	waitForResult()
}

// waitForResult: In this pattern, the parent goroutine waits for the child
// goroutine to finish some work to signal the result.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
