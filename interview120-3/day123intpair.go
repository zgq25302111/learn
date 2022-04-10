package main

import "fmt"

func main() {
	nums := []int{3, 9, 1, 7, 2}
	findPairs(nums, 11)
}

func findPairs(nums []int, sum int){
    m := make(map[int]int)
	for i, numi := range nums {
		value, ok := m[numi]
		if ok {
			fmt.Println(numi, value)
		} else {
			m[sum-nums[i]] = nums[i]
		}
	}
}