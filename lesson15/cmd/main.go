package main

import (
	"fmt"
	"sync"
)

// func makeCoffee() {
// 	fmt.Println("Coffee is ready")
// }
// func makeDish() {
// 	fmt.Println("Dish is ready")
// }

// func main() {
// 	now := time.Now()
// 	var wg sync.WaitGroup
// 	wg.Add(2)

// 	go func() {
// 		makeCoffee()
// 		wg.Done()
// 	}()

// 	go func() {
// 		makeDish()
// 		wg.Done()
// 	}()

// 	wg.Wait()
// 	fmt.Println(time.Since(now))
// }

//////////////////////////////////////

/////////////////////////////////////

// func calculateSquare(wg *sync.WaitGroup, start int, end int, channnel chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := start; i <= end; i += 2 {
// 		sum += i * i
// 	}
// 	channnel <- sum

// }

// func main() {

// 	var wg sync.WaitGroup
// 	start := 1
// 	end := 4
// 	ch1 := make(chan int)
// 	// ch2 := make(chan int)

// 	wg.Add(2)
// 	go calculateSquare(&wg, start, end, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	go calculateSquare(&wg, start+1, end, ch1)

// 	sum2 := <-ch1
// 	fmt.Println(sum2)
// 	wg.Wait()

// }

////////////////////////////////////

///////////////////////////////////

var counter = 0

// var mutex = sync.Mutex{}

func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	// mutex.Lock()
	counter += 12
	// mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go incrementCounter(&wg)
	go incrementCounter(&wg)
	go incrementCounter(&wg)
	go incrementCounter(&wg)

	fmt.Println(counter)
	wg.Wait()
}
