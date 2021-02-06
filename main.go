package main

import (
	"fmt"
	"sync"
)

func fib(i int) uint64 {
	if i < 2 {
		return (uint64)(i)
	}
	return fib(i-1) + fib(i-2)
}

func pfib(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("fib(", i, ")=", fib(i))
}

func main() {
	var fibCalcGroup sync.WaitGroup

	for i := 0; i < 100; i++ {
		fibCalcGroup.Add(1)
		go pfib(&fibCalcGroup, i)
	}

	fibCalcGroup.Wait()
}
