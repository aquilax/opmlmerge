package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/gilliek/go-opml/opml"
)

func walkOutlines(outlines []opml.Outline, cb func(opml.Outline)) {
	for _, outline := range outlines {
		cb(outline)
		if len(outline.Outlines) > 0 {
			walkOutlines(outline.Outlines, cb)
		}
	}
}

func main() {
	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "Usage: %s [FILE]...\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	fileNames := flag.Args()
	var first *opml.OPML
	var err error
	var existing []string
	if len(fileNames) < 1 {
		flag.Usage()
		os.Exit(0)
	}
	for i, fileName := range fileNames {
		if i == 0 {
			first, err = opml.NewOPMLFromFile(fileName)
			if err != nil {
				log.Fatal(fmt.Errorf("error reading file %s, %w", fileName, err))
			}
			walkOutlines(first.Outlines(), func(outline opml.Outline) {
				if outline.XMLURL != "" {
					existing = append(existing, outline.XMLURL)
				}
			})
		} else {
			rest, err := opml.NewOPMLFromFile(fileName)
			if err != nil {
				log.Fatal(fmt.Errorf("error reading file %s, %w", fileName, err))
			}
			walkOutlines(rest.Outlines(), func(outline opml.Outline) {
				if outline.XMLURL != "" {
					if slices.Contains(existing, outline.XMLURL) {
						return
					}
					first.Body.Outlines = append(first.Body.Outlines, outline)
				}
			})
		}
	}
	xml, err := first.XML()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xml)

}
