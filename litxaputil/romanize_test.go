package litxaputil

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanizeIPA(t *testing.T) {
	table := []struct {
		curr     string
		expected [][][]string
		stress   [][]int
	}{
		// One syllable
		{"ɛ", [][][]string{{{"e"}}}, [][]int{{-1}}},
		{"ʔawk'", [][][]string{{{"'awkx"}}}, [][]int{{-1}}},
		{"fko", [][][]string{{{"fko"}}}, [][]int{{-1}}},
		{"mo", [][][]string{{{"mo"}}}, [][]int{{-1}}},
		{"t'on", [][][]string{{{"txon"}}}, [][]int{{-1}}},
		{"t͡sam", [][][]string{{{"tsam"}}}, [][]int{{-1}}},
		{"fpom", [][][]string{{{"fpom"}}}, [][]int{{-1}}},
		{"kṛ", [][][]string{{{"krr"}}}, [][]int{{-1}}},
		{"k'ḷ", [][][]string{{{"kxll"}}}, [][]int{{-1}}},
		{"po", [][][]string{{{"po"}}}, [][]int{{-1}}},
		{"sk'awŋ", [][][]string{{{"skxawng"}}}, [][]int{{-1}}},
		{"tʃok'", [][][]string{{{"chokx"}}}, [][]int{{-1}}},
		{"ʃ·awm", [][][]string{{{"shawm"}}}, [][]int{{-1}}},
		{"sko", [][][]string{{{"sko"}}}, [][]int{{-1}}},
		{"ftɛ", [][][]string{{{"fte"}}}, [][]int{{-1}}},

		// Multi syllable
		{"tɪ.ˈfmɛ.tok̚", [][][]string{{{"tì", "fme", "tok"}}}, [][]int{{1}}},
		{"u.ˈvan", [][][]string{{{"u", "van"}}}, [][]int{{1}}},
		{"ˈu.ɾan", [][][]string{{{"u", "ran"}}}, [][]int{{0}}},
		{"ˈt·a.ɾ·on", [][][]string{{{"ta", "ron"}}}, [][]int{{0}}},
		{"ˈʔɛ.koŋ", [][][]string{{{"'e", "kong"}}}, [][]int{{0}}},
		{"ɛ.ˈjawɾ", [][][]string{{{"e", "yawr"}}}, [][]int{{1}}},
		{"sæ.ˈfɾɪp̚", [][][]string{{{"sä", "frìp"}}}, [][]int{{1}}},
		{"ˈuk̚.jom", [][][]string{{{"uk", "yom"}}}, [][]int{{0}}},
		{"ˈmo.ʔat̚", [][][]string{{{"mo", "'at"}}}, [][]int{{0}}},
		{"t͡suʔ.ˈtɛj", [][][]string{{{"tsu'", "tey"}}}, [][]int{{1}}},
		{"ɾi.ˈnaʔ", [][][]string{{{"ri", "na'"}}}, [][]int{{1}}},
		{"ˈfk'a.ɾa", [][][]string{{{"fkxa", "ra"}}}, [][]int{{0}}},
		{"ˌmɛ.o.a.u.ni.a.ˈɛ.a", [][][]string{{{"me", "o", "a", "u", "ni", "a", "e", "a"}}}, [][]int{{6}}},

		// Flexible syllable stress
		{"aj.ˈfo] or [ˈaj.fo", [][][]string{{{"ay", "fo"}}, {{"ay", "fo"}}}, [][]int{{1}, {0}}},
		{"ˈɪ.læ] or [ɪ.ˈlæ", [][][]string{{{"ì", "lä"}}, {{"ì", "lä"}}}, [][]int{{0}, {1}}},
		{"ˈmɪ.fa] or [mɪ.ˈfa", [][][]string{{{"mì", "fa"}}, {{"mì", "fa"}}}, [][]int{{0}, {1}}},
		{"ˈt͡sa.kɛm] or [t͡sa.ˈkɛm", [][][]string{{{"tsa", "kem"}}, {{"tsa", "kem"}}}, [][]int{{0}, {1}}},
		{"t͡sa.ˈt͡sɛŋ] or [ˈt͡sa.t͡sɛŋ", [][][]string{{{"tsa", "tseng"}}, {{"tsa", "tseng"}}}, [][]int{{1}, {0}}},

		// Multiple pronunciation
		{"nɪ.aw.ˈno.mʊm] or [naw.ˈno.mʊm", [][][]string{{{"nì", "aw", "no", "mum"}}, {{"naw", "no", "mum"}}}, [][]int{{2}, {1}}},
		{"nɪ.aj.ˈwɛŋ] or [naj.ˈwɛŋ", [][][]string{{{"nì", "ay", "weng"}}, {{"nay", "weng"}}}, [][]int{{2}, {1}}},
		{"tɪ.sjɪ.maw.nʊn.ˈʔi] or [t͡sjɪ.maw.nʊn.ˈʔi", [][][]string{{{"tì", "syì", "maw", "nun", "'i"}}, {{"tsyì", "maw", "nun", "'i"}}}, [][]int{{4}, {3}}},
		{"tɪ.sæ.ˈfpɪl.jɛwn] or [t͡sæ.ˈfpɪl.jɛwn", [][][]string{{{"tì", "sä", "fpìl", "yewn"}}, {{"tsä", "fpìl", "yewn"}}}, [][]int{{2}, {1}}},

		// Multiple words
		{"ˈut.ɾa.ja ˈmok.ɾi", [][][]string{{{"ut", "ra", "ya"}, {"mok", "ri"}}}, [][]int{{0, 0}}},
		{"t͡sa.ˈhɛjl s·i", [][][]string{{{"tsa", "heyl"}, {"si"}}}, [][]int{{1, -1}}},
		{"ˈnɪ.ˌju ˈjoɾ.kɪ", [][][]string{{{"nì", "yu"}, {"yor", "kì"}}}, [][]int{{0, 0}}},
		{"t͡sawl sl·u", [][][]string{{{"tsawl"}, {"slu"}}}, [][]int{{-1, -1}}},
		{"o.ˈɪsss s·i", [][][]string{{{"o", "ìsss"}, {"si"}}}, [][]int{{1, -1}}},

		// Empty string
		{"", [][][]string{}, [][]int{}},
		{"  ", [][][]string{}, [][]int{}},
	}

	for _, row := range table {
		t.Run(row.curr, func(t *testing.T) {
			spelling, stress := RomanizeIPA(row.curr)
			assert.Equal(t, row.expected, spelling)
			assert.Equal(t, row.stress, stress)
		})
	}
}

