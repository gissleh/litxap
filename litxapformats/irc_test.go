package litxapformats

import (
	"fmt"
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestIRC(t *testing.T) {
	table := []struct {
		input      litxap.Line
		output     string
		selections map[int]int
	}{
		{lineOelNgatiKameie, `Oel ngati 01,08kameie.`, nil},
		{lineFikemIlaFyao, `Fìkem ìlä fya'o!`, map[int]int{2: 2}},
		{lineVolaSkeynven, `Vola 01,04skeynven.`, nil},
	}

	for i, row := range table {
		t.Run(fmt.Sprintf("row_%d", i), func(t *testing.T) {
			assert.Equal(t, row.output, row.input.Format(IRCDefaultColors(), row.selections))
		})
	}
}
