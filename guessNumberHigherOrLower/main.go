package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int

func guessNumber(n int) int {
	left := 1
	right := n
	for left <= right {
		mid := (right-left)/2 + left
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1

}
