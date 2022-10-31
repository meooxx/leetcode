package main

func candy(ratings []int) int {

	candyOut := make([]int, len(ratings))
	for i := range candyOut {
		candyOut[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyOut[i] = candyOut[i-1] + 1
		}
	}

	// case:                
	// candyOut: 1 2 ?       
	//           1 2 1 1 1 1 1
	//                   3 2 1	<- 
	// ratings:  1 8 6 5 4 3 2
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candyOut[i] <= candyOut[i+1] {
				candyOut[i] = candyOut[i+1] + 1
			}
		}
	}
	sum := 0
	for i := range candyOut {
		sum += candyOut[i]
	}
	return sum

}
