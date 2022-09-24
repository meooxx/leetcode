package main

import (
	"fmt"
	"strings"
	"strconv"
)

// [[1,2], [3]].reduce((a,b)=>{
//  let result = []	 
//	a.forEach(i=>{
//     b.forEach(i=>{
//				result.push(`${ab}`)
//			})    
//		}
//	})
// return result
// }) 


func main () {
	a:=[][]string{
		[]string{"a", "b"},
		[]string{"c", "d"}}
	fmt.Println(reduce(&a, genKey ))
	b:="3"
	c:=letterCombinations(b)
	fmt.Println(c)
}


func genKey(a, b []string) []string{
	r:=[]string{}
	for i:=0;i<len(a);i++ {
	
		for j:=0;j<len(b);j++ {
			r = append(r, a[i]+b[j])
		}
	}
	return r

}

type MapString []string


func letterCombinations(s string) []string {
	 EMPTY := MapString{}
	if len(s) < 1 {
		return EMPTY
	}
	 dic := map[int64]MapString {
		 1:EMPTY,
		 2:MapString{"a", "b", "c"},
		 3:MapString{"d", "e", "f"},
		 4:MapString{"g", "h", "i"},
		 5:MapString{"j", "k", "l"},
		 6:MapString{"m", "n", "o"},
		 7:MapString{"p", "q", "r", "s"},
		 8:MapString{"t", "u", "v"},
		 9:MapString{"w", "x", "y", "z"},
	 }

	a:=strings.Split(s, "")
	if len(a)==1 {
		index, _:= strconv.ParseInt(a[0], 10, 0) 
		return dic[index]
	}
	array:=[][]string{}
	for i:=0;i<len(a);i++ {
		num, _:=strconv.ParseInt(a[i], 10, 0)
		arr := dic[num]
		array = append(array, arr)
	}

	r:= reduce(&array, genKey)


	return r
}


func reduce(a *[][]string,
	f func([]string, []string) []string ) []string {
	pre:=(*a)[0]
	for i:=1; i<len(*a);i++ {
		pre= f(pre, (*a)[i])
	}

	return pre
}