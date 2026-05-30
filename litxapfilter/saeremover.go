package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// SaeRemover turns an unstressed sä syllable into a pre-onset s if it is possible.
// There are no examples of stressed sä syllables in the dictionary, but it may occur
// in names or in the future.
func SaeRemover(curr, next *FilterTarget) (*string, *string) {
	// Check if there's two syllables where the first is an unstressed sä
	if next == nil || curr.Stressed || !strings.EqualFold(curr.Syllable, "sä") {
		return nil, nil
	}

	// Don't do this across any breaks, though.
	if strings.Trim(curr.After, "  \t\r") != "" {
		return nil, nil
	}

	// I don't believe it can happen across word boundaries
	if curr.PartIndex != next.PartIndex {
		return nil, nil
	}

	// Only continue if the first letter can follow a pre-onset s
	r1, l1 := utf8.DecodeRuneInString(next.Syllable)
	if !slices.Contains(canFollowSae, unicode.ToLower(r1)) {
		return nil, nil
	}

	// Check for 's' since "st" is allowed, but "sts" is not.
	// "sng" and "stx" are allowed, though.
	r2, _ := utf8.DecodeRuneInString(next.Syllable[l1:])
	if r2 == 's' {
		return nil, nil
	}

	// Delete the current, expand the next.
	newCurr := ""
	newNext := curr.Syllable[:len("s")] + next.Syllable
	return &newCurr, &newNext
}

var canFollowSae = []rune{
	'p', 't', 'k', 'm', 'n', 'r', 'l', 'w', 'y',
}
