package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

/*
  These filters do not a translator make because:
    1. The ù/u distinction is not currently possible to make happen with filters
    2. There are vocabulary differences, like shawm (RN) vs omum (FN)
    3. There are grammatical differences, like topical being allowed at the end of the sentence in RN
*/

func ReefUnstressedAeAsE(curr, _ *FilterTarget) (*string, *string) {
	if curr.Stressed {
		return nil, nil
	}

	newCurr := aeToEReplacer.Replace(curr.Syllable)
	if newCurr == curr.Syllable {
		return nil, nil
	}

	return &newCurr, nil
}

func ReefEjectiveToVoiced(curr, next *FilterTarget) (*string, *string) {
	newCurr := curr.Syllable

	for _, ejective := range ejectives {
		// Check for syllable-initial
		if ejectiveInSyllable, has := hasPrefixFold(newCurr, ejective); has {
			efr, _ := utf8.DecodeRuneInString(ejectiveInSyllable)
			newCurr = voicedConsonantMap[efr] + newCurr[len(ejectiveInSyllable):]
		}

		// Change at the end if immediately followed by another ejective.
		if ejectiveInSyllable, has := hasSuffixFold(newCurr, ejective); has && next != nil && next.SyllableIndex != 0 {
			if prefix := findPrefix(next.Syllable, voiceCodaIf[ejectiveInSyllable]); prefix != nil {
				efr, _ := utf8.DecodeRuneInString(ejectiveInSyllable)
				newCurr = newCurr[:len(newCurr)-len(ejectiveInSyllable)] + voicedConsonantMap[efr]
			}
		}
	}

	if newCurr == curr.Syllable {
		return nil, nil
	}

	return &newCurr, nil
}

func ReefDropGlottalStopsBetweenVowels(curr, next *FilterTarget) (*string, *string) {
	if next == nil || !strings.HasPrefix(next.Syllable, "'") {
		return nil, nil
	}

	// This cannot happen between words.
	if next.SyllableIndex == 0 {
		return nil, nil
	}

	clr, _ := utf8.DecodeLastRuneInString(curr.Syllable)
	nsr, _ := utf8.DecodeRuneInString(next.Syllable[len("'"):])
	if !slices.Contains(vowels, unicode.ToLower(clr)) || !slices.Contains(vowels, unicode.ToLower(nsr)) {
		return nil, nil
	}

	newNext := next.Syllable[len("'"):]
	return nil, &newNext
}

func ReefApplyChSh(curr, _ *FilterTarget) (*string, *string) {
	prefix := findPrefix(curr.Syllable, []string{"sy", "tsy"})
	if prefix == nil {
		return nil, nil
	}

	cfr, cfrLen := utf8.DecodeRuneInString(*prefix)
	firstLetter := ""
	switch cfr {
	case 't':
		firstLetter = "c"
	case 'T':
		firstLetter = "C"
	case 's':
		firstLetter = "s"
	case 'S':
		firstLetter = "S"
	}

	csr, _ := utf8.DecodeRuneInString((*prefix)[cfrLen:])
	secondLetter := ""
	switch csr {
	case 's', 'y':
		secondLetter = "h"
	case 'S', 'Y':
		secondLetter = "H"
	}

	newCurr := firstLetter + secondLetter + curr.Syllable[len(*prefix):]
	return &newCurr, nil
}

var aeToEReplacer = strings.NewReplacer("ä", "e", "Â", "E")

var voiceCodaIf = map[string][]string{
	"kx": {"px", "tx", "d", "p"},
	"px": {"kx", "tx", "g", "p"},
	"tx": {"kx", "px", "g", "b"},
}

var voicedConsonantMap = map[rune]string{
	'k': "g", 'K': "G",
	'p': "b", 'P': "B",
	't': "d", 'T': "D",
}
