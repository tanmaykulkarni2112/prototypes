package main

import (
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var count int64 = 0

func main() {
	var wg sync.WaitGroup

	numWorkers := 50 // increase concurrency

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go getResponse(0, &wg) // no artificial delay
	}

	wg.Wait()
	log.Println("Final Count:", count)
}

func getResponse(millisec int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	log.Println("Worker started")

	for i := 0; i < 1000; i++ {

		resp, err := client.Get("http://localhost:3000/home")
		if err != nil {
			log.Println("Request failed:", err)
			continue
		}

		// Close immediately (NO defer inside loop)
		resp.Body.Close()

		// Thread-safe increment
		newVal := atomic.AddInt64(&count, 1)

		if newVal%100 == 0 {
			log.Println("Requests sent:", newVal)
		}

		if millisec > 0 {
			time.Sleep(time.Duration(millisec) * time.Millisecond)
		}
	}
}