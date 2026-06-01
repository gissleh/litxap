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
			nfr, nfrLen := utf8.DecodeRuneInString(next.Syllable)

			switch unicode.ToLower(nfr) {
			case 'i', 'ì':
				y := "y"
				if nfr == 'I' || nfr == 'Ì' {
					y = "Y"
				}

				currChange := curr.Syllable + y + next.Syllable[nfrLen:]
				nextChange := ""
				return &currChange, &nextChange
			case 'u', 'ù':
				w := "w"
				if nfr == 'U' || nfr == 'Ù' {
					w = "W"
				}

				currChange := curr.Syllable + w + next.Syllable[nfrLen:]
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

	clr, clrLen := utf8.DecodeLastRuneInString(curr.Syllable)
	clr = unicode.ToLower(clr)
	if clr != 'y' && clr != 'w' {
		return nil, nil
	}

	// Ay.ye.rik => A.ye.rik
	nfr, _ := utf8.DecodeRuneInString(next.Syllable)
	nfr = unicode.ToLower(nfr)
	if clr == nfr {
		currChange := curr.Syllable[:len(curr.Syllable)-clrLen]
		return &currChange, nil
	}
	if !slices.Contains(vowels, nfr) {
		return nil, nil
	}

	currChange := curr.Syllable[:len(curr.Syllable)-clrLen]
	nextChange := curr.Syllable[len(curr.Syllable)-clrLen:] + next.Syllable
	return &currChange, &nextChange
}

var vowels = []rune{'i', 'ì', 'ù', 'u', 'e', 'é', 'o', 'ô', 'a', 'ä'}
