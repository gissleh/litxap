package litxapfilter

import (
	"slices"
	"strings"

	"github.com/gissleh/litxap"
)

// A Filter takes in a sliding window of three syllables and return modification to
// the current and the next.
type Filter func(curr, next *FilterTarget) (*string, *string)

type FilterTarget struct {
	PartIndex     int
	MatchIndex    int
	SyllableIndex int
	Syllable      string
	Stressed      bool
	After         string
}

// ApplyFilters is just a wrapper for running one filter after another. They'll each make a full pass
// so the next filter will be dealing with the output of the previous filter.
func ApplyFilters(line litxap.Line, filters ...Filter) litxap.Line {
	for _, filter := range filters {
		line = ApplyFilter(line, filter)
	}

	return line
}

// ApplyFilter runs the filtering logic. It will not copy any more than necessary, up to returning the
// passed litxap.Line if the filter ends up changing nothing.
func ApplyFilter(line litxap.Line, filter Filter) litxap.Line {
	newLine := line

	copiedLine := false
	copiedParts := make(map[int]bool)
	copiedSyllables := make(map[[2]int]bool)

	setSyllable := func(pi, mi, si int, value string) {
		if !copiedSyllables[[2]int{pi, mi}] {
			if !copiedParts[pi] {
				if !copiedLine {
					newLine = slices.Clone(line)
					copiedLine = true
				}

				newLine[pi].Matches = slices.Clone(newLine[pi].Matches)
				copiedParts[pi] = true
			}

			newLine[pi].Matches[mi].Syllables = slices.Clone(newLine[pi].Matches[mi].Syllables)
			copiedSyllables[[2]int{pi, mi}] = true
		}

		newLine[pi].Matches[mi].Syllables[si] = value
	}

	for pi := range newLine {
		after := nonWordAfter(newLine, pi)
		piNext := nextPartAfter(newLine, pi)

		for mi := range newLine[pi].Matches {
			for si, syllable := range newLine[pi].Matches[mi].Syllables {
				if syllable == "" {
					continue
				}

				var curr *FilterTarget
				if si < len(newLine[pi].Matches[mi].Syllables)-1 { // Within word
					afterSecondSyllable := ""
					if si == len(newLine[pi].Matches[mi].Syllables)-2 {
						afterSecondSyllable = after
					}

					curr = &FilterTarget{
						PartIndex:     pi,
						MatchIndex:    mi,
						SyllableIndex: si,
						Syllable:      syllable,
						Stressed:      newLine[pi].Matches[mi].Stress == si,
						After:         "",
					}
					next := &FilterTarget{
						PartIndex:     pi,
						MatchIndex:    mi,
						SyllableIndex: si + 1,
						Syllable:      newLine[pi].Matches[mi].Syllables[si+1],
						Stressed:      newLine[pi].Matches[mi].Stress == si+1,
						After:         afterSecondSyllable,
					}

					currChange, nextChange := filter(curr, next)
					if currChange != nil {
						setSyllable(pi, mi, si, *currChange)
					}
					if nextChange != nil {
						setSyllable(pi, mi, si+1, *nextChange)
					}
				} else { // Across word boundaries
					curr = &FilterTarget{
						PartIndex:     pi,
						MatchIndex:    mi,
						SyllableIndex: si,
						Syllable:      syllable,
						Stressed:      newLine[pi].Matches[mi].Stress == si,
						After:         after,
					}

					// Run the filter on the next word's beginning first.
					if piNext != -1 && len(newLine[piNext].Matches) > 0 {
						for miNext, matchNext := range newLine[piNext].Matches {
							afterNextFirstSyllable := ""
							if len(matchNext.Syllables) == 1 {
								afterNextFirstSyllable = nonWordAfter(newLine, piNext)
							}

							next := &FilterTarget{
								PartIndex:     pi,
								MatchIndex:    mi,
								SyllableIndex: 0,
								Syllable:      matchNext.Syllables[0],
								Stressed:      matchNext.Stress == 0,
								After:         afterNextFirstSyllable,
							}

							currChange, nextChange := filter(curr, next)
							if nextChange != nil {
								setSyllable(piNext, miNext, 0, *nextChange)
							}
							if currChange != nil {
								setSyllable(pi, mi, si, *currChange)
							}
						}
					} else {
						// Then run it on the current word
						currChange, _ := filter(curr, nil)
						if currChange != nil {
							setSyllable(pi, mi, si, *currChange)
						}
					}
				}
			}
		}
	}

	// Remove cleared syllables
	if copiedLine {
		for pi, part := range newLine {
			if !copiedParts[pi] {
				continue
			}

			for mi, match := range part.Matches {
				if !copiedSyllables[[2]int{pi, mi}] {
					continue
				}

				n := 0
				for si, syllable := range match.Syllables {
					if syllable != "" {
						match.Syllables[n] = syllable
						n += 1
					} else if match.Stress >= si {
						// Omitted syllable left of stress should move it back.
						newLine[pi].Matches[mi].Stress -= 1
					}
				}
				newLine[pi].Matches[mi].Syllables = match.Syllables[:n]

				if mi == 0 {
					newLine[pi].Raw = strings.Join(newLine[pi].Matches[mi].Syllables, "")
				}
			}
		}
	}

	return newLine
}

func nextPartAfter(line litxap.Line, i int) int {
	for j := i + 1; j < len(line); j++ {
		if line[j].IsWord {
			return j
		}
	}

	return -1
}

func nonWordAfter(line litxap.Line, i int) string {
	raw := ""
	for j := i + 1; j < len(line); j++ {
		if line[j].IsWord {
			break
		}

		raw += line[j].Raw
	}

	return raw
}
