package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	res := make([]int, 0, size)
	src := rand.NewSource(int64(1048576))
	var randomNumber int64

	for range size {
		randomNumber = (src.Int63() + 1)
		res = append(res, int(randomNumber))
	}

	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, num := range data {
		if num >= max {
			max = num
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}else if len(data) < CHUNKS {
 	   return maximum(data)
	}

	maxAr := make([]int, CHUNKS)
	var wg sync.WaitGroup

	cut := len(data) / CHUNKS
	for i := range CHUNKS {
		wg.Add(1)
		go func(i int) { 
			defer wg.Done()
			if i != CHUNKS-1 {
				maxAr[i] = maximum(data[(i * cut):(i*cut + cut)])
			} else {
				maxAr[i] = maximum(data[(CHUNKS-1)*cut:])
			}
		}(i)
	}
	wg.Wait()

	return maximum(maxAr)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	initTime := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(initTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	initTime = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(initTime)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed.Microseconds())
}
