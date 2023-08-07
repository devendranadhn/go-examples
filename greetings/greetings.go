package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person

// In go A function whose name starts with capital letter can be called by a funtion which is not in the same package.
// This is know ain Go as an exported name
// := is shortcut for declaring and initializing
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Empty Name")
	}

	//var message string
	//message = fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	message = fmt.Sprintf(randomFormat()[rand.Intn(3)], name)
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {

	// A map to associate names with messages
	messages := make(map[string]string)

	// Loop thorough the received names of messages and prepare the Hello message
	for _, name := range names {
		message, error := Hello(name)
		if error != nil {
			return nil, error
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() []string {

	// A slice of message formats
	formats := []string{
		"Hi %v welcome!",
		"Great to see you %v!",
		"Hail %v well met!",
	}

	return formats

}
