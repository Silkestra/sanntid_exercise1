// Use `go run foo.go` to run your program

package main

import (
	"fmt"
	"time"
)

/*


Shared vaiable / mutex -> not recomended for Golang
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	i  int
}


func (c *SafeCounter) incrementing() {
	for j := 0; j < 1000000; j++ {
        c.mu.Lock()
		c.i++
        c.mu.Unlock()
	}
}

func (c *SafeCounter) decrementing() {
	for k := 0; k < 1000000; k++ {
        c.mu.Lock()
		c.i--
        c.mu.Unlock()
	}
}

func main() {
	// What does GOMAXPROCS do? What happens if you set it to 1?
	runtime.GOMAXPROCS(2)
	c := SafeCounter{i:0}
	for i := 0; i < 1000; i++ {
		go c.incrementing()
        go c.decrementing()
	}
	// TODO: Spawn both functions as goroutines

	// We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
	// We will do it properly with channels soon. For now: Sleep.
	time.Sleep(500 * time.Millisecond)
	fmt.Println("The magic number is:", c.i)
}
*/

//Number server method -> idomatic for Golang

// Request struct to define the operations on the counter
type Request struct {
	action string   // "increment", "decrement", or "get"
	result chan int // Channel to return the result for "get"
}

func Server() chan<- Request {
	reqChan := make(chan Request)

	go func() {
		var counter int // Internal count
		for req := range reqChan {
			switch req.action {
			case "increment":
				counter++
			case "decrement":
				counter--
			case "get":
				req.result <- counter
			}
		}
	}()
	return reqChan
}

func main() {
	reqChan := Server()

	for i := 0; i < 1000; i++ {
		go func() { reqChan <- Request{action: "increment"} }()
		go func() { reqChan <- Request{action: "decrement"} }()
	}

	time.Sleep(2 * time.Second)

	resultChan := make(chan int)
	reqChan <- Request{action: "get", result: resultChan}
	finalValue := <-resultChan
	fmt.Println("The magic number is:", finalValue)
}
