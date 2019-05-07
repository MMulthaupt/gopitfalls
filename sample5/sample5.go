

package stringutil

func ShortString(s string, leadingCount int, trailingCount int) string {
	maxLen := leadingCount + trailingCount + 3
	sLen := len(s)
	if sLen > maxLen {
		s = s[:leadingCount] + "..." + s[sLen-trailingCount:]
	}
	return s
}
