package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	workQueue := make(chan work, 10)

	// start 5 worker goroutine pool
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for range workQueue {
				w := <-workQueue
				w.Func(w.Arg)
				//reflect.ValueOf(w.Func).Call(w.Args)
			}
		}()
	}

	// distribute 100 pieces of work among the pool
	for i := 0; i < 100; i++ {
		workQueue <- work{
			Func: func(n int) {
				fmt.Printf("goroutine %d - %d\n", n, n*n)
			},
			//Args: []reflect.Value{reflect.ValueOf(i)},
			Arg: i,
		}
	}

	close(workQueue)

	wg.Wait()
}

type work struct {
	//Func any
	//Args []reflect.Value
	Func func(n int)
	Arg  int
}
