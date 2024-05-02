// package main

// import (
// 	"fmt"
// 	"time"
// )

// func makeCoffee(ch chan string) {
// 	time.Sleep(time.Second * 2)
// 	ch <- "coffee is ready..."

// }
// func cookDish(ch chan string) {
// 	msg := <-ch
// 	fmt.Println("dish and ", msg)
// 	// time.Sleep(time.Second * 3)
// }

// func main() {
// 	ch := make(chan string)

// 	go makeCoffee(ch)
// 	go cookDish(ch)
// 	// go cookDish()
// 	// msg := <-ch
// 	// fmt.Println(msg)
// 	time.Sleep(time.Second * 5)
// }

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func produce(ch chan string) {
	for {
		newNum := strconv.Itoa(rand.Intn(100))
		ch <- "data" + newNum
		fmt.Println("Produced data = ", newNum)
		time.Sleep(time.Second)
	}
}

func consumer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("Consumed ", data)
		time.Sleep(time.Second)
	}
}

func consumer2(ch chan string) {
	for {
		data := <-ch
		fmt.Println("Consumed 2 ", data)
		time.Sleep(2 * time.Second)
	}
}
func main() {
	ch := make(chan string)
	go produce(ch)
	go consumer(ch)
	// go consumer2(ch)
	time.Sleep(time.Second * 5)
}
