package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropGlottalStopsBetweenUnlikeVowels(t *testing.T) {
	table := []struct {
		currSyllable string
		currAfter    string
		nextSyllable string
		nextChange   string
	}{
		{"ve", "", "'O", "O"},
		{"fra", "", "'u", "u"},
		{"fra", "", "'Ur", "Ur"},
		{"ep", "", "'ang", ""},
		{"rä", "", "'ä", ""},
		{"rä", "", "'Ä", ""},
		{"rÄ", "", "'ä", ""},
		{"rÄ", "", "'Ä", ""},
		{"ve", "", "", ""},
		{"ma", " ", "'e", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s-%s-%s", row.currSyllable, row.currAfter, row.nextSyllable), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.currSyllable, After: row.currAfter}
			var next *FilterTarget
			if row.nextSyllable != "" {
				next = &FilterTarget{Syllable: row.nextSyllable}
				if curr.After == "" {
					next.SyllableIndex = 1
				}
			}

			currChange, nextChange := DropGlottalStopBetweenUnlikeVowels(curr, next)
			assert.Nil(t, currChange)
			if row.nextChange == "" {
				assert.Nil(t, nextChange)
			} else {
				if assert.NotNil(t, nextChange) {
					assert.Equal(t, row.nextChange, *nextChange)
				}
			}
		})
	}
}
