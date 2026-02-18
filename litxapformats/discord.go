package litxapformats

import "github.com/gissleh/litxap"

func DiscordMarkdown() litxap.LineFormatter {
	return &discordFormatter{}
}

type discordFormatter struct{}

func (f *discordFormatter) LinePartTags(_ litxap.LinePart, stress int) (string, string) {
	if stress == litxap.LPSNoMatches {
		return "\\*", "(NO MATCHES)"
	}
	if stress == litxap.LPSAmbiguousMatches {
		return "\\*", "(AMBIGUOUS)"
	}

	return "", ""
}

func (f *discordFormatter) StressedSyllableTags() (string, string) {
	return "__", "__"
}
