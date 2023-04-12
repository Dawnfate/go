package main

import (
	"fmt"
	"sync"
)

func main() {
	var x int32
	num := 1000
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func() {
			m.Lock()
			x = x + 1
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(x)
}
