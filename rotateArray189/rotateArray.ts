/**
 Do not return anything, modify nums in-place instead.
 */
function rotate(nums: number[], k: number): void {
	if (nums.length === 1) return
	k %= nums.length


	const reverse = (arr: number[], start: number, end: number) => {
		while (start < end) {
			[arr[start], arr[end]] = [arr[end], arr[start]];
			start++
			end--
		}
	}
	nums.reverse()
	reverse(nums, 0, k - 1)
	reverse(nums, k, nums.length - 1)
};