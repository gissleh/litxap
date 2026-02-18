package litxapformats

import (
	"testing"

	"github.com/gissleh/litxap"
	"github.com/stretchr/testify/assert"
)

func TestDiscordMarkdown(t *testing.T) {
	table := []struct {
		input      litxap.Line
		output     string
		selections map[int]int
	}{
		{lineOelNgatiKameie, "Oel __nga__ti \\*kameie(AMBIGUOUS).", nil},
		{lineOelNgatiKameie, "Oel __nga__ti __ka__meie.", map[int]int{4: 0}},
		{lineKaltxiMaFmetokyu, "Kal__txì__, ma __fme__tokyu!", nil},
		{lineVolaSkeynven, "__Vo__la \\*skeynven(NO MATCHES).", nil},
	}

	for _, row := range table {
		t.Run(row.output, func(t *testing.T) {
			assert.Equal(t, row.output, row.input.Format(DiscordMarkdown(), row.selections))
		})
	}
}
