function multiply(num1: string, num2: string): string {

	const n1 = num1.length
	const n2 = num2.length
	const digits: number[] = Array(n1 + n2).fill(0)

	const zero = '0'.charCodeAt(0)

	//    9 9 9
	//    9 9 9

	//      8 1
	//    8 1
	//  8 1 
	//    8 1  
	for (let i = n1 - 1; i >= 0; i--) {
		for (let j = n2 - 1; j >= 0; j--) {
			const sum = digits[i + j + 1] + (num1[i].charCodeAt(0) - zero) * (num2[j].charCodeAt(0) - zero)
			digits[i + j + 1] = (sum % 10)
			digits[i + j] += (~~(sum / 10))
		}
	}
	while (digits[0] === 0) {
		digits.shift()
	}
	return digits.join('')
};

multiply('12', '12')