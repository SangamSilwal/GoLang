package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	p.views += 1
	p.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var1 := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go var1.inc(&wg)
	}

	wg.Wait()
	fmt.Println(var1.views)
}
