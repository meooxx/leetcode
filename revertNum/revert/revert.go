package revert
import "fmt"
import "strings"
import "strconv"


func RevertNum (n int) int {

	s:=fmt.Sprint(n)
	if len(s) == 1 {
		return n
	}
	a:= strings.Split(s, "")
	var str string 
	
	if n<0 {
		a=a[1:]
	}
	for i:=0;i<len(a);i++ {
		if a[i] == "0" && i==0  {
			i ++
		}
		str = a[i] + str
	}
	
	v, _:= strconv.ParseInt(str, 10, 0 )
	if v > 1<<31-1 || -1<<31 > v {
		return 0
	}
	if n<0 {
		return int(-v)
	}
	return int(v)
}