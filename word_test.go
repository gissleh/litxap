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
			Raw: "Ayoe", Entry: "o.e: ay-",
			Res: "A.yo.e", ResStress: 1,
		},
		{
			Raw: "Ayoe", Entry: "ay.*o.e",
			Res: "Ay.o.e", ResStress: 1,
		},
		{
			Raw: "Ayoeti", Entry: "o.e: ay- -ti",
			Res: "Ay.oe.ti", ResStress: 1,
		},
		{
			Raw: "Ayoeteri", Entry: "ay.*o.e: -teri",
			Res: "Ay.oe.te.ri", ResStress: 1,
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
		{
			Raw: "Teykìran", Entry: "t·ì.*r·an: <eyk>",
			Res: "Tey.kì.ran", ResStress: 2,
		},
		{
			Raw: "Erinan", Entry: "·i.*n·an: <er>",
			Res: "E.ri.nan", ResStress: 1,
		},
		{
			Raw: "Inatsan", Entry: "·i.*n·an: <ats>",
			Res: "I.na.tsan", ResStress: 0,
		},
		{
			Raw: "Eraho", Entry: "·a.*h·o: <er>",
			Res: "E.ra.ho", ResStress: 2,
		},
		{
			Raw: "TSUKinan", Entry: "·i.*n·an: tsuk-",
			Res: "TSU.Ki.nan", ResStress: 1,
		},
		{
			Raw: "inanTSWO", Entry: "·i.*n·an: -tswo",
			Res: "i.nan.TSWO", ResStress: 1,
		},
		{
			Raw: "Inanyu", Entry: "·i.*n·an: -yu",
			Res: "I.nan.yu", ResStress: 1,
		},
		{
			Raw: "pxeylan", Entry: "'ey.lan: pxe-",
			Res: "pxey.lan", ResStress: 0,
		},
		{
			Raw: "ayeylan", Entry: "'ey.lan: ay-",
			Res: "a.yey.lan", ResStress: 1,
		},
		{
			Raw: "oengal", Entry: "o.*eng: -l",
			Res: "oe.ngal", ResStress: 0,
		},
		{
			Raw: "meoeng", Entry: "o.*eng: me-",
			Res: "me.o.eng", ResStress: 2,
		},
		{
			Raw: "ayoengal", Entry: "ay.*o.eng: -l",
			Res: "ay.oe.ngal", ResStress: 1,
		},
		{
			Raw: "oengteri", Entry: "o.*eng: -teri",
			Res: "oeng.te.ri", ResStress: 0,
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
