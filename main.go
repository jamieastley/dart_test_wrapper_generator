package main

import (
	"fmt"
	"github.com/hoisie/mustache"
	"log"
	"regexp"
	"strings"
)

func main() {

	args := parseArgs()
	consoleLogger := NewConsoleLogger(args.Verbose)
	f := fileManager{
		Logger: consoleLogger,
	}
	defer consoleLogger.logger.Sync()

	imports, err := parseTestFiles(f, args, consoleLogger)
	if err != nil {
		consoleLogger.logger.Error(err)
		log.Fatal(err)
	}

	result := mustache.RenderFile("wrapper_test.dart.mustache", CreateMustacheData(imports))

	fileErr := f.writeFile(result, fmt.Sprintf("%s/%s.dart", args.Path, args.OutputFilename))
	if fileErr != nil {
		consoleLogger.logger.Error(err)
		log.Fatal(err)
	}

	fmt.Println("")
	consoleLogger.Info("Added", len(imports), "test imports to", args.OutputFilename)
}

func parseTestFiles(fr fileManager, args *CliArgs, logger *ConsoleLogger) ([]DartImport, error) {
	var dartImports []DartImport
	pathPattern := regexp.MustCompile("^/\\S+/test/")
	filenamePattern := regexp.MustCompile("[a-zA-Z_]+_test")

	files, err := fr.getTestFiles(args.Path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if strings.Contains(f, "_test.dart") {
			relPath := pathPattern.ReplaceAllString(f, "")
			importAlias := filenamePattern.FindString(relPath)

			if importAlias != args.OutputFilename {
				logger.Debug(fmt.Sprintf("Adding %s", importAlias))
				dartImports = append(dartImports, DartImport{
					Alias: importAlias,
					Path:  relPath,
				})
			}
		}
	}

	return dartImports, nil
}
