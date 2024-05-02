package main

/*

/////////////////////////////////11111111111111111111111

*/

// func printTen(wg *sync.WaitGroup, temp int) {
// 	defer wg.Done()
// 	fmt.Println(temp)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i <= 10; i++ {
// 		wg.Add(1)
// 		go printTen(&wg, i)
// 	}
// 	wg.Wait()

// }

/*

/////////////////////////////////22222222222222222222222

*/

// func calculateSum(wg *sync.WaitGroup, channnel chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 1; i <= 100; i++ {
// 		sum += i
// 	}
// 	channnel <- sum

// }

// func main() {

// 	var wg sync.WaitGroup

// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go calculateSum(&wg, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

/////////////////////////////////33333333333333333333333333

*/

// func calculateFactorial(wg *sync.WaitGroup, givenNum int, channnel chan int) {
// 	defer wg.Done()
// 	sum := 1
// 	for i := 1; i <= givenNum; i++ {
// 		sum *= i
// 	}
// 	channnel <- sum

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenNum = 4
// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go calculateFactorial(&wg, givenNum, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

/////////////////////////////////444444444444444444444444

*/

// func checkConc(wg *sync.WaitGroup, givenNum int, channnel chan bool) {
// 	defer wg.Done()
// 	sum := 0
// 	check := false
// 	for i := 1; i < givenNum; i++ {
// 		if givenNum%i == 0 {
// 			sum++
// 		}
// 	}
// 	if sum < 2 {
// 		check = true
// 	}
// 	channnel <- check

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenNum = 2
// 	ch1 := make(chan bool)

// 	wg.Add(1)
// 	go checkConc(&wg, givenNum, ch1)
// 	sum1 := <-ch1
// 	if sum1 {
// 		fmt.Println("prostoe")
// 	} else {
// 		fmt.Println("ne prostoe")
// 	}

// 	wg.Wait()

// }

/*

/////////////////////////////////5555555555555555555

*/

// func findLargest(wg *sync.WaitGroup, givenSlice []int, channnel chan int) {
// 	defer wg.Done()
// 	sort.Slice(givenSlice, func(i, j int) bool {
// 		return givenSlice[i] > givenSlice[j]
// 	})

// 	channnel <- givenSlice[0]

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenSlice = []int{1, 52, 33, 41, 1}
// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go findLargest(&wg, givenSlice, ch1)
// 	largest := <-ch1

// 	fmt.Println(largest)

// 	wg.Wait()

// }

/*

/////////////////////////////////6666666666666666666666

*/

// func reverseString(wg *sync.WaitGroup, givenString string, channnel chan string) {
// 	defer wg.Done()
// 	runes := []rune(givenString)

// 	n := len(runes)

// 	for i := 0; i < n/2; i++ {
// 		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
// 	}
// 	runes1 := string(runes)
// 	channnel <- runes1

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenString = "hajime"
// 	ch1 := make(chan string)

// 	wg.Add(1)
// 	go reverseString(&wg, givenString, ch1)
// 	largest := <-ch1

// 	fmt.Println(largest)

// 	wg.Wait()

// }

/*

/////////////////////////////////7777777777777777777777777

*/

// func multipleToTwo(wg *sync.WaitGroup, givenSlice []int, channnel chan []int) {
// 	defer wg.Done()
// 	for i := range givenSlice {
// 		givenSlice[i] *= 2
// 	}

// 	channnel <- givenSlice

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenSlice = []int{1, 52, 33, 41, 1}
// 	ch1 := make(chan []int)

// 	wg.Add(1)
// 	go multipleToTwo(&wg, givenSlice, ch1)
// 	largest := <-ch1

// 	fmt.Println(largest)

// 	wg.Wait()

// }

/*

/////////////////////////////////888888888888888888888888

*/

// func calculateAvg(wg *sync.WaitGroup, givenSlice []int, channnel chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := range givenSlice {
// 		sum += givenSlice[i]
// 	}
// 	avg := sum / len(givenSlice)
// 	channnel <- avg

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenSlice = []int{1, 2, 3, 4}
// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go calculateAvg(&wg, givenSlice, ch1)
// 	largest := <-ch1

// 	fmt.Println(largest)

// 	wg.Wait()

// }

/*

////////////////////////////////9999999999999999999999

*/

