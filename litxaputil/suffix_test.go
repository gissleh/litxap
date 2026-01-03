package litxaputil

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplySuffixes(t *testing.T) {
	table := []struct {
		curr     string
		isVerb   bool
		suffixes string
		expected string
	}{
		{"tì.fme.tok", false, "teri", "tì.fme.tok.te.ri"},
		{"uvan", false, "ti", "uvan.ti"},
		{"uran", false, "it", "ura.nit"},
		{"awkx", false, "ìl", "aw.kxìl"},
		{"fko", false, "l", "fkol"},
		{"mo", false, "t", "mot"},
		{"mo", false, "o", "mo.o"},
		{"txon", false, "ejectiveReplacer", "txon.ejectiveReplacer"}, // Ungrammatical, but coverage is coverage.
		{"tì.fme.tok", false, "ur", "tì.fme.to.kur"},
		{"tsam", false, "o,ti", "tsa.mo.ti"},
		{"tsa.mo", false, "ti", "tsa.mo.ti"},
		{"fpom", false, "ka", "fpom.ka"},
		{"ta.ron", true, "tswo,tsyìp,o,teri", "ta.ron.tswo.tsyì.po.te.ri"},
		{"'e.kong", false, "o", "'e.ko.ngo"},
		{"e.yawr", false, "a", "e.yaw.ra"},
		{"u.van", false, "ä", "u.va.nä"},
		{"krr", false, "o", "krr.o"},
		{"kxll", false, "ä", "kxll.ä"},
		{"po", false, "r", "por"},
		{"si", true, "", "si"},
		{"'e.kong", false, "ä", "'e.ko.ngä"},
		{"'e.kong", false, "teri", "'e.kong.te.ri"},
		{"te.li.si", false, "t", "te.li.sit"},
		{"u.van.si", true, "", "u.van.si"},
		{"u.van.si", true, "yu", "u.van.si.yu"},
		{"u.van.si", true, "yu,o,ti", "u.van.si.yu.o.ti"},
		{"u.van.si", true, "yu,o,t", "u.van.si.yu.ot"},
		{"u.van.si", true, "tswo", "u.van.tswo"},
		{"u.van.su.si", false, "a", "u.van.su.si.a"},
		{"u.van.si", true, "tsyìp,yu,tsyìp,it", "u.van.tsyìp.si.yu.tsyì.pit"}, // grammatically dubious
		{"u.van.si", true, "teri,yu,l", "u.van.te.ri.si.yul"},                 // not grammatically correct
		{"u.van.si", true, "teri,yu", "u.van.te.ri.si.yu"},                    // not grammatically correct
		{"u.van.si", true, "o,tswo,ti", "u.va.no.tswo.ti"},                    // not grammatically correct
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s -%s", row.curr, row.suffixes), func(t *testing.T) {
			curr := strings.Split(row.curr, ".")
			suffixes := strings.Split(row.suffixes, ",")
			if row.suffixes == "" {
				suffixes = []string{}
			}

			next := ApplySuffixes(curr, suffixes, row.isVerb)

			assert.Equal(t, row.expected, strings.Join(next, "."))
		})
	}
}

func TestApplySuffixes_Panic(t *testing.T) {
	badSuffix := Suffix{
		reanalysis:    -19392,
		syllableSplit: []string{"blarg"},
	}
	assert.Panics(t, func() { badSuffix.Apply([]string{"stuff"}) })
	assert.Panics(t, func() { badSuffix.Apply([]string{}) })
	assert.Panics(t, func() { findSuffix("teri").Apply([]string{}) })
}
