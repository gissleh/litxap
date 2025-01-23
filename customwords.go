package litxap

import (
	"fmt"
	"strings"
)

// CustomWords generates a dictionary for custom words like names and so on. It has a hardcoded set of
// adpositions, so that must be changed if we get a new one.
func CustomWords(names []string) Dictionary {
	table := make(map[string][]string)

	for _, name := range names {
		beforeTrim := len(name)
		name := strings.TrimPrefix(name, "-")
		noStress := ""
		if len(name) != beforeTrim {
			noStress = "no_stress"
		}

		name = strings.ReplaceAll(strings.ToLower(name), "-", ".")
		key := strings.ReplaceAll(strings.ReplaceAll(name, "*", ""), ".", "")

		table[key] = append(table[key], fmt.Sprintf("%s: %s", name, noStress))
		for _, suffix := range customWordSuffixes {
			key := key + suffix
			table[key] = append(table[key], fmt.Sprintf("%s: -%s %s", name, suffix, noStress))
		}
		for _, adposition := range customWordAdpositions {
			key := key + adposition
			table[key] = append(table[key], fmt.Sprintf("%s: -%s %s", name, adposition, noStress))
		}

		if possibleLoanWord := strings.TrimSuffix(name, "ì"); possibleLoanWord != name {
			key := strings.Replace(strings.Replace(possibleLoanWord, "*", "", -1), ".", "", -1)
			table[key] = append(table[key], fmt.Sprintf("%s: %s", name, noStress))
			for _, suffix := range customWordLoanWordSuffixes {
				key := key + suffix
				table[key] = append(table[key], fmt.Sprintf("%s: -%s %s", name, suffix, noStress))
			}
		}
	}

	return &customWordDictionary{table: table, definition: "Custom Name"}
}

type customWordDictionary struct {
	table      map[string][]string
	definition string
}

func (n *customWordDictionary) LookupEntries(word string) ([]Entry, error) {
	entryStrs, ok := n.table[word]
	if !ok {
		return nil, ErrEntryNotFound
	}

	entries := make([]Entry, len(entryStrs))
	for i, entryStr := range entryStrs {
		entries[i] = *ParseEntry(entryStr + ": " + n.definition)
	}

	return entries, nil
}

var customWordSuffixes = []string{
	"l", "ìl",
	"t", "ti", "it",
	"r", "ur", "ru",
	"ri", "ìri",
	"yä", "ä", "ye",
}

var customWordLoanWordSuffixes = []string{
	"ìl",
	"it",
	"ur",
	"ìri",
	"ä",
}

// Hardcoded adpositions for custom words. Generated using this command in node.js repl.
//
// child_process.execSync("fwew \"/list pos has adp\"").toString().split("\n").filter(l => l).map(l => l.split(" ")[1].replace("+", ""));
var customWordAdpositions = []string{
	"äo", "eo", "fa", "fpi", "ftu", "hu", "ìlä", "ka", "kip", "mì", "mungwrr", "na", "ne", "ta",
	"teri", "vay", "fkip", "io", "kxamlä", "lok", "maw", "mìkam", "nemfa", "pxaw", "pxel",
	"sìn", "takip", "uo", "luke", "tafkip", "ro", "sre", "wä", "pxisre", "pximaw", "rofa",
	"few", "lisre", "kam", "kay", "nuä", "sko", "talun", "ftumfa", "yoa", "krrka", "ftuopa",
	"raw", "ken",
}
