package plaintext

import (
	"github.com/russross/blackfriday"
)

// Options specifies options for formatting.
type MarkdownText struct {
	Extensions int
	Extractor  Extractor
}

// MarkdownText returns plaintext from a markdown file
func NewMarkdownText(options ...func(*MarkdownText) error) (*MarkdownText, error) {
	// Default is GitHub Flavored Markdown-like extensions.
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS

	processor := MarkdownText{}
	processor.Extensions = extensions

	for _, option := range options {
		err := option(&processor)
		if err != nil {
			return nil, err
		}
	}

	// if an HTML extractor isn't set up, use default
	if processor.Extractor == nil {
		e, err := NewHTMLText()
		if err != nil {
			return nil, err
		}
		processor.Extractor = e
	}

	return &processor, nil
}

func (p *MarkdownText) Text(text []byte) []byte {
	mp := blackfriday.HtmlRenderer(0, "", "")
	out := blackfriday.Markdown(text, mp, p.Extensions)
	return p.Extractor.Text(out)
}
