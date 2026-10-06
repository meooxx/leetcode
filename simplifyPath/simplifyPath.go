package main

import "strings"

func simplifyPath(path string) string {

	pArr := strings.Split(path, "/")
	answer := []string{}
	for _, p := range pArr {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(answer) > 0 {
				answer = answer[:len(answer)-1]
			}
			continue
		}
		answer = append(answer, p)
	}
	return "/" + strings.Join(answer, "/")

}
