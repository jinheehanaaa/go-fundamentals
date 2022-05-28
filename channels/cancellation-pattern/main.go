/*NOTE
- We need to cancel some work to send data to the client when the client request something from web.
- Timeout is important in this case.
- We call cancel function with timeout function.
- Call cancel function at least once to avoid memory leak.
- Buffered Channel of 1 = No guarantee of 1 signal send.
*/

/* COMMON MISTAKE
- If you use Unbuffered channel => send & recv have to come together and recv happens before send
- "send" cannot finish without "recv" when work cancelled => Goroutine Leak
*/

// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cancellation()
}

// cancellation: In this pattern, the parent goroutine creates a child
// goroutine to perform some work. The parent goroutine is only willing to
// wait 150 milliseconds for that work to be completed. After 150 milliseconds
// the parent goroutine walks away.
func cancellation() {
	duration := 150 * time.Millisecond                                 // give 150ms to complete
	ctx, cancel := context.WithTimeout(context.Background(), duration) // Get WithTimeout with context & duration to get new context (ctx)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	// Blocking Call
	select {
	case d := <-ch:
		fmt.Println("work complete", d)

	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
