package litxap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummyDictionary = DummyDictionary{
	"kaltxì":        *ParseEntry("kal.*txì"),
	"ma":            *ParseEntry("ma"),
	"fmetokyu":      *ParseEntry("fme.tok: -yu"),
	"ayhapxìtu":     *ParseEntry("ha.*pxì.tu: ay-"),
	"soaiä":         *ParseEntry("so.*a.i.a: -ä"),
	"ngeyä":         *ParseEntry("nga: -yä"),
	"lu":            *ParseEntry("lu"),
	"oeru":          *ParseEntry("o.e: -ru"),
	"let'eylan":     *ParseEntry("let.*'ey.lan"),
	"nìwotx":        *ParseEntry("nì.*wotx"),
	"oel":           *ParseEntry("o.e: -l"),
	"ngati":         *ParseEntry("nga: -ti"),
	"kameie":        *ParseEntry("k·a.m·e: <ei>: see, see into, understand, know (spiritual sense)"),
	"kameie:0":      *ParseEntry("k··ä: <am,ei>: go"),
	"säkeynven":     *ParseEntry("sä.keyn.*ven"),
	"vola":          *ParseEntry("vol: -a"),
	"tsafneioanghu": *ParseEntry("i.*o.ang: tsa-fne- -hu"),
	"rä'ä":          *ParseEntry("rä.*'ä"),
	"tsaheyl si":    *ParseEntry("tsa.heyl.*s··i"),
	"'eylan":        *ParseEntry("'ey.lan"),
	"tsaheyl":       *ParseEntry("tsa.heyl: no_stress"),
	"soli":          *ParseEntry("s··i: <ol>"),
	"po":            *ParseEntry("po"),
	"ikranhu":       *ParseEntry("ik.ran: -hu"),
	"ke":            *ParseEntry("ke"),
	"a":             *ParseEntry("a"),
	"uvan":          *ParseEntry("u.*van"),
	"uvan soli":     *ParseEntry("u.*van si: <ol>"),
	"po soli":       *ParseEntry("po"),
	"po soli:0":     *ParseEntry("si: <ol>"),
}

