package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func sumWithoutParallelism(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func sumWithParallelism(arr []int, numGoroutines int) int {
	var wg sync.WaitGroup
	ch := make(chan int)
	chunkSize := len(arr) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(startIndex int) {
			defer wg.Done()
			partialSum := 0
			endIndex := startIndex + chunkSize
			for j := startIndex; j < endIndex; j++ {
				partialSum += arr[j]
			}
			ch <- partialSum
		}(i * chunkSize)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for partialSum := range ch {
		totalSum += partialSum
	}

	return totalSum
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <numGoroutines> <arraySize>")
		return
	}

	numGoroutines, err := strconv.Atoi(os.Args[1])
	if err != nil || numGoroutines <= 0 {
		fmt.Println("Invalid number of goroutines:", os.Args[1])
		return
	}

	arraySize, err := strconv.Atoi(os.Args[2])
	if err != nil || arraySize <= 0 {
		fmt.Println("Invalid array size:", os.Args[2])
		return
	}

	arr := make([]int, arraySize)
	for i := range arr {
		arr[i] = 1
	}

	start := time.Now()
	sum := sumWithoutParallelism(arr)
	elapsed := time.Since(start)
	fmt.Printf("Sum without parallelism: %d, Time taken: %s\n", sum, elapsed)

	start = time.Now()
	sum = sumWithParallelism(arr, numGoroutines)
	elapsed = time.Since(start)
	fmt.Printf("Sum with parallelism (%d goroutines): %d, Time taken: %s\n", numGoroutines, sum, elapsed)
}
