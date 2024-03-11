package longest

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	var i int
	for _, s := range strs {
		for i < len(s) && i < len(prefix) && prefix[i] == s[i] {
			i++
		}

		prefix = prefix[:i]
	}
	
	return prefix
}
