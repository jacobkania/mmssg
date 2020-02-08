package main

import (
	"bytes"
	"fmt"
	model "github.com/jacobkania/mmssg/models"
	"github.com/jacobkania/mmssg/utils"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
)

func main() {
	flags := utils.ReadFlags()

	fmt.Println("#################")
	fmt.Println("MMSSG - Most Minimal Static Site Generator")
	fmt.Println("(c) 2020 Jacob R Kania -- https://github.com/jacobkania")
	fmt.Println("#################\n")

	// get locations and read list of files
	inputLocation := utils.ParsePathFromUserInput(flags.InputDir, false)
	outputLocation := utils.ParsePathFromUserInput(flags.OutputDir, false)
	indexTemplateLocation := utils.ParsePathFromUserInput(flags.IndexTemplateFilename, true)
	pageTemplateLocation := utils.ParsePathFromUserInput(flags.PageTemplateFilename, true)

	if flags.PreURL != "" {
		flags.PreURL += "/"
	}

	fmt.Printf("Read posts from directory: %s\n", inputLocation)
	fmt.Printf("Output posts to directory: %s\n", outputLocation)

	if _, err := os.Stat(outputLocation); err != nil {
		fmt.Printf("Note: Output directory does not exist. Creating now...\n")
		err = os.MkdirAll(outputLocation, os.ModePerm)
		utils.HandleErr(&err, "Couldn't create the output directory", true)
	}

	var inputFiles []os.FileInfo
	inputFiles, err := ioutil.ReadDir(inputLocation)
	utils.HandleErr(&err, fmt.Sprintf("Couldn't read the input files from %v", inputLocation), true)

	// generate templates
	indexContents, err := ioutil.ReadFile(indexTemplateLocation)
	utils.HandleErr(&err, "Couldn't read the index template file", true)
	indexTemplate, err := template.New("index").Parse(string(indexContents))
	utils.HandleErr(&err, "Couldn't parse the index template file", true)

	pageContents, err := ioutil.ReadFile(pageTemplateLocation)
	utils.HandleErr(&err, "Couldn't read the page template file", true)
	pageTemplate, err := template.New("page").Parse(string(pageContents))
	utils.HandleErr(&err, "Couldn't parse the page template file", true)

	// extract entries from files
	var pages []model.Entry

	for i := len(inputFiles) - 1; i >= 0; i-- {
		file := inputFiles[i]
		if file.IsDir() {
			continue
		}

		var outputURL string = fmt.Sprintf("%s", strings.TrimSuffix(file.Name(), path.Ext(file.Name())))

		fmt.Printf("Processing file: %s\n", file.Name())
		fileContents, err := ioutil.ReadFile(inputLocation + file.Name())
		utils.HandleErr(&err, fmt.Sprintf("Failed to load: %v", file.Name()), true)

		entry := utils.ParseMeta(string(fileContents), outputURL, file.Name(), flags.PreURL)
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
		utils.HandleErr(&err, fmt.Sprintf("Failed to parse into template: %v", entry.OriginalFilename), true)

		err = os.MkdirAll(outputPath, os.ModePerm)
		utils.HandleErr(&err, "Couldn't create the output directory", true)

		writeLocation := outputFileName
		err = ioutil.WriteFile(writeLocation, generatedPage.Bytes(), 0644)
		utils.HandleErr(&err, fmt.Sprintf("Failed to write to page file to: %v", writeLocation), true)
	}

	// parse contents into index from template

	var generatedIndex bytes.Buffer
	err = indexTemplate.Execute(&generatedIndex, model.IndexData{pages, "/" + flags.PreURL})
	utils.HandleErr(&err, "Failed to parse index page", true)
	err = ioutil.WriteFile(outputLocation+"index.html", generatedIndex.Bytes(), 0644)
	utils.HandleErr(&err, fmt.Sprintf("Failed to write index file to: %v", outputLocation+"index.html"), true)
}
