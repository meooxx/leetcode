/**
 Do not return anything, modify nums in-place instead.
 */
// 4 3 2 8 7 4 2
//       8 7 4 2, find the 2 which is right prior to 8
// swap 2, 4 as 4 is first element that is great than 2
// reverse 8 7 2 2
function nextPermutation(nums: number[]): void {
	// find the longest non-increase suffix
	let left = nums.length - 1
	while (left > 0 && nums[left - 1] >= nums[left]) {
		left--
	}
	if (left === 0) {
		nums.reverse()
		return
	}
	let right = nums.length - 1
	// find the successor
	while (right > 0 && nums[right] <= nums[left - 1]) right--;
	// swap
	[nums[left - 1], nums[right]] = [nums[right], nums[left - 1]]
	let l = left
	let r = nums.length - 1
	while (l < r) {
		[nums[l], nums[r]] = [nums[r], nums[l]];
		l++
		r--
	}
};
nextPermutation([1, 3, 2])