package litxapfilter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestElideUnstressedEWordEndings(t *testing.T) {
	table := []struct {
		curr, after, next string
		changeNext        string
	}{
		{"me", ", ", "ul", "mul"},
		{"te", " ", "Ey", "tEy"},
		{"te", " ", "Naw", ""},
		{"*mä", " ", "Ey", ""},
		{"na", " ", "A", ""},
		{"me", "", "u", ""},
		{"me", "", "-u", ""},
		{"ne", "", "", ""},
		{"ne", "! ", "A", ""},
		{"pe", "? ", "Oel", ""},
		{"me", ". ", "Ey", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s%s%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			if strings.HasPrefix(curr.Syllable, "*") {
				curr.Syllable = strings.TrimPrefix(curr.Syllable, "*")
				curr.Stressed = true
			}
			if next != nil && strings.HasPrefix(next.Syllable, "-") {
				next.Syllable = strings.TrimPrefix(next.Syllable, "-")
				next.SyllableIndex = 1
			}

			changeCurr, changeNext := ElideUnstressedEWordEndings(curr, next)

			if row.changeNext == "" {
				assert.Nil(t, changeCurr)
				assert.Nil(t, changeNext)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, "", *changeCurr)
				}
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.changeNext, *changeNext)
				}
			}
		})
	}
}

func TestElideMiSiNiBeforeAy(t *testing.T) {
	table := []struct {
		curr, after, next, word string
		changeNext              string
	}{
		{"nì", "", "ay", "ayfo", "nay"},
		{"sì", " ", "ay", "sì", "say"},
		{"sì", "", "", "aynga", ""},
		{"mì", "", "fa", "mìfa", ""},
		{"mì", " ", "ay", "kelku", ""},
		{"nì", "", "mun", "nìmun", ""},
		{"pe", "", "lun", "pelun", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s%s%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}
			curr.Entry = &litxap.Entry{
				Word:      row.word,
				Syllables: []string{row.word},
			}

			changeCurr, changeNext := ElideMiSiNiBeforeAy(curr, next)

			if row.changeNext == "" {
				assert.Nil(t, changeCurr)
				assert.Nil(t, changeNext)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, "", *changeCurr)
				}
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.changeNext, *changeNext)
				}
			}
		})
	}
}

func TestElideAdvPrefixAndE(t *testing.T) {
	table := []struct {
		curr, after, next, word string
		changeNext              string
	}{
		{"nì", "", "*et", "etrìp", "net"},
		{"nì", "", "e", "eyawr", "nì"},
		{"nì", "", "ke", "keyawr", ""},
		{"nì", "", "e", "ean", ""},
		{"nì", " ", "ye", "stä'nì", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s%s%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
				if strings.HasPrefix(next.Syllable, "*") {
					next.Syllable = strings.TrimPrefix(next.Syllable, "*")
					next.Stressed = true
				}
			}
			if row.after != "" {
				next.PartIndex = 1
				next.SyllableIndex = 0
			} else {
				next.SyllableIndex = 1
			}
			curr.Entry = &litxap.Entry{Word: row.word}

			changeCurr, changeNext := ElideAdvPrefixAndE(curr, next)

			if row.changeNext == "" {
				assert.Nil(t, changeCurr)
				assert.Nil(t, changeNext)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, "", *changeCurr)
				}
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.changeNext, *changeNext)
				}
			}
		})
	}
}
