package greetings

import (
    "errors"
	"fmt"
    "math/rand"
    "time"
)

// init sets initial values for variables userd in the function
func init() {
    rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
    formats := []string{
        "Hi, %v. Welcome!",
        "Greet to see you, %v!",
        "Hail, %v! Well met!",
    }

    // return a randomly selected message format by specifying a random index of the slice of formats.
    return formats[rand.Intn(len(formats))]
}

// Hello returns a greetings for the name person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if "" == name {
        return "", errors.New("empty name")
    }

    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }

        messages[name] = message
    }
    return messages, nil
}
