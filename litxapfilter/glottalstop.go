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
	if next == nil {
		return nil, nil
	}

	inCurr := strings.HasSuffix(curr.Syllable, "'")
	inNext := strings.HasPrefix(next.Syllable, "'")
	if !inCurr && !inNext {
		return nil, nil
	}

	// This cannot happen between words.
	if next.SyllableIndex == 0 {
		return nil, nil
	}

	var withoutStop string
	var clr, nsr rune
	if inCurr {
		withoutStop = curr.Syllable[:len(curr.Syllable)-len("'")]
		clr, _ = utf8.DecodeLastRuneInString(withoutStop)
		nsr, _ = utf8.DecodeRuneInString(next.Syllable)
	} else {
		withoutStop = next.Syllable[len("'"):]
		clr, _ = utf8.DecodeLastRuneInString(curr.Syllable)
		nsr, _ = utf8.DecodeRuneInString(withoutStop)
	}
	clr = unicode.ToLower(clr)
	nsr = unicode.ToLower(nsr)

	if unlikeOnly && clr == nsr {
		return nil, nil
	}
	if !slices.Contains(vowels, clr) || !slices.Contains(vowels, nsr) {
		return nil, nil
	}

	if inCurr {
		return &withoutStop, nil
	}
	return nil, &withoutStop
}
