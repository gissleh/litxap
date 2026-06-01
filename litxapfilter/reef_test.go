package litxapfilter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReefUnstressedAeAsE(t *testing.T) {
	table := []struct {
		currSyllable string
		currChange   string
	}{
		{"sä", "se"},
		{"speng", ""},
		{"ä", "e"},
		{"*ä", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s", row.currSyllable), func(t *testing.T) {
			syllable := row.currSyllable
			stressed := false
			if strings.HasPrefix(syllable, "*") {
				stressed = true
				syllable = syllable[len("*"):]
			}
			curr := &FilterTarget{Syllable: syllable, Stressed: stressed}
			next := &FilterTarget{Syllable: "ngä", Stressed: false}

			currChange, nextChange := ReefUnstressedAeAsE(curr, next)
			if row.currChange == "" {
				assert.Nil(t, currChange)
			} else {
				if assert.NotNil(t, currChange) {
					assert.Equal(t, row.currChange, *currChange)
				}
			}
			assert.Nil(t, nextChange)
		})
	}
}

func TestReefDropGlottalStopsBetweenVowels(t *testing.T) {
	table := []struct {
		currSyllable string
		currAfter    string
		nextSyllable string
		nextChange   string
	}{
		{"rä", "", "'ä", "ä"},
		{"rä", "", "'Ä", "Ä"},
		{"rÄ", "", "'ä", "ä"},
		{"rÄ", "", "'Ä", "Ä"},
		{"ve", "", "'O", "O"},
		{"ep", "", "'ang", ""},
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

			currChange, nextChange := ReefDropGlottalStopsBetweenVowels(curr, next)
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

func TestReefEjectiveToVoiced(t *testing.T) {
	table := []struct {
		currSyllable string
		currAfter    string
		nextSyllable string
		currChange   string
	}{
		{"kxitx", "", "pxaw", "gid"},
		{"kxitx", "", "maw", "gitx"},
		{"PxA", "", "", "BA"},
		{"tXa", "", "", "da"},
		{"ngay", "", "txo", ""},
		{"kit", "", "", ""},
		{"nopx", " ", "kxo", ""},
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

			currChange, nextChange := ReefEjectiveToVoiced(curr, next)
			if row.currChange == "" {
				assert.Nil(t, currChange)
			} else {
				if assert.NotNil(t, currChange) {
					assert.Equal(t, row.currChange, *currChange)
				}
			}
			assert.Nil(t, nextChange)
		})
	}
}

func TestReefApplyChSh(t *testing.T) {
	table := []struct {
		currSyllable string
		currChange   string
	}{
		{"sä", ""},
		{"syo", "sho"},
		{"sYo", "sHo"},
		{"Sye", "She"},
		{"SYu", "SHu"},
		{"TsYä", "Chä"},
		{"tSYä", "cHä"},
		{"speng", ""},
		{"ä", ""},
		{"*ä", ""},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s", row.currSyllable), func(t *testing.T) {
			curr := &FilterTarget{Syllable: row.currSyllable, Stressed: false}
			next := &FilterTarget{Syllable: "tsya", Stressed: false}

			currChange, nextChange := ReefApplyChSh(curr, next)
			if row.currChange == "" {
				assert.Nil(t, currChange)
			} else {
				if assert.NotNil(t, currChange) {
					assert.Equal(t, row.currChange, *currChange)
				}
			}
			assert.Nil(t, nextChange)
		})
	}
}
