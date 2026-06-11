package litxapformats

import (
	"fmt"

	"github.com/gissleh/litxap"
)

// IRCDefaultColors returns the same as IRC with the default colors (black text, yellow ambiguities and red no-matches)
func IRCDefaultColors() litxap.LineFormatter {
	return IRC(1, 8, 4)
}

// IRC formats the line with IRC carets
func IRC(textColor, amColor, nmColor int) litxap.LineFormatter {
	return &ircFormatter{
		amColor: fmt.Sprintf("\u0003%02d,%02d", textColor, amColor),
		nmColor: fmt.Sprintf("\u0003%02d,%02d", textColor, nmColor),
	}
}

type ircFormatter struct {
	amColor string
	nmColor string
}

func (f *ircFormatter) LinePartTags(lp litxap.LinePart, stress int) (string, string) {
	switch stress {
	case litxap.LPSAmbiguousMatches:
		return f.amColor, "\u0003"
	case litxap.LPSNoMatches:
		return f.nmColor, "\u0003"
	case litxap.LPSAnyStress:
		return "\u001F", "\u001F"
	default:
		return "", ""
	}
}

func (f *ircFormatter) StressedSyllableTags() (string, string) {
	return "\u001F", "\u001F"
}
