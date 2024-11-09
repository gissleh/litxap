package litxaputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanize(t *testing.T) {
	table := []struct {
		curr     string
		expected []string
		stress   [][]int
	}{
		// Ordinary words
		{"tɪ.ˈfmɛ.tok̚", []string{"tìfmetok"}, [][]int{{1}}},
		{"u.ˈvan", []string{"uvan"}, [][]int{{1}}},
		{"ˈu.ɾan", []string{"uran"}, [][]int{{0}}},
		{"ʔawk'", []string{"'awkx"}, [][]int{{-1}}},
		{"fko", []string{"fko"}, [][]int{{-1}}},
		{"mo", []string{"mo"}, [][]int{{-1}}},
		{"t'on", []string{"txon"}, [][]int{{-1}}},
		{"tɪ.ˈfmɛ.tok̚", []string{"tìfmetok"}, [][]int{{1}}},
		{"t͡sam", []string{"tsam"}, [][]int{{-1}}},
		{"fpom", []string{"fpom"}, [][]int{{-1}}},
		{"ˈt·a.ɾ·on", []string{"taron"}, [][]int{{0}}},
		{"ˈʔɛ.koŋ", []string{"'ekong"}, [][]int{{0}}},
		{"ɛ.ˈjawɾ", []string{"eyawr"}, [][]int{{1}}},
		{"kṛ", []string{"krr"}, [][]int{{-1}}},
		{"k'ḷ", []string{"kxll"}, [][]int{{-1}}},
		{"po", []string{"po"}, [][]int{{-1}}},
		// Flexible syllable stress
		{"aj.ˈfo] or [ˈaj.fo", []string{"ayfo"}, [][]int{{-1}}},
		{"ˈɪ.læ] or [ɪ.ˈlæ", []string{"ìlä"}, [][]int{{-1}}},
		{"ˈmɪ.fa] or [mɪ.ˈfa", []string{"mìfa"}, [][]int{{-1}}},
		{"ˈt͡sa.kɛm] or [t͡sa.ˈkɛm", []string{"tsakem"}, [][]int{{-1}}},
		{"t͡sa.ˈt͡sɛŋ] or [ˈt͡sa.t͡sɛŋ", []string{"tsatseng"}, [][]int{{-1}}},
		// Multiple pronunciation
		{"nɪ.aw.ˈno.mʊm] or [naw.ˈno.mʊm", []string{"nìawnomum", "nawnomum"}, [][]int{{2}, {1}}},
		{"nɪ.aj.ˈwɛŋ] or [naj.ˈwɛŋ", []string{"nìayweng", "nayweng"}, [][]int{{2}, {1}}},
		{"tɪ.sjɪ.maw.nʊn.ˈʔi] or [t͡sjɪ.maw.nʊn.ˈʔi", []string{"tìsyìmawnun'i", "tsyìmawnun'i"}, [][]int{{4}, {3}}},
		{"tɪ.sæ.ˈfpɪl.jɛwn] or [t͡sæ.ˈfpɪl.jɛwn", []string{"tìsäfpìlyewn", "tsäfpìlyewn"}, [][]int{{2}, {1}}},
		// Multiple words
		{"ˈut.ɾa.ja ˈmok.ɾi", []string{"utraya mokri"}, [][]int{{0, 0}}},
		{"t͡sa.ˈhɛjl s·i", []string{"tsaheyl si"}, [][]int{{1, -1}}},
		{"ˈnɪ.ˌju ˈjoɾ.kɪ", []string{"nìyu yorkì"}, [][]int{{0, 0}}},
		{"t͡sawl sl·u", []string{"tsawl slu"}, [][]int{{-1, -1}}},
	}

	for _, row := range table {
		t.Run(row.curr, func(t *testing.T) {
			spelling, stress := RomanizeIPA(row.curr)
			assert.Equal(t, row.expected, spelling)
			assert.Equal(t, row.stress, stress)
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
