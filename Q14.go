package main

func Q14LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	mixPrefix := 0
	for {
		for i := 1; i < len(strs); i++ {
			if mixPrefix >= len(strs[i-1]) || mixPrefix >= len(strs[i]) || strs[i-1][mixPrefix] != strs[i][mixPrefix] {
				return strs[0][:mixPrefix]
			}
		}
		mixPrefix++
	}
	return strs[0][:mixPrefix]
}
