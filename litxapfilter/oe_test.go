package litxapfilter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpellOeAsWe(t *testing.T) {
	table := []struct {
		curr, changeCurr string
	}{
		{"oeng", "weng"},
		{"oel", "wel"},
		{"oet", "wet"},
		{"oe", "we"},
		{"oE", "wE"},
		{"Oe", "We"},
		{"OE", "WE"},
		{"he", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s", row.curr), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.curr}
			next := &FilterTarget{Syllable: "ta"}

			changeCurr, changeNext := SpellOeAsWe(curr, next)

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
