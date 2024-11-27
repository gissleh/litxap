package litxap

import (
	"fmt"
	"github.com/gissleh/litxap/litxaputil"
	"strings"
)

type NumberDictionary struct{}

func (n *NumberDictionary) LookupEntries(word string) ([]Entry, error) {
	res := litxaputil.ParseNumber(strings.ToLower(word))
	if res == nil {
		return nil, ErrEntryNotFound
	}

	syllables, stress, ok := res.GenerateSyllables(false)
	if !ok {
		return nil, ErrEntryNotFound
	}

	numberKind := "Number"
	if res.Ordinal {
		numberKind = "Ordinal number"
	}

	var translation string
	if res.Value < 0o10 {
		translation = fmt.Sprintf("%s %d", numberKind, res.Value)
	} else {
		translation = fmt.Sprintf("%s °%o (%d)", numberKind, res.Value, res.Value)
	}

	var prefixes, suffixes []string
	if res.Prefix != "" {
		prefixes = append(prefixes, res.Prefix)
	}
	if res.Suffix != "" {
		suffixes = append(suffixes, res.Suffix)
	}

	return []Entry{{
		Word:        strings.Join(syllables, ""),
		Translation: translation,
		Syllables:   syllables,
		Stress:      stress,
		Prefixes:    prefixes,
		Suffixes:    suffixes,
	}}, nil
}
