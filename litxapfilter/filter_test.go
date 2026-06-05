package litxapfilter

import (
	"slices"
	"strings"
	"testing"
	"unsafe"

	"github.com/gissleh/litxap"
	"github.com/gissleh/litxap/litxapformats"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLine_ApplyFilter(t *testing.T) {
	table := []struct {
		input    string
		expected litxap.Line
		filters  []Filter
	}{
		{
			input: "Kaltxì, ma kxitx.",
			expected: litxap.Line{
				{Raw: "Kaltì", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Kal", "tì"}, 1, dummyDictionary.entry("kaltxì", 0), false},
				}},
				{Raw: ", "},
				{Raw: "ma", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"ma"}, 0, dummyDictionary.entry("ma", 0), false},
				}},
				{Raw: " "},
				{Raw: "kit", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"kit"}, 0, dummyDictionary.entry("kxitx", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{dummyFilterEjectiveHater},
		},
		{
			input: "Oel ngati kameie, ma RumaUt.",
			expected: litxap.Line{
				{Raw: "Wel", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Wel"}, 0, dummyDictionary.entry("oel", 0), false},
				}},
				{Raw: " "},
				{Raw: "ngati", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"nga", "ti"}, 0, dummyDictionary.entry("ngati", 0), false},
				}},
				{Raw: " "},
				{Raw: "kameye", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"ka", "me", "ye"}, 0, dummyDictionary.entry("kameie", 0), false},
				}},
				{Raw: ", "},
				{Raw: "ma", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"ma"}, 0, dummyDictionary.entry("ma", 0), false},
				}},
				{Raw: " "},
				{Raw: "RumaWt", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Ru", "maWt"}, 0, dummyDictionary.entry("rumaut", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{
				DiphthongFromWeakVowel,
				ReanalyzeDiphthongs,
				SpellOeAsWe,
			},
		},
		{
			input: "fmetokyu fmeretok.",
			expected: litxap.Line{
				{Raw: "fmetokyu", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"fme", "tok", "yu"}, 0, dummyDictionary.entry("fmetokyu", 0), false},
				}},
				{Raw: " "},
				{Raw: "retok", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"re", "tok"}, 0, dummyDictionary.entry("fmeretok", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{dummyFilterNextEliminator("fme")},
		},

		{
			input: "Oe tìng nari.",
			expected: litxap.Line{
				{Raw: "Oe", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"O", "e"}, 0, dummyDictionary.entry("oe", 0), false},
				}},
				{Raw: " "},
				{Raw: "tì", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"tì"}, 0, dummyDictionary.entry("tìng", 0), false},
				}},
				{Raw: " "},
				{Raw: "nari", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"na", "ri"}, 0, dummyDictionary.entry("nari", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{NasalAssimilation},
		},
		{
			input: "Fmetan mal lu!",
			expected: litxap.Line{
				{Raw: "Fmeta", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Fme", "ta"}, 0, dummyDictionary.entry("fmetan", 0), false},
					{[]string{"Fme", "ta"}, 1, dummyDictionary.entry("fmetan", 1), false},
				}},
				{Raw: " "},
				{Raw: "mal", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"mal"}, 0, dummyDictionary.entry("mal", 0), false},
				}},
				{Raw: " "},
				{Raw: "lu", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"lu"}, 0, dummyDictionary.entry("lu", 0), false},
				}},
				{Raw: "!"},
			},
			filters: []Filter{NasalAssimilation},
		},
		{
			input: "Fmetan?",
			expected: litxap.Line{
				{Raw: "Fmetan", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Fme", "tan"}, 0, dummyDictionary.entry("fmetan", 0), false},
					{[]string{"tan"}, 0, dummyDictionary.entry("fmetan", 1), false},
				}},
				{Raw: "?"},
			},
			filters: []Filter{NasalAssimilation, dummyFilterCurrEliminatorAtIndex(1, "Fme")},
		},
		{
			input: "Sänume säpeyki.",
			expected: litxap.Line{
				{Raw: "Snume", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Snu", "me"}, 0, dummyDictionary.entry("sänume", 0), false},
				}},
				{Raw: " "},
				{Raw: "speyki", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"spey", "ki"}, 1, dummyDictionary.entry("säpeyki", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{SaeRemover},
		},
		{
			input: "Pori fpomtoKX sì fpomroN yo'.",
			expected: litxap.Line{
				{Raw: "Pori", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Po", "ri"}, 0, dummyDictionary.entry("pori", 0), false},
				}},
				{Raw: " "},
				{Raw: "fpomtoK", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"fpom", "toK"}, 1, dummyDictionary.entry("fpomtokx", 0), false},
				}},
				{Raw: " "},
				{Raw: "sì", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"sì"}, 0, dummyDictionary.entry("sì", 0), false},
				}},
				{Raw: " "},
				{Raw: "fpomroN", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"fpom", "roN"}, 1, dummyDictionary.entry("fpomron", 0), false},
				}},
				{Raw: " "},
				{Raw: "yo'", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"yo'"}, 0, dummyDictionary.entry("yo'", 0), false},
				}},
				{Raw: "."},
			},
			filters: []Filter{
				NasalAssimilation,
				DemoteEjectivesBeforeConsonants,
			},
		},
		{
			input: "Sunu oer aymauti, sì ayspxam nìayfo!",
			expected: litxap.Line{
				{Raw: "Sunu", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"Su", "nu"}, 0, dummyDictionary.entry("sunu", 0), false},
				}},
				{Raw: " "},
				{Raw: "oer", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"oer"}, 0, dummyDictionary.entry("oer", 0), false},
				}},
				{Raw: " "},
				{Raw: "aymauti", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"ay", "ma", "u", "ti"}, 1, dummyDictionary.entry("aymauti", 0), false},
				}},
				{Raw: ", "},
				{Raw: "sayspxa", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"say", "spxa"}, 1, dummyDictionary.entry("ayspxam", 0), false},
				}},
				{Raw: " "},
				{Raw: "nayfo", IsWord: true, Matches: []litxap.LinePartMatch{
					{[]string{"nay", "fo"}, 1, dummyDictionary.entry("nìayfo", 0), false},
				}},
				{Raw: "!"},
			},
			filters: []Filter{
				ElideMiSiNiBeforeAy,
				NasalAssimilation,
			},
		},
	}

	for _, row := range table {
		t.Run(row.input, func(t *testing.T) {
			line, err := litxap.RunLine(row.input, dummyDictionary)
			require.NoError(t, err)

			line = ApplyFilters(line, row.filters...)
			assert.Equal(t, row.expected, line)

			if t.Failed() {
				t.Log("Result as markdown:", line.Format(litxapformats.DiscordMarkdown(), nil))
			}
		})
	}
}

