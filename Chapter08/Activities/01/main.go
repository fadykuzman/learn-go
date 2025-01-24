package main

import "fmt"

func findMin[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	n := []int{1, 32, 5, 8, 10, 11}
	result := findMin(n)
	f := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}
	flResult := findMin(f)
	fmt.Println(result)
	fmt.Println(flResult)
}
