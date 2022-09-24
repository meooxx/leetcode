// -41 ==> 41
// ____31xxx =>31
//   000-32 ==> 0
function myAtoi(s: string): number {
	let flag: '+' | '-' | '' = ''; // +-
	let shouldBreak = false; // 已经检测到至少一个合法0-9 || +-
	let num: number = 0;
	for (let i = 0; i < s.length; i++) {
		if (s[i] === ' ') {
			if (shouldBreak) {
				break;
			}
			continue;
		}

		if (s[i] === '+' || s[i] === '-') {
			if (shouldBreak) {
				break;
			}
			flag = s[i] as '+' | '-';
			shouldBreak = true;
			continue;
		}
		// 0-9 ASCII 编码
		if (s.charCodeAt(i) >= 48 && s.charCodeAt(i) <= 57) {
			// r += s[i];
			num = num * 10 + Number(s[i]);
			shouldBreak = true;
			continue;
		}
		// other character break
		break;
	}
	num = flag === '-' ? -num : num;

	const maxInt32 = Math.pow(2, 31);
	if (num > maxInt32 - 1 || num < -maxInt32) {
		return num > 0 ? maxInt32 - 1 : -maxInt32;
	}
	return num;
}

// console.log(myAtoi('-43'));
console.log(myAtoi(' w322wwww'));

myAtoi('-43');
myAtoi('    322wwww');

