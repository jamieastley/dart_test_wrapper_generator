package main

import (
	"errors"
	"fmt"
	"github.com/hoisie/mustache"
	"log"
	"regexp"
	"strings"
)

const defaultTemplate = `{{#imports}}
   import '{{Path}}' as {{Alias}};
{{/imports}}

void main() {
{{#imports}}
    {{Alias}}.main();
{{/imports}}
}`

func main() {

	args := parseArgs()
	consoleLogger := NewConsoleLogger(args.Verbose)
	f := fileManager{
		Logger: consoleLogger,
	}
	defer consoleLogger.logger.Sync()

	imports, err := ParseTestFiles(&f, args, consoleLogger)
	if err != nil {
		consoleLogger.logger.Error(err)
		log.Fatal(err)
	}

	result := RenderMustache(args.Template, consoleLogger, &imports)

	fileErr := f.writeFile(result, fmt.Sprintf("%s/%s.dart", args.Path, args.OutputFilename))
	if fileErr != nil {
		consoleLogger.logger.Error(err)
		log.Fatal(err)
	}

	fmt.Println("")
	fmt.Printf("Added %d test imports to %s", len(imports), args.OutputFilename)
}

func ParseTestFiles(fr FileManager, args *CliArgs, logger Logger) ([]DartImport, error) {
	var dartImports []DartImport
	pathPattern := regexp.MustCompile("^/\\S+/test/")
	filenamePattern := regexp.MustCompile("[a-zA-Z_]+_test")

	files, err := fr.getTestFiles(args.Path)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return []DartImport{}, errors.New("no files found")
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
	if dartImports == nil {
		return []DartImport{}, errors.New("no test files found")
	}

	return dartImports, nil
}

func RenderMustache(templatePath string, logger *ConsoleLogger, imports *[]DartImport) string {
	data := map[string]*[]DartImport{
		"imports": imports,
	}

	if templatePath != "" {
		logger.Debug("Using mustache template located at", templatePath)
		return mustache.RenderFile(templatePath, data)
	}

	logger.Debug("Using default Dart mustache template")
	return mustache.Render(defaultTemplate, data)
}
