package litxap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomWords(t *testing.T) {
	res := CustomWords([]string{"nor", "no", "ta-*mu", "-ke.a.fkxa.ra", "kel.nì", "te.li.si"}, "")

	sort.Strings(res.(*customWordDictionary).table["nor"]) // Just for the test, because Go maps randomize

	assert.Equal(t, []string{"no: -r ", "nor: "}, res.(*customWordDictionary).table["nor"])
	assert.Equal(t, []string{"ta.*mu: -ri "}, res.(*customWordDictionary).table["tamuri"])

	entries, err := res.LookupEntries("nor")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("no: -r: Custom Word/Name"),
		*ParseEntry("nor: : Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("tamul")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("ta.*mu: -l: Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("kelnur")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("kel.nì: -ur: Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("Kelnìl")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("kel.nì: -l: Custom Word/Name"),
		*ParseEntry("kel.nì: -ìl: Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("keafkxarateri")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("ke.a.fkxa.ra: -teri no_stress: Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("telisit")
	assert.Equal(t, []Entry{
		*ParseEntry("te.li.si: -t: Custom Word/Name"),
	}, entries)

	entries, err = res.LookupEntries("neytiriti")
	assert.ErrorIs(t, err, ErrEntryNotFound)
	assert.Nil(t, entries)
}

func TestCustomWords_WithIDs(t *testing.T) {
	res := CustomWordsWithIDs(map[string]string{
		"nor":    "1",
		"ta-*mu": "2",
		"kel.nì": "3",
	}, "")

	assert.Equal(t, "1", ParseEntry(res.(*customWordDictionary).table["nor"][0]).ID)
	assert.Equal(t, "2", ParseEntry(res.(*customWordDictionary).table["tamul"][0]).ID)
	assert.Equal(t, "3", ParseEntry(res.(*customWordDictionary).table["kelnur"][0]).ID)
}
