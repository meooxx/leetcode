package main

import "strings"

// every "#", minus a slot
// every none '#', add a slot
//    
//       9(next)
//  slot1 slot2

// next node is 9, it occupies a slot , and create two slot
// so the final slot is 2 - 1, gain a slot
// if next node is null, it occupies a slot, 


func isValidSerialization(preorder string) bool {
	slot := 1
	nodes := strings.Split(preorder, ",")
	for i := range nodes {
		if slot == 0 {
			return false
		}
		if nodes[i] == "#" {
			slot -= 1
		} else {
			slot += 1
		}
	}
	return slot == 0
}
