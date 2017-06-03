package leetcode_go_mayulu

func longestPalindrome(s string) string {
	start := 0
	maxlen := 0
	length := len(s)

	if length < 2 {
		return s
	}

	for i := 0; i < length-1; i++ {
		start, maxlen = extendPalindrome(s, i, i, start, maxlen)
		start, maxlen = extendPalindrome(s, i, i+1, start, maxlen)
	}

	return string([]byte(s)[start : start+maxlen])

}

func extendPalindrome(s string, j int, k int, lo int, max int) (int, int) {
	for j >= 0 && k < len(s) && (s[j] == s[k]) {
		j -= 1
		k += 1
	}

	if max < (k - j - 1) {
		lo = j + 1
		max = k - j - 1
	}
	return lo, max
}
