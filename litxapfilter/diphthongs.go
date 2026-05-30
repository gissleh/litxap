package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// DiphthongFromWeakVowel turns syllables like a.u*, e.u*, a.ù*, e.ù*, a.i*, e.i*, a.ì*, e.ì* into
// one syllable aw*, ew*, ay*, or ey*. The * here means any coda would follow. Examples
// KA.me.i.e => KA.mey.e, TSKXE.ma.u.ti => TSKXE.maw.ti and RU.ma.ut => RU.mawt.
//
// You should follow it up with ReanalyzeDiphthongs to make the syllables more standard.
func DiphthongFromWeakVowel(curr, next *FilterTarget) (*string, *string) {
	if next == nil || curr.Stressed || next.Stressed {
		return nil, nil
	}

	// Don't do this across any kind of breaks.
	if strings.Trim(curr.After, "  \t\r") != "" {
		return nil, nil
	}

	// *ka.me.i.e = ka.mey.e
	if !curr.Stressed && !next.Stressed {
		plr, _ := utf8.DecodeLastRuneInString(curr.Syllable)
		plr = unicode.ToLower(plr)
		if plr == 'a' || plr == 'e' {
			if iVariant := findPrefix(next.Syllable, iVariants); iVariant != nil {
				y := "y"
				if *iVariant == "I" || *iVariant == "Ì" {
					y = "Y"
				}

				currChange := curr.Syllable + y + next.Syllable[len(*iVariant):]
				nextChange := ""
				return &currChange, &nextChange
			}
			if uVariant := findPrefix(next.Syllable, uVariants); uVariant != nil {
				w := "w"
				if *uVariant == "U" || *uVariant == "Ù" {
					w = "W"
				}

				currChange := curr.Syllable + w + next.Syllable[len(*uVariant):]
				nextChange := ""
				return &currChange, &nextChange
			}
		}
	}

	return nil, nil
}

// ReanalyzeDiphthongs moves the w or y from a diphthong into the next syllable if the next syllable
// can accept an onset.
func ReanalyzeDiphthongs(curr, next *FilterTarget) (*string, *string) {
	if next == nil {
		return nil, nil
	}

	// Don't do this across any kind of breaks.
	if strings.Trim(curr.After, "  \t\r") != "" {
		return nil, nil
	}

	nfr, _ := utf8.DecodeRuneInString(next.Syllable)
	nfr = unicode.ToLower(nfr)
	if !slices.Contains(vowels, nfr) {
		return nil, nil
	}

	suffix := findSuffix(curr.Syllable, diphthongVariants)
	if suffix == nil {
		return nil, nil
	}

	_, runeLength := utf8.DecodeLastRuneInString(*suffix)
	secondLetterInDiphthong := curr.Syllable[len(curr.Syllable)-runeLength:]

	currChange := strings.TrimSuffix(curr.Syllable, secondLetterInDiphthong)
	nextChange := secondLetterInDiphthong + next.Syllable
	return &currChange, &nextChange
}

var diphthongVariants = []string{"ay", "ey", "aw", "ey"}
var vowels = []rune{'i', 'ì', 'ù', 'u', 'e', 'é', 'o', 'ô', 'a', 'ä'}
var iVariants = []string{"i", "ì"}
var uVariants = []string{"u", "ù"}
