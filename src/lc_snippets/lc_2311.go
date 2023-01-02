package lc_snippets

import "strings"

func longestSubsequence(s string, k int) int {
	cnt, val, pow := 0, 0, 1
	for i := len(s) - 1; i >= 0 && val+pow <= k; i-- {
		if s[i] == '1' {
			val += pow
			cnt++
		}
		pow <<= 1
	}

	return strings.Count(s, "0") + cnt
}
