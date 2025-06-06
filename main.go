package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"

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
OuterLoop:
	for _, file := range files {
		if file != "" {
			for _, pattern := range cli.Flags.IgnorePatterns {
				matched, err := filepath.Match(pattern, file)
				if err != nil {
					fmt.Printf("Warning: Invalid ignore pattern '%s': %v\n", pattern, err)
					continue
				}
				if matched {
					continue OuterLoop // Skip this file if it matches an ignore pattern.
				}
			}

			dir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			filePath := filepath.Join(dir, file)
			basename := filepath.Base(filePath)
			var fileContent string
			if helpers.IsTextFileExtension(filepath.Ext(basename)) {
				content, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Println(err)
				}
				fileContent = string(content)
				if cli.Flags.PrependFileName {
					fileContent = "<" + basename + ">" + " " + fileContent
				}
			} else {
				fileContent = basename
			}
			results = append(results, FileData{
				Text: fileContent,
				Name: "<" + basename + ">",
			})
		}
	}
	return results
}

func ProcessFiles(files []FileData) {
	var final string

	if cli.Flags.WithTree {
		final += helpers.GetTreeOutput(cli.Flags.IgnorePatterns)
	}

	for _, content := range files {
		final += content.Text
	}

	if cli.Flags.Minify {
		re := regexp.MustCompile(`\s+`)
		final = re.ReplaceAllString(final, " ")
		final = strings.TrimSpace(final)
	}

	// New OSC52 logic takes precedence
	if cli.Flags.OSC52 {
		helpers.CopyToClipboardOSC52(final)
		fmt.Printf("Sent %d characters to clipboard via OSC52\n", len(final))
	} else if constants.IsCloudEnvironment {
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
	cli.SetupRootCommand()
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	var filePaths []string
	for scanner.Scan() {
		filePaths = append(filePaths, scanner.Text())
	}

	if len(filePaths) > 0 {
		content := GetFiles(filePaths)
		ProcessFiles(content)
	}
}
