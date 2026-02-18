package litxapformats

import (
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestBBCode(t *testing.T) {
	table := []struct {
		input      litxap.Line
		output     string
		selections map[int]int
	}{
		{lineOelNgatiKameie, "Oel [u]nga[/u]ti [color=yellow]kameie[/color].", nil},
		{lineOelNgatiKameie, "Oel [u]nga[/u]ti [u]ka[/u]meie.", map[int]int{4: 0}},
		{lineKaltxiMaFmetokyu, "Kal[u]txì[/u], ma [u]fme[/u]tokyu!", nil},
		{lineVolaSkeynven, "[u]Vo[/u]la [color=red]skeynven[/color].", nil},
		{lineFikemIlaFyao, "Fì[u]kem[/u] [color=skyblue]ìlä[/color] [u]fya[/u]'o!", map[int]int{2: 2}},
	}

	for _, row := range table {
		t.Run(row.output, func(t *testing.T) {
			assert.Equal(t, row.output, row.input.Format(BBCode(), row.selections))
		})
	}
}
