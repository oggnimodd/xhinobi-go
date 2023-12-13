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
	"xhinobi-go/cli"
	"xhinobi-go/constants"
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

				// Prepend file name if flag is set to true
				if cli.Flags.PrependFileName {
					fileContent = "<" + fileName + ">" + " " + fileContent
				}

			} else {
				fileContent = fileName
			}

			results = append(results, FileData{
				Text: fileContent,
				Name: "<" + fileName + ">",
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

	// Minify
	if cli.Flags.Minify {
		re := regexp.MustCompile(`\s+`)
		final = re.ReplaceAllString(final, " ")
		final = strings.TrimSpace(final)
	}

	if constants.IsCloudEnvironment {
		tempfilename, err := helpers.CreateTempFile(final)
		if err != nil {
			fmt.Println(err)
		}

		cmd, err := helpers.OpenTempFileInCode(tempfilename)
		if err != nil {
			fmt.Println(err)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		err := clipboard.WriteAll(final)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Copied %d characters to clipboard\n", len(final))
		}
	}

}

func main() {
	// Setup root command
	cli.SetupRootCommand()

	// Execute the root command
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Get files
	var filePaths []string

	for scanner.Scan() {
		filePaths = append(filePaths, scanner.Text())
	}

	if len(filePaths) > 0 {
		content := GetFiles(filePaths)
		ProcessFiles(content)
	}
}
