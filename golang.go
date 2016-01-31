package plaintext

import (
	"bytes"
	"text/scanner"
)

// GolangText extracts plaintext from Golang and other similar C or Java like files
type GolangText struct {
}

// NewGolangText creates a new extractor
func NewGolangText() (*GolangText, error) {
	return &GolangText{}, nil
}

// Text satifies the Extractor interface
//
//ReplaceGo is a specialized routine for correcting Golang source
// files.  Currently only checks comments, not identifiers for
// spelling.
//
// Other items:
//   - check strings, but need to ignore
//      * import "statements" blocks
//      * import ( "blocks" )
//   - skip first comment (line 0) if build comment
//
func (p *GolangText) Text(raw []byte) []byte {
	out := bytes.Buffer{}
	s := scanner.Scanner{}
	s.Init(bytes.NewReader(raw))
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanChars | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments
	for {
		switch s.Scan() {
		case scanner.Comment:
			out.WriteString(s.TokenText())
			out.WriteByte('\n')
		case scanner.EOF:
			return out.Bytes()
		}
	}
}
