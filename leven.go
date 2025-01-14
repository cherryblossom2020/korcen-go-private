package korcen

import (
	"unicode"
)

func SplitHangul(s string) []rune {
	var result []rune
	for _, r := range s {
		if unicode.Is(unicode.Hangul, r) {
			code := r - 0xAC00
			chun := code / (21 * 28)
			jung := (code % (21 * 28)) / 28
			jong := code % 28
			if chun != 0 {
				result = append(result, rune(chun+0x1100))
			}
			if jung != 0 {
				result = append(result, rune(jung+0x1161))
			}
			if jong != 0 {
				result = append(result, rune(jong+0x11A7))
			}
		} else {
			result = append(result, r)
		}
	}
	return result
}

func Levenshtein(a, b string) int {
	m := len(SplitHangul(a))
	n := len(SplitHangul(b))
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+cost))
		}
	}
	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
