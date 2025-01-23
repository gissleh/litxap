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

		newLine[i].Matches = make([]LinePartMatch, 0, len(results))
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

func (line Line) UnStressSiVerbParts(dict Dictionary) Line {
	isCopied := false
	newLine := line

	for i, part := range line {
		if part.matchesWord("si") {
			if p := line.prevWord(i); p != -1 {
				p2 := p
				isKe := line[p].matchesWord("ke")
				isRaa := line[p].matchesWord("rä'ä")
				if isKe || isRaa {
					p2 = line.prevWord(p)
					if p2 == -1 {
						continue
					}
				}

				// Check if it's a si-verb
				entries, _ := dict.LookupEntries(line[p2].Raw + " " + line[i].Raw)
				foundSiVerb := false
				for _, entry := range entries {
					if strings.HasSuffix(entry.Word, " si") || strings.HasSuffix(entry.Word, " säpi") || strings.HasSuffix(entry.Word, " seyki") || strings.HasSuffix(entry.Word, " säpeyki") {
						foundSiVerb = true
					}
				}
				if !foundSiVerb {
					continue
				}

				if !isCopied {
					isCopied = true
					newLine = append(newLine[:0:0], newLine...)
					for j := range newLine {
						newLine[j].Matches = append(newLine[j].Matches[:0:0], newLine[j].Matches...)
					}
				}

				// Unstress the word (rä'ä might need this as well, but it's not explicitly defined)
				if isKe {
					for j := range line[p2].Matches {
						newLine[p2].Matches[j].Stress = -1
					}

					// Stress the "ke"
					for j, match := range line[p].Matches {
						if match.Entry.Word == "ke" {
							newLine[p].Matches[j].StressedWord = true
							break
						}
					}
				}

				// Unstress the si
				for j := range line[i].Matches {
					newLine[i].Matches[j].Stress = -1
				}
			}
		}
	}

	return newLine
}

func (line Line) prevWord(i int) int {
	for j := i - 1; j >= 0; j-- {
		if line[j].IsWord {
			return j
		} else if len(strings.Trim(line[j].Raw, "  ")) > 0 {
			// sentence or clause boundary, most likely.
			break
		}
	}

	return -1
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

func (part *LinePart) matchesWord(word string) bool {
	for _, match := range part.Matches {
		if match.Entry.Word == word {
			return true
		}
	}

	return false
}

type LinePartMatch struct {
	Syllables    []string `json:"syllables"`
	Stress       int      `json:"stress"`
	Entry        Entry    `json:"entry"`
	StressedWord bool     `json:"stressedWord,omitempty"`
}
