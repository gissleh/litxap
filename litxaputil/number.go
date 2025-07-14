package litxaputil

import (
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
		stress = 0 // mune, pukap and kinä have all stress on the first syllable
		ok = true
		return
	}

	for i := 4; i >= 1; i-- {
		powerValue := numberPowerValues[i]
		if number >= powerValue {
			digit := number / powerValue
			number %= powerValue

			if numberPrefixes[digit] != "" {
				syllables = append(syllables, numberPrefixes[digit])
			}

			if number >= 1 && number < 0o10 {
				suffixes := numberSuffixes
				if ordinal {
					suffixes = numberSuffixesOrdinal
				}

				syllables = append(syllables, numberPowers[i]...)
				if number == 1 {
					syllables[len(syllables)-1] += "aw"
					if ordinal {
						syllables = append(syllables, "ve")
					}
				} else {
					syllables = append(syllables[:len(syllables)-1], suffixes[number]...)
				}

				stress = len(syllables) - 1
				if ordinal {
					stress -= 1
				}

				break
			} else {
				syllables = append(syllables, numberPowersClosed[i]...)

				if number == 0 && ordinal {
					lastSyllable := syllables[len(syllables)-1]
					if strings.HasSuffix(lastSyllable, "zam") {
						syllables[len(syllables)-1] = strings.TrimSuffix(lastSyllable, "m")
					}

					syllables = append(syllables, "ve")
					break
				}
			}
		}
	}

	ok = true
	return
}

func ParseNumber(s string) *ParseNumberResult {
	lastPower := 100000
	number := 0
	ordinal := false

	prefix := ""
	if strings.HasPrefix(s, "a") {
		prefix = "a"
		s = s[len("a"):]
	}
	suffix := ""
	if strings.HasSuffix(s, "a") {
		suffix = "a"
		s = s[:len(s)-len("a")]
	}
	if prefix != "" && suffix != "" {
		return nil
	}

	for len(s) > 0 {
		part, next := ParseNumberPart(s)
		if part == nil {
			return nil
		}
		s = next

		// Do not allow lenition
		if part.Lenited && (number > 0 || prefix == "a") {
			return nil
		}

		// Ensure we only get descending powers
		if part.Power >= lastPower {
			return nil
		}
		lastPower = part.Power
		number += part.Value()
		ordinal = part.Ordinal

		if part.Terminal() && len(s) > 0 {
			return nil
		}
	}

	return &ParseNumberResult{
		Value:   number,
		Ordinal: ordinal,
		Prefix:  prefix,
		Suffix:  suffix,
	}
}

type NumberPart struct {
	Multiplier int
	Power      int
	Remainder  int
	Ordinal    bool
	Lenited    bool
}

func (np *NumberPart) Value() int {
	return (np.Multiplier * np.Power) + np.Remainder
}

func (np *NumberPart) Terminal() bool {
	return np.Ordinal || np.Remainder > 0 || np.Power == 1
}

