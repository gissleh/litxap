package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestNasalAssimilationCasual(t *testing.T) {
	table := []struct {
		curr, after, next, changeTo string
	}{
		{"seyng", "", "pe", "seym"},
		{"seyn", "", "pe", "seym"},
		{"seyN", "", "pe", "seyM"},
		{"seyNg", "", "pe", "seyM"},
		{"seynG", "", "pe", "seym"},
		{"seyNG", "", "pe", "seyM"},
		{"tseng", "", "ne", "tse"},
		{"kan", "", "-gan", "kang"},
		{"seyn", "", "tsyìp", ""},
		{"kem", "", "tsyìp", ""},
		{"kem", "", "tSyìp", ""},
		{"tan", " ", "na", "ta"},
		{"tìng", " ", "na", "tì"},
		{"zeN", "", "KE", "zeNG"},
		{"lum", "", "pe", ""},
		{"lun", "", "pe", "lum"},
		{"hol", "", "pxay", ""},
		{"fme", "", "tok", ""},
		{"fti", "", "a", ""},
		{"syen", "", "", ""},
		{"tseng", ".", "pe", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			changeCurr, changeNext := NasalAssimilationCasual(curr, next)

			if row.changeTo == "" {
				assert.Nil(t, changeCurr)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, row.changeTo, *changeCurr)
				}
			}

			assert.Nil(t, changeNext, "Next syllable should be untouched")
		})
	}
}

func TestNasalAssimilation(t *testing.T) {
	table := []struct {
		curr, next, expected string
		currEntry, nextEntry string
		currSI, nextSI       int
	}{
		{"tseng", "ne", "", "tseng: -ne", "tseng: -ne", 0, 1},
		{"seyn", "pe", "", "seyn: -pe", "seyn: -pe", 0, 1},
		{"tìng", "na", "tì", "tìng", "nari", 0, 0},
		{"tìng", "mi", "tì", "tìng", "mikyun", 0, 0},
		{"tìng", "po", "", "tìng", "po", 0, 0},
		{"tìng", "na", "", "tìng", "nari: -t", 0, 0},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s(%s)-%s(%s)", row.curr, row.currEntry, row.next, row.nextEntry), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, SyllableIndex: row.currSI}
			if row.currEntry != "" {
				curr.Entry = litxap.ParseEntry(row.currEntry)
			}
			next := &FilterTarget{Syllable: row.next, SyllableIndex: row.nextSI}
			if row.nextEntry != "" {
				next.Entry = litxap.ParseEntry(row.nextEntry)
			}
			changeCurr, changeNext := NasalAssimilation(curr, next)

			if row.expected == "" {
				assert.Nil(t, changeCurr)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, row.expected, *changeCurr)
				}
			}

			assert.Nil(t, changeNext, "Next syllable should be untouched")
		})
	}
}
