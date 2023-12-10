package main


type valWithIndex struct {
	val         int
	originIndex int
}

func countSmaller(nums []int) []int {

	vals := make([]valWithIndex, len(nums))
	for i := range nums {
		vals[i] = valWithIndex{
			val:         nums[i],
			originIndex: i,
		}
	}
	anwser := make([]int, len(nums))

	mergeSort(vals, 0, len(nums)-1, anwser)
  return anwser
}

func mergeSort(vals []valWithIndex, start, end int, anwser []int) {
	if start >= end {
		return
	}

	mid := (end+start)/2
	mergeSort(vals, start, mid, anwser)
	mergeSort(vals, mid+1, end, anwser)

	i := start
	j := mid + 1
	merged := []valWithIndex{}
	numsOfValLessLeft := 0
	// 1 5
	// 2 4
	for i <= mid && j <= end {
		if vals[i].val > vals[j].val {
			merged = append(merged, vals[j])
			numsOfValLessLeft++
			j++
		} else {
			merged = append(merged, vals[i])
			anwser[vals[i].originIndex] += numsOfValLessLeft
      i++
		}
	}
	for i <= mid {
		merged = append(merged, vals[i])
    anwser[vals[i].originIndex] += numsOfValLessLeft
		i++
	}
	for j <= end {
		merged = append(merged, vals[j])
		j++
	}
	for i := range merged {
		vals[start+i] = merged[i]
	}
}
