package main

/**

case 1:               case 2                     case 3:                         case 4: fallback to case 1
			 i-2                 i-3                  i-4 													     
			 _____              ______  +             ________												------ 
			|			|	           |      | |  i-4       |				| i        						 |         ______
	i-1	|			|i-3     i-2 |      | +        i-3 |        |--------|             |        |     | 
			|_____|__ i		     |      | |  i         |							   | i-1				 |		    |____ |
						|							i-1	 								 |_________________|             |______________|
						                                       i-2
*/ 

func isSelfCrossing(distance []int) bool {

	if len(distance) < 4 {
		return false
	}
	for l := 3; l < len(distance); l++ {
		if l >= 3 && distance[l] >= distance[l-2] &&
			distance[l-1] <= distance[l-3] {
			return true
		}

		if l >= 4 && distance[l-1] == distance[l-3] &&
			distance[l-2] <= distance[l]+distance[l-4] {
			return true
		}

		if l >= 5 && distance[l-2] >= distance[l-4] && distance[l-3] >= distance[l-1] && distance[l-1]+distance[l-5] >= distance[l-3] && distance[l]+distance[l-4] >= distance[l-2] {
			return true
		}
	}

	return false

}
