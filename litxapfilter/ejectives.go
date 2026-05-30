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
			// Omit double ejectives. They are pronounced differently ("srätx txo" has a long tx)
			if prefix, has := hasPrefixFold(next.Syllable, ejective); has {
				newNext := strings.TrimPrefix(next.Syllable, prefix)
				return nil, &newNext
			}

			// Allow repeating ejectives
			if prefix := findPrefix(next.Syllable, ejectives); prefix != nil {
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

var ejectives = []string{"kx", "tx", "px"}
var consonantOnsets = []string{"ng", "n", "m", "p", "t", "k", "f", "s"}
