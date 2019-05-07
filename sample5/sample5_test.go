package stringutil_test

import (
	"testing"

	stringutil "github.com/MMulthaupt/gopitfalls/sample5"
)

func TestShortString(t *testing.T) {
	tests := []struct {
		text                        string
		leadingCount, trailingCount int
		expected                    string
	}{
		{"what a stupid text", 7, 5, "what a ... text"},
		{"what a great text", 7, 7, "what a great text"},
		{"what a great text", 6, 7, "what a...at text"},
		{"日本語偽善者", 1, 2, "日本語偽善者"},
		{"日本語偽善者", 1, 1, "日...者"},
		{"ab日本語皮を被る", 1, 2, "a...被る"},
		{"", 0, 0, ""},
		{"0123456789", 4, 0, "0123..."},
		{"0123456789", 0, 4, "...6789"},
	}
	for i, test := range tests {
		result := stringutil.ShortString(test.text, test.leadingCount, test.trailingCount)
		if result != test.expected {
			t.Errorf("Test #%d failed: ShortString(\"%s\", %d, %d) yielded %s. Expected %s.", i+1, test.text, test.leadingCount, test.trailingCount, result, test.expected)
		}
	}
}
