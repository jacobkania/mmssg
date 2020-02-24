package filehandling

import (
	"encoding/json"
	"fmt"
	"github.com/jacobkania/mmssg/errors"
	model "github.com/jacobkania/mmssg/models"
	"io"
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
)

// ProcessPlugins will take a list of input files and convert them to entries
func ProcessPlugins(entries []model.Entry, pluginDir string) {
	// fmt.Printf("run process plugins on entries")

	// find all plugins to apply
	pluginFiles, err := ioutil.ReadDir(pluginDir)
	errors.HandleErr(&err, fmt.Sprintf("failed to read plugins from: %s", pluginDir), false)
	if err != nil {
		return
	}

	// prepare JSON to be passed to plugin scripts
	entriesJSON, err := json.Marshal(entries)
	errors.HandleErr(&err, "Failed to marshal entries to JSON when running plugins", true)

	// apply plugins to each entry's `Plugins` field
	for i := len(pluginFiles) - 1; i >= 0; i-- {
		file := pluginFiles[i]
		if file.IsDir() || path.Ext(file.Name()) != ".py" {
			continue
		}

		pluginName := strings.TrimSuffix(file.Name(), path.Ext(file.Name()))

		pluginCmd := exec.Command("python3", pluginDir+file.Name())
		stdin, err := pluginCmd.StdinPipe()

		go handlePipe(stdin, string(entriesJSON))

		result, err := pluginCmd.Output()
		errors.HandleErr(&err, fmt.Sprintf("failed to run plugin: %s", pluginName), true)

		err = json.Unmarshal(result, entries)
		errors.HandleErr(&err, fmt.Sprintf("Failed to unmarshal JSON from plugin: %s . This indicates a problem with the plugin.\nPlease remove or fix this plugin. It will be skipped.", pluginName), false)
	}
}

func handlePipe(stdin io.WriteCloser, entriesJSON string) {
	defer stdin.Close()
	io.WriteString(stdin, entriesJSON)
}
