package main

import (
    "fmt"
    "github.com/Bearox/go-test/greetings"
    "log"
)

func main() {
    log.SetPrefix("greetings:")
    log.SetFlags(0)

    names := []string{
        "weiwen",
        "Bearox",
        "Davin",
    }

    messages, err := greetings.Hellos(names)

    if nil != err {
        log.Fatal(err)
    }
    fmt.Println(messages)
}
