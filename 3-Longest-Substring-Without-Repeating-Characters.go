package leetcode_go_mayulu

func LengthOfLongestSubstring(s string) int {

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
			curLength = curLength + 1 - (repeated - start)
			resetStatForStringIndexRange(stats, s, start, repeated)
			stats[s[end]] = end
			start = repeated + 1
			end++
		}

	}
	return maxLength

}

func resetStatForStringIndexRange(m map[byte]int, s string, start int, repeated int) {
	size := len(s)
	for start < size && end < repeated && start <= repeated {
		delete(m, s[start])
		start++
	}

}
