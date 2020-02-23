package utils

import (
	"flag"
	model "github.com/jacobkania/mmssg/models"
)

/*ReadFlags reads the command flags into a model
 */
func ReadFlags() model.Flags {
	optionInputDir := flag.String("i", "", "Input directory")
	optionOutputDir := flag.String("o", "docs", "Output directory")
	optionIndexTemplateFilename := flag.String("t", "index.html", "File to use for index template")
	optionPageTemplateFilename := flag.String("p", "page.html", "File to use for page template")

	optionPreURL := flag.String("u", "", "Leading url path after domain name to use for links")

	optionPluginDir := flag.String("x", "plugins", "Directory that contains plugins, if you choose to use them")

	flag.Parse()

	if *optionPreURL != "" {
		*optionPreURL += "/"
	}

	return model.Flags{
		*optionInputDir,
		*optionOutputDir,
		*optionIndexTemplateFilename,
		*optionPageTemplateFilename,
		*optionPreURL,
		*optionPluginDir,
	}
}
