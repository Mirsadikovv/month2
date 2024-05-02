package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{1, 2, 3, 4}))
}
func smallerNumbersThanCurrent(nums []int) []int {
	var num_of_smaller []int
	var count int
	for _, val := range nums {
		for _, val1 := range nums {
			if val > val1 {
				count++
			}
		}
		num_of_smaller = append(num_of_smaller, count)
		count = 0
	}
	return num_of_smaller
}

func fizzBuzz(n int) []string {
	var list []string
	for in := 1; in <= n; in++ {
		if in%15 == 0 {
			list = append(list, "FizzBuzz")
		} else if in%5 == 0 {
			list = append(list, "Buzz")
		} else if in%3 == 0 {
			list = append(list, "Fizz")
		} else {
			list = append(list, strconv.Itoa(in))
		}
	}
	return list
}

func missingNumber(nums []int) int {
	m := -1
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for in, val := range nums {
		fmt.Println(in, " ", val)
		if in != val {
			m = in
			break
		}
		if m < 0 {
			m = in + 1
		}

	}

	return m
}
