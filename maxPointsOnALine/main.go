package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{
		{1, 1},
		{3, 2},
		{5, 3},
		{4, 1},
		{2, 3},
		{1, 4},
	}))
	fmt.Println(maxPoints([][]int{
		{0, 0},
		{-1, 1},
		{1, 1},
	}))
}

// y=kx+b所有点俩俩组合将可能的 k 遍历出来
// 找到相同 K 数量
// 统计最多点数目在线段上?
// 假如 1 2 3 点在容易线段上
// 最先出现的点会记录所有同其他点连接的记录, 最后加1即可
// eg.1 2 3 在一条线时, 就有 1-2 1-3俩种+1==3(从2的时候只有2-3的路径小于1开始的路径)
// eg.2 3 4 5 在一条线上时,就有 2-3, 2-4, 2-5 +1 == 4
// 后续遍历过程中会有从 3开始的情况, 3-4 3-5 总数小于从2开始的情况.同理4
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	// GCD(15 6) == GCD(6 15)
	// 15%6 ==3 , 6%3 == 0 return 3
	var GCD func(a, b int) int
	GCD = func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	max := 1
	for i := 0; i < len(points); i++ {
		mp := map[string]int{}
		for j := i + 1; j < len(points); j++ {
			dx := points[i][0] - points[j][0]
			dy := points[i][1] - points[j][1]
			// 使用 greatest common divisor
			// 防止浮点型数据问题
			// 也可计算出 K 值, leetcode 上反而使用更少的运行时间
			// ?? 竟然没有平行线的问题case 
			gcd := GCD(dx, dy)
			key := fmt.Sprint(dx/gcd, "/", dy/gcd)
			mp[key]++
			if mp[key] > max {
				max = mp[key]
			}
		}
	}

	return max + 1
}
