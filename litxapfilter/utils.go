package litxapfilter

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func findPrefix(s string, prefixes []string) *string {
	for _, prefix := range prefixes {
		if prefixInS, ok := hasPrefixFold(s, prefix); ok {
			return &prefixInS
		}
	}

	return nil
}

func hasPrefixFold(s string, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return prefix, true
	}

	currS := s
	currPrefix := prefix
	for {
		if currS == "" {
			return "", false
		}

		lrS, lrnS := utf8.DecodeRuneInString(currS)
		lrPrefix, lrnPrefix := utf8.DecodeRuneInString(currPrefix)
		if unicode.ToLower(lrS) != unicode.ToLower(lrPrefix) {
			return "", false
		}

		currS = currS[lrnS:]
		currPrefix = currPrefix[lrnPrefix:]
		if currPrefix == "" {
			return s[:len(s)-len(currS)], true
		}
	}
}

func hasSuffixFold(s string, suffix string) (string, bool) {
	if strings.HasSuffix(s, suffix) {
		return suffix, true
	}

	currS := s
	currSuffix := suffix
	for {
		if currS == "" {
			return "", false
		}

		lrS, lrnS := utf8.DecodeLastRuneInString(currS)
		lrSuffix, lrnSuffix := utf8.DecodeLastRuneInString(currSuffix)
		if unicode.ToLower(lrS) != unicode.ToLower(lrSuffix) {
			return "", false
		}

		currS = currS[:len(currS)-lrnS]
		currSuffix = currSuffix[:len(currSuffix)-lrnSuffix]
		if currSuffix == "" {
			return s[len(currS):], true
		}
	}
}
