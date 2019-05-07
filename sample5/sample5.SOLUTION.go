//+build ignore

package stringutil

import (
	"unicode/utf8"
)

func ShortString(s string, leadingCount int, trailingCount int) string {
	ellipsis := "..."
	maxLen := leadingCount + trailingCount + utf8.RuneCountInString(ellipsis)
	runeCount := utf8.RuneCountInString(s)
	if runeCount > maxLen {
		firstByteIndex := RuneIndexToByteIndex(s, leadingCount)
		omitByteIndex := RuneIndexToByteIndex(s, runeCount-trailingCount)
		s = s[:firstByteIndex] + ellipsis + s[omitByteIndex:]
	}
	return s
}

func RuneIndexToByteIndex(s string, runeIndex int) int {
	currentRuneIndex := 0
	for i := range s {
		if currentRuneIndex == runeIndex {
			return i
		}
		currentRuneIndex++
	}
	if currentRuneIndex == runeIndex {
		return len(s)
	}
	return -1
}
