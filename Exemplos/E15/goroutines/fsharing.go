package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type paddedInt struct {
	value int
	_     [56]byte
}

func main() {
	var numGoroutines int
	var numIterations int
	flag.IntVar(&numGoroutines, "g", 4, "number of goroutines")
	flag.IntVar(&numIterations, "s", 10000000, "number of iterations per goroutine")
	flag.Parse()

	// Sem padding
	start := time.Now()
	counters := make([]int, numGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				counters[index]++
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Sem padding: %v\n", time.Since(start))

	// Com padding
	start = time.Now()
	paddedCounters := make([]paddedInt, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				paddedCounters[index].value++
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Com padding: %v\n", time.Since(start))
}
