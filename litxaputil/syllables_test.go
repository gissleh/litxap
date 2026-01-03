package litxaputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSyllables(t *testing.T) {
	table := []struct {
		Input    string
		Expected string
		First    Syllable
	}{
		{"tok", "tok", Syllable{Onset: "t", Body: "o", Coda: "k"}},
		{"fmi", "fmi", Syllable{PreOnset: "f", Onset: "m", Body: "i"}},
		{"fmetok", "fme-tok", Syllable{PreOnset: "f", Onset: "m", Body: "e"}},
		{"tskxet", "tskxet", Syllable{PreOnset: "ts", Onset: "kx", Body: "e", Coda: "t"}},
		{"tskxeti", "tskxe-ti", Syllable{PreOnset: "ts", Onset: "kx", Body: "e"}},
		{"sivako", "si-va-ko", Syllable{Onset: "s", Body: "i"}},
		{"prrkxentrrkrr", "prr-kxen-trr-krr", Syllable{Onset: "p", Body: "rr"}},
		{"trr'ong", "trr-'ong", Syllable{Onset: "t", Body: "rr"}},
		{"nga'prrnen", "nga'-prr-nen", Syllable{Onset: "ng", Body: "a", Coda: "'"}},
		{"kxllyu", "kxll-yu", Syllable{Onset: "kx", Body: "ll"}},
		{"var", "var", Syllable{Onset: "v", Body: "a", Coda: "r"}},
		{"vll", "vll", Syllable{Onset: "v", Body: "ll"}},
		{"vol", "vol", Syllable{Onset: "v", Body: "o", Coda: "l"}},
		{"tswayon", "tswa-yon", Syllable{PreOnset: "ts", Onset: "w", Body: "a"}},
		{"meoauniaea", "me-o-a-u-ni-a-e-a", Syllable{Onset: "m", Body: "e"}},
		{"tlalim", "tla-lim (irregulars: tl)", Syllable{Onset: "t", Irregular: "l", Body: "a"}},
		{"mangkwan", "mang-kwan (irregulars: kw)", Syllable{Onset: "m", Body: "a", Coda: "ng"}},
		{"kreytu'um", "krey-tu-'um (irregulars: kr)", Syllable{Onset: "k", Irregular: "r", Body: "ey"}},
		{"kramtran", "kram-tran (irregulars: kr,tr)", Syllable{Onset: "k", Irregular: "r", Body: "a", Coda: "m"}},
		{"", "", Syllable{}},
		{"ehanis", "<nil>", Syllable{}},
		{"keln", "<nil>", Syllable{}},
		{"ftspa", "<nil>", Syllable{}},
		{"xr", "<nil>", Syllable{}},
		{"rr", "<nil>", Syllable{}},
		{"ll", "<nil>", Syllable{}},
		{"fha", "<nil>", Syllable{}},
		{"svane", "<nil>", Syllable{}},
	}

	for _, row := range table {
		t.Run(row.Input, func(t *testing.T) {
			res := SplitSyllables(row.Input)
			assert.Equal(t, row.Expected, res.String())
			if row.First != (Syllable{}) {
				assert.Equal(t, res[0], row.First)
			}
		})
	}
}
