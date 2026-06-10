package litxap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dummyDictionary = DummyDictionary{
	"kaltxì":        *ParseEntry("kal.*txì"),
	"ma":            *ParseEntry("ma"),
	"fmetok":        *ParseEntry("fme.tok"),
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
	"oe":            *ParseEntry("*o.e"),
	"tsaktap":       *ParseEntry("*tsak.tap"),
	"uvan":          *ParseEntry("u.*van"),
	"si":            *ParseEntry("s··i"),
	"uvan soli":     *ParseEntry("u.*van s··i: <ol>"),
	"po soli":       *ParseEntry("po"),
	"po soli:0":     *ParseEntry("si: <ol>"),
	"tslolam":       *ParseEntry("tsl··am: <ol>"),
	"futa":          *ParseEntry("*fu.ta"),
	"frapo":         *ParseEntry("*fra.po"),
	"frapo:0":       *ParseEntry("po: fra-"),
	"fmetan":        *ParseEntry("*fme.tan"),
	"fmetan:0":      *ParseEntry("fme.*tan"),
	"'efu":          *ParseEntry("*'·e.f·u"),
	"nitram":        *ParseEntry("nit.*ram"),
	"mì":            *ParseEntry("mì"),
	"oer":           *ParseEntry("*o.e: -r"),
	"tìnitram":      *ParseEntry("tì.nit.*ram"),
	"ngeyn":         *ParseEntry("ngeyn"),
	"talun":         *ParseEntry("ta.*lun"),
	"talun:0":       *ParseEntry("ta.*lun"),
	"holahaw":       *ParseEntry("*h·a.h·aw: <ol>"),
	"nìtam":         *ParseEntry("nì.*tam"),
}

var lineOelNgatiKameie = Line{
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
	}},
	LinePart{Raw: "."},
}

var lineKaltxiMaFmetokyu = Line{
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
}

var lineKaltxiMaFmetan = Line{
	LinePart{Raw: "Kaltxì", IsWord: true, Matches: []LinePartMatch{
		{[]string{"Kal", "txì"}, 1, dummyDictionary["kaltxì"], false},
	}},
	LinePart{Raw: ", "},
	LinePart{Raw: "ma", IsWord: true, Matches: []LinePartMatch{
		{[]string{"ma"}, 0, dummyDictionary["ma"], false},
	}},
	LinePart{Raw: " "},
	LinePart{Raw: "Fmetan", IsWord: true, Matches: []LinePartMatch{
		{[]string{"Fme", "tan"}, 0, dummyDictionary["fmetan"], false},
		{[]string{"Fme", "tan"}, 1, dummyDictionary["fmetan:0"], false},
	}},
	LinePart{Raw: "!"},
}

var lineVolaSkeynven = Line{
	LinePart{Raw: "Vola", IsWord: true, Matches: []LinePartMatch{
		{[]string{"Vo", "la"}, 0, dummyDictionary["vola"], false},
	}},
	LinePart{Raw: " "},
	LinePart{Raw: "skeynven", IsWord: true},
	LinePart{Raw: "."},
}

var lineFmetokBad = Line{
	LinePart{Raw: "Vola", IsWord: true, Matches: []LinePartMatch{
		{[]string{"Fme", "tök"}, 0, dummyDictionary["fmetok"], false},
	}},
}

