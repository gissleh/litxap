package litxapfilter

import (
	"unicode"
	"unicode/utf8"
)

func SpellOeAsWe(curr, _ *FilterTarget) (*string, *string) {
	oe, hasOe := hasPrefixFold(curr.Syllable, "oe")
	if !hasOe {
		return nil, nil
	}

	w := "w"
	o, ol := utf8.DecodeRuneInString(oe)
	if o == unicode.ToUpper(o) {
		w = "W"
	}

	we := w + curr.Syllable[ol:]
	return &we, nil
}
