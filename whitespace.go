package plaintext

// Collapse merges duplicative whitespace
//
// It's not very smart but can be useful to clean up some
// output.
//
// TODO: deal with tabs
func CollapseWhitespace(raw []byte) []byte {
	raw = bytes.Replace(raw, []byte{' ', ' '}, -1)
	raw = bytes.Replace(raw, []byte{'\n', '\n'}, -1)
	return raw
}
