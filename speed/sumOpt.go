package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var (
	total float64 = 0
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func calculateSum(start, end float64) {
	//	fmt.Println(start, " ", end)
	tempTotal := 0.0
	for ; start < end; start++ {
		tempTotal += start
	}
	mu.Lock()
	total += tempTotal
	mu.Unlock()

	wg.Done()

}

func main() {
	upto, _ := strconv.ParseFloat(os.Args[1], 64)
	numChunks, _ := strconv.ParseFloat(os.Args[2], 64)
	chunkSize := upto / numChunks

	for numChunk := 0.0; numChunk < numChunks; numChunk++ {
		wg.Add(1)
		start := upto / numChunks * numChunk
		end := start + chunkSize
		go calculateSum(start, end)
	}

	wg.Wait()

	fmt.Println(total)
}
