package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

const usage = `<FILENAME>
Convert the json contents of FILENAME to yaml and print it on stdout. If
FILENAME is '-' or omitted, stdin is read instead.
The input is read completely before re-encoding begins.`

func main() {
	prog := filepath.Base(os.Args[0])

	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s %s\n", prog, usage)
	}

	flag.Parse()
	args := flag.Args()
	if len(args) > 1 {
		log.Fatal("expected at most one argument; the name of a json file")
	}

	b, err := readInput(args)
	if err != nil {
		log.Fatal(err)
	}

	b, err = yaml.JSONToYAML(b)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
	os.Stdout.Write([]byte{'\n'})
}

func readInput(args []string) ([]byte, error) {
	if len(args) > 0 && args[0] != "-" {
		return ioutil.ReadFile(args[0])
	} else {
		return ioutil.ReadAll(os.Stdin)
	}
}
