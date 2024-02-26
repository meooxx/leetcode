package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	left := max(ax1, bx1)
	right := min(ax2, bx2)

	width := right - left
	top := min(ay2, by2)
	bottom := max(by1, ay1)

	height := top - bottom
	area := width * height
	total := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	if width < 0 || height < 0 {
		area = 0
	}
	return total - area

}
