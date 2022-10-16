package problems

func lengthOfLongestSubstringQ3(s string) int {
	ans, start := 0, 0
	m := make(map[string]int)
	for end := 0; end < len(s); end++ {
		if val, ok := m[string(s[end])]; ok {
			start = max(start, val+1)
		}
		m[string(s[end])] = end
		ans = max(ans, end-start+1)
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
