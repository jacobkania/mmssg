package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"flag"
	"strings"
	"gopkg.in/russross/blackfriday.v2"
	"text/template"
)

type Entry struct {
	Body	[]byte
	Meta	map[string]interface{}
}

func handleErr(err *error, message string, doPanic bool) {
	if *err != nil {
		fmt.Println(message)
		if doPanic {
			panic(message)
		}
	}
}

func parseMeta(content string) Entry {
	byLine := strings.Split(content, "\n")
	// if there is no metadata, just return an Entry with the Body
	if byLine[0] != "---" {
		return Entry{blackfriday.Run([]byte(content)), nil}
	}
	// read off metadata
	var entry Entry

	var index int
	for index, line := range byLine {
		if index == 0 { continue }
		if line == "---" { break }
		lineContents := strings.Split(line, ":")
		if len(lineContents) < 2 { panic("^ ^ File's meta is formatted wrong, please view") }

		metaName := strings.TrimSpace(lineContents[0])
		metaContent := strings.TrimSpace(strings.Join(lineContents[1:], ""))

		entry.Meta[metaName] = metaContent;
	}

	remainingText := strings.TrimSpace(strings.Join(byLine[index:], "\n"))
	entry.Body = blackfriday.Run([]byte(remainingText))

	return entry
}

func parseDirectoryFromUserInput(userInput string) string {
	pwd, err := os.Getwd()
	handleErr(&err, "Couldn't get current working directory", true)
	
	if userInput == "" {
		return pwd
	} else if string((userInput)[0]) == "/" {
		return userInput
	}
	return pwd + userInput
}

func main() {
	fmt.Println("#################")
	fmt.Println("MMSSG - Most Minimal Static Site Generator")
	fmt.Println("(c) 2019 Jacob R Kania -- https://github.com/jacobkania")
	fmt.Println("#################\n")

	// options
	optionInputDir := flag.String("i", "", "Input directory")
	optionOutputDir := flag.String("o", "docs", "Output directory")
	optionIndexTemplateFilename := flag.String("t", "index.html", "File to use for index template")
	optionPageTemplateFilename := flag.String("p", "page.html", "File to use for page template")

	flag.Parse()

	// get locations and read list of files
	inputLocation := parseDirectoryFromUserInput(*optionInputDir)
	outputLocation := parseDirectoryFromUserInput(*optionOutputDir)

	fmt.Printf("Reading from directory: %s\n", inputLocation)

	var inputFiles []os.FileInfo
	inputFiles, err := ioutil.ReadDir(inputLocation)
	handleErr(&err, fmt.Sprintf("Couldn't read the input files from %v", inputLocation), true)

	// generate templates
	indexTemplate, err := template.New("index").Parse()

	// parse contents into pages from template
	for _, file := range inputFiles {
		if file.IsDir() { continue }

		fmt.Printf("Processing file: %s\n", file.Name())
		fileContents, err := ioutil.ReadFile(inputLocation + file.Name());
		handleErr(&err, fmt.Sprintf("Failed to load %v", file.Name()), true)

		entry := parseMeta(string(fileContents))
	}

}