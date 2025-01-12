package litxaputil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestApplySuffixes(t *testing.T) {
	table := []struct {
		curr     string
		suffixes string
		expected string
	}{
		{"tì.fme.tok", "teri", "tì.fme.tok.te.ri"},
		{"uvan", "ti", "uvan.ti"},
		{"uran", "it", "ura.nit"},
		{"awkx", "ìl", "aw.kxìl"},
		{"fko", "l", "fkol"},
		{"mo", "t", "mot"},
		{"mo", "o", "mo.o"},
		{"txon", "ejectiveReplacer", "txon.ejectiveReplacer"}, // Ungrammatical, but coverage is coverage.
		{"tì.fme.tok", "ur", "tì.fme.to.kur"},
		{"tsam", "o,ti", "tsa.mo.ti"},
		{"tsa.mo", "ti", "tsa.mo.ti"},
		{"fpom", "ka", "fpom.ka"},
		{"ta.ron", "tswo,tsyìp,o,teri", "ta.ron.tswo.tsyì.po.te.ri"},
		{"'e.kong", "o", "'e.ko.ngo"},
		{"e.yawr", "a", "e.yaw.ra"},
		{"u.van", "ä", "u.va.nä"},
		{"krr", "o", "krr.o"},
		{"kxll", "ä", "kxll.ä"},
		{"po", "r", "por"},
		{"si", "", "si"},
		{"'e.kong", "ä", "'e.ko.ngä"},
		{"'e.kong", "teri", "'e.kong.te.ri"},
		{"u.van.si", "", "u.van.si"},
		{"u.van.si", "yu", "u.van.si.yu"},
		{"u.van.si", "yu,o,ti", "u.van.si.yu.o.ti"},
		{"u.van.si", "yu,o,t", "u.van.si.yu.ot"},
		{"u.van.si", "tswo", "u.van.tswo"},
		{"u.van.su.si", "a", "u.van.su.si.a"},
		{"u.van.si", "tsyìp,yu,tsyìp,it", "u.van.tsyìp.si.yu.tsyì.pit"}, // grammatically dubious
		{"u.van.si", "teri,yu,l", "u.van.te.ri.si.yul"},                 // not grammatically correct
		{"u.van.si", "teri,yu", "u.van.te.ri.si.yu"},                    // not grammatically correct
		{"u.van.si", "o,tswo,ti", "u.va.no.tswo.ti"},                    // not grammatically correct
	}

	for _, row := range table {
		t.Run(fmt.Sprintf("%s -%s", row.curr, row.suffixes), func(t *testing.T) {
			curr := strings.Split(row.curr, ".")
			suffixes := strings.Split(row.suffixes, ",")
			if row.suffixes == "" {
				suffixes = []string{}
			}

			next := ApplySuffixes(curr, suffixes)

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
