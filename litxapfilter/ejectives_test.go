package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDemoteEjectivesBeforeConsonants(t *testing.T) {
	table := []struct {
		curr, after, next string
		changeCurr        string
	}{
		{"tokx", "", "nga'", "tok"},
		{"sräTX", "", "maw", "sräT"},
		{"srätx", " ", "txo", ""},
		{"srätx", " ", "do", ""},
		{"atx", "", "kxe", ""},
		{"atx", "", "e", ""},
		{"tokx", ".", "nga'", ""},
		{"tor", ".", "", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			changeCurr, changeNext := DemoteEjectivesBeforeConsonants(curr, next)

			if row.changeCurr == "" {
				assert.Nil(t, changeCurr)
			} else {
				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, row.changeCurr, *changeCurr)
				}
			}
			assert.Nil(t, changeNext)
		})
	}
}

func TestRemoveRepeatedEjective(t *testing.T) {
	table := []struct {
		curr, after, next string
		changeNext        string
	}{
		{"srätx", " ", "txo", "o"},
		{"srätx", " ", "do", "o"},
		{"srätx", "", "pxaw", ""},
		{"tokx", ". ", "Pelun", ""},
		{"Kxitx", "!", "", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%s", row.curr, row.after, row.next), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr, After: row.after}
			var next *FilterTarget
			if row.next != "" {
				next = &FilterTarget{Syllable: row.next}
			}

			changeCurr, changeNext := RemoveRepeatedEjective(curr, next)

			assert.Nil(t, changeCurr)
			if row.changeNext == "" {
				assert.Nil(t, changeNext)
			} else {
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.changeNext, *changeNext)
				}
			}
		})
	}
}
