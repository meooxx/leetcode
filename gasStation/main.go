package main

import "fmt"

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 3, 2, 4, 1, 5, 3, 2, 4}, []int{
		1, 1, 1, 3, 2, 4, 3, 6, 7, 4, 3, 1,
	}))
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{
		3, 4, 5, 1, 2,
	}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{
		3, 4, 3,
	}))
}

func canCompleteCircuit(gas []int, cost []int) int {
	total := 0
	tank := 0
	result := 0
	for i := range gas {
		remaind := gas[i] - cost[i]
		total += remaind
		tank += remaind
		if tank < 0 {
			tank = 0
			result = i + 1
		}
	}

	if total >= 0 {
		return result
	}
	return -1
}

// leetcode pass fail
// gas  = [1, 2, 3, 4, 5]
// cost = [3, 4, 5, 1, 2]
//        -2 -2 -2  3  3
func canCompleteCircuit2(gas []int, cost []int) int {
	for i := range gas {
		remaind := gas[i] - cost[i]

		if remaind >= 0 {
			sum := remaind
			for j := i + 1; j%len(gas) != i; j++ {
				sum += gas[j%len(gas)] - cost[j%len(gas)]
				if sum < 0 {
					break
				}
			}
			if sum >= 0 {
				return i
			}
		}
	}

	return -1

}
