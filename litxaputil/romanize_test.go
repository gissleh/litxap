package litxaputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanize(t *testing.T) {
	table := []struct {
		curr     string
		expected []string
	}{
		// Ordinary words
		{"tɪ.ˈfmɛ.tok̚", []string{"tì__fme__tok"}},
		{"u.ˈvan", []string{"u__van__"}},
		{"ˈu.ɾan", []string{"__u__ran"}},
		{"ʔawk'", []string{"'awkx"}},
		{"fko", []string{"fko"}},
		{"mo", []string{"mo"}},
		{"t'on", []string{"txon"}},
		{"tɪ.ˈfmɛ.tok̚", []string{"tì__fme__tok"}},
		{"t͡sam", []string{"tsam"}},
		{"fpom", []string{"fpom"}},
		{"ˈt·a.ɾ·on", []string{"__ta__ron"}},
		{"ˈʔɛ.koŋ", []string{"__'e__kong"}},
		{"ɛ.ˈjawɾ", []string{"e__yawr__"}},
		{"kṛ", []string{"krr"}},
		{"k'ḷ", []string{"kxll"}},
		{"po", []string{"po"}},
		// Flexible syllable stress
		{"aj.ˈfo] or [ˈaj.fo", []string{"ayfo"}},
		{"ˈɪ.læ] or [ɪ.ˈlæ", []string{"ìlä"}},
		{"ˈmɪ.fa] or [mɪ.ˈfa", []string{"mìfa"}},
		{"ˈt͡sa.kɛm] or [t͡sa.ˈkɛm", []string{"tsakem"}},
		{"t͡sa.ˈt͡sɛŋ] or [ˈt͡sa.t͡sɛŋ", []string{"tsatseng"}},
		// Multiple pronunciation
		{"nɪ.aw.ˈno.mʊm] or [naw.ˈno.mʊm", []string{"nìaw__no__mum", "naw__no__mum"}},
		{"nɪ.aj.ˈwɛŋ] or [naj.ˈwɛŋ", []string{"nìay__weng__", "nay__weng__"}},
		{"tɪ.sjɪ.maw.nʊn.ˈʔi] or [t͡sjɪ.maw.nʊn.ˈʔi", []string{"tìsyìmawnun__'i__", "tsyìmawnun__'i__"}},
		{"tɪ.sæ.ˈfpɪl.jɛwn] or [t͡sæ.ˈfpɪl.jɛwn", []string{"tìsä__fpìl__yewn", "tsä__fpìl__yewn"}},
	}

	for _, row := range table {
		t.Run(row.curr, func(t *testing.T) {
			next := RomanizeIPA(row.curr)
			assert.Equal(t, row.expected, next)
		})
	}
}

func TestRomanize_Panic(t *testing.T) {
	badSuffix := Suffix{
		reanalysis:    -19392,
		syllableSplit: []string{"blarg"},
	}
	assert.Panics(t, func() { badSuffix.Apply([]string{"stuff"}) })
	assert.Panics(t, func() { findSuffix("teri").Apply([]string{}) })
}
