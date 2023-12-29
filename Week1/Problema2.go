package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minSum, maxSum int64

	for i := 0; i < 4; i++ {
		minSum += int64(arr[i])
		maxSum += int64(arr[i+1])
	}

	fmt.Println(minSum, maxSum)
}

func main() {
	arr := []int32{1, 3, 5, 7, 9}
	miniMaxSum(arr)
}
