func lengthOfLongestSubstring(s string) int {
    stats := make(map[byte]int)
	maxLength := 0
	curLength := 0
	size := len(s)
	start := 0
	end := 0
	for start < size && end < size {
		value, ok := stats[s[end]]
		if !ok {
			stats[s[end]] = end
			curLength++
			if curLength > maxLength {
				maxLength = curLength
			}
			end++
		} else {
			repeated := value
			resetStatForStringIndexRange(stats, s, start, repeated-1)
			stats[s[end]] = end
			curLength -= repeated - start
			start = repeated+1
			end++
		}
	}
	return maxLength
    
}

func resetStatForStringIndexRange(m map[byte]int, s string, start int, repeatedIndex int) {
	size := len(s)
	for start < size && end < repeatedIndex && start <= repeatedIndex {
		delete(m, s[start])
		start++
	}

}