package plaintext

type Extractor interface {
	Text([]byte) []byte
}
