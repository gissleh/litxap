package litxapfilter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiphthongFromWeakVowel(t *testing.T) {
	table := []struct {
		curr, next, after, changeCurr string
		clearNext                     bool
	}{
		{"ma", "ut", "", "mawt", true},
		{"ma", "u", " ", "maw", true},
		{"ke", "U", "", "keW", true},
		{"me", "i", "", "mey", true},
		{"ha", "Ìng", "", "haYng", true},
		{"hA", "Ìng", "", "hAYng", true},
		{"me", "kxa", "", "", false},
		{"me", "", "", "", false},
		{"me", "*i", "", "", false},
		{"*me", "i", "", "", false},
		{"*me", "*i", "", "", false},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			for _, target := range []*FilterTarget{curr, next} {
				if target != nil && strings.HasPrefix(target.Syllable, "*") {
					target.Syllable = target.Syllable[len("*"):]
					target.Stressed = true
				}
			}

			changeCurr, changeNext := DiphthongFromWeakVowel(curr, next)

			if row.changeCurr == "" {
				assert.Nil(t, changeCurr)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, row.changeCurr, *changeCurr)
				}
			}
			if row.clearNext {
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, "", *changeNext)
				}
			} else {
				assert.Nil(t, changeNext)
			}
		})
	}
}

func TestReanalyzeDiphthongs(t *testing.T) {
	table := []struct {
		curr, next, changeCurr, changeNext string
	}{
		{"ey", "e", "e", "ye"},
		{"ay", "on", "a", "yon"},
		{"ay", "nga", "", ""},
		{"ay", "", "", ""},
		{"a", "e", "", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s", row.curr, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			changeCurr, changeNext := ReanalyzeDiphthongs(curr, next)

			if row.changeCurr == "" {
				assert.Nil(t, changeCurr)
				assert.Nil(t, changeNext)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, row.changeCurr, *changeCurr)
				}
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.changeNext, *changeNext)
				}
			}
		})
	}
}
