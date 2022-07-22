/*
 [1 2 3 9]  sum 8 - no
 [1 2 4 4]  sum 8 - yes
*/

// type Node struct {
// 	Next *Node
// 	Val  string
// }

func hasFoundPair(input []int, target int) bool {
	comp := make(map[int]bool, len(input)-1)

	for _, value := range input {
		if value < target {
			d := target - value
			if _, ok := comp[d]; ok {
				return true
			} else {
				comp[value] = true
			}
		}
	}
	return false
}