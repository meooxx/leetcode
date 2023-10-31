package main

/**
	 								1  -5   2     range [-2, 2]
 	sums 				0  	1  -4  -2
	mergeSort     [-4, 1], [-2]
	//             i         j
  //            find the first j that satisfied sum[j] - sum[i] >= lower
	//            find the first k that satisfied sum[k] - sum[i] > upper
	//            counts = k-j         
	//            sum all (k,j)
	1 why sums is sorted and works fine?
		we use the sum of the range, 
		it's already be caculated before being used, so the sort is ok
*/
func countRangeSum(nums []int, lower int, upper int) int {
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = nums[i] + sum[i]
	}
	return mergeSortAndCount(nums, 0, len(sum)-1, lower, upper, sum)
}

func mergeSortAndCount(nums []int, start, end, lower, upper int, sum []int) int {
	if start >= end {
		return 0
	}
	midIndex := (end-start)/2 + start
	leftCount := mergeSortAndCount(sum, start, midIndex, lower, upper, sum)
	rightCount := mergeSortAndCount(sum, midIndex+1, end, lower, upper, sum)

	count := leftCount + rightCount

	lessIndex := midIndex + 1
	leftIndex := midIndex + 1
	rightIndex := midIndex + 1
	orderedSums := make([]int, end-start+1)
	index := 0

	for i := start; i <= midIndex; i++ {
		for leftIndex <= end && sum[leftIndex]-sum[i] < lower {
			leftIndex++
		}

		for rightIndex <= end && sum[rightIndex]-sum[i] <= upper {
			rightIndex++
		}
		for lessIndex <= end && sum[lessIndex] < sum[i] {
			orderedSums[index] = sum[lessIndex]
			lessIndex++
			index++
		}
		count += rightIndex - leftIndex
		orderedSums[index] = sum[i]
		index++
	}
	for lessIndex <= end {
		orderedSums[index] = sum[lessIndex]
		index++
		lessIndex++
	}
	for i := range orderedSums {
		sum[start+i] = orderedSums[i]
	}
	return count
}
