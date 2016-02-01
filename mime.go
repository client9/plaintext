package plaintext

import (
	"strings"
)

// returns the mimetype of the full filename if none
func getSuffix(filename string) string {
	idx := strings.LastIndex(filename, ".")
	if idx == -1 || idx+1 == len(filename) {
		return filename
	}
	return filename[idx+1:]
}

// ExtractorByFilename returns an plaintext extractor based on
// filename heuristic
func ExtractorByFilename(filename string) (Extractor, error) {
	var e Extractor
	var err error
	switch getSuffix(filename) {
	case "md":
		e, err = NewMarkdownText()
	case "html":
		e, err = NewHTMLText()
	case "go", "h", "c", "java", "hxx", "cxx":
		e, err = NewGolangText()
	case "py", "sh", "pl", "Makefile", "Dockerfile":
		e, err = NewScriptText()
	default:
		e, err = NewIdentity()
	}
	if err != nil {
		return nil, err
	}
	return e, nil
}
