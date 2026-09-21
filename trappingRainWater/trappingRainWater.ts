function trap(height: number[]): number {

	let leftMax = 0
	let rightMax = 0
	let left = 0
	let right = height.length - 1
	let answer = 0
	while (left < right) {
		leftMax = Math.max(leftMax, height[left])
		rightMax = Math.max(rightMax, height[right])

		if (leftMax <= rightMax) {
			answer += leftMax - height[left]
			left++
		} else {
			answer += rightMax - height[right]
			right--
		}


	}
	return answer
};