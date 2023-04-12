package main

import (
	"fmt"
	"sync"
)

func main() {
	x := 0
	num := 1000
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			x = x + 1
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
}
