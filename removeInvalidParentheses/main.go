package main

// ( ) )
//
//	) ) ,   // first level
//
// (   )     	 at first finded (minium step) and stop the procedure
// ( )       	 dupalicate
//
//	      // second level
//	  )
//	)
//
// (   ) ...
func removeInvalidParentheses0(s string) []string {
	isValid := func(s string) bool {
		stack := 0
		for i := range s {
			if s[i] == '(' {
				stack++
			}
			if s[i] == ')' {
				if stack == 0 {
					return false
				}
				stack--
			}
		}
		return stack == 0
	}

	queue := []string{s}
	visited := map[string]bool{}
	anwser := []string{}
	found := false
	for len(queue) != 0 {
		qSize := len(queue)
		for i := 0; i < qSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			if isValid(cur) {
				anwser = append(anwser, cur)
				found = true
			}
			if found {
				continue
			}
			for j := range cur {
				if cur[j] != '(' && cur[j] != ')' {
					continue
				}
				newS := cur[:j] + cur[j+1:]
				// () ())
				// )  ()
				// (  ()
				//    ()
				//    ()
				if !visited[newS] {
					queue = append(queue, newS)
					visited[newS] = true
				}

			}
		}
	}
	return anwser

}

// () )   (()
func removeInvalidParentheses(s string) []string {
	var remove func(s string, lastIndex, lastInvalid int, p0, p1 byte)
	anwser := []string{}
	remove = func(s string, lastIndex, lastInvalid int, p0, p1 byte) {
		stack := 0
		for i := lastIndex; i < len(s); i++ {
			if s[i] == p0 {
				stack++
			}
			if s[i] == p1 {
				stack--
			}
			// ((() stack > 0
			// ())) stack < 0
			if stack >= 0 {
				continue
			}
			for j := lastInvalid; j < len(s); j++ {
				//                  第一个就是非法字符或者重复的闭括号
				//                        )   || ))
				if s[j] == p1 && (j == lastInvalid || s[j-1] != p1) {
					remove(s[:j]+s[j+1:], i, j, p0, p1)
				}
			}
			return
		}
		// if (() -> )((
		bts := []byte(s)
		for i, j := 0, len(bts)-1; i < j; {
			bts[i], bts[j] = bts[j], bts[i]
			i++
			j--
		}
		// case: (()
		if p0 == '(' {
			remove(string(bts), 0, 0, ')', '(')
		} else {
			anwser = append(anwser, string(bts))
		}
	}
	remove(s, 0, 0, '(', ')')
	return anwser
}

func main() {
	removeInvalidParentheses("))()")
}
