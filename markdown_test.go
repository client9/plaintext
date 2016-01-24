package plaintext

import (
	"testing"
)

func TestMD(t *testing.T) {
	cases := []struct {
		text string
		want string
	}{
		{"\nfoo bar\n", "foo bar"},
		{"\nfoo\nbar\n", "foo\nbar"},
		{"*italic*", "italic"},
		{"**bold**", "bold"},
		{"_emphasis_", "emphasis"},
		{"**combo _text_**", "combo text"},
		{"~~strike~~", "strike"},
		{"# heading1", "heading1"},

		// in-line code should be ignored
		{"first `middle` last", "first last"},

		// auto-links really should be ignore, but they get removed in plain-text tokenizer
		{"first http://foobar.com/apple last ", "first http://foobar.com/apple last"},

		// links
		{"foo\n[hello world](http://foobar.com/apple) foo ", "foo hello world foo"},
		{"[Visit GitHub!](https://www.github.com)", "Visit GitHub!"},

		// images
		{"![Image of Yaktocat](https://octodex.github.com/images/yaktocat.png)", "Image of Yaktocat"},
		{"![GitHub Logo](/images/logo.png)", "GitHub Logo"},

		// code blocks
		{"```\ncode\n```\nfoo", "foo"},

		// blockquote
		{"> foo\n> bar\n", "foo\nbar"},

		// entity
		{"&lt;", "<"},
	}

	for pos, tt := range cases {
		mt, err := NewMarkdownText()
		if err != nil {
			t.Fatalf("Unable to run test: %s", err)
		}
		got := string(mt.Text([]byte(tt.text)))
		if tt.want != got {
			t.Errorf("Test %d failed:  want %q, got %q", pos, tt.want, got)
		}
	}
}
