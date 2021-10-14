package main

import (
	"sync"
)

type MathFunction struct {
	F func(int) int
}

func Generator(mfunc MathFunction) <-chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			out <- mfunc.F(i)
		}
		close(out)
	}()

	return out
}

func Merge(cs... <-chan int) <-chan int{
	wg := sync.WaitGroup{}
	out := make(chan int)

	wg.Add(1)

	go func() {
		for _, c := range(cs) {
			wg.Add(1)
			go func(ch <-chan int) {
				for x := range(ch) {
					out <- x
				}
				wg.Done()
			}(c)
		}
		wg.Done()
	} ()

	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

func MergeResult(funcs []MathFunction) []int {
	var chs []<-chan int
	for _, f := range(funcs) {
		chs = append(chs, Generator(f))
	}
	ch := Merge(chs...)
	var slice []int 
	for i := range ch {
		slice = append(slice, i)
	}
	return slice
}