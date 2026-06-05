package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ElideUnstressedEWordEndings changes "Ki.ye.va.me ul.te Ey.wa nga.hu" to "Ki.ye.va mul tEy.wa nga.hu"
func ElideUnstressedEWordEndings(curr, next *FilterTarget) (*string, *string) {
	if curr.Stressed || next == nil || next.SyllableIndex != 0 || curr.After == "" || strings.Trim(curr.After, " ,-; \t\r\n") != "" {
		return nil, nil
	}

	clr, clrLen := utf8.DecodeLastRuneInString(curr.Syllable)
	if clr != 'e' && clr != 'E' {
		return nil, nil
	}

	nfr, _ := utf8.DecodeRuneInString(next.Syllable)
	if !slices.Contains(vowels, unicode.ToLower(nfr)) {
		return nil, nil
	}

	currChange := ""
	nextChange := curr.Syllable[:len(curr.Syllable)-clrLen] + next.Syllable
	return &currChange, &nextChange
}

// ElideMiSiNiBeforeAy combines nì, mì and sì with ay- words (2.3.7.1)
func ElideMiSiNiBeforeAy(curr, next *FilterTarget) (*string, *string) {
	if next == nil {
		return nil, nil
	}

	for i, elideSyllable := range elideSyllables {
		if strings.EqualFold(curr.Syllable, elideSyllable) {
			if elideSyllablesWord[i] != "" && strings.TrimSuffix(curr.Entry.Word, "+") != elideSyllablesWord[i] {
				return nil, nil
			}

			if _, has := hasPrefixFold(next.Syllable, "ay"); !has {
				return nil, nil
			}

			_, clrLen := utf8.DecodeLastRuneInString(curr.Syllable)
			currChange := ""
			nextChange := curr.Syllable[:len(curr.Syllable)-clrLen] + next.Syllable
			return &currChange, &nextChange
		}
	}

	return nil, nil
}

// ElideAdvPrefixAndE combines nì- with words starting with e, respecting the exception provided in Horen.. (2.3.7.2)
func ElideAdvPrefixAndE(curr, next *FilterTarget) (*string, *string) {
	// Only do it on the first syllable if it's nì and the word is not "ean"
	if next == nil ||
		curr.SyllableIndex != 0 ||
		next.SyllableIndex != 1 ||
		!strings.EqualFold(curr.Syllable, "nì") ||
		slices.ContainsFunc(elisionIAndEBlacklist, func(s string) bool {
			return strings.EqualFold(curr.Entry.Word, s)
		}) {
		return nil, nil
	}

	_, has := hasPrefixFold(next.Syllable, "e")
	if !has {
		return nil, nil
	}

	// Always 1, but we can never be too careful, right?
	_, cfrLen := utf8.DecodeRuneInString(curr.Syllable)

	if next.Stressed {
		newCurr := ""
		newNext := curr.Syllable[:cfrLen] + next.Syllable
		return &newCurr, &newNext
	}

	newCurr := ""
	newNext := curr.Syllable
	return &newCurr, &newNext
}

var elisionIAndEBlacklist = []string{"ean"}
var elideSyllables = []string{"mì", "sì", "nì"}
var elideSyllablesWord = []string{"mì", "sì", ""}
