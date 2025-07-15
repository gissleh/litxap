package litxaputil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateNumber(t *testing.T) {
	table := []struct {
		number       int
		ordinal      bool
		resSyllables string
		resStress    int
	}{
		{
			number: 0o100000, ordinal: false,
			resSyllables: "", resStress: 0,
		},
		{
			number: 0o6, ordinal: false,
			resSyllables: "pu.kap", resStress: 0,
		},
		{
			number: 0o10, ordinal: false,
			resSyllables: "vol", resStress: 0,
		},
		{
			number: 0o10, ordinal: true,
			resSyllables: "vol.ve", resStress: 0,
		},
		{
			number: 0o11, ordinal: true,
			resSyllables: "vo.law.ve", resStress: 1,
		},
		{
			number: 0o11, ordinal: false,
			resSyllables: "vo.law", resStress: 1,
		},
		{
			number: 0o30, ordinal: false,
			resSyllables: "pxe.vol", resStress: 0,
		},
		{
			number: 0o100, ordinal: false,
			resSyllables: "zam", resStress: 0,
		},
		{
			number: 0o100, ordinal: true,
			resSyllables: "za.ve", resStress: 0,
		},
		{
			number: 0o3000, ordinal: false,
			resSyllables: "pxe.vo.zam", resStress: 0,
		},
		{
			number: 0o34, ordinal: false,
			resSyllables: "pxe.vo.sìng", resStress: 2,
		},
		{
			number: 0o3004, ordinal: false,
			resSyllables: "pxe.vo.za.sìng", resStress: 3,
		},
		{
			number: 0o3004, ordinal: true,
			resSyllables: "pxe.vo.za.sì.ve", resStress: 3,
		},
		{
			number: 0o43270, ordinal: false,
			resSyllables: "tsì.za.zam.pxe.vo.zam.me.zam.ki.vol", resStress: 0,
		},
		{
			number: 0o43272, ordinal: false,
			resSyllables: "tsì.za.zam.pxe.vo.zam.me.zam.ki.vo.mun", resStress: 10,
		},
		{
			number: 0o63217, ordinal: true,
			resSyllables: "pu.za.zam.pxe.vo.zam.me.zam.vo.hi.ve", resStress: 9,
		},
		{
			number: 0o5010, ordinal: false,
			resSyllables: "mrr.vo.zam.vol", resStress: 0,
		},
		{
			number: 0o5020, ordinal: false,
			resSyllables: "mrr.vo.zam.me.vol", resStress: 0,
		},
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s(%d)", row.resSyllables, row.resStress), func(t *testing.T) {
			syllables, stress, ok := GenerateNumber(row.number, row.ordinal)
			assert.Equal(t, row.resSyllables != "", ok)
			assert.Equal(t, row.resSyllables, strings.Join(syllables, "."))
			assert.Equal(t, row.resStress, stress)
		})
	}
}

func TestParseNumberPart(t *testing.T) {
	table := []struct {
		Input string
		Res   *NumberPart
		Next  string
	}{
		{"", nil, ""},
		{"mune", &NumberPart{Multiplier: 2, Power: 1}, ""},
		{"pukap", &NumberPart{Multiplier: 6, Power: 1}, ""},
		{"sìng", &NumberPart{Multiplier: 4, Power: 1, Lenited: true}, ""},
		{"kive", &NumberPart{Multiplier: 7, Power: 1, Ordinal: true}, ""},
		{"hive", &NumberPart{Multiplier: 7, Power: 1, Ordinal: true, Lenited: true}, ""},
		{"vol", &NumberPart{Multiplier: 1, Power: 0o10}, ""},
		{"volaw", &NumberPart{Multiplier: 1, Power: 0o10, Remainder: 1}, ""},
		{"vomun", &NumberPart{Multiplier: 1, Power: 0o10, Remainder: 2}, ""},
		{"zam", &NumberPart{Multiplier: 1, Power: 0o100}, ""},
		{"vozam", &NumberPart{Multiplier: 1, Power: 0o1000}, ""},
		{"zazam", &NumberPart{Multiplier: 1, Power: 0o10000}, ""},
		{"volve", &NumberPart{Multiplier: 1, Power: 0o10, Ordinal: true}, ""},
		{"zave", &NumberPart{Multiplier: 1, Power: 0o100, Ordinal: true}, ""},
		{"mevol", &NumberPart{Multiplier: 2, Power: 0o10}, ""},
		{"mevolaw", &NumberPart{Multiplier: 2, Power: 0o10, Remainder: 1}, ""},
		{"kizam", &NumberPart{Multiplier: 7, Power: 0o100, Remainder: 0}, ""},
		{"puzave", &NumberPart{Multiplier: 6, Power: 0o100, Remainder: 0, Ordinal: true}, ""},
		{"hivozave", &NumberPart{Multiplier: 7, Power: 0o1000, Remainder: 0, Ordinal: true, Lenited: true}, ""},
		{"hivozamawve", &NumberPart{Multiplier: 7, Power: 0o1000, Remainder: 1, Ordinal: true, Lenited: true}, ""},
		{"kizazamkivozamkizamkivohin", &NumberPart{Multiplier: 7, Power: 0o10000, Remainder: 0}, "kivozamkizamkivohin"},
		{"kivozamkizamkivohin", &NumberPart{Multiplier: 7, Power: 0o1000, Remainder: 0}, "kizamkivohin"},
		{"kizamkivohin", &NumberPart{Multiplier: 7, Power: 0o100, Remainder: 0}, "kivohin"},
		{"kivohin", &NumberPart{Multiplier: 7, Power: 0o10, Remainder: 7}, ""},
	}

	for _, row := range table {
		t.Run(row.Input, func(t *testing.T) {
			res, next := ParseNumberPart(row.Input)
			assert.Equal(t, row.Res, res)
			assert.Equal(t, row.Next, next)
		})
	}
}

