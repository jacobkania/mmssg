package filehandling

import (
	"fmt"
	"github.com/jacobkania/mmssg/errors"
	model "github.com/jacobkania/mmssg/models"
	"github.com/jacobkania/mmssg/utils"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

// ExtractEntries will take a list of input files and convert them to entries
func ExtractEntries(inputFiles []os.FileInfo, inputDir, preURL string) []model.Entry {
	var pages []model.Entry

	for i := len(inputFiles) - 1; i >= 0; i-- {
		file := inputFiles[i]
		if file.IsDir() {
			continue
		}

		var outputURL string = fmt.Sprintf("%s", strings.TrimSuffix(file.Name(), path.Ext(file.Name())))

		fmt.Printf("Processing file: %s\n", file.Name())
		fileContents, err := ioutil.ReadFile(inputDir + file.Name())
		errors.HandleErr(&err, fmt.Sprintf("Failed to load: %v", file.Name()), true)

		entry := utils.ParseMeta(string(fileContents), outputURL, file.Name(), preURL)
		pages = append(pages, entry)
	}

	// sort entries by published date
	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Meta["PublishedDate"] > pages[j].Meta["PublishedDate"]
	})

	return pages
}
