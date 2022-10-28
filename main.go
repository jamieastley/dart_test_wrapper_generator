package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {

	args := parseArgs()

	f := fileReader{}

	files := parseTestFiles(f, args.Path)
	fmt.Print(files)
}

func parseTestFiles(fr fileReader, path string) []string {
	var relPaths []string
	pattern := regexp.MustCompile("^/\\S+/test/")

	files, err := fr.getTestFiles(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f, "_test.dart") {
			relPath := pattern.ReplaceAllString(f, "")
			relPaths = append(relPaths, relPath)
		}
	}

	return relPaths
}
