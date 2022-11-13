package main

import "fmt"

func maxPoints(points [][]int) int {

	// 15 9
	getGCD := func(a, b int) int {
		for b != 0 {
			r := a % b
			a, b = b, r
		}

		return a
	}
	max := 0
	// k = y2-y1 / (x2-x1)
	for i := 0; i < len(points); i++ {
		// 从一个开始找所有的
		kMap := map[string]int{}
		for j := i + 1; j < len(points); j++ {
			y := points[j][1] - points[i][1]
			x := points[j][0] - points[i][0]
			gcd := getGCD(x, y)

			key := fmt.Sprintf("%d/%d", y/gcd, x/gcd)
			kMap[key]++
			if kMap[key] > max {
				max = kMap[key]
			}
		}
	}

	return max + 1
}
