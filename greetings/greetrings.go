package greetings

import (
    "errors"
	"fmt"
)

// Hello returns a greetings for the name person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if "" == name {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf("Hi %v. Welcome!", name)
    return message, nil
}
