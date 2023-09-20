package main


// ax + by = m * d
// x, y for pour or fill water
// d is max gcd of a, b

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {

	if jug1Capacity == targetCapacity || jug2Capacity == targetCapacity || jug1Capacity+jug2Capacity == targetCapacity {
		return true
	}

	return targetCapacity % GCD(jug1Capacity, jug2Capacity) == 0

}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}