func TestSyllableToIPA(t *testing.T) {
	table := []struct {
		input    string
		expected string
	}{
		{"tskxe", "t͡sk'ɛ"},
		{"keng", "kɛŋ"},
		{"fme", "fmɛ"},
		{"tok", "tok̚"},
		{"lawr", "lawɾ"},
		{"ran", "ɾan"},
		{"syon", "sjon"},
		{"tsway", "t͡swaj"},
		{"on", "on"},
		{"u", "u"},
		{"van", "van"},
	}

	for _, row := range table {
		t.Run(row.expected, func(t *testing.T) {
			res, err := SyllableToIPA(row.input)
			assert.NoError(t, err)
			assert.Equal(t, row.expected, res)
		})
	}
}

func TestSyllablesToIPA(t *testing.T) {
	table := []struct {
		input            string
		delimiter        string
		strongEmphasises []int
		weakEmphasises   []int
		expected         string
	}{
		{"fme.tok", ".", []int{0}, []int{}, "ˈfmɛ.tok̚"},
		{"tsak.tap", ".", []int{0}, []int{}, "ˈt͡sak̚.tap̚"},
		{"tal.i.o.ang", ".", []int{0}, []int{2}, "ˈtal.i.ˌo.aŋ"},
		{"shawm", ".", []int{}, []int{}, "ʃawm"},
		{"syu.ra", ".", []int{1}, []int{}, "sju.ˈɾa"},
		{"tsyey.tsyìp", ".", []int{0}, []int{}, "ˈt͡sjɛj.t͡sjɪp̚"},
		{"chey.chìp", ".", []int{0}, []int{}, "ˈtʃɛj.tʃɪp̚"},
		{"ad.ge", ".", []int{1}, []int{}, "ad.ˈgɛ"},
		{"a.ba", ".", []int{1}, []int{}, "a.ˈba"},
		{"zaw.prr.te'", "-", []int{1}, []int{}, "zaw-ˈpṛ-tɛʔ"},
		{"me.o.a.u.ni.a.e.a", "", []int{6}, []int{0}, "ˌmɛoauniaˈɛa"},
	}

	for _, row := range table {
		t.Run(row.expected, func(t *testing.T) {
			res, err := SyllablesToIPA(
				strings.Split(row.input, "."),
				row.delimiter,
				row.strongEmphasises,
				row.weakEmphasises,
			)

			assert.NoError(t, err)
			assert.Equal(t, row.expected, res)
		})
	}
}

func TestWriteSyllablesAsIPATo_Errors(t *testing.T) {
	set := []string{"ˌ", "k", "a", "l", ".", "ˈ", "t'", "ɪ", "t", "f", "m", "ɛ", "t", "o", "k", "̚"}

	for i := range set {
		assert.Error(t, WriteSyllablesAsIPATo(
			&badStringWriter{whitelist: set[:i]},
			[]string{"kal", "txì", "ma", "tì", "fme", "tok"},
			".", []int{1, 4}, []int{0},
		))
	}

	assert.NoError(t, WriteSyllablesAsIPATo(
		&badStringWriter{whitelist: set},
		[]string{"kal", "txì"},
		".", []int{1}, []int{0},
	))

	assert.ErrorContains(t, WriteSyllablesAsIPATo(
		&badStringWriter{whitelist: set},
		[]string{"txxap"},
		".", []int{}, []int{},
	), "unknown symbols [\"x\", \"xa\"] in syllable txxap")
}

type badStringWriter struct {
	whitelist []string
}

func (b *badStringWriter) WriteString(s string) (n int, err error) {
	if slices.Contains(b.whitelist, s) {
		return len(s), nil
	}

	return 0, fmt.Errorf("badStringWriter-ìl ke tung fìpamrelit alu: %s", s)
}