func TestLine_ApplyFilter_NoChange(t *testing.T) {
	line, err := litxap.RunLine("Kaltxì, ma kxitx!", dummyDictionary)
	assert.NoError(t, err)

	line2 := ApplyFilter(line, func(_, _ *FilterTarget) (*string, *string) {
		return nil, nil
	})

	assert.Same(t, unsafe.SliceData(line), unsafe.SliceData(line2))
}

var dummyFilterEjectiveHater Filter = func(curr, next *FilterTarget) (currChange *string, nextChange *string) {
	if strings.ContainsRune(curr.Syllable, 'x') {
		ejectiveLess := strings.ReplaceAll(curr.Syllable, "x", "")
		currChange = &ejectiveLess
	}

	return
}

func dummyFilterNextEliminator(blacklist ...string) Filter {
	return func(curr, next *FilterTarget) (currChange *string, nextChange *string) {
		if next != nil && slices.Contains(blacklist, next.Syllable) {
			return nil, new(string)
		}

		return nil, nil
	}
}

func dummyFilterCurrEliminatorAtIndex(matchIndex int, blacklist ...string) Filter {
	return func(curr, next *FilterTarget) (currChange *string, nextChange *string) {
		if curr.MatchIndex == matchIndex && slices.Contains(blacklist, curr.Syllable) {
			return new(string), nil
		}

		return nil, nil
	}
}

type DummyDictionary map[string]string

func (dictionary DummyDictionary) LookupEntries(word string) ([]litxap.Entry, error) {
	if entries, ok := dictionary[strings.ToLower(word)]; ok {
		lines := strings.Split(entries, "\n")
		res := make([]litxap.Entry, 0, len(lines))
		for _, line := range lines {
			res = append(res, *litxap.ParseEntry(line))
		}

		return res, nil
	}

	return nil, litxap.ErrEntryNotFound
}

func (dictionary DummyDictionary) entry(word string, i int) litxap.Entry {
	res, _ := dictionary.LookupEntries(word)
	return res[i]
}

var dummyDictionary = DummyDictionary{
	"kaltxì":   "kal.*txì",
	"ma":       "ma",
	"fmetokyu": "fme.tok: -yu",
	"fmeretok": "fm·e.t·ok: <er>",
	"lu":       "lu",
	"oel":      "o.e: -l",
	"ngati":    "nga: -ti",
	"kameie":   "k·a.m·e: <ei>: see, see into, understand, know (spiritual sense)\nk··ä: <am,ei>: go",
	"rumaut":   "ru.ma.ut",
	"oe":       "*o.e",
	"si":       "s··i",
	"fmetan":   "*fme.tan\nfme.*tan",
	"tìng":     "t··ìng",
	"nari":     "na.ri",
	"mal":      "mal",
	"kxitx":    "kxitx",
	"sänume":   "sä.*nu.me",
	"säpeyki":  "s··i: <äp,eyk>",
	"pori":     "po: -ri",
	"fpomtokx": "fpom.*tokx",
	"sì":       "sì",
	"fpomron":  "fpom.*ron",
	"yo'":      "y··o'",
	"sunu":     "su.nu",
	"oer":      "*o.e: -r",
	"aymauti":  "*ma.u.ti: ay-",
	"ayspxam":  "spxam: ay-",
	"nìayfo":   "ay.*fo: nì-",
}
