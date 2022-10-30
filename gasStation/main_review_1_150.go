package main

//  1, 2, 3, 4, 5
//  3, 4, 5, 1, 2
// -2 -2 -2  3  3
// 上次剩余加本次消耗 > 0 继续
// 否则 尝试从下一个站点开始 start = i + 1
// tank = gas[i] - cost[i]
// tank < 0 -> tank = 0, start = i+1
func canCompleteCircuit(gas []int, cost []int) int {
	total := 0

	tank := 0
	start := 0
	for i := range gas {
		tank += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if tank < 0 {
			tank = 0
			start = i + 1
		}
	}
	if total < 0 {
		return -1
	}
	return start
}
