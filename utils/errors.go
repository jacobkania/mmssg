package utils

import "fmt"

/*HandleErr will handle an error in a simple way
 *err: the error object
 *message: the error message to print in case of an error
 *doPanic: whether or not the error should cause the program to panic
 */
func HandleErr(err *error, message string, doPanic bool) {
	if *err != nil {
		if doPanic {
			fmt.Println("FATAL ERROR: \n" + message)
			panic(message)
		} else {
			fmt.Println(message)
		}
	}
}
