package main


//  [[3,15] [7,0]  [5,12] [12,0]]
//     i               j
//  	left			 right										pre = 0, rightH = 0, leftH = 0
//  compare(left[0], right[0])
//  [3,15] vs [5, 12] =>   pick left  =>		leftH = 15, x = 3, (maxH = amx(15, 0)) != pre, push(x, maxH)
//  [7, 0] vs [5,12], =>   pick right =>		rightH == 12 maxH= max(leftH, rightH) = max(15,12) == 15 continue
//  [7, 0] vs [12, 12]=>   pick left  =>    leftH = 0, x= 7, maxH =max(0, 12) push(x,maxH) push(7,12)
//             [12,0]         push(rest)     push(12, 0)
func getSkyline(buildings [][]int) [][]int {

	return getSkylineImpl(buildings)
}

func getSkylineImpl(buildings [][]int) [][]int {

	if len(buildings) <= 1 {
		return [][]int{
			{buildings[0][0], buildings[0][2]},
			{buildings[0][1], 0},
		}
	}
	midIndex := len(buildings) / 2
	left := getSkylineImpl(buildings[:midIndex])
	right := getSkylineImpl(buildings[midIndex:])
	i, j := 0, 0
	// left 数组当前高度
	leftH := 0
	// right 数组当前高度
	rightH := 0
	// 当前高度的x坐标
	x := 0
	// 之前最高点
	pre := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := [][]int{}
	for i < len(left) && j < len(right) {
		if left[i][0] < right[j][0] {
			leftH = left[i][1]
			x = left[i][0]
			i++
		} else if right[j][0] < left[i][0] {
			rightH = right[j][1]
			x = right[j][0]
			j++
		} else {
			leftH = left[i][1]
			rightH = right[j][1]
			x = left[i][0]
			i++
			j++
		}
		maxH := max(leftH, rightH)
		if maxH != pre {
			result = append(result, []int{x, maxH})
			pre = maxH
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result

}
