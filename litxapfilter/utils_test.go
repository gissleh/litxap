package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPrefixFold(t *testing.T) {
	table := []struct {
		s              string
		prefix         string
		expectedPrefix string
		expectedOk     bool
	}{
		{"ŋœ®ſ™ŊŒ", "ŋœ", "ŋœ", true},
		{"ngay", "ng", "ng", true},
		{"lRrtok", "lrr", "lRr", true},
		{"ngop", "m", "", false},
		{"a", "ay", "", false},
		{"'a", "a", "", false},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s_%s", row.s, row.prefix), func(t *testing.T) {
			prefixInS, ok := hasPrefixFold(row.s, row.prefix)
			assert.Equal(t, row.expectedPrefix, prefixInS)
			assert.Equal(t, row.expectedOk, ok)
		})
	}
}

func TestHasSuffixFold(t *testing.T) {
	table := []struct {
		s              string
		suffix         string
		expectedSuffix string
		expectedOk     bool
	}{
		{"ŋœ®ſ™ŊŒ", "ŋœ", "ŊŒ", true},
		{"ngay", "ay", "ay", true},
		{"ngAy", "ay", "Ay", true},
		{"ngAY", "ey", "", false},
		{"e'", "te'", "", false},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s_%s", row.s, row.suffix), func(t *testing.T) {
			suffixInS, ok := hasSuffixFold(row.s, row.suffix)
			assert.Equal(t, row.expectedSuffix, suffixInS)
			assert.Equal(t, row.expectedOk, ok)
		})
	}
}
