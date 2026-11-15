function canCompleteCircuit(gas: number[], cost: number[]): number {
	let tankSum = 0

	// check if there is answer by computing the total gas
	for (let i = 0; i < gas.length; i++) {
		tankSum += gas[i] - cost[i]
	}
	if (tankSum < 0) return -1
	let start = 0
	tankSum = 0
	// find the answer 
	for (let i = 0; i < gas.length; i++) {
		tankSum += gas[i] - cost[i]
		// no sufficient gas from previous station, move to start station
		if (tankSum < 0) {
			tankSum = 0
			start = i + 1
		}
	}
	return start

};