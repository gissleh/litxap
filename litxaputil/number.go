package litxaputil

import (
	"strconv"
	"strings"
)

func GenerateNumber(number int, ordinal bool) (syllables []string, stress int, ok bool) {
	if number <= 0 || number > 0o77777 {
		ok = false
		return
	}

	if number < 0o10 {
		if ordinal {
			syllables = append(syllables, numberRootsOrdinal[number]...)
		} else {
			syllables = append(syllables, numberRoots[number]...)
		}

		ok = true
		return
	}

	for i := 4; i >= 1; i-- {
		powerValue := numberPowerValues[i]
		if number >= powerValue {
			digit := number / powerValue
			number %= powerValue

			stress = len(syllables)
			if len(syllables) > 0 {
				syllables = syllables[:len(syllables)-1]
			}
			if len(numberPowers[i]) == 3 {
				stress += 1
			}
			if numberPrefixes[digit] != "" {
				syllables = append(syllables, numberPrefixes[digit])
			}
			syllables = append(syllables, numberPowers[i]...)
		}
	}

	if number > 0 {
		suffixes := numberSuffixes
		if ordinal {
			suffixes = numberSuffixesOrdinal
		}

		stress = len(syllables) - 1
		if number == 1 {
			syllables = append(syllables[:len(syllables)-1], syllables[len(syllables)-1]+suffixes[number][0])
			syllables = append(syllables, suffixes[number][1:]...)
		} else {
			syllables = append(syllables[:len(syllables)-1], suffixes[number]...)
		}
	} else {
		if len(syllables[len(syllables)-1]) == 1 {
			syllables = append(syllables[:len(syllables)-2], syllables[len(syllables)-2]+syllables[len(syllables)-1])
		}
		if ordinal {
			if syllables[len(syllables)-1] == "zam" {
				syllables = append(syllables[:len(syllables)-1], "za", "ve")
			} else {
				syllables = append(syllables, "ve")
			}
		}
	}

	ok = true
	return
}

func ParseNumber(s string) *ParseNumberResult {
	res := ParseNumberResult{}
	if strings.HasPrefix(s, "a") && !strings.HasPrefix(s, "aw") {
		res.Prefix = "a"
		s = s[len("a"):]
	} else if strings.HasSuffix(s, "a") {
		res.Suffix = "a"
		s = s[:len(s)-len("a")]
	}

	// Parse a32, 32, 32ve, etc...
	if n, err := strconv.Atoi(strings.TrimSuffix(s, "ve")); err == nil {
		s2 := strconv.Itoa(n)
		without := strings.TrimPrefix(s, s2)
		if without == "ve" {
			res.Ordinal = true
			res.Value = n
			return &res
		} else {
			res.Value = n
			return &res
		}
	}

	// Parse a°40, °40ve, etc...
	if st := strings.TrimPrefix(strings.TrimPrefix(s, "0o"), "°"); st != s {
		n, err := strconv.ParseInt(strings.TrimSuffix(st, "ve"), 8, 32)
		if err != nil {
			return nil
		}

		s2 := strconv.FormatInt(n, 8)
		without := strings.TrimPrefix(st, s2)
		if without == "ve" {
			res.Ordinal = true
			res.Value = int(n)
			return &res
		} else {
			res.Value = int(n)
			return &res
		}
	}

	for n, root := range numberRootsCombined {
		if s == root {
			res.Value = n
			return &res
		}
	}
	for n, root := range numberRootsCombinedOrdinal {
		if s == root {
			res.Value = n
			res.Ordinal = true
			return &res
		}
	}
	if res.Prefix != "a" {
		for n, root := range numberRootsCombinedLenition {
			if s == root {
				res.Value = n
				return &res
			}
		}
		for n, root := range numberRootsCombinedOrdinalLenition {
			if s == root {
				res.Value = n
				res.Ordinal = true
				return &res
			}
		}
	}

	prevPower := numberPowerValues[len(numberPowerValues)-1] * 10 // Ensure it starts on an unattainable power.

	for {
		if res.Value > 0 {
			numberSuffixes := numberSuffixesZam
			ordinalSuffixes := numberSuffixesZamOrdinal
			if prevPower == 0o10 {
				numberSuffixes = numberSuffixesVol
				ordinalSuffixes = numberSuffixesVolOrdinal
			}

			for n, suff := range numberSuffixes {
				if s == suff {
					res.Value += n
					return &res
				}
			}
			for n, suff := range ordinalSuffixes {
				if s == suff {
					res.Ordinal = true
					res.Value += n
					return &res
				}
			}
		}

		multiplier := 1
		for n, pref := range numberPrefixes {
			if pref != "" && strings.HasPrefix(s, pref) {
				multiplier = n
				s = s[len(pref):]
				break
			}
		}

		power := 0
		for i := len(numberPowersCombined) - 1; i >= 0; i-- {
			powerName := numberPowersCombined[i]
			if len(powerName) == 0 {
				continue
			}

			if strings.HasPrefix(s, powerName) {
				n := numberPowerValues[i]
				if n >= prevPower {
					return nil
				}

				power = n
				prevPower = n
				s = s[len(powerName):]
				break
			}
		}
		if power > 0 {
			res.Value += multiplier * power
			continue
		}

		return nil // If something is found, it's continued
	}
}

