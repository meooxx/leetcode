
function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
	const sorted: number[] = []
	let index1 = 0
	let index2 = 0
	while (index1 < nums1.length || index2 < nums2.length) {
		if (index1 >= nums1.length) {
			sorted.push(...nums2.slice(index2))
			break
		}
		if (index2 >= nums2.length) {
			sorted.push(...nums1.slice(index1))
			break
		}

		if (nums1[index1] < nums2[index2]) {
			sorted.push(nums1[index1])
			index1++
		} else {
			sorted.push(nums2[index2])
			index2++
		}
	}
	if ((sorted.length & 1) === 1) {
		return sorted[Math.floor(sorted.length / 2)]
	}
	return (sorted[sorted.length / 2] + sorted[sorted.length / 2 - 1]) / 2

}


console.log(findMedianSortedArrays([1, 2], [3, 4]));
console.log(findMedianSortedArrays([1], [2, 3]));

