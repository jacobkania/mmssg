package utils

import (
	model "github.com/jacobkania/mmssg/models"
	"gopkg.in/russross/blackfriday.v2"
	"strings"
)

/*ParseMeta takes in a .md file and processes all of the data from it
 *content: the contents of the .md file
 *outputFileURL: URL of the output HTML file that the entry will be generated to
 *originalFilename: File name of the original .md file
 *URLBase: the URL where the base website is located (ie for https://jacobkania.github.io/mmssg, this would be "mmssg")
 */
func ParseMeta(content, outputFileURL, originalFilename, URLBase string) model.Entry {
	byLine := strings.Split(content, "\n")
	// if there is no metadata, just return an Entry with the Body
	if byLine[0] != "---" {
		return model.Entry{string(blackfriday.Run([]byte(content))), nil, nil, "/" + URLBase + outputFileURL, "/" + URLBase, originalFilename}
	}
	// read off metadata
	var entry model.Entry
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
