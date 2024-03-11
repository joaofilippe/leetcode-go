package longest

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 { // handle only 1 element
        return strs[0]
    }
    
    p := strs[0]
	var i int
    for _, s := range strs {
        for i = 0 ; i < len(s) && i < len(p) && p[i] == s[i]; i++ {}
        p = p[:i]
    }
    return p
}
