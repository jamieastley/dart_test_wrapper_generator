package main

import (
	"os"
	"path/filepath"
)

type FileManager interface {
	getTestFiles(path string) ([]string, error)
	writeFile(output string, path string) error
}

type fileManager struct {
	Logger Logger
}

func (t *fileManager) getTestFiles(path string) ([]string, error) {
	var files []string
	t.Logger.Debug("Looking for test files in", path)
	fileErr := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, currentPath)
		return nil
	})

	return files, fileErr
}

func (t *fileManager) writeFile(output string, filePath string) error {
	t.Logger.Debug("Writing output file to", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, writeErr := file.WriteString(output)
	if writeErr != nil {
		return writeErr
	}
	t.Logger.Debug("Successfully wrote output file to", filePath)

	return nil
}
