package main

import (
	"fmt"
	"github.com/jacobkania/mmssg/errors"
	"github.com/jacobkania/mmssg/filehandling"
	model "github.com/jacobkania/mmssg/models"
	"github.com/jacobkania/mmssg/utils"
	"io/ioutil"
	"os"
)

func main() {
	flags := utils.ReadFlags()

	fmt.Println("#################")
	fmt.Println("MMSSG - Most Minimal Static Site Generator")
	fmt.Println("(c) 2020 Jacob R Kania")
	fmt.Println("Original source can be found at: https://github.com/jacobkania/mmssg")
	fmt.Println("#################\n")

	// get locations from user flags
	inputLocation := utils.ParsePathFromUserInput(flags.InputDir, false)
	outputLocation := utils.ParsePathFromUserInput(flags.OutputDir, false)
	indexTemplateLocation := utils.ParsePathFromUserInput(flags.IndexTemplateFilename, true)
	pageTemplateLocation := utils.ParsePathFromUserInput(flags.PageTemplateFilename, true)

	pluginLocation := utils.ParsePathFromUserInput(flags.PluginDir, false)

	fmt.Printf("Read posts from directory: %s\n", inputLocation)
	fmt.Printf("Output posts to directory: %s\n", outputLocation)
	fmt.Printf("---------------\n")

	if _, err := os.Stat(outputLocation); err != nil {
		fmt.Printf("Note: Output directory does not exist. Creating now...\n")
		err = os.MkdirAll(outputLocation, os.ModePerm)
		errors.HandleErr(&err, "Couldn't create the output directory", true)
	}

	var inputFiles []os.FileInfo
	inputFiles, err := ioutil.ReadDir(inputLocation)
	errors.HandleErr(&err, fmt.Sprintf("Couldn't read the input files from %v", inputLocation), true)

	//

	var pages []model.Entry = filehandling.ExtractEntries(inputFiles, inputLocation, flags.PreURL)

	filehandling.ProcessPlugins(pages, pluginLocation)

	filehandling.WritePages(pages, outputLocation, pageTemplateLocation)

	filehandling.WriteIndex(model.IndexData{pages, "/" + flags.PreURL}, outputLocation, indexTemplateLocation)

}
