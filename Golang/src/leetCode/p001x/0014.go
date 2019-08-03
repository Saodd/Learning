package p001x

func longestCommonPrefix(strs []string) string {
	var l int = len(strs)
	if l == 0  {
		return ""
	}
	var pos int = 0
	for c, first, lfirst := byte(0), strs[0], len(strs[0]); pos < lfirst; pos++ {
		c = first[pos]
		for i := 1; i < l; i++ {
			if pos > len(strs[i])-1 || c != strs[i][pos] {
				goto STOPLOOP
			}
		}
	}
STOPLOOP:
	return strs[0][0:pos]
}
