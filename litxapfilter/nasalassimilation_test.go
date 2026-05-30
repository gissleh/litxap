package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNasalAssimilation(t *testing.T) {
	table := []struct {
		curr, after, next, changeTo string
	}{
		{"seyng", "", "pe", "seym"},
		{"seyn", "", "pe", "seym"},
		{"seyN", "", "pe", "seyM"},
		{"seyNg", "", "pe", "seyM"},
		{"seynG", "", "pe", "seym"},
		{"seyNG", "", "pe", "seyM"},
		{"kan", "", "-gan", "kang"},
		{"seyn", "", "tsyìp", ""},
		{"kem", "", "tsyìp", ""},
		{"kem", "", "tSyìp", ""},
		{"tan", " ", "na", "ta"},
		{"tìng", " ", "na", "tì"},
		//{"taM", "", "tey", "taN"},
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

			changeCurr, changeNext := NasalAssimilation(curr, next)

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
