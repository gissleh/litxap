package litxapfilter

import "strings"

func DemoteEjectivesBeforeConsonants(curr, next *FilterTarget) (*string, *string) {
	if next == nil {
		return nil, nil
	}

	// Keep this within clauses as a even a small break ought to give space for it.
	after := strings.Trim(curr.After, "  \t\r")
	if after != "" {
		return nil, nil
	}

	for _, ejective := range ejectives {
		if _, has := hasSuffixFold(curr.Syllable, ejective); has {
			// Allow repeating ejectives
			if prefix := findPrefix(next.Syllable, ejectivesAndVoiced); prefix != nil {
				return nil, nil
			}

			// Omit the ejective if it's any of these
			if prefix := findPrefix(next.Syllable, consonantOnsets); prefix != nil {
				newCurr := curr.Syllable[:len(curr.Syllable)-len("x")]
				return &newCurr, nil
			}
		}
	}

	return nil, nil
}

func RemoveRepeatedEjective(curr, next *FilterTarget) (*string, *string) {
	if next == nil {
		return nil, nil
	}

	// Keep this within clauses as a even a small break ought to give space for it.
	after := strings.Trim(curr.After, "  \t\r")
	if after != "" {
		return nil, nil
	}

	for i, ejective := range ejectives {
		if _, has := hasSuffixFold(curr.Syllable, ejective); has {
			// Omit double ejectives. They are pronounced differently ("srätx txo" has a long tx)
			if prefix := findPrefix(next.Syllable, []string{ejective, ejectivesVoiced[i]}); prefix != nil {
				newNext := strings.TrimPrefix(next.Syllable, *prefix)
				return nil, &newNext
			}

			break
		}
	}

	return nil, nil
}

var ejectives = []string{"px", "kx", "tx"}
var ejectivesVoiced = []string{"b", "g", "d"}
var ejectivesAndVoiced = append(ejectives, ejectivesVoiced...)
var consonantOnsets = []string{"f", "k", "ng", "n", "m", "p", "t", "s"}
