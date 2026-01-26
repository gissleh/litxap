package litxaputil

import (
	"slices"
	"strings"
)

var preOnsets = []string{"f", "ts", "s"}
var onsetsAfterPreF = []string{"k", "kx", "l", "m", "n", "ng", "p", "px", "t", "tx", "w", "y"}
var onsetsAfterPreS = []string{"h", "k", "kx", "l", "m", "n", "ng", "p", "px", "t", "tx", "w", "y"}
var onsets = []string{"'", "h", "k", "kx", "l", "m", "n", "ng", "p", "px", "r", "t", "ts", "tx", "s", "v", "w", "y", "z", "b", "d", "g"}
var bodies = []string{"a", "ä", "e", "é", "i", "ì", "o", "õ", "u", "ù", "rr", "ll", "ay", "ey", "aw", "ew"}
var codas = []string{"'", "k", "kx", "l", "m", "n", "ng", "p", "px", "r", "t", "tx", "d", "b", "g"}

type Syllable struct {
	PreOnset  string
	Onset     string
	Irregular string
	Body      string
	Coda      string
}

// SplitSyllables uses predictable reanalysis rules to split a na'vi word into syllables. It will handle some
// irregular words (including: tlalim, mangkwan, kreytu'um) but will log their irregularities.
func SplitSyllables(s string) Syllables {
	res := make(Syllables, 0, len(s))
	for len(s) > 0 {
		var curr Syllable

		// Try adding a coda
		for _, coda := range codas {
			if strings.HasSuffix(s, coda) {
				if coda == "r" && strings.HasSuffix(s, "rr") && !strings.HasSuffix(s, "rrr") {
					break
				}
				if coda == "l" && strings.HasSuffix(s, "ll") && !strings.HasSuffix(s, "lll") {
					break
				}

				curr.Coda = coda
				s = strings.TrimSuffix(s, coda)
				break
			}
		}

		// Add a body
		foundBody := false
		for _, body := range bodies {
			if strings.HasSuffix(s, body) {
				curr.Body = body
				s = strings.TrimSuffix(s, body)
				foundBody = true
				break
			}
		}

		// Handle edge cases of colloquial pronunciations: Kreytu'um, Mangkwan, Tlalim
		if !foundBody {
			if len(res) > 0 {
				last := &res[len(res)-1]
				if last.PreOnset != "" {
					return nil
				}

				last.Irregular = last.Onset
				last.Onset = curr.Coda
				continue
			} else {
				return nil
			}
		}

		// Try adding an onset
		for _, onset := range onsets {
			if strings.HasSuffix(s, onset) {
				curr.Onset = onset
				s = strings.TrimSuffix(s, onset)
				break
			}
		}

		// Try adding a pre-onset
		for _, preOnset := range preOnsets {
			if strings.HasSuffix(s, preOnset) {
				if preOnset == "f" {
					if !slices.Contains(onsetsAfterPreF, curr.Onset) {
						return nil
					}
				} else {
					if !slices.Contains(onsetsAfterPreS, curr.Onset) {
						return nil
					}
				}

				curr.PreOnset = preOnset
				s = strings.TrimSuffix(s, preOnset)
				break
			}
		}

		// rr and ll must have an onset. Even lenition won't break this rule.
		if (curr.Body == "rr" || curr.Body == "ll") && curr.Onset == "" {
			return nil
		}

		res = append(res, curr)
	}

	// We've been working backwards, so flip it.
	slices.Reverse(res)

	return res
}

type Syllables []Syllable

func (s Syllables) Irregulars() []string {
	var res []string
	for _, syl := range s {
		if syl.Irregular != "" {
			res = append(res, syl.PreOnset+syl.Onset+syl.Irregular)
		}
	}

	return res
}

func (s Syllables) String() string {
	if s == nil {
		return "<nil>"
	}

	var res strings.Builder
	res.Grow(len(s) * 8)

	for i, syl := range s {
		if i > 0 {
			res.WriteByte('-')
		}

		for _, part := range [5]string{syl.PreOnset, syl.Onset, syl.Irregular, syl.Body, syl.Coda} {
			if len(part) > 0 {
				res.WriteString(part)
			}
		}
	}

	if irregulars := s.Irregulars(); irregulars != nil {
		res.WriteString(" (irregulars: ")
		for i, irr := range irregulars {
			if i > 0 {
				res.WriteString(",")
			}

			res.WriteString(irr)
		}
		res.WriteString(")")
	}

	return res.String()
}
