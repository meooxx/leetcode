

// merge greater one from right to left
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
	if (n == 0) return
	let size = m + n - 1
	m--
	n--
	while (m >= 0 && n >= 0) {
		if (nums1[m] >= nums2[n]) {
			nums1[size] = nums1[m]
			size--
			m--
		} else {
			nums1[size] = nums2[n]
			n--
			size--
		}
	}
	while (m >= 0) {
		nums1[size] = nums1[m]
		size--
		m--
	}
	while (n >= 0) {
		nums1[size] = nums2[n]
		size--
		n--
	}


};