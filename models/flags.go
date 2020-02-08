package models

// Flags contains the flag information from the program start
type Flags struct {
	InputDir              string
	OutputDir             string
	IndexTemplateFilename string
	PageTemplateFilename  string
	PreURL                string
}
