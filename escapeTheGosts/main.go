package main	

import (
	"fmt"
	"math"
)

// (a + b)^2 > (a1 + b1)^2

// exp1:(a+b)^2 -2ab > exp2:(a1 + b1)^2 -2a1b1

//  sqrt((exp1)) > sqrt(exp2)

// if a + b > a1 + b1, then ac > a1c1







func main() {
	gosts1 := [][]int{[]int{1, 0}}
	gosts2:= [][]int{[]int{1, 0}, []int{0, 3}}
	gosts3:= [][]int{[]int{2, 0}}


	fmt.Println(escapeGhosts(gosts1, []int{2 ,0}))
	fmt.Println(escapeGhosts(gosts2, []int{0 ,1}))
	fmt.Println(escapeGhosts(gosts3, []int{1 ,0}))
	
}


func escapeGhosts(gosts [][]int, target[]int)  bool{
	for _,v := range gosts {
		if math.Abs(float64(v[0] - target[0])) + math.Abs(float64(v[1]-target[1])) <= math.Abs(float64(target[0])) +math.Abs(float64(target[1])) {
			return false
		} 
	}

	return true
}