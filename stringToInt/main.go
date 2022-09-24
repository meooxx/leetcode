package main

import (
	"fmt"
	"math"
	// "strconv"
	// "./matchString"
)

// -41 ==> 41
// ____31xxx =>31
//   000-32 ==> 0
func myAtoi(s string) int {
	var flag rune
	shouldBreak := false
	num := 0
	for _, v := range s {
		if v == ' ' {
			if shouldBreak {
				break
			}
			continue
		}
		if v == '+' || v == '-' {
			if shouldBreak {
				break
			}
			flag = v
			shouldBreak = true
			continue
		}
		if v >= '0' && v <= '9' {

			num = num*10 + int(v-'0')
			if num > math.MaxInt32 || num < math.MinInt32 {
				break
			}

			shouldBreak = true
			continue
		}
		break
	}
	if flag == '-' {
		num = -num
	}
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		return math.MinInt32
	}
	return num

}

func main() {
	// golang x>>y means x / 2^x, 1>>2 == 0 , not -4
	// a := matchString.MatchStr1(" 4193 6 with words")
	// n1 := myAtoi(" 00-00005000")
	// n2 := myAtoi(fmt.Sprint(1<<31 - 1))
	// fmt.Println(n1)
	// fmt.Println(n2)


	n3 := myAtoi2("9223372036854775808")
	fmt.Println(n3)
}

func myAtoi2(s string) int {
    
    r := []int{}
    for index := range s {
        if s[index] != ' ' {
            s = s[index:]
            break
        }
    }
    neg := false
    if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
        if s[0] == '-' {
            neg = true
        }
        s = s[1:]
    }
    for index:= range s {
        if s[index] >= '0' && s[index] <= '9' {
            r = append(r, int(s[index] - '0'))
        }else {
            break
        }
    }
    anwser := 0
    min := -(1 << 31)
    max := -min-1
    for index:= range r {
        anwser = anwser * 10 + r[index]
    }
    if neg  {
        anwser = -anwser
    }
    if anwser > max {
        return max 
    }
    if anwser < min {
        return min
    }
    return anwser
    
}