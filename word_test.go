package litxap

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRunWord(t *testing.T) {
	table := []struct {
		Raw       string
		Entry     string
		Res       string
		ResStress int
	}{
		{
			Raw: "Fmetok", Entry: "fme.tok",
			Res: "Fme.tok", ResStress: 0,
		},
		{
			Raw: "Tìtusìranìri", Entry: "t·ì.*r·an: tì- <us> -ìri",
			Res: "Tì.tu.sì.ra.nì.ri", ResStress: 3,
		},
		{
			Raw: "fneUvanur", Entry: "u.*van: fne- -ur",
			Res: "fne.U.va.nur", ResStress: 2,
		},
		{
			Raw: "täpeykìyeverkeiup", Entry: "t·er.k·up: <äp,eyk,ìyev,ei>",
			Res: "tä.pey.kì.ye.ver.ke.i.up", ResStress: 4,
		},
		{
			Raw: "tìlamteri", Entry: "tì.*lam: -teri",
			Res: "tì.lam.te.ri", ResStress: 1,
		},
		{
			Raw: "tsukanom", Entry: "k·a.n·om: tsuk-",
			Res: "tsu.ka.nom", ResStress: 1,
		},
		{
			Raw: "tsukkanom", Entry: "k·a.n·om: tsuk-",
			Res: "tsuk.ka.nom", ResStress: 1,
		},
		{
			Raw: "tskxekeng sìsyi", Entry: "tskxe.keng.s··i: <ìsy>",
			Res: "tskxe.keng. .sì.syi", ResStress: 0,
		},
		{
			Raw: "ayskxe", Entry: "tskxe: ay-",
			Res: "ay.skxe", ResStress: 1,
		},
		{
			Raw: "tanlokxe", Entry: "txan.lo.*kxe",
			Res: "tan.lo.kxe", ResStress: 2,
		},
		{
			Raw: "Tsaheyl", Entry: "tsa.heyl: no_stress",
			Res: "Tsa.heyl", ResStress: -1,
		},
		{
			Raw: "taronyu", Entry: "ta.ron: -yu",
			Res: "ta.ron.yu", ResStress: 0,
		},
		{
			Raw: "uvantswo", Entry: "u.*van.si: -tswo: gamer cred",
			Res: "u.van.tswo", ResStress: 1,
		},
		{
			Raw: "narisiyu", Entry: "*na.ri.si: -yu",
			Res: "na.ri.si.yu", ResStress: 0,
		},
		{
			Raw: "nari seyki", Entry: "*na.ri.s··i: <eyk>",
			Res: "na.ri. .sey.ki", ResStress: 0,
		},
		{
			Raw: "srung-susia", Entry: "*srung.s··i: <us> -a",
			Res: "srung.-su.si.a", ResStress: 0,
		},
		{
			Raw: "kemtsyìposiyu", Entry: "*kem.si: -tsyìp-o-yu",
			Res: "kem.tsyì.po.si.yu", ResStress: 0,
		},
		{
			Raw: "kemtsyìposiyuo", Entry: "*kem.si: -tsyìp-o-yu-o",
			Res: "kem.tsyì.po.si.yu.o", ResStress: 0,
		},
		{
			Raw: "si", Entry: "s··i: ",
			Res: "si", ResStress: 0,
		},
		{
			Raw: "rììrmì", Entry: "rì.'ìr: -mì",
			Res: "rì.ìr.mì", ResStress: 0,
		},
		{
			Raw: "ngey", Entry: "nga: -y",
			Res: "ngey", ResStress: 0,
		},
	}

	for _, row := range table {
		t.Run(row.Entry, func(t *testing.T) {
			res, resStress := RunWord(row.Raw, *ParseEntry(row.Entry))

			syllables, stress, root := ParseEntry(row.Entry).GenerateSyllables()
			t.Log("Generated Syllables", syllables)
			t.Log("Generated Stress", stress)
			t.Log("Generated Root", root)

			assert.Equal(t, row.Res, strings.Join(res, "."))
			assert.Equal(t, row.ResStress, resStress)
		})
	}
}
