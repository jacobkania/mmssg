package main

import (
	"bytes"
	"flag"
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
)

// Entry is a blog post entry
type Entry struct {
	Body             string
	Meta             map[string]string
	URL              string
	HomeURL          string
	OriginalFilename string
}

// IndexData is a list of entries
type IndexData struct {
	Entries []Entry
	HomeURL string
}

func handleErr(err *error, message string, doPanic bool) {
	if *err != nil {
		if doPanic {
			fmt.Println("FATAL ERROR: \n" + message)
			panic(message)
		} else {
			fmt.Println(message)
		}
	}
}

func parseMeta(content, outputFileURL, originalFilename, URLBase string) Entry {
	byLine := strings.Split(content, "\n")
	// if there is no metadata, just return an Entry with the Body
	if byLine[0] != "---" {
		return Entry{string(blackfriday.Run([]byte(content))), nil, "/" + URLBase + outputFileURL, "/" + URLBase, originalFilename}
	}
	// read off metadata
	var entry Entry
	entry.Meta = make(map[string]string)

	var indexCount int = 0
	for index, line := range byLine {
		indexCount++
		if index == 0 {
			continue
		}
		if line == "---" {
			break
		}
		lineContents := strings.Split(line, ":")
		if len(lineContents) < 2 {
			panic("^ ^ File's meta is formatted wrong, please view")
		}

		metaName := strings.TrimSpace(lineContents[0])
		metaContent := strings.TrimSpace(strings.Join(lineContents[1:], ""))

		entry.Meta[metaName] = metaContent
	}

	remainingText := strings.TrimSpace(strings.Join(byLine[indexCount:], "\n"))
	entry.Body = string(blackfriday.Run([]byte(remainingText)))

	entry.URL = "/" + URLBase + outputFileURL
	entry.HomeURL = "/" + URLBase
	entry.OriginalFilename = originalFilename

	return entry
}

func parseDirectoryFromUserInput(userInput string, isFile bool) string {
	pwd, err := os.Getwd()
	handleErr(&err, "Couldn't get current working directory", true)

	pwd += "/"
	if !isFile {
		userInput += "/"
	}

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

	optionPreURL := flag.String("u", "", "Leading url path after domain name to use for links")

	flag.Parse()

	// get locations and read list of files
	inputLocation := parseDirectoryFromUserInput(*optionInputDir, false)
	outputLocation := parseDirectoryFromUserInput(*optionOutputDir, false)
	indexTemplateLocation := parseDirectoryFromUserInput(*optionIndexTemplateFilename, true)
	pageTemplateLocation := parseDirectoryFromUserInput(*optionPageTemplateFilename, true)

	if *optionPreURL != "" {
		*optionPreURL += "/"
	}

	fmt.Printf("Read posts from directory: %s\n", inputLocation)
	fmt.Printf("Output posts to directory: %s\n", outputLocation)

	if _, err := os.Stat(outputLocation); err != nil {
		fmt.Printf("Note: Output directory does not exist. Creating now...\n")
		err = os.MkdirAll(outputLocation, os.ModePerm)
		handleErr(&err, "Couldn't create the output directory", true)
	}

	var inputFiles []os.FileInfo
	inputFiles, err := ioutil.ReadDir(inputLocation)
	handleErr(&err, fmt.Sprintf("Couldn't read the input files from %v", inputLocation), true)

	// generate templates
	indexContents, err := ioutil.ReadFile(indexTemplateLocation)
	handleErr(&err, "Couldn't read the index template file", true)
	indexTemplate, err := template.New("index").Parse(string(indexContents))
	handleErr(&err, "Couldn't parse the index template file", true)

	pageContents, err := ioutil.ReadFile(pageTemplateLocation)
	handleErr(&err, "Couldn't read the page template file", true)
	pageTemplate, err := template.New("page").Parse(string(pageContents))
	handleErr(&err, "Couldn't parse the page template file", true)

	// extract entries from files
	var pages []Entry

	for i := len(inputFiles) - 1; i >= 0; i-- {
		file := inputFiles[i]
		if file.IsDir() {
			continue
		}

		var outputURL string = fmt.Sprintf("%s", strings.TrimSuffix(file.Name(), path.Ext(file.Name())))

		fmt.Printf("Processing file: %s\n", file.Name())
		fileContents, err := ioutil.ReadFile(inputLocation + file.Name())
		handleErr(&err, fmt.Sprintf("Failed to load: %v", file.Name()), true)

		entry := parseMeta(string(fileContents), outputURL, file.Name(), *optionPreURL)
		pages = append(pages, entry)
	}

	// sort entries by published date
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Meta["PublishedDate"] > pages[j].Meta["PublishedDate"]
	})

	// parse contents into pages from template
	for _, entry := range pages {
		var outputPath string = outputLocation + strings.TrimSuffix(entry.OriginalFilename, path.Ext(entry.OriginalFilename))
		var outputFileName string = fmt.Sprintf("%s/index.html", outputPath)

		var generatedPage bytes.Buffer
		err = pageTemplate.Execute(&generatedPage, entry)
		handleErr(&err, fmt.Sprintf("Failed to parse into template: %v", entry.OriginalFilename), true)

		err = os.MkdirAll(outputPath, os.ModePerm)
		handleErr(&err, "Couldn't create the output directory", true)

		writeLocation := outputFileName
		err = ioutil.WriteFile(writeLocation, generatedPage.Bytes(), 0644)
		handleErr(&err, fmt.Sprintf("Failed to write to page file to: %v", writeLocation), true)
	}

	// parse contents into index from template

	var generatedIndex bytes.Buffer
	err = indexTemplate.Execute(&generatedIndex, IndexData{pages, "/" + *optionPreURL})
	handleErr(&err, "Failed to parse index page", true)
	err = ioutil.WriteFile(outputLocation+"index.html", generatedIndex.Bytes(), 0644)
	handleErr(&err, fmt.Sprintf("Failed to write index file to: %v", outputLocation+"index.html"), true)
}
