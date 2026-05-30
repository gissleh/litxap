package litxapfilter

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// NasalAssimilation is a Filter that replaces a nasal at the end of the current syllable
// if there is a more easily pronounced one for the onset of the next. In Na'vi, the incorrect
// nasal is often kept in the spelling to make the etymology more clear.
func NasalAssimilation(curr, next *FilterTarget) (*string, *string) {
	if next == nil {
		return nil, nil
	}

	// Keep nasal assimilation within sentences.
	after := strings.Trim(curr.After, "  \t\r")
	if after != "" && after != "," {
		return nil, nil
	}

	// TS does not trigger nasal assimilation
	if _, has := hasPrefixFold(next.Syllable, "ts"); has {
		return nil, nil
	}

	for _, nr := range nasalAssimilationTable {
		if _, has := hasPrefixFold(next.Syllable, nr[1]); has {
			for _, nasal := range nasals {
				if nasalInSyllable, has := hasSuffixFold(curr.Syllable, nasal); has {
					if slices.Contains(nasals, nr[1]) {
						// tìng nari => tì nari (omit first nasal)
						changeTo := strings.TrimSuffix(curr.Syllable, nasalInSyllable)
						return &changeTo, nil
					} else if nasal == nr[0] {
						// lumpe => lumpe (no change)
						return nil, nil
					}

					r, _ := utf8.DecodeRuneInString(nasalInSyllable)

					var changeTo string
					if unicode.IsUpper(r) {
						changeTo = strings.TrimSuffix(curr.Syllable, nasalInSyllable) + strings.ToUpper(nr[0])
					} else {
						changeTo = strings.TrimSuffix(curr.Syllable, nasalInSyllable) + nr[0]
					}
					return &changeTo, nil
				}
			}
		}
	}

	return nil, nil
}

var nasals = []string{"n", "m", "ng"}
var nasalAssimilationTable = [][2]string{
	{"ng", "ng"},
	{"ng", "kx"},
	{"ng", "k"},
	{"ng", "g"},
	{"ng", "-g"},
	{"m", "px"},
	{"m", "p"},
	{"m", "b"},
	{"m", "m"},
	//{"n", "tx"}, This assimilation doesn't seem to happen.
	//{"n", "t"},
	//{"n", "d"},
	{"n", "n"},
}
