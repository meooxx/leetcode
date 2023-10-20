package main

import "strings"

func isValidSerialization(preorder string) bool {
	slot := 1
	nodes := strings.Split(preorder, ",")
	for i := range nodes {
		if slot == 0 {
			return false
		}
		if nodes[i] == "#" {
			slot -= 1
		}else {
			slot += 1 
		}
	}
	return slot == 0
}
