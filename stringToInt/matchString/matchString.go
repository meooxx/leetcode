package matchString
import (
	"strings"
	"strconv"
	"regexp"
	"math"
)

// 4 | 8 ms
func MatchStr1(str string) int {
	strSlice := strings.Split(str, "")
	var s string
	for i:=0;i<len(strSlice);i++{
		v:=strSlice[i]
		if v >= "0" && v<="9" {
			s+=v
		}else if  s=="" {
			// "  -4", "+-33", "-  44", "-44 444" 
			if   v =="-" || v =="+" {
					s= v+s
			}else if v != " "  {
				//"   333"
				break
		}         
		}else {
				
			break
		}
	}
	n,_:=strconv.ParseInt(s,10, 0)
	return ValidNum(n)
}


func MatchStr (str string) int {
	
// 使用捕获组去匹配需要的内容 24ms

	reg:=regexp.MustCompile(`^[\s]*([-+]?[0-9]+)`)
	s:=reg.FindStringSubmatch(str)
	if s!=nil {
		n,_:=strconv.ParseInt(s[1],10, 64)
		return ValidNum(n)

	}
	return 0
}  

func ValidNum(n int64) int {
	if n > math.MaxInt32 {
		return math.MaxInt32
	} else if n< math.MinInt32 {
	
		return math.MinInt32
	}
	return int(n)
}