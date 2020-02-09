package utils

import (
	"github.com/jacobkania/mmssg/errors"
	"os"
)

/*ParsePathFromUserInput will turn user input into an OS path
 *userInput: the text that the user gives
 *isFile: whether the result will be a file or directory
 */
func ParsePathFromUserInput(userInput string, isFile bool) string {
	pwd, err := os.Getwd()
	errors.HandleErr(&err, "Couldn't get current working directory", true)

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
