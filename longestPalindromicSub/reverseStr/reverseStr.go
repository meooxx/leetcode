package reverseStr
import "fmt"

func Reverse(s string) string {
	var str string
	for i:=len(s)-1; i>=0; i-- {
		str= fmt.Sprint(str, string(s[i]))
	}
	return string(str)
}