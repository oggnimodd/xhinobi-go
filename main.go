package main

import (
	// std
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	// 3rd party
	"github.com/atotto/clipboard"

	// local
	"xhinobi-go/helpers"
)

type FileData struct {
	Text string
	Name string
}

func GetFiles(files []string) []FileData {
	var results []FileData

	// Filter files
	for _, file := range files {
		if file != "" {
			dir, err := os.Getwd()
			if err != nil {
				// Panic
				panic(err)
			}

			// Join current directory with the file name
			filePath := filepath.Join(dir, file)
			// Get the file name basename
			fileName := filepath.Base(filePath)

			var fileContent string
			if helpers.IsTextFileExtension(filepath.Ext(fileName)) {
				content, err := os.ReadFile(filePath)

				if err != nil {
					fmt.Println(err)
				}

				fileContent = string(content)
			} else {
				fileContent = fileName
			}

			results = append(results, FileData{
				Text: fileContent,
				Name: fileName,
			})
		}
	}

	return results
}

func ProcessFiles(files []FileData) {
	var final string

	for _, content := range files {
		final += content.Text
	}

	re := regexp.MustCompile(`\s+`)
	final = re.ReplaceAllString(final, " ")
	final = strings.TrimSpace(final)

	// Copy final to clipboard
	err := clipboard.WriteAll(final)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Copied %d characters to clipboard\n", len(final))
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	var filePaths []string

	for scanner.Scan() {
		filePaths = append(filePaths, scanner.Text())
	}

	content := GetFiles(filePaths)
	ProcessFiles(content)
}
