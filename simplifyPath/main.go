package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/.../"))

	fmt.Println(simplifyPath("/../..ga/b/.f..d/..../e.baaeeh./.a"))

	// a := strings.Split("/a/b/c", "/")

}

func simplifyPath(path string) string {
	p := strings.Split(path, "/")
	result := []string{}

	for i := 1; i < len(p); i++ {
		switch true {
		case p[i] == "":
		case p[i][0] == '.':
			if len(p[i]) >= 2 {
				if p[i][1] == '.' && len(p[i]) == 2 {
					// must contain the root
					if len(result) > 0 {
						result = result[:len(result)-1]
					}
				} else {
					result = append(result, p[i])
				}
			}
		default:
			result = append(result, p[i])
		}

	}

	return "/" + strings.Join(result, "/")
}
