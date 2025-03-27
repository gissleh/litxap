package litxaputil

import "strings"

func prefix(lt string, s ...string) Prefix {
	return Prefix{loseTail: lt, syllableSplit: s}
}

type Prefix struct {
	// Trigger reanalysis, e.g. awkx + -ìl => aw.kxìl
	loseTail string
	// syllableSplit describes how the prefix will be added.
	syllableSplit []string
	// hasLenition indicates that it has lenition.
	hasLenition bool
}

func (p Prefix) withLenition() Prefix {
	p.hasLenition = true
	return p
}

func (p Prefix) Apply(curr []string) ([]string, int) {
	if p.hasLenition {
		_, curr[0] = ApplyLenition(curr[0])
	}

	curr = append(p.syllableSplit[:len(p.syllableSplit):len(p.syllableSplit)], curr...)

	if p.loseTail != "" {
		lostTail := false
		for _, core := range attachableCores {
			if strings.HasPrefix(curr[len(p.syllableSplit)], core) {
				lostTail = true
				curr[len(p.syllableSplit)] = p.loseTail + curr[len(p.syllableSplit)]
			}
		}

		if !lostTail {
			curr[len(p.syllableSplit)-1] += p.loseTail
		}
	}

	return curr, len(p.syllableSplit)
}

func ApplyPrefixes(curr []string, prefixNames []string) ([]string, int) {
	totalOffset := 0
	for i := len(prefixNames) - 1; i >= 0; i-- {
		prefixName := prefixNames[i]
		prefix := findPrefix(prefixName)

		next, n := prefix.Apply(curr)
		curr = next
		totalOffset += n
	}

	return curr, totalOffset
}

func findPrefix(name string) Prefix {
	if prefix, ok := prefixMap[name]; ok {
		return prefix
	}

	return Prefix{loseTail: "", syllableSplit: []string{name}}
}

// tsuk.fmong -> tsuk.fmong
// tsuk.inan -> tsu.ki.nan

var prefixMap = map[string]Prefix{
	"tsuk":   prefix("k", "tsu"),
	"ketsuk": prefix("k", "ke", "tsu"),
	"pe":     prefix("", "pe").withLenition(),
	"me":     prefix("", "me").withLenition(),
	"pxe":    prefix("", "pxe").withLenition(),
	"ay":     prefix("y", "a").withLenition(),
	"pay":    prefix("y", "pa").withLenition(),
	"fay":    prefix("y", "fa").withLenition(),
}
