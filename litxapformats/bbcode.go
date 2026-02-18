package litxapformats

import (
	"github.com/gissleh/litxap"
)

func BBCode() litxap.LineFormatter {
	return &bbCodeFormatter{}
}

type bbCodeFormatter struct{}

func (f *bbCodeFormatter) LinePartTags(lp litxap.LinePart, stress int) (string, string) {
	switch stress {
	case litxap.LPSNotWord:
		return "", ""
	case litxap.LPSAmbiguousMatches:
		return "[color=yellow]", "[/color]"
	case litxap.LPSNoMatches:
		return "[color=red]", "[/color]"
	case litxap.LPSAnyStress:
		return "[color=skyblue]", "[/color]"
	default:
		return "", ""
	}
}

func (f *bbCodeFormatter) StressedSyllableTags() (string, string) {
	return "[u]", "[/u]"
}
