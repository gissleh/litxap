package litxap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	res := CustomWords([]string{"nor", "no", "ta-*mu", "-ke.a.fkxa.ra", "kel.nì"})
	assert.Equal(t, []string{"nor: ", "no: -r "}, res.(*customWordDictionary).table["nor"])
	assert.Equal(t, []string{"ta.*mu: -ri "}, res.(*customWordDictionary).table["tamuri"])

	entries, err := res.LookupEntries("nor")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("nor: : Custom Name"),
		*ParseEntry("no: -r: Custom Name"),
	}, entries)

	entries, err = res.LookupEntries("tamul")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("ta.*mu: -l: Custom Name"),
	}, entries)

	entries, err = res.LookupEntries("kelnur")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("kel.nì: -ur: Custom Name"),
	}, entries)

	entries, err = res.LookupEntries("keafkxarateri")
	assert.NoError(t, err)
	assert.Equal(t, []Entry{
		*ParseEntry("ke.a.fkxa.ra: -teri no_stress: Custom Name"),
	}, entries)

	entries, err = res.LookupEntries("neytiriti")
	assert.ErrorIs(t, err, ErrEntryNotFound)
	assert.Nil(t, entries)
}
