package plaintext

import (
	"testing"
)

func TestScript(t *testing.T) {
	cases := []struct {
		text string
		want string
	}{
		{"\nfoo1\n# line2bar\nfoo3", "\nfoo1\n\nfoo3"},
		{"\nfoo1\nline2# bar\nfoo3", "\nfoo1\nline2\nfoo3"},
	}
	mt, err := NewScriptText()
	if err != nil {
		t.Fatalf("Unable to run test: %s", err)
	}
	for pos, tt := range cases {
		got := string(mt.Text([]byte(tt.text)))
		if tt.want != got {
			t.Errorf("Test %d failed:  want %q, got %q", pos, tt.want, got)
		}
	}
}
