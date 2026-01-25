package litxap

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type DummyDictionary map[string]Entry

func (t DummyDictionary) LookupEntries(word string) ([]Entry, error) {
	word = strings.ToLower(word)

	res := make([]Entry, 0, 1)
	if entry, ok := t[word]; ok {
		res = append(res, entry)

		for i := 0; i < 10; i++ {
			if entry, ok := t[word+":"+fmt.Sprint(i)]; ok {
				res = append(res, entry)
			} else {
				break
			}
		}
	} else {
		return nil, ErrEntryNotFound
	}

	return res, nil
}

type BrokenDictionary struct{}

func (b BrokenDictionary) LookupEntries(_ string) ([]Entry, error) {
	return nil, errors.New("500 something something")
}

func TestParseEntry(t *testing.T) {
	table := []string{
		"tskxe",
		"lo.ran: pe-fne- -ìri",
		"u.*van: -ti",
		"t·ì.*r·an: <äpeyk,ol>: walk",
		"t··el: <ei>: get, receive",
		"t·a.r·on: tì- <us> -ti: hunt",
		"t·ì.*r·an: tì- <us> -ìri: walk",
		"te.li.*si: -t: whirlwind",
		"sä.*pxor: : explosion",
		"tsa.heyl: no_stress: (part of tsaheyl si)",
	}

	for _, row := range table {
		t.Run(row, func(t *testing.T) {
			parsed := ParseEntry(row)
			assert.Equal(t, row, parsed.String())
		})
	}
}

func TestMultiDictionary_LookupEntries(t *testing.T) {
	mdGood := MultiDictionary{
		dummyDictionary,
		DummyDictionary{
			"kameie":   *ParseEntry("k·a.m·e: <ei>: see into, understand"),
			"sa'nokur": *ParseEntry("sa'.nok: -ur: nother"),
		},
	}
	mdEmpty := MultiDictionary{}
	mdBad := MultiDictionary{dummyDictionary, BrokenDictionary{}}
	mdBad2 := MultiDictionary{BrokenDictionary{}}

	res, err := mdEmpty.LookupEntries("sa'nokur")
	assert.ErrorIs(t, err, ErrEntryNotFound)
	assert.Nil(t, res)

	res, err = mdBad.LookupEntries("sa'nokur")
	assert.Error(t, err)
	assert.Nil(t, res)

	res, err = mdBad2.LookupEntries("tìfmetok")
	assert.Error(t, err)
	assert.Nil(t, res)

	res, err = mdGood.LookupEntries("mìfa")
	assert.ErrorIs(t, err, ErrEntryNotFound)
	assert.Nil(t, res)

	res, err = mdGood.LookupEntries("kameie")
	assert.NoError(t, err)
	assert.Equal(t, res, []Entry{
		*ParseEntry("k·a.m·e: <ei>: see, see into, understand, know (spiritual sense)"),
		*ParseEntry("k··ä: <am,ei>: go"),
		*ParseEntry("k·a.m·e: <ei>: see into, understand"),
	})

	res, err = mdGood.LookupEntries("sa'nokur")
	assert.NoError(t, err)
	assert.Equal(t, res, []Entry{*ParseEntry("sa'.nok: -ur: nother")})
}

func TestEntry_GenerateSyllables(t *testing.T) {
	table := []struct {
		Entry     string
		Syllables string
		Stress    int
		Offset    int
	}{
		{"fme.tok", "fme.tok", 0, 0},
		{"fm·e.t·ok: <ìm>", "fmì.me.tok", 1, 0},
		{"em.*k··ä: pe-pxe-tì- <us> -tsyìp-ìl", "pe.pe.sì.em.ku.sä.tsyì.pìl", 5, 3},
		{"*o.e: -ti", "oe.ti", 0, 0},
		{"*o.e: -ti", "oe.ti", 0, 0},
		{"o.*eng: -ti", "oeng.ti", 0, 0},
		{"ay.*oe: -r", "ay.oer", 1, 0},
		{"*o.e: ay-", "a.yo.e", 1, 1},
		{"ay.*oeng: -ìl", "ay.oe.ngìl", 1, 0},
		{"·o.*m·um: tsuk-", "tsu.ko.mum", 1, 1},
		{"·i.*n·an: -tswo", "i.nan.tswo", 1, 0},
		{"·i.*n·an: <äp,eyk,us>", "ä.pey.ku.si.nan", 3, 0},
		{"te.li.*si: -t", "te.li.sit", 2, 0},
		{"pxo.*eng: -ru", "pxo.eng.ru", 1, 0},
	}

	for _, row := range table {
		t.Run(row.Entry, func(t *testing.T) {
			entry := ParseEntry(row.Entry)
			if !assert.NotNil(t, entry) {
				return
			}

			syllables, stress, offset := entry.GenerateSyllables()

			assert.Equal(t, row.Syllables, strings.Join(syllables, "."))
			assert.Equal(t, row.Stress, stress)
			assert.Equal(t, row.Offset, offset)
		})
	}
}
