package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaeRemover(t *testing.T) {
	table := []struct {
		currSyllable  string
		currAfter     string
		currPartIndex int
		currStressed  bool
		nextSyllable  string
		nextPartIndex int
		newNext       string
	}{
		{"sä", "", 0, false, "tspang", 0, ""},
		{"sä", "", 0, false, "ta", 1, ""},
		{"sä", "", 0, false, "vll", 0, ""},
		{"sä", "", 0, false, "", 0, ""},
		{"sä", "", 0, true, "ta", 0, ""},
		{"sä", "", 0, false, "ta", 0, "sta"},
		{"sä", "", 0, false, "ta", 0, "sta"},
		{"sä", "", 0, false, "keyn", 0, "skeyn"},
		{"sä", "", 0, false, "leym", 0, "sleym"},
		{"sä", "", 0, false, "po", 0, "spo"},
		{"sÄ", "", 0, false, "rawn", 0, "srawn"},
		{"Sä", "", 0, false, "pxor", 0, "Spxor"},
		{"sä", "", 0, false, "Pxor", 0, "sPxor"},
		{"sä", ". ", 0, false, "Pxor", 0, ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%d-%s-%d", row.currSyllable, row.currAfter, row.currPartIndex, row.nextSyllable, row.nextPartIndex), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.currSyllable, After: row.currAfter, PartIndex: row.currPartIndex, Stressed: row.currStressed}
			var next *FilterTarget
			if row.nextSyllable != "" {
				next = &FilterTarget{Syllable: row.nextSyllable, PartIndex: row.nextPartIndex}
			}

			changeCurr, changeNext := SaeRemover(curr, next)
			if row.newNext != "" {
				t.Log("Expected:", row.newNext)

				if assert.NotNil(t, changeCurr) {
					assert.Equal(t, "", *changeCurr)
				}
				if assert.NotNil(t, changeNext) {
					assert.Equal(t, row.newNext, *changeNext)
				}
			} else {
				assert.Nil(t, changeCurr)
				assert.Nil(t, changeNext)
			}
		})
	}
}