func TestRunLine(t *testing.T) {
	table := []struct {
		input    string
		expected Line
	}{
		{
			input: "Kaltxì, ma fmetokyu!",
			expected: Line{
				LinePart{Raw: "Kaltxì", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Kal", "txì"}, 1, dummyDictionary["kaltxì"], false},
				}},
				LinePart{Raw: ", "},
				LinePart{Raw: "ma", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ma"}, 0, dummyDictionary["ma"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "fmetokyu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"fme", "tok", "yu"}, 0, dummyDictionary["fmetokyu"], false},
				}},
				LinePart{Raw: "!"},
			},
		},
		{
			input: "Oel ngati kameie.",
			expected: Line{
				LinePart{Raw: "Oel", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Oel"}, 0, dummyDictionary["oel"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ngati", IsWord: true, Matches: []LinePartMatch{
					{[]string{"nga", "ti"}, 0, dummyDictionary["ngati"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "kameie", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ka", "me", "i", "e"}, 0, dummyDictionary["kameie"], false},
					{[]string{"ka", "me", "i", "e"}, 3, dummyDictionary["kameie:0"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Ayhapxìtu soaiä ngeyä lu oeru let'eylan nìwotx.",
			expected: Line{
				LinePart{Raw: "Ayhapxìtu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Ay", "ha", "pxì", "tu"}, 2, dummyDictionary["ayhapxìtu"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "soaiä", IsWord: true, Matches: []LinePartMatch{
					{[]string{"so", "a", "i", "ä"}, 1, dummyDictionary["soaiä"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ngeyä", IsWord: true, Matches: []LinePartMatch{
					{[]string{"nge", "yä"}, 0, dummyDictionary["ngeyä"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "lu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"lu"}, 0, dummyDictionary["lu"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "oeru", IsWord: true, Matches: []LinePartMatch{
					{[]string{"oe", "ru"}, 0, dummyDictionary["oeru"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "let'eylan", IsWord: true, Matches: []LinePartMatch{
					{[]string{"let", "'ey", "lan"}, 1, dummyDictionary["let'eylan"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "nìwotx", IsWord: true, Matches: []LinePartMatch{
					{[]string{"nì", "wotx"}, 1, dummyDictionary["nìwotx"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Vola skeynven.",
			expected: Line{
				LinePart{Raw: "Vola", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Vo", "la"}, 0, dummyDictionary["vola"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "skeynven", IsWord: true},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Vola säkeynven|skeynven.",
			expected: Line{
				LinePart{Raw: "Vola", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Vo", "la"}, 0, dummyDictionary["vola"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "skeynven", Lookup: "säkeynven", IsWord: true, Matches: []LinePartMatch{
					{[]string{"skeyn", "ven"}, 1, dummyDictionary["säkeynven"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Po tsaheyl soli ikranhu.",
			expected: Line{
				LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Po"}, 0, dummyDictionary["po"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "tsaheyl", IsWord: true, Matches: []LinePartMatch{
					{[]string{"tsa", "heyl"}, -1, dummyDictionary["tsaheyl"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
					{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
				}},
				LinePart{Raw: "."},
			},
		},
	}

	for _, row := range table {
		t.Run(row.input, func(t *testing.T) {
			var res Line
			var err error
			res, err = RunLine(row.input, dummyDictionary)
			assert.NoError(t, err)
			assert.Equal(t, row.expected, res)
		})
	}
}

func TestRunLine_Fail(t *testing.T) {
	line, err := RunLine("Kaltxì, ma kifkey!", BrokenDictionary{})

	assert.Error(t, err)
	assert.Nil(t, line)
	assert.NotErrorIs(t, err, ErrEntryNotFound)
}

func TestParseLine(t *testing.T) {
	table := []struct {
		input    string
		expected Line
	}{
		{
			input: "Ftuea tìfmetok",
			expected: Line{
				LinePart{Raw: "Ftuea", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "tìfmetok", IsWord: true},
			},
		},
		{
			input: "spono-o aean-na-pay",
			expected: Line{
				LinePart{Raw: "spono-o", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "a", IsWord: true},
				LinePart{Raw: "ean", IsWord: true},
				LinePart{Raw: "-"},
				LinePart{Raw: "na", IsWord: true},
				LinePart{Raw: "-"},
				LinePart{Raw: "pay", IsWord: true},
			},
		},
		{
			input: "ean-na-ta'lenga tute",
			expected: Line{
				LinePart{Raw: "ean", IsWord: true},
				LinePart{Raw: "-"},
				LinePart{Raw: "na", IsWord: true},
				LinePart{Raw: "-"},
				LinePart{Raw: "ta'leng", IsWord: true},
				LinePart{Raw: "a", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "tute", IsWord: true},
			},
		},
		{
			input: "Ngäzìka tìkenong-o",
			expected: Line{
				LinePart{Raw: "Ngäzìka", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "tìkenong-o", IsWord: true},
			},
		},
		{
			input: "Fìtìfmetok lu nì'it ngäzìk to pum aham.",
			expected: Line{
				LinePart{Raw: "Fìtìfmetok", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "lu", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "nì'it", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "ngäzìk", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "to", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "pum", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "aham", IsWord: true},
				LinePart{Raw: "."},
			},
		},
		{
			input: "'Awa säkeynven|skeynven angim",
			expected: Line{
				LinePart{Raw: "'Awa", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "skeynven", Lookup: "säkeynven", IsWord: true},
				LinePart{Raw: " "},
				LinePart{Raw: "angim", IsWord: true},
			},
		},
	}

	for _, row := range table {
		t.Run(row.input, func(t *testing.T) {
			assert.Equal(t, row.expected, ParseLine(row.input))
		})
	}
}

func TestLine_UnStressSiVerbParts(t *testing.T) {
	line := Line{
		LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"u", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}

	assert.Equal(t, Line{
		LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"u", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, -1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}, line.UnStressSiVerbParts(dummyDictionary))
}

func TestLine_UnStressSiVerbParts_Negated(t *testing.T) {
	line := Line{
		LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"u", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ke", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ke"}, 0, dummyDictionary["ke"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}

	assert.Equal(t, Line{
		LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"u", "van"}, -1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ke", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ke"}, 0, dummyDictionary["ke"], true},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, -1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}, line.UnStressSiVerbParts(dummyDictionary))

	line[4] = LinePart{Raw: "rä'ä", IsWord: true, Matches: []LinePartMatch{
		{[]string{"rä'ä"}, 1, dummyDictionary["rä'ä"], false},
	}}

	assert.Equal(t, Line{
		LinePart{Raw: "Po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"u", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "rä'ä", IsWord: true, Matches: []LinePartMatch{
			{[]string{"rä'ä"}, 1, dummyDictionary["rä'ä"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, -1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}, line.UnStressSiVerbParts(dummyDictionary))
}

func TestLine_UnStressSiVerbParts_SubClause(t *testing.T) {
	line := Line{
		LinePart{Raw: "Uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"U", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "a", IsWord: true, Matches: []LinePartMatch{
			{[]string{"a"}, 0, dummyDictionary["a"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}

	assert.Equal(t, line, line.UnStressSiVerbParts(dummyDictionary))
}

func TestLine_UnStressSiVerbParts_SentenceBoundary(t *testing.T) {
	line := Line{
		LinePart{Raw: "Uvan", IsWord: true, Matches: []LinePartMatch{
			{[]string{"U", "van"}, 1, dummyDictionary["uvan"], false},
		}},
		LinePart{Raw: ", "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "po", IsWord: true, Matches: []LinePartMatch{
			{[]string{"po"}, 0, dummyDictionary["po"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "ikranhu", IsWord: true, Matches: []LinePartMatch{
			{[]string{"ik", "ran", "hu"}, 0, dummyDictionary["ikranhu"], false},
		}},
		LinePart{Raw: "."},
	}

	assert.Equal(t, line, line.UnStressSiVerbParts(dummyDictionary))
}

func TestLine_UnStressSiVerbParts_NoWordBeforeKe(t *testing.T) {
	line := Line{
		LinePart{Raw: "Ke", IsWord: true, Matches: []LinePartMatch{
			{[]string{"Ke"}, 1, dummyDictionary["ke"], false},
		}},
		LinePart{Raw: " "},
		LinePart{Raw: "soli", IsWord: true, Matches: []LinePartMatch{
			{[]string{"so", "li"}, 1, dummyDictionary["soli"], false},
		}},
		LinePart{Raw: "."},
	}

	assert.Equal(t, line, line.UnStressSiVerbParts(dummyDictionary))
}
