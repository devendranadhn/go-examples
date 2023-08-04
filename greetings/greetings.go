package greetings

import "fmt"

// Hello returns a greeting for the named person

// In go A function whose name starts with capital letter can be called by a funtion which is not in the same package.
// This is know ain Go as an exported name
// := is shortcut for declaring and initializing
func Hello(name string) string {

	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message

}
