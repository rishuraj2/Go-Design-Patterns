package main

import (
	"fmt"
	"sync"
	"threadsafecounter/counter"
)

func incN(wg *sync.WaitGroup, c *counter.Counter, n int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		c.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	c1 := counter.GetCounter()
	c2 := counter.GetCounter()

	incVal := []int{9980, 20}

	wg.Add(2)

	go incN(&wg, c1, incVal[0])
	go incN(&wg, c2, incVal[1])

	wg.Wait()

	total := 0
	for _, val := range incVal {
		total += val
		if total > 9999 {
			total = 9999
			break
		}
	}

	fmt.Printf("Expected count: %d\n", total)
	fmt.Printf("Got Count: %d\n", c1.GetCount())
}
