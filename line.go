package litxap

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func RunLine(line string, dictionary Dictionary) (Line, error) {
	return ParseLine(line).Run(dictionary)
}

type Line []LinePart

func (line Line) Run(dict Dictionary) (Line, error) {
	newLine := append(line[:0:0], line...)

	for i, part := range newLine {
		if !part.IsWord {
			continue
		}

		lookup := part.Raw
		if part.Lookup != "" {
			lookup = part.Lookup
		}

		lookup = strings.ToLower(lookup)

		results, err := dict.LookupEntries(lookup)
		if err != nil {
			if errors.Is(err, ErrEntryNotFound) {
				continue
			}

			return nil, fmt.Errorf("failed to lookup \"%s\": %w", lookup, err)
		}

		for _, result := range results {
			syllables, stress := RunWord(part.Raw, result)
			if syllables != nil {
				newLine[i].Matches = append(newLine[i].Matches, LinePartMatch{
					Syllables: syllables,
					Stress:    stress,
					Entry:     result,
				})
			}
		}
	}

	return newLine, nil
}

// ParseLine splits out the words from a line of text.
func ParseLine(s string) Line {
	wordMode := false
	lastPos := 0
	lastPipe := 0
	currentPos := 0
	res := make(Line, 0, (len(s)/5)+1)

	s = strings.NewReplacer("’", "'", "‘", "'").Replace(s) + "\n"

	for _, ch := range s {
		if ch == '|' {
			lastPipe = currentPos
		} else if ch == '\n' || wordMode != (unicode.IsLetter(ch) || ch == '\'' || ch == '-') {
			if lastPos != currentPos {
				// Colors
				if wordMode && strings.Contains(s[lastPos:currentPos], "-") && strings.Contains(s[lastPos:currentPos], "na") {
					split := strings.Split(s[lastPos:currentPos], "-")
					for i, token := range split {
						if i > 0 {
							res = append(res, LinePart{Raw: "-"})
						}

						if i == 0 && strings.HasPrefix(token, "a") {
							res = append(res, LinePart{Raw: "a", IsWord: true})
							token = strings.TrimPrefix(token, "a")
						}

						hasAttrSuffix := false
						if i == len(split)-1 && strings.HasSuffix(token, "a") {
							token = strings.TrimSuffix(token, "a")
							hasAttrSuffix = true
						}

						res = append(res, LinePart{Raw: token, IsWord: true})
						if hasAttrSuffix {
							res = append(res, LinePart{Raw: "a", IsWord: true})
						}
					}
				} else {
					raw := s[lastPos:currentPos]
					lookup := s[lastPos:lastPos]
					if lastPipe != lastPos {
						lookup = s[lastPos:lastPipe]
						raw = s[lastPipe+1 : currentPos]
					}

					res = append(res, LinePart{
						Raw:     raw,
						Lookup:  lookup,
						IsWord:  wordMode,
						Matches: nil,
					})
				}

				lastPos = currentPos
				lastPipe = currentPos
			}

			wordMode = !wordMode
		}

		currentPos += utf8.RuneLen(ch)
	}

	return res
}

type LinePart struct {
	Raw     string          `json:"raw"`
	Lookup  string          `json:"lookup,omitempty"`
	IsWord  bool            `json:"isWord,omitempty"`
	Matches []LinePartMatch `json:"matches,omitempty"`
}

type LinePartMatch struct {
	Syllables []string `json:"syllables"`
	Stress    int      `json:"stress"`
	Entry     Entry    `json:"entry"`
}
