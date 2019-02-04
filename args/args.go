package main

import "fmt"

func main() {
	x, _ := curriedMult(4, 2, 2)

	_, timesFour := curriedMult(4)
	y := timesFour(2, 2)

	fmt.Println(x, y)
}

func curriedMult(num int, nums ...int) (int, func(...int) int) {
	if len(nums) == 0 {
		return 0, func(n ...int) int {
			var result int = num

			for _, v := range n {
				result *= v
			}

			return result
		}
	}

	var result int = num

	for _, v := range nums {
		result *= v
	}

	return result, nil
}
