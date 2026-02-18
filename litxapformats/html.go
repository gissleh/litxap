package litxapformats

import (
	"github.com/gissleh/litxap"
)

// CompactHTML formats using a compact HTML with <span></span> around words and two-letter class names for special cases.
// `am`=ambiguous matches, nm=no matches, us=unstressed[deprecated], as=any stress (like ìlä).
// The <u> tag is used for the stressed syllable.
func CompactHTML() litxap.LineFormatter {
	return &compactHtmlFormatter{}
}

type compactHtmlFormatter struct{}

func (f *compactHtmlFormatter) LinePartTags(lp litxap.LinePart, stress int) (string, string) {
	switch stress {
	case litxap.LPSNotWord:
		return "", ""
	case litxap.LPSAmbiguousMatches:
		return "<span class=\"am\">", "</span>"
	case litxap.LPSNoMatches:
		return "<span class=\"nm\">", "</span>"
	case litxap.LPSAnyStress:
		return "<span class=\"as\">", "</span>"
	default:
		return "<span>", "</span>"
	}
}

func (f *compactHtmlFormatter) StressedSyllableTags() (string, string) {
	return "<u>", "</u>"
}
