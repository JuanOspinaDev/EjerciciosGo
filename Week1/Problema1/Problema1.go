package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var positiveCount, negativeCount, zeroCount float64

	for _, value := range arr {
		if value > 0 {
			positiveCount += 1
		} else if value < 0 {
			negativeCount += 1
		} else {
			zeroCount += 1
		}
	}
	total := float64(len(arr))
    
	fmt.Printf("%.6f\n", positiveCount/total)
	fmt.Printf("%.6f\n", negativeCount/total)
	fmt.Printf("%.6f\n", zeroCount/total)
}

func main() {
	arr := []int32{1, 1, 0, -1, -1}

	plusMinus(arr)
}
