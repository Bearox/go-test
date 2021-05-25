package main

import (
	"fmt"
	"math"
)

func main() {
	ret := judgeSquareSum(1000)
	fmt.Println(ret)
}

func judgeSquareSum(c int) bool {
	temp := math.Sqrt(float64(c))
	left := 0
	right := int(math.Floor(temp))

	for left <= right {
		if left * left + right * right == c {
			return true
		} else if left * left + right * right > c {
			right -= 1
		} else {
			left += 1
		}
	}
	return false
}
