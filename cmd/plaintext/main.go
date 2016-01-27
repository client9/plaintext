package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/client9/plaintext"
)

func main() {
	flag.Parse()
	args := flag.Args()

	// stdin support
	if len(args) == 0 {
		raw, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("Unable to read Stdin: %s", err)
		}
		md, err := plaintext.ExtractorByFilename("stdin")
		if err != nil {
			log.Fatalf("Unable to create parser: %s", err)
		}

		raw = plaintext.StripTemplate(raw)
		os.Stdout.Write(md.Text(raw))
	}

	for _, arg := range args {
		raw, err := ioutil.ReadFile(arg)
		if err != nil {
			log.Fatalf("Unable to read %q: %s", arg, err)
		}
		md, err := plaintext.ExtractorByFilename(arg)
		if err != nil {
			log.Fatalf("Unable to create parser: %s", err)
		}

		raw = plaintext.StripTemplate(raw)
		os.Stdout.Write(md.Text(raw))
		os.Stdout.Write([]byte{'\n'})
	}
}
