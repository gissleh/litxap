package litxap

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNumberDictionary_LookupEntries(t *testing.T) {
	table := []struct {
		Lookup  string
		Results string
	}{
		{"azave", "za.ve: a-: Ordinal number °100 (64)"},
		{"amrr", "mrr: a-: Number 5"},
		{"mezavea", "me.za.ve: -a: Ordinal number °200 (128)"},
		{"Tsìvol", "tsì.vol: : Number °40 (32)"},
		{"mevozam", "me.vo.zam: : Number °2000 (1024)"},
		{"mevozave", "me.vo.za.ve: : Ordinal number °2000 (1024)"},
		{"mrrvomrr", "mrr.vo.*mrr: : Number °55 (45)"},
		{"mrrvozam", "mrr.vo.zam: : Number °5000 (2560)"},
		{"mrrvozamvol", "mrr.vo.zam.vol: : Number °5010 (2568)"},
		{"mrrvozammevol", "mrr.vo.zam.me.vol: : Number °5020 (2576)"},
		{"mezamvolaw", "me.zam.vo.*law: : Number °211 (137)"},
		{"mrrvomrrr", ""},
		{"amrra", ""},
	}

	for _, row := range table {
		t.Run(row.Lookup, func(t *testing.T) {
			res, _ := (&NumberDictionary{}).LookupEntries(row.Lookup)
			resStr := strings.Builder{}
			for i, res := range res {
				if i > 0 {
					resStr.WriteString(";")
				}

				resStr.WriteString(res.String())
			}

			assert.Equal(t, row.Results, resStr.String())
		})
	}
}