type ParseNumberResult struct {
	Value   int
	Ordinal bool
	Prefix  string
	Suffix  string
}

func (r *ParseNumberResult) GenerateSyllables(affixed bool) (syllables []string, stress int, ok bool) {
	syllables, stress, ok = GenerateNumber(r.Value, r.Ordinal)
	if !ok {
		return
	}

	if affixed {
		if r.Prefix != "" {
			syllables = append(syllables[:1], syllables...)
			syllables[0] = r.Prefix
		}
		if r.Suffix != "" {
			syllables = append(syllables, r.Suffix)
		}
	}

	return
}

var numberPowerValues = [5]int{0o1, 0o10, 0o100, 0o1000, 0o10000}
var numberPowers = [5][]string{{}, {"vo", "l"}, {"za", "m"}, {"vo", "za", "m"}, {"za", "za", "m"}}
var numberRoots = [8][]string{{}, {"'aw"}, {"mu", "ne"}, {"pxey"}, {"tsìng"}, {"mrr"}, {"pu", "kap"}, {"ki", "nä"}}
var numberRootsOrdinal = [8][]string{{}, {"'aw", "ve"}, {"mu", "ve"}, {"pxey", "ve"}, {"tsì", "ve"}, {"mrr", "ve"}, {"pu", "ve"}, {"ki", "ve"}}
var numberPrefixes = [8]string{"", "", "me", "pxe", "tsì", "mrr", "pu", "ki"}
var numberSuffixes = [8][]string{{}, {"aw"}, {"mun"}, {"pey"}, {"sìng"}, {"mrr"}, {"fu"}, {"hin"}}
var numberSuffixesOrdinal = [8][]string{{"ve"}, {"aw", "ve"}, {"mu", "ve"}, {"pey", "ve"}, {"sì", "ve"}, {"mrr", "ve"}, {"fu", "ve"}, {"hi", "ve"}}

var numberPowersCombined = [5]string{"", "vo", "za", "voza", "zaza"}
var numberSuffixesVol = [8]string{"l", "law", "mun", "pey", "sìng", "mrr", "fu", "hin"}
var numberSuffixesVolOrdinal = [8]string{"lve", "lawve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive"}
var numberSuffixesZam = [8]string{"m", "maw", "mun", "pey", "sìng", "mrr", "fu", "hin"}
var numberSuffixesZamOrdinal = [8]string{"ve", "mawve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive"}
var numberRootsCombined = [8]string{"kew", "'aw", "mune", "pxey", "tsìng", "mrr", "pukap", "kinä"}
var numberRootsCombinedLenition = [8]string{"hew", "aw", "mune", "pey", "sìng", "mrr", "fukap", "hinä"}
var numberRootsCombinedOrdinal = [9]string{"", "'awve", "muve", "pxeyve", "tsìve", "mrrve", "puve", "kive", "volve"}
var numberRootsCombinedOrdinalLenition = [9]string{"", "awve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive", "volve"}
