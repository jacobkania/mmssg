package filehandling

import (
	"bytes"
	"fmt"
	"github.com/jacobkania/mmssg/errors"
	model "github.com/jacobkania/mmssg/models"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

// WritePages will create the pages output
func WritePages(pages []model.Entry, outputLocation, pageTemplateLocation string) {

	pageContents, err := ioutil.ReadFile(pageTemplateLocation)
	errors.HandleErr(&err, "Couldn't read the page template file", true)

	pageTemplate, err := template.New("page").Parse(string(pageContents))
	errors.HandleErr(&err, "Couldn't parse the page template file", true)

	// parse contents into pages from template
	for _, entry := range pages {
		var outputPath string = outputLocation + strings.TrimSuffix(entry.OriginalFilename, path.Ext(entry.OriginalFilename))
		var outputFileName string = fmt.Sprintf("%s/index.html", outputPath)

		var generatedPage bytes.Buffer
		err = pageTemplate.Execute(&generatedPage, entry)
		errors.HandleErr(&err, fmt.Sprintf("Failed to parse into template: %v", entry.OriginalFilename), true)

		err = os.MkdirAll(outputPath, os.ModePerm)
		errors.HandleErr(&err, "Couldn't create the output directory", true)

		writeLocation := outputFileName
		err = ioutil.WriteFile(writeLocation, generatedPage.Bytes(), 0644)
		errors.HandleErr(&err, fmt.Sprintf("Failed to write to page file to: %v", writeLocation), true)
	}
}
