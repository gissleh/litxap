package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

func DropGlottalStopBetweenUnlikeVowels(curr, next *FilterTarget) (*string, *string) {
	return dropGlottalStop(curr, next, true)
}

func dropGlottalStop(curr, next *FilterTarget, unlikeOnly bool) (*string, *string) {
	if next == nil || !strings.HasPrefix(next.Syllable, "'") {
		return nil, nil
	}

	// This cannot happen between words.
	if next.SyllableIndex == 0 {
		return nil, nil
	}

	clr, _ := utf8.DecodeLastRuneInString(curr.Syllable)
	nsr, _ := utf8.DecodeRuneInString(next.Syllable[len("'"):])
	if unlikeOnly && unicode.ToLower(clr) == unicode.ToLower(nsr) {
		return nil, nil
	}
	if !slices.Contains(vowels, unicode.ToLower(clr)) || !slices.Contains(vowels, unicode.ToLower(nsr)) {
		return nil, nil
	}

	newNext := next.Syllable[len("'"):]
	return nil, &newNext
}