func TestParseNumber(t *testing.T) {
	table := []struct {
		Input  string
		Output *ParseNumberResult
	}{
		{"'awa", &ParseNumberResult{Value: 1, Suffix: "a"}},
		{"a'awa", nil},
		{"amune", &ParseNumberResult{Value: 2, Prefix: "a"}},
		{"mrrve", &ParseNumberResult{Value: 5, Ordinal: true}},
		{"akive", &ParseNumberResult{Value: 7, Ordinal: true, Prefix: "a"}},
		{"volvea", &ParseNumberResult{Value: 8, Ordinal: true, Suffix: "a"}},
		{"avol", &ParseNumberResult{Value: 0o10, Prefix: "a"}},
		{"zama", &ParseNumberResult{Value: 0o100, Suffix: "a"}},
		{"vozama", &ParseNumberResult{Value: 0o1000, Suffix: "a"}},
		{"azazam", &ParseNumberResult{Value: 0o10000, Prefix: "a"}},
		{"volaw", &ParseNumberResult{Value: 0o11}},
		{"volawve", &ParseNumberResult{Value: 0o11, Ordinal: true}},
		{"vomuve", &ParseNumberResult{Value: 0o12, Ordinal: true}},
		{"vomun", &ParseNumberResult{Value: 0o12}},
		{"vopey", &ParseNumberResult{Value: 0o13}},
		{"mevomuna", &ParseNumberResult{Value: 0o22, Suffix: "a"}},
		{"avopey", &ParseNumberResult{Value: 0o13, Prefix: "a"}},
		{"avofuve", &ParseNumberResult{Value: 0o16, Prefix: "a", Ordinal: true}},
		{"mezamaw", &ParseNumberResult{Value: 0o201}},
		{"mezapey", &ParseNumberResult{Value: 0o203}},
		{"mezampxevol", &ParseNumberResult{Value: 0o230}},
		{"tsìzampxevol", &ParseNumberResult{Value: 0o430}},
		{"mezazamkivozamtsìzamvomrr", &ParseNumberResult{Value: 0o27415}},
		{"kivozamezazatsìzavomrr", nil},
		{"pxevoltsìzam", nil},
		{"pxevol kaltxì, oer syaw fko kelsara aylì'u", nil},
		{"kizamawpxevohin", nil},
		{"", nil},
		{"a", nil},
		{"aa", nil},

		{"ahive", nil},
		{"afukap", nil},
		{"hive", &ParseNumberResult{Value: 7, Ordinal: true}},
		{"fukap", &ParseNumberResult{Value: 6}},
	}

	for _, row := range table {
		t.Run(row.Input, func(t *testing.T) {
			assert.Equal(t, row.Output, ParseNumber(row.Input))
		})
	}
}

func TestParseNumber_Exhaustive(t *testing.T) {
	results := [6]ParseNumberResult{
		{Value: 0},
		{Value: 0, Ordinal: true},
		{Value: 0, Prefix: "a"},
		{Value: 0, Suffix: "a"},
		{Value: 0, Ordinal: true, Prefix: "a"},
		{Value: 0, Ordinal: true, Suffix: "a"},
	}

	for n := 1; n <= 0o77777; n++ {
		for i, result := range results {
			result.Value = n

			syllables, _, ok := result.GenerateSyllables(true)
			assert.True(t, ok)
			assert.Equal(t, &result, ParseNumber(strings.Join(syllables, "")))

			if t.Failed() {
				t.Log("Failed at:", n, i, strings.Join(syllables, "."))
				return
			}
		}
	}
}

func TestParseNumber_NotOk(t *testing.T) {
	result := ParseNumberResult{Value: 0}
	syllables, _, ok := result.GenerateSyllables(true)
	assert.False(t, ok)
	assert.Empty(t, syllables)

	result2 := ParseNumberResult{Value: 10000000000}
	syllables, _, ok = result2.GenerateSyllables(true)
	assert.False(t, ok)
	assert.Empty(t, syllables)
}
