package plaintext

import (
	"strings"
)

var toStraight *strings.Replacer
var toCurly *strings.Replacer

func init() {
	conversions := []string{
		"\u2013", "-", // en dash
		"\u2014", "-", // em dash
		"\u2018", "'", // left single quotation mark
		"\u2019", "'", // right single quotation mark
		"\u201C", "\"", // left double quotation mark
		"\u201D", "\"", // right double quotation mark
	}
	toStraight = strings.NewReplacer(conversions...)
}

// StraightQuotes converts maybe fancy typographical characters into their
// ASCII equivalent
func StraightQuotes(s string) string {
	return toStraight.Replace(s)
}
