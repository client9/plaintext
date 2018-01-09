package plaintext

import (
	"regexp"
)

// Collapse merges duplicative whitespace
//
// It's not very smart but can be useful to clean up some
// output.
//
func CollapseWhitespace(raw []byte) []byte {
	re2 := regexp.MustCompile(`[ \t]+`)
	re1 := regexp.MustCompile(`\n `)
	re3 := regexp.MustCompile(`\n+`)

	raw = re2.ReplaceAll(raw, []byte{' '})
	raw = re1.ReplaceAll(raw, []byte{'\n'})
	raw = re3.ReplaceAll(raw, []byte{'\n'})

	return raw
}
