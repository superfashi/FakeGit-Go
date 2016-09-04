package fakegit

func IsIn(s string, n []string) bool {
	for _, k := range n {
		if s == k {
			return true
		}
	}
	return false
}

func IndexOf(s string, n []string) int {
	for i, k := range n {
		if s == k {
			return i
		}
	}
	return -1
}

func Pop(s int, n *[]string) string {
	ret := *n[s]
	*n = append(*n[:s], *n[s+1:])
	return ret
}
