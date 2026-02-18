package litxapformats

import (
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestCompactHTML(t *testing.T) {
	table := []struct {
		input      litxap.Line
		output     string
		selections map[int]int
	}{
		{lineOelNgatiKameie, `<span>Oel</span> <span><u>nga</u>ti</span> <span class="am">kameie</span>.`, nil},
		{lineFikemIlaFyao, `<span>Fì<u>kem</u></span> <span class="as">ìlä</span> <span><u>fya</u>'o</span>!`, map[int]int{2: 2}},
		{lineVolaSkeynven, `<span><u>Vo</u>la</span> <span class="nm">skeynven</span>.`, nil},
	}

	for _, row := range table {
		t.Run(row.output, func(t *testing.T) {
			assert.Equal(t, row.output, row.input.Format(CompactHTML(), row.selections))
		})
	}
}
