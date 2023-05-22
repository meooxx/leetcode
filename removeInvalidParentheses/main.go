package main
//  ( ) ) 
//    ) ) ,   // first level 
//  (   )     	 at first finded (minium step) and stop the procedure
//  ( )       	 dupalicate
//          // second level
//      )
//    )
//  (   ) ... 
func removeInvalidParentheses(s string) []string {
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
