package main

import "fmt"

func main() {
	input := []int{2, 3, 4}
	target := 6
	output := twoSum(input, target)
	fmt.Println(output)
}

func twoSum(numbers []int, target int) []int {
	prevsMap := map[int]int{}
	result := make([]int, 2)

	for i := 0; i < len(numbers); i++ {
		if idxPrev, ok := prevsMap[target-numbers[i]]; ok {
			result[0] = idxPrev + 1
			result[1] = i + 1
			break
		}
		prevsMap[numbers[i]] = i
	}

	return result
}
