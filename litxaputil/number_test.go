package litxaputil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNaviNumber(t *testing.T) {
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
			resSyllables: "tsì.za.za.pxe.vo.za.me.za.ki.vol", resStress: 0,
		},
		{
			number: 0o43272, ordinal: false,
			resSyllables: "tsì.za.za.pxe.vo.za.me.za.ki.vo.mun", resStress: 10,
		},
		{
			number: 0o63217, ordinal: true,
			resSyllables: "pu.za.za.pxe.vo.za.me.za.vo.hi.ve", resStress: 9,
		},
		{
			number: 0o5010, ordinal: false,
			resSyllables: "mrr.vo.za.vol", resStress: 0,
		},
		{
			number: 0o5020, ordinal: false,
			resSyllables: "mrr.vo.za.me.vol", resStress: 0,
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
		{"mezazakivozatsìzavomrr", &ParseNumberResult{Value: 0o27415}},
		{"kivozamezazatsìzavomrr", nil},

		{"1a", &ParseNumberResult{Value: 1, Suffix: "a"}},
		{"a17", &ParseNumberResult{Value: 17, Prefix: "a"}},
		{"312", &ParseNumberResult{Value: 312}},
		{"312ve", &ParseNumberResult{Value: 312, Ordinal: true}},
		{"a1942ve", &ParseNumberResult{Value: 1942, Ordinal: true, Prefix: "a"}},
		{"a1942vea", nil},
		{"a1942vvea", nil},

		{"°5a", &ParseNumberResult{Value: 0o5, Suffix: "a"}},
		{"a°13", &ParseNumberResult{Value: 0o13, Prefix: "a"}},
		{"°312", &ParseNumberResult{Value: 0o312}},
		{"°312ve", &ParseNumberResult{Value: 0o312, Ordinal: true}},
		{"°1742vea", &ParseNumberResult{Value: 0o1742, Ordinal: true, Suffix: "a"}},
		{"a°1742vea", nil},
		{"a°1742vvea", nil},

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

	for n := 1; n < 0o77777; n++ {
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