func TestRunLine(t *testing.T) {
	table := []struct {
		input    string
		expected Line
	}{
		{
			input:    "Kaltxì, ma fmetokyu!",
			expected: lineKaltxiMaFmetokyu,
		},
		{
			input:    "Oel ngati kameie.",
			expected: lineOelNgatiKameie,
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
			input:    "Vola skeynven.",
			expected: lineVolaSkeynven,
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
			input: "Lu oer tìnitram.", // This line crashes 1.13.2
			expected: Line{
				LinePart{Raw: "Lu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Lu"}, 0, dummyDictionary["lu"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "oer", IsWord: true, Matches: []LinePartMatch{
					{[]string{"oer"}, 0, dummyDictionary["oer"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "tìnitram", IsWord: true, Matches: []LinePartMatch{
					{[]string{"tì", "nit", "ram"}, 2, dummyDictionary["tìnitram"], false},
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
		{
			input: "Tslolam oel futa ke frapo ke tslolam.",
			expected: Line{
				LinePart{Raw: "Tslolam", IsWord: true, Matches: []LinePartMatch{
					{[]string{"Tslo", "lam"}, 1, dummyDictionary["tslolam"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "oel", IsWord: true, Matches: []LinePartMatch{
					{[]string{"oel"}, 0, dummyDictionary["oel"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "futa", IsWord: true, Matches: []LinePartMatch{
					{[]string{"fu", "ta"}, 0, dummyDictionary["futa"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ke", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ke"}, 0, dummyDictionary["ke"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "frapo", IsWord: true, Matches: []LinePartMatch{
					{[]string{"fra", "po"}, 0, dummyDictionary["frapo"], false},
					{[]string{"fra", "po"}, 1, dummyDictionary["frapo:0"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ke", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ke"}, 0, dummyDictionary["ke"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "tslolam", IsWord: true, Matches: []LinePartMatch{
					{[]string{"tslo", "lam"}, 1, dummyDictionary["tslolam"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Oe tsaktap si.",
			expected: Line{
				LinePart{Raw: "Oe", IsWord: true, Matches: []LinePartMatch{
					{[]string{"O", "e"}, 0, dummyDictionary["oe"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "tsaktap", IsWord: true, Matches: []LinePartMatch{
					{[]string{"tsak", "tap"}, 0, dummyDictionary["tsaktap"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "si", IsWord: true, Matches: []LinePartMatch{
					{[]string{"si"}, 0, dummyDictionary["si"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "Oe uvan si.",
			expected: Line{
				LinePart{Raw: "Oe", IsWord: true, Matches: []LinePartMatch{
					{[]string{"O", "e"}, 0, dummyDictionary["oe"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "uvan", IsWord: true, Matches: []LinePartMatch{
					{[]string{"u", "van"}, 1, dummyDictionary["uvan"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "si", IsWord: true, Matches: []LinePartMatch{
					{[]string{"si"}, 0, dummyDictionary["si"], false},
				}},
				LinePart{Raw: "."},
			},
		},
		{
			input: "'EFU OE NITRAM!",
			expected: Line{
				LinePart{Raw: "'EFU", IsWord: true, Matches: []LinePartMatch{
					{[]string{"'E", "FU"}, 0, dummyDictionary["'efu"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "OE", IsWord: true, Matches: []LinePartMatch{
					{[]string{"OE"}, 0, dummyDictionary["oe"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "NITRAM", IsWord: true, Matches: []LinePartMatch{
					{[]string{"NIT", "RAM"}, 1, dummyDictionary["nitram"], false},
				}},
				LinePart{Raw: "!"},
			},
		},
		{
			input: "'Efu oe ngeyn talun oe ke holahaw nìtam.",
			expected: Line{
				LinePart{Raw: "'Efu", IsWord: true, Matches: []LinePartMatch{
					{[]string{"'E", "fu"}, 0, dummyDictionary["'efu"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "oe", IsWord: true, Matches: []LinePartMatch{
					{[]string{"oe"}, 0, dummyDictionary["oe"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ngeyn", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ngeyn"}, 0, dummyDictionary["ngeyn"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "talun", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ta", "lun"}, 1, dummyDictionary["talun"], false},
					{[]string{"ta", "lun"}, 1, dummyDictionary["talun:0"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "oe", IsWord: true, Matches: []LinePartMatch{
					{[]string{"o", "e"}, 0, dummyDictionary["oe"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "ke", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ke"}, 0, dummyDictionary["ke"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "holahaw", IsWord: true, Matches: []LinePartMatch{
					{[]string{"ho", "la", "haw"}, 1, dummyDictionary["holahaw"], false},
				}},
				LinePart{Raw: " "},
				LinePart{Raw: "nìtam", IsWord: true, Matches: []LinePartMatch{
					{[]string{"nì", "tam"}, 1, dummyDictionary["nìtam"], false},
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

			if t.Failed() {
				t.Log("dummyLineFormatter preview:")
				t.Log("  Expected:", row.expected.Format(&dummyLineFormatter{}, nil))
				t.Log("  Actual  :", res.Format(&dummyLineFormatter{}, nil))
			}
		})
	}

	t.Run("_all_as_RunLines", func(t *testing.T) {
		input := make([]string, 0)
		expected := make([]Line, 0)
		for _, row := range table {
			input = append(input, row.input)
			expected = append(expected, row.expected)
		}

		res, err := RunLines(input, dummyDictionary)
		assert.NoError(t, err)
		assert.Equal(t, expected, res)
	})
}

func TestRunLine_Fail(t *testing.T) {
	line, err := RunLine("Kaltxì, ma kifkey!", BrokenDictionary{})

	assert.Error(t, err)
	assert.Nil(t, line)
	assert.NotErrorIs(t, err, ErrEntryNotFound)
}

func TestRunLines_Fail(t *testing.T) {
	line, err := RunLines([]string{"Kaltxì, ma kifkey!"}, BrokenDictionary{})

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

type dummyLineFormatter struct{}

func (f *dummyLineFormatter) LinePartTags(lp LinePart, stress int) (string, string) {
	switch stress {
	case LPSNotWord:
		return "", ""
	case LPSAmbiguousMatches:
		return "[AM]", "[/AM]"
	case LPSNoMatches:
		return "[NM]", "[/NM]"
	case LPSAnyStress:
		return "[AS]", "[/AS]"
	default:
		return "[S]", "[/S]"
	}
}

func (f *dummyLineFormatter) StressedSyllableTags() (string, string) {
	return "{", "}"
}

var lineFikemIlaFyao = Line{
	LinePart{Raw: "Fìkem", IsWord: true, Matches: []LinePartMatch{
		{[]string{"Fì", "kem"}, 1, dummyDictionary["fìkem"], false},
		{[]string{"Fì", "kem"}, 1, dummyDictionary["fìkem:0"], false},
	}},
	LinePart{Raw: " "},
	LinePart{Raw: "ìlä", IsWord: true, Matches: []LinePartMatch{
		{[]string{"ì", "lä"}, 0, dummyDictionary["ìlä"], false},
		{[]string{"ì", "lä"}, 1, dummyDictionary["ìlä:0"], false},
	}},
	LinePart{Raw: " "},
	LinePart{Raw: "fya'o", IsWord: true, Matches: []LinePartMatch{
		{[]string{"fya", "'o"}, 0, dummyDictionary["fya'o"], false},
	}},
	LinePart{Raw: "!"},
}

func TestLine_Format(t *testing.T) {
	table := []struct {
		input      Line
		output     string
		selections map[int]int
	}{
		{lineOelNgatiKameie, "[S]Oel[/S] [S]{nga}ti[/S] [S]{ka}meie[/S].", nil},
		{lineKaltxiMaFmetokyu, "[S]Kal{txì}[/S], [S]ma[/S] [S]{fme}tokyu[/S]!", nil},
		{lineKaltxiMaFmetan, "[S]Kal{txì}[/S], [S]ma[/S] [AM]Fmetan[/AM]!", map[int]int{999999: 1}},
		{lineKaltxiMaFmetan, "[S]Kal{txì}[/S], [S]ma[/S] [S]{Fme}tan[/S]!", map[int]int{4: 0}},
		{lineKaltxiMaFmetan, "[S]Kal{txì}[/S], [S]ma[/S] [S]Fme{tan}[/S]!", map[int]int{4: 1}},
		{lineVolaSkeynven, "[S]{Vo}la[/S] [NM]skeynven[/NM].", nil},
		{lineFikemIlaFyao, "[S]Fì{kem}[/S] [AS]ìlä[/AS] [S]{fya}'o[/S]!", map[int]int{2: 2}},
	}

	for _, row := range table {
		t.Run(row.output, func(t *testing.T) {
			assert.Equal(t, row.output, row.input.Format(&dummyLineFormatter{}, row.selections))
		})
	}
}

func TestLine_IPA(t *testing.T) {
	table := []struct {
		input      Line
		delim      string
		output     string
		selections map[int]int
		err        string
	}{
		{lineOelNgatiKameie, "", "wɛl ˈŋati ˈkamɛiɛ.", nil, ""},
		{lineKaltxiMaFmetokyu, "", "kalˈtʼɪ, ma ˈfmɛtok̚ju!", nil, ""},
		{lineFikemIlaFyao, ".", "fɪ.ˈkɛm ɪ.ˈlæ ˈfja.ʔo!", map[int]int{2: 1}, ""},
		{lineFikemIlaFyao, ".", "fɪ.ˈkɛm ˈɪ.læ ˈfja.ʔo!", map[int]int{2: 2}, ""},
		{lineKaltxiMaFmetan, "", "kalˈtʼɪ, ma ˈfmɛtan!", map[int]int{4: 0}, ""},
		{lineKaltxiMaFmetan, "", "kalˈtʼɪ, ma fmɛˈtan!", map[int]int{4: 1}, ""},
		{lineVolaSkeynven, "", "", nil, fmt.Sprintf("no matches for line[%d] (%#+v)", 2, "skeynven")},
		{lineFmetokBad, "", "", nil, "unknown symbols [\"ö\", \"ök\"] in syllable tök"},
	}

	for _, row := range table {
		t.Run(row.output, func(t *testing.T) {
			output, err := row.input.IPA(row.selections, row.delim)

			if row.err != "" {
				assert.ErrorContains(t, err, row.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, row.output, output)
			}
		})
	}
}