func ParseNumberPart(s string) (*NumberPart, string) {
	for i, roots := range [][]string{numberPowersCombinedClosed[:], numberPowersCombinedClosedOrdinal[:]} {
		for j, root := range roots {
			if j == 0 {
				continue
			}

			if root != "" && root == s {
				return &NumberPart{
					Multiplier: 1,
					Power:      numberPowerValues[j],
					Ordinal:    i == 1,
					Lenited:    false,
				}, ""
			}
		}
	}

	for i, roots := range [][]string{numberRootsCombined[:], numberRootsCombinedLenition[:], numberRootsCombinedOrdinal[:], numberRootsCombinedOrdinalLenition[:]} {
		for j, root := range roots {
			if root != "" && root == s {
				return &NumberPart{
					Multiplier: j,
					Power:      1,
					Ordinal:    i >= 2,
					Lenited:    i%2 == 1,
				}, ""
			}
		}
	}

	for prefixesIndex, prefixes := range [2][8]string{numberPrefixes, numberPrefixesLenited} {
		for i, pre := range prefixes {
			if pre == "" && (i != 1 || prefixesIndex != 0) {
				continue
			}

			if strings.HasPrefix(s, pre) {
				s := s[len(pre):]

				for j, power := range numberPowersCombined {
					if j == 0 {
						continue
					}

					if strings.HasPrefix(s, power) {
						s := s[len(power):]

						suffixes := numberSuffixesZam
						ordinalSuffixes := numberSuffixesZamOrdinal
						if j == 1 {
							suffixes = numberSuffixesVol
							ordinalSuffixes = numberSuffixesVolOrdinal
						}

						for suffixesIndex, suffixes := range [][8]string{ordinalSuffixes, suffixes} {
							for k := len(suffixes) - 1; k >= 0; k-- {
								suf := suffixes[k]
								if strings.HasPrefix(s, suf) {
									s := s[len(suf):]

									return &NumberPart{
										Multiplier: i,
										Power:      numberPowerValues[j],
										Remainder:  k,
										Ordinal:    suffixesIndex == 0,
										Lenited:    prefixesIndex == 1,
									}, s
								}
							}
						}
					}
				}
			}
		}
	}

	return nil, s
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
var numberPowersClosed = [5][]string{{}, {"vol"}, {"zam"}, {"vo", "zam"}, {"za", "zam"}}
var numberRoots = [8][]string{{}, {"'aw"}, {"mu", "ne"}, {"pxey"}, {"tsìng"}, {"mrr"}, {"pu", "kap"}, {"ki", "nä"}}
var numberRootsOrdinal = [8][]string{{}, {"'aw", "ve"}, {"mu", "ve"}, {"pxey", "ve"}, {"tsì", "ve"}, {"mrr", "ve"}, {"pu", "ve"}, {"ki", "ve"}}
var numberPrefixes = [8]string{"", "", "me", "pxe", "tsì", "mrr", "pu", "ki"}
var numberPrefixesLenited = [8]string{"", "", "", "pe", "sì", "", "fu", "hi"}
var numberSuffixes = [8][]string{{}, {"aw"}, {"mun"}, {"pey"}, {"sìng"}, {"mrr"}, {"fu"}, {"hin"}}
var numberSuffixesOrdinal = [8][]string{{"ve"}, {"aw", "ve"}, {"mu", "ve"}, {"pey", "ve"}, {"sì", "ve"}, {"mrr", "ve"}, {"fu", "ve"}, {"hi", "ve"}}
var numberPowersCombinedClosed = [5]string{"", "vol", "zam", "vozam", "zazam"}
var numberPowersCombinedClosedOrdinal = [5]string{"", "volve", "zave", "vozave", "zazave"}
var numberPowersCombined = [5]string{"", "vo", "za", "voza", "zaza"}
var numberSuffixesVol = [8]string{"l", "law", "mun", "pey", "sìng", "mrr", "fu", "hin"}
var numberSuffixesVolOrdinal = [8]string{"lve", "lawve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive"}
var numberSuffixesZam = [8]string{"m", "maw", "mun", "pey", "sìng", "mrr", "fu", "hin"}
var numberSuffixesZamOrdinal = [8]string{"ve", "mawve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive"}
var numberRootsCombined = [8]string{"kew", "'aw", "mune", "pxey", "tsìng", "mrr", "pukap", "kinä"}
var numberRootsCombinedLenition = [8]string{"hew", "aw", "mune", "pey", "sìng", "mrr", "fukap", "hinä"}
var numberRootsCombinedOrdinal = [9]string{"", "'awve", "muve", "pxeyve", "tsìve", "mrrve", "puve", "kive", "volve"}
var numberRootsCombinedOrdinalLenition = [9]string{"", "awve", "muve", "peyve", "sìve", "mrrve", "fuve", "hive", "volve"}
