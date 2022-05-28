/* NOTE
- We drop everything after 100 requests
- Useful for having too many pending requests
- select/case allows us to perform multiple channel operations at the same time on the same goroutine
- select/case allows us to detect that we're at capacity
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
	drop()
}

// drop: In this pattern, the parent goroutine signals 2000 pieces of work to
// a single child goroutine that can't handle all the work. If the parent
// performs a send and the child is not ready, that work is discarded and dropped.
func drop() {
	const cap = 100 // capacity
	ch := make(chan string, cap)

	// goroutine is waiting on a recv, this could be a pool of goroutines to handle DNS request
	// But we made just 1 for simplicity
	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	// Let's send 2000 DNS Request through server
	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
