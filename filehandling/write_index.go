package filehandling

import (
	"bytes"
	"fmt"
	"github.com/jacobkania/mmssg/errors"
	model "github.com/jacobkania/mmssg/models"
	"io/ioutil"
	"text/template"
)

// WriteIndex will write out the index file
func WriteIndex(data model.IndexData, outputLocation, indexTemplateLocation string) {
	indexContents, err := ioutil.ReadFile(indexTemplateLocation)
	errors.HandleErr(&err, "Couldn't read the index template file", true)

	indexTemplate, err := template.New("index").Parse(string(indexContents))
	errors.HandleErr(&err, "Couldn't parse the index template file", true)

	var generatedIndex bytes.Buffer
	err = indexTemplate.Execute(&generatedIndex, data)
	errors.HandleErr(&err, "Failed to parse index page", true)

	err = ioutil.WriteFile(outputLocation+"index.html", generatedIndex.Bytes(), 0644)
	errors.HandleErr(&err, fmt.Sprintf("Failed to write index file to: %v", outputLocation+"index.html"), true)
}
