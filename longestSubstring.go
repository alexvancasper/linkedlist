func lengthOfLongestSubstring(s string) int {
if len(s) == 1 {
		return 1
	}
	if len(s) == 0 {
		return 0
	}
	totalLen := len(s)
	start := -1
	end := -1
	idx := 0
	left := 0
	maxLen := 0
	empty := true
	hash := make(map[byte]bool, totalLen)
	for left < totalLen {
		if idx == totalLen {
			end = idx
			break
		}
		if _, ok := hash[s[idx]]; !ok {
			if empty {
				start = idx
			} else {
				end = idx
			}
			hash[s[idx]] = true
			empty = false
			idx++
		} else {
			end = idx
			if end-start > maxLen {
				maxLen = end - start
			}
			hash = make(map[byte]bool, totalLen-idx)
			empty = true
			left++
			idx = left
		}
	}
	if end-start > maxLen {
		maxLen = end - start
	}
	return maxLen
}