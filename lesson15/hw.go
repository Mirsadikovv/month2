package main

import (
	"fmt"
	"sync"
)

func plus(wg *sync.WaitGroup, givenSlice []int, index int, channnel chan []int) {
	defer wg.Done()
	givenSlice[index] = givenSlice[index] + 2
	channnel <- givenSlice

}

func multiple(wg *sync.WaitGroup, givenSlice []int, index int, channnel chan []int) {
	defer wg.Done()
	givenSlice[index] = givenSlice[index] * 2
	channnel <- givenSlice

}

func devide(wg *sync.WaitGroup, givenSlice []int, index int, channnel chan []int) {
	defer wg.Done()
	givenSlice[index] = givenSlice[index] / 2
	channnel <- givenSlice

}

func minus(wg *sync.WaitGroup, givenSlice []int, index int, channnel chan []int) {
	defer wg.Done()
	givenSlice[index] = givenSlice[index] - 2
	channnel <- givenSlice

}

func main() {

	var wg sync.WaitGroup
	var mx sync.Mutex
	var givenSlice = []int{1, 2, 3, 4}
	ch1 := make(chan []int)

	wg.Add(4)
	mx.Lock()
	go plus(&wg, givenSlice, 0, ch1)
	sum1 := <-ch1
	mx.Unlock()
	mx.Lock()
	go minus(&wg, sum1, 1, ch1)
	sum2 := <-ch1
	mx.Unlock()
	mx.Lock()
	go multiple(&wg, sum2, 2, ch1)
	sum3 := <-ch1
	mx.Unlock()
	mx.Lock()
	go devide(&wg, sum3, 3, ch1)
	sum4 := <-ch1
	mx.Unlock()

	fmt.Println(sum4)

	wg.Wait()

}
