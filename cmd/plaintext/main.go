package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/client9/plaintext"
)

func main() {
	extension := flag.String("s", "", "over-ride file suffix to determine parser")
	collapse := flag.Bool("c", false, "collapse whitespace")

	flag.Parse()
	ext := *extension
	if ext != "" && ext[0] != '.' {
		ext = "." + ext
	}
	args := flag.Args()

	// stdin support
	if len(args) == 0 {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Unable to read Stdin: %s", err)
		}
		md, err := plaintext.ExtractorByFilename("stdin" + *extension)
		if err != nil {
			log.Fatalf("Unable to create parser: %s", err)
		}

		raw = plaintext.StripTemplate(raw)
		raw = md.Text(raw)
		if *collapse {
			raw = plaintext.CollapseWhitespace(raw)
		}
		os.Stdout.Write(raw)
	}

	for _, arg := range args {
		raw, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatalf("Unable to read %q: %s", arg, err)
		}
		md, err := plaintext.ExtractorByFilename(arg + *extension)
		if err != nil {
			log.Fatalf("Unable to create parser: %s", err)
		}

		raw = plaintext.StripTemplate(raw)
		raw = md.Text(raw)
		if *collapse {
			raw = plaintext.CollapseWhitespace(raw)
		}
		os.Stdout.Write(raw)
		os.Stdout.Write([]byte{'\n'})
	}
}
