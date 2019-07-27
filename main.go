package main

import (
	"fmt"
	"sync"
        "runtime"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
        fmt.Println("Num of CPUs:", runtime.NumCPU())
	wg.Add(2)
	go firstProc()
	go secondProc()
	wg.Wait()
}

func firstProc() {
	for i := 0; i < 50; i++ {
		fmt.Println("Proc01:", i)
	}
	wg.Done()
}

func secondProc() {
	for i := 0; i < 50; i++ {
		fmt.Println("  Proc02:", i)
	}
	wg.Done()
}
