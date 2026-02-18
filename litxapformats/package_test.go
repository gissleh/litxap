package litxapformats

import "github.com/gissleh/litxap"

var dummyDictionary = map[string]litxap.Entry{
	"kaltxì":    *litxap.ParseEntry("kal.*txì"),
	"ma":        *litxap.ParseEntry("ma"),
	"fmetokyu":  *litxap.ParseEntry("fme.tok: -yu"),
	"oel":       *litxap.ParseEntry("o.e: -l"),
	"nga":       *litxap.ParseEntry("nga"),
	"ngati":     *litxap.ParseEntry("nga: -ti"),
	"kameie":    *litxap.ParseEntry("k·a.m·e: <ei>: see, see into, understand, know (spiritual sense)"),
	"kameie:0":  *litxap.ParseEntry("k··ä: <am,ei>: go"),
	"säkeynven": *litxap.ParseEntry("sä.keyn.*ven"),
	"vola":      *litxap.ParseEntry("vol: -a"),
	"fìkem":     *litxap.ParseEntry("fì.*kem"),
	"fìkem:0":   *litxap.ParseEntry("kem: fì-"),
	"ìlä":       *litxap.ParseEntry("*ì.lä"),
	"ìlä:0":     *litxap.ParseEntry("ì.*lä"),
	"fya'o":     *litxap.ParseEntry("*fya.'o"),
}

var lineOelNgatiKameie = litxap.Line{
	litxap.LinePart{Raw: "Oel", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"Oel"}, 0, dummyDictionary["oel"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "ngati", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"nga", "ti"}, 0, dummyDictionary["ngati"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "kameie", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"ka", "me", "i", "e"}, 0, dummyDictionary["kameie"], false},
		{[]string{"ka", "me", "i", "e"}, 3, dummyDictionary["kameie:0"], false},
	}},
	litxap.LinePart{Raw: "."},
}

var lineFikemIlaFyao = litxap.Line{
	litxap.LinePart{Raw: "Fìkem", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"Fì", "kem"}, 1, dummyDictionary["fìkem"], false},
		{[]string{"Fì", "kem"}, 1, dummyDictionary["fìkem:0"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "ìlä", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"ì", "lä"}, 0, dummyDictionary["ìlä"], false},
		{[]string{"ì", "lä"}, 1, dummyDictionary["ìlä:0"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "fya'o", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"fya", "'o"}, 0, dummyDictionary["fya'o"], false},
	}},
	litxap.LinePart{Raw: "!"},
}

var lineKaltxiMaFmetokyu = litxap.Line{
	litxap.LinePart{Raw: "Kaltxì", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"Kal", "txì"}, 1, dummyDictionary["kaltxì"], false},
	}},
	litxap.LinePart{Raw: ", "},
	litxap.LinePart{Raw: "ma", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"ma"}, 0, dummyDictionary["ma"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "fmetokyu", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"fme", "tok", "yu"}, 0, dummyDictionary["fmetokyu"], false},
	}},
	litxap.LinePart{Raw: "!"},
}

var lineVolaSkeynven = litxap.Line{
	litxap.LinePart{Raw: "Vola", IsWord: true, Matches: []litxap.LinePartMatch{
		{[]string{"Vo", "la"}, 0, dummyDictionary["vola"], false},
	}},
	litxap.LinePart{Raw: " "},
	litxap.LinePart{Raw: "skeynven", IsWord: true},
	litxap.LinePart{Raw: "."},
}