// func sortSlice(wg *sync.WaitGroup, givenSlice []int, channnel chan []int) {
// 	defer wg.Done()
// 	sort.Slice(givenSlice, func(i, j int) bool {
// 		return givenSlice[i] < givenSlice[j]
// 	})

// 	channnel <- givenSlice

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenSlice = []int{1, 52, 33, 41, 1}
// 	ch1 := make(chan []int)

// 	wg.Add(1)
// 	go sortSlice(&wg, givenSlice, ch1)
// 	largest := <-ch1

// 	fmt.Println(largest)

// 	wg.Wait()

// }

/*

////////////////////////////////10 10 10 10 10 10

*/

// func calculateSumOfSquares(wg *sync.WaitGroup, channnel chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i * i
// 	}
// 	channnel <- sum

// }

// func main() {

// 	var wg sync.WaitGroup

// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go calculateSumOfSquares(&wg, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

////////////////////////////////11 11 11 11 11 11 11

*/

// func calculateFibonacci(wg *sync.WaitGroup, givenNum int, channnel chan []int) {
// 	defer wg.Done()
// 	fibonacciSequence := []int{0, 1}

// 	for i := 2; ; i++ {
// 		next := fibonacciSequence[i-1] + fibonacciSequence[i-2]
// 		if next > givenNum {
// 			break
// 		}
// 		fibonacciSequence = append(fibonacciSequence, next)
// 	}

// 	channnel <- fibonacciSequence

// }

// func main() {

// 	var wg sync.WaitGroup
// 	givenNum := 100
// 	ch1 := make(chan []int)

// 	wg.Add(1)
// 	go calculateFibonacci(&wg, givenNum, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

////////////////////////////////12 12 12 12 12 12 12

*/

// func calculateSum(wg *sync.WaitGroup, channnel chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 1; i <= 10; i++ {
// 		sum += i
// 	}
// 	channnel <- sum

// }

// func main() {

// 	var wg sync.WaitGroup

// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go calculateSum(&wg, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

////////////////////////////////13 13 13 13 13 13 13

*/

// func findLenOfLongest(wg *sync.WaitGroup, givenSlice []string, channnel chan int) {
// 	defer wg.Done()
// 	sort.Slice(givenSlice, func(i, j int) bool {
// 		return len(givenSlice[i]) > len(givenSlice[j])
// 	})

// 	channnel <- len(givenSlice[0])

// }

// func main() {

// 	var wg sync.WaitGroup
// 	var givenSlice = []string{"asd", "czxzfz", "rqwqwqweqweeww", "qwe"}
// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go findLenOfLongest(&wg, givenSlice, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	wg.Wait()

// }

/*

////////////////////////////////14 14 14 14 14 14 14 14

*/

// func findCountOfLines(wg *sync.WaitGroup, givenFile os.File, channnel chan int) {
// 	defer wg.Done()
// 	defer givenFile.Close()
// 	scanner := bufio.NewScanner(&givenFile)
// 	lineCount := 0

// 	for scanner.Scan() {
// 		lineCount++
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Ошибка при сканировании файла:", err)
// 		return
// 	}
// 	channnel <- lineCount

// }

// func main() {

// 	var wg sync.WaitGroup
// 	file, err := os.Open("first.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ch1 := make(chan int)

// 	wg.Add(1)
// 	go findCountOfLines(&wg, *file, ch1)

// 	sum1 := <-ch1
// 	fmt.Println(sum1)

// 	file2, err := os.Open("second.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ch2 := make(chan int)

// 	wg.Add(1)
// 	go findCountOfLines(&wg, *file2, ch2)

// 	sum2 := <-ch2
// 	fmt.Println(sum2)

// 	wg.Wait()

// }

/*

////////////////////////////////15 15 15 15 15 15 15

*/

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func produce(ch chan string) {
	for {
		newNum := strconv.Itoa(rand.Intn(100))
		ch <- "data " + newNum
		fmt.Println("Produced data = ", newNum)
		time.Sleep(time.Second)
	}
}

func consumer(ch chan string) {
	for {
		data := <-ch
		fmt.Println("Consumed: ", data)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan string, 1)

	go produce(ch)
	go consumer(ch)

	time.Sleep(10 * time.Second)
}